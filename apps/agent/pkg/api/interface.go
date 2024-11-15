package api

import "github.com/getghost/ghost/apps/agent/pkg/clickhouse/schema"

type EventBuffer interface {
	BufferApiRequest(schema.ApiRequestV1)
}
