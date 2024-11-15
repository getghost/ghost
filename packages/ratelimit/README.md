<div align="center">
    <h1 align="center">@ghost/ratelimit</h1>
    <h5>@ghost/ratelimit is a library for fast global ratelimiting in serverless functions.</h5>
</div>

<div align="center">
  <a href="https://www.ghost.com/docs/libraries/ts/ratelimit">Documentation</a>
</div>
<br/>

## Installation

```bash
npm install @ghost/ratelimit
```

## Quickstart

1.  Configure your ratelimiter:

```ts
import { Ratelimit } from "@ghost/ratelimit";

const ghost = new Ratelimit({
  rootKey: process.env.GHOST_ROOT_KEY,
  namespace: "my-app",
  limit: 10,
  duration: "30s",
  async: true,
});
```

2.  Use it:

```ts
async function handler(request) {
  const identifier = request.getUserId(); // or IP or anything else you want

  const ratelimit = await ghost.limit(identifier);
  if (!ratelimit.success) {
    return new Response("try again later", { status: 429 });
  }

  // handle the request here
}
```

## Making it Bullet Proof

To ensure reliability, you can configure timeout and error handling:

```ts
import { Ratelimit } from "@ghost/ratelimit";

const fallback = (identifier: string) => ({
  success: true,
  limit: 0,
  reset: 0,
  remaining: 0,
});

const ghost = new Ratelimit({
  // ... standard configuration
  timeout: {
    ms: 3000, // only wait 3s at most before returning the fallback
    fallback,
  },
  onError: (err, identifier) => {
    console.error(`${identifier} - ${err.message}`);
    return fallback(identifier);
  },
});
```

## API Overview

Create a new instance for ratelimiting by providing the necessary configuration.

```ts
new Ratelimit(config: RatelimitConfig)
```

Check whether a specific identifier is currently allowed to do something or if they have currently exceeded their limit.

```ts
.limit(identifier: string, opts: LimitOptions): Promise<RatelimitResponse>
```

### Documentation

[Read the full documentation](https://www.ghost.com/docs/libraries/ts/ratelimit)
