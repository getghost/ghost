import { SchemaError } from "@ghost/error";
import type { PermissionQuery } from "@ghost/rbac";
import type { Context } from "hono";
import { GhostApiError } from "../errors";
import type { HonoEnv } from "../hono/env";
import { DisabledWorkspaceError } from "../keys/service";

/**
 * rootKeyAuth takes the bearer token from the request and verifies the key
 *
 * if the key doesnt exist, isn't valid or isn't a root key, an error is thrown, which gets handled
 * automatically by hono
 */
export async function rootKeyAuth(c: Context<HonoEnv>, permissionQuery?: PermissionQuery) {
  const authorization = c.req.header("authorization")?.replace("Bearer ", "");
  if (!authorization) {
    throw new GhostApiError({ code: "UNAUTHORIZED", message: "key required" });
  }

  const { keyService, analytics } = c.get("services");
  const { val: rootKey, err } = await keyService.verifyKey(c, {
    key: authorization,
    permissionQuery,
  });

  if (err) {
    switch (true) {
      case err instanceof SchemaError:
        throw new GhostApiError({
          code: "BAD_REQUEST",
          message: err.message,
        });
      case err instanceof DisabledWorkspaceError:
        throw new GhostApiError({
          code: "FORBIDDEN",
          message: "workspace is disabled",
        });
    }
    throw new GhostApiError({
      code: "INTERNAL_SERVER_ERROR",
      message: err.message,
    });
  }

  if (!rootKey.key) {
    throw new GhostApiError({
      code: "UNAUTHORIZED",
      message: "key not found",
    });
  }

  // if we have identified the key, we can send the analytics event
  // otherwise, they likely sent garbage to us and we can't associate it with anything

  c.executionCtx.waitUntil(
    analytics.insertKeyVerification({
      workspace_id: rootKey.key.workspaceId,
      key_id: rootKey.key.id,
      time: Date.now(),
      outcome: rootKey.code ?? "VALID",
      key_space_id: rootKey.key.keyAuthId,
      // @ts-expect-error - the cf object will be there on cloudflare
      region: c.req.cf?.region,
      request_id: c.get("requestId"),
    }),
  );

  if (!rootKey.valid) {
    throw new GhostApiError({
      code: rootKey.code,
      message: "message" in rootKey && rootKey.message ? rootKey.message : "unauthorized",
    });
  }
  if (!rootKey.isRootKey) {
    throw new GhostApiError({
      code: "UNAUTHORIZED",
      message: "root key required",
    });
  }

  return rootKey;
}