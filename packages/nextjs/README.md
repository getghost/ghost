<div align="center">
    <h1 align="center">@ghost/nextjs</h1>
    <h5>`@ghost/nextjs` the official SDK for Next.js. Just use it in your route handlers a direct and type-safe method to verify API keys.</h5>
</div>

<div align="center">
  <a href="https://www.ghost.com/docs/libraries/ts/nextjs">Documentation</a>
</div>
<br/>

## Installation

```bash
npm install @ghost/nextjs
```

## Quickstart

Protecting API routes is as simple as wrapping them with the ```withGhost``` handler:

```ts
import { NextRequestWithGhostContext, withGhost } from '@ghost/nextjs';

export const POST = withGhost(async (req) => {

  // Process the request here
  // You have access to the verification response using `req.ghost`
  console.log(req.ghost);

  return new Response('Your API key is valid!');
});
```

That’s it! Since this is just handling key verification, there’s no need to specify a root key as an environment variable.

If you want to customize how withGhost processes incoming requests, you can do so as follows:

## ```getKey```
By default, withGhost will look for a bearer token located in the ```authorization``` header. If you want to customize this, you can do so by passing a getter in the configuration object:

```ts
export const GET = withGhost(async (req) => {
  // ...
}, {
	getKey: (req) => new URL(req.url).searchParams.get("key"),
});
```

## ```onError```

You can specify custom error handling. By default errors will be logged to the console, and ```withGhost``` will return a NextResponse with status 500.

```ts
export const GET = withGhost(async (req) => {
  // ...
}, {
	onError: async (req, res) => {
      await analytics.trackEvent(`Error ${res.code}: ${res.message}`)
      return new NextResponse("Ghost error", { status: 500 })
    }
});
```

## ```handleInvalidKey```

Specify what to do if Ghost reports that your key is invalid.

```ts
export const GET = withGhost(async (req) => {
  // ...
}, {
	handleInvalidKey: (req, res) => {
      return new Response("Unauthorized", { status: 401 })
    }
});
```


## Disable telemetry

By default, Ghost collects anonymous telemetry data to help us understand how our SDKs are used.

If you wish to disable this, you can do so by passing a boolean flag to the constructor:

```ts
export const GET = withGhost(async (req) => {
  // ...
}, { disableTelemetry: true });
```

### Documentation

[Read the full documentation](https://www.ghost.com/docs/libraries/ts/nextjs)