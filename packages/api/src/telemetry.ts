import { version } from "../package.json";
import type { GhostOptions } from "./client";

export type Telemetry = {
  /**
   * Ghost-Telemetry-Sdk
   * @example @ghost/api@v1.1.1
   */
  sdkVersions: string[];
  /**
   * Ghost-Telemetry-Platform
   * @example cloudflare
   */
  platform?: string;
  /**
   * Ghost-Telemetry-Runtime
   * @example node@v18
   */
  runtime?: string;
};

export function getTelemetry(opts: GhostOptions): Telemetry | null {
  let platform: string | undefined;
  let runtime: string | undefined;
  const sdkVersions = [`@ghost/api@${version}`];

  try {
    if (typeof process !== "undefined") {
      if (process.env.GHOST_DISABLE_TELEMETRY) {
        return null;
      }
      platform = process.env.VERCEL ? "vercel" : process.env.AWS_REGION ? "aws" : undefined;

      // @ts-ignore
      if (typeof EdgeRuntime !== "undefined") {
        runtime = "edge-light";
      } else {
        runtime = `node@${process.version}`;
      }
    }

    if (opts.wrapperSdkVersion) {
      sdkVersions.push(opts.wrapperSdkVersion);
    }
  } catch (_error) {}

  return { platform, runtime, sdkVersions };
}
