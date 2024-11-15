FROM golang:1.22-alpine AS builder



WORKDIR /go/src/github.com/getghost/ghost/apps/agent
COPY go.sum go.mod ./
RUN go mod download

COPY . .
ARG VERSION
RUN go build -o bin/ghost -ldflags "-X 'github.com/getghost/ghost/apps/agent/pkg/version.Version=${VERSION}'"  ./cmd/main.go

FROM golang:1.22-alpine
WORKDIR  /usr/local/bin
COPY --from=builder /go/src/github.com/getghost/ghost/apps/agent/bin/ghost .
COPY --from=builder /go/src/github.com/getghost/ghost/apps/agent/config.production.json .
COPY --from=builder /go/src/github.com/getghost/ghost/apps/agent/config.staging.json .
COPY --from=builder /go/src/github.com/getghost/ghost/apps/agent/config.docker.json .
COPY --from=builder /go/src/github.com/getghost/ghost/apps/agent/pkg/openapi/openapi.json ./pkg/openapi/openapi.json

CMD [ "/usr/local/bin/ghost", "agent"]
