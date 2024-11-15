import { env } from "@/lib/env";
import { Ratelimit as GhostRatelimit } from "@ghost/ratelimit";
import { Ratelimit as UpstashRatelimit } from "@upstash/ratelimit";
import { Redis } from "@upstash/redis";
import { cookies } from "next/headers";
import { z } from "zod";

export const runtime = "edge";

const GHOST_RATELIMIT_COOKIE = "GHOST_RATELIMIT";

export const POST = async (req: Request): Promise<Response> => {
  const { limit, duration } = z
    .object({
      limit: z.number().int(),
      duration: z.enum(["1s", "10s", "60s", "5m"]),
    })
    .parse(await req.json());

  const ghostSync = new GhostRatelimit({
    namespace: "ratelimit-demo-sync",
    rootKey: env().RATELIMIT_DEMO_ROOT_KEY!,
    limit,
    duration,
    async: false,
  });
  const ghostAsync = new GhostRatelimit({
    namespace: "ratelimit-demo-async",
    rootKey: env().RATELIMIT_DEMO_ROOT_KEY!,
    limit,
    duration,
    async: true,
  });
  const upstash = new UpstashRatelimit({
    redis: Redis.fromEnv(),
    limiter: UpstashRatelimit.fixedWindow(limit, duration),
  });

  let id: string = crypto.randomUUID();
  const c = cookies().get(GHOST_RATELIMIT_COOKIE);
  if (c) {
    id = c.value;
  } else {
    cookies().set(GHOST_RATELIMIT_COOKIE, id, {
      maxAge: 60 * 60 * 24,
    });
  }

  const t1 = performance.now();
  const [ghostSyncResponse, ghostAsyncResponse, upstashResponse] = await Promise.all([
    ghostSync
      .limit(`${id}-ghost-sync`)
      .then((res) => ({ ...res, latency: performance.now() - t1 })),
    ghostAsync
      .limit(`${id}-ghost-async`)
      .then((res) => ({ ...res, latency: performance.now() - t1 })),
    upstash.limit(id).then((res) => ({ ...res, latency: performance.now() - t1 })),
  ]);

  return Response.json({
    time: Date.now(),
    ghostSync: ghostSyncResponse,
    ghostAsync: ghostAsyncResponse,
    upstash: upstashResponse,
  });
};
