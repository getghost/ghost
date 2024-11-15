FROM  node:lts AS compile

WORKDIR /app
RUN npm install -g workerd@latest pnpm

COPY pnpm-lock.yaml ./
RUN pnpm fetch


COPY . .
RUN rm .dev.vars || true


RUN pnpm install --offline -r

WORKDIR /app/apps/api
RUN pnpm wrangler deploy --outdir=dist --dry-run

RUN pnpm workerd compile ./worker.capnp > ghost


FROM ubuntu:latest
COPY --from=compile /app/apps/api/ghost /usr/bin/ghost

EXPOSE 8787
CMD ["ghost"]


# FROM --platform=linux/amd64 node:lts AS compile

# WORKDIR /app
# RUN npm install -g workerd pnpm

# COPY pnpm-lock.yaml ./
# RUN pnpm fetch


# COPY . .
# RUN rm .dev.vars || true



# RUN pnpm install --offline -r

# WORKDIR /app/apps/api
# RUN pnpm wrangler deploy --dry-run --outdir .wrangler/dist

# FROM jacoblincool/workerd

# COPY --from=compile /app/apps/api/worker.capnp /app/worker.capnp
# COPY --from=compile /app/apps/api/.wrangler/dist /app/.wrangler/dist

# WORKDIR /app

# EXPOSE 8787
