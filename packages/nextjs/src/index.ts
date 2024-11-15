import { type ErrorResponse, Ghost } from "@ghost/api";
import { type NextRequest, NextResponse } from "next/server";

import { version } from "../package.json";

export type WithGhostConfig = {
  /**
   * The apiId to verify against.
   *
   * This will be required soon.
   */
  apiId?: string;

  /**
   *
   * By default telemetry data is enabled, and sends:
   * runtime (Node.js / Edge)
   * platform (Node.js / Vercel / AWS)
   * SDK version
   */
  disableTelemetry?: boolean;

  /**
   * How to get the key from the request
   * Usually the key is provided in an `Authorization` header, but you can do what you want.
   *
   * Return the key as string, or null if it doesn't exist.
   *
   * You can also override the response given to the caller by returning a `NextResponse`
   *
   * @default `req.headers.get("authorization")?.replace("Bearer ", "") ?? null`
   */
  getKey?: (req: NextRequest) => string | null | Response | NextResponse;

  /**
   * Automatically return a custom response when a key is invalid
   */
  handleInvalidKey?: (
    req: NextRequest,
    result: GhostContext,
  ) => Response | NextResponse | Promise<Response> | Promise<NextResponse>;

  /**
   * What to do if things go wrong
   */
  onError?: (
    req: NextRequest,
    err: ErrorResponse["error"],
  ) => Response | NextResponse | Promise<Response> | Promise<NextResponse>;
};

type VerifyResponse = Awaited<ReturnType<InstanceType<typeof Ghost>["keys"]["verify"]>>;
export type GhostContext = VerifyResponse["result"];

export type NextContext = { params: Promise<Record<string, string | string[]>> };

export type NextRequestWithGhostContext = NextRequest & { ghost: GhostContext };

export function withGhost<TContext extends NextContext = NextContext>(
  handler: (
    req: NextRequestWithGhostContext,
    context: TContext,
  ) => Response | NextResponse | Promise<Response | NextResponse>,
  config?: WithGhostConfig,
) {
  return async (req: NextRequest, context: TContext) => {
    /**
     * Get key from request and return a response early if not found
     */
    const key = config?.getKey
      ? config.getKey(req)
      : req.headers.get("authorization")?.replace("Bearer ", "") ?? null;
    if (key === null) {
      return NextResponse.json({ error: "unauthorized" }, { status: 401 });
    }
    if (typeof key !== "string") {
      return key;
    }

    const ghost = new Ghost({
      rootKey: "public",
      wrapperSdkVersion: `@ghost/nextjs@${version}`,
      disableTelemetry: config?.disableTelemetry,
    });

    const res = await ghost.keys.verify(config?.apiId ? { key, apiId: config.apiId } : { key });

    if (res.error) {
      if (config?.onError) {
        return config.onError(req, res.error);
      }
      console.error(
        `ghost error: [CODE: ${res.error.code}] - [TRACE: ${res.error.requestId}] - ${res.error.message} - read more at ${res.error.docs}`,
      );
      return new NextResponse("Internal Server Error", { status: 500 });
    }

    if (!res.result.valid) {
      if (config?.handleInvalidKey) {
        return config.handleInvalidKey(req, res.result);
      }

      return new NextResponse("Unauthorized", { status: 500 });
    }

    // @ts-ignore
    req.ghost = res.result;

    return handler(req as NextRequestWithGhostContext, context);
  };
}
