# Contributing to Ghost


## Developing

The development branch is `main`. This is the branch that all pull
requests should be made against.

To develop locally:

1. [Fork](https://help.github.com/articles/fork-a-repo/) this repository to your
   own GitHub account and then
   [clone](https://help.github.com/articles/cloning-a-repository/) it to your local device.
2. Create a new branch:

   ```sh
   git switch -c MY_BRANCH_NAME
   ```
3. Follow our getting started guide in our [documentation](https://ghost.com/docs/contributing/getting-started)

## Installing

Ghost uses [Corepack](https://nodejs.org/api/corepack.html) and [PNPM](https://pnpm.io/) for package management.

To set the correct version of PNPM, run `corepack enable` from the monorepo root. This will set your PNPM
version correctly. To install the project's dependencies, run `pnpm install`.

## Building

You can build the project with:

```bash
pnpm build
```
## Linting

To check the formatting of your code:

```sh
pnpm fmt
```
