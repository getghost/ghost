import { createConnection } from "@/pkg/db";
import { ConsoleLogger } from "@ghost/worker-logging";

import { newId } from "@ghost/id";
import type { MiddlewareHandler } from "hono";
import type { HonoEnv } from "../hono/env";
import { type Metrics, NoopMetrics } from "../metrics";
import { LogdrainMetrics } from "../metrics/logdrain";

/**
 * Initialize all services.
 *
 * Call this once before any hono handlers run.
 */
export function init(): MiddlewareHandler<HonoEnv> {
  return async (c, next) => {
    const requestId = newId("request");
    c.res.headers.set("Ghost-Request-Id", requestId);
    c.set("requestId", requestId);

    const db = createConnection(c.env);

    const metrics: Metrics = c.env.EMIT_METRICS_LOGS
      ? new LogdrainMetrics({ requestId, environment: c.env.ENVIRONMENT })
      : new NoopMetrics();

    const logger = new ConsoleLogger({
      requestId,
      application: "semantic-cache",
      environment: c.env.ENVIRONMENT,
    });

    c.set("services", {
      db,
      metrics,
      logger,
    });

    await next();
  };
}
