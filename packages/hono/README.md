
<div align="center">
    <h1 align="center">@ghost/hono</h1>
    <h5>Hono.js middleware for authenticating API keys</h5>
</div>

<div align="center">
  Inspired by <a href="https://www.openstatus.dev/blog/secure-api-with-ghost">openstatus.dev/blog/secure-api-with-ghost</a>
</div>
<br/>



Check out the docs at [ghost.dev/docs](https://ghost.com/docs/libraries/ts/hono).


Here's just an example:

```ts
import { Hono } from "hono"
import { GhostContext, ghost } from "@ghost/hono";

const app = new Hono<{ Variables: { ghost: GhostContext } }>();

app.use("*", ghost());


app.get("/somewhere", (c) => {
  // access the ghost response here to get metadata of the key etc
  const ... = c.get("ghost")

  return c.text("yo")
})
``
