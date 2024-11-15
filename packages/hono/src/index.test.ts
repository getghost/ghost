import { Hono } from "hono";
import { http, HttpResponse } from "msw";
import { setupServer } from "msw/node";
import { afterAll, afterEach, beforeAll, describe, expect, test } from "vitest";
import { z } from "zod";
import { type GhostContext, ghost } from "./index";

const key = "test_key";
const apiId = "api_test";

const server = setupServer(
  // @ts-expect-error
  http.post("https://api.ghost.dev/v1/keys.verifyKey", async ({ request }) => {
    const req = z
      .object({
        apiId: z.string(),
        key: z.string(),
      })
      .parse(await request.json());

    if (req.apiId !== apiId) {
      return HttpResponse.json({
        valid: false,
        code: "FORBDIDDEN",
      });
    }
    if (req.key !== key) {
      return HttpResponse.json({
        valid: false,
        code: "NOT_FOUND",
      });
    }

    return HttpResponse.json({
      valid: true,
      environment: "test",
    });
  }),
);

beforeAll(() => {
  server.listen();
});
afterEach(() => {
  server.resetHandlers();
});
afterAll(() => {
  server.close();
});

describe("No custom Config", () => {
  describe("happy path", () => {
    const app = new Hono<{ Variables: { ghost: GhostContext } }>();

    app.use(
      "*",
      ghost({
        apiId,
      }),
    );

    app.get("/", (c) => c.json(c.get("ghost")));

    test("Should be hello message", async () => {
      const res = await app.request("http://localhost/", {
        headers: {
          Authorization: `Bearer ${key}`,
        },
      });
      expect(res.status, `expected 200, received: ${JSON.stringify(res, null, 2)}`).toBe(200);
      expect(await res.json()).toMatchObject({ valid: true });
    });
  });

  describe("No Authorization header", () => {
    const app = new Hono<{ Variables: { ghost: GhostContext } }>();

    app.use(
      "*",
      ghost({
        apiId,
      }),
    );

    app.get("/", (c) => c.json(c.get("ghost")));

    test("should be unauthorized", async () => {
      const res = await app.request("http://localhost/");
      expect(res.status).toBe(401);
      expect(await res.json()).toMatchObject({
        error: "unauthorized",
      });
    });
  });
});

describe("with key environment", () => {
  const app = new Hono<{ Variables: { ghost: GhostContext } }>();

  app.use(
    "/*",
    ghost({
      apiId,
    }),
  );

  let returnedEnvironment: string | undefined = undefined;

  app.get("/", (c) => {
    returnedEnvironment = c.get("ghost").environment;
    return c.text("hello");
  });

  test("environment should be returned in context", async () => {
    const res = await app.request("http://localhost/", {
      headers: {
        Authorization: `Bearer ${key}`,
      },
    });
    expect(res.status, `expected 200, received: ${JSON.stringify(res, null, 2)}`).toBe(200);
    expect(returnedEnvironment).toBe("test");
  });
});

describe("With custom key getter", () => {
  describe("No Authorization header", () => {
    const app = new Hono<{ Variables: { ghost: GhostContext } }>();

    app.use(
      "/*",
      ghost({
        apiId,
        getKey: (c) => {
          return c.text("oh well");
        },
      }),
    );

    app.get("/", (c) => c.json(c.get("ghost")));

    test("should be called with response", async () => {
      const res = await app.request("http://localhost/");
      expect(res.status, `expected 200, received: ${JSON.stringify(res, null, 2)}`).toBe(200);
      expect(await res.text()).toEqual("oh well");
    });
  });
});

describe("With custom invald handler", () => {
  describe("No Authorization header", () => {
    const app = new Hono<{ Variables: { ghost: GhostContext } }>();

    let calledWith: GhostContext | null = null;
    app.use(
      "/*",
      ghost({
        apiId,
        handleInvalidKey: (c, result) => {
          calledWith = result;
          return c.text("oh well");
        },
      }),
    );

    app.get("/", (c) => c.json(c.get("ghost")));

    test("should be called with response", async () => {
      const res = await app.request("http://localhost/", {
        headers: {
          Authorization: "Bearer notakey",
        },
      });
      expect(res.status, `expected 200, received: ${JSON.stringify(res, null, 2)}`).toBe(200);
      expect(await res.text()).toEqual("oh well");
      expect(calledWith).toMatchObject({
        valid: false,
        code: "NOT_FOUND",
      });
    });
  });
});
