import { describe, expect, test } from "vitest";
import { Ghost } from "./client";

describe("client", () => {
  test("fetch can encode the params without throwing", async () => {
    const ghost = new Ghost({ token: "rawr" });
    expect(() => {
      ghost.apis.listKeys({
        apiId: "meow",
        cursor: undefined,
      });
    }).not.toThrow();
  });
});
