import { cookies } from "next/headers";
export const runtime = "edge";

const GHOST_RATELIMIT_COOKIE = "GHOST_RATELIMIT";

export const POST = async (_req: Request): Promise<Response> => {
  cookies().delete(GHOST_RATELIMIT_COOKIE);
  return new Response("ok");
};
