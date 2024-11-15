package clickhouse

import (
	"github.com/getghost/ghost/apps/agent/pkg/clickhouse/schema"
)

type Bufferer interface {
	BufferApiRequest(schema.ApiRequestV1)
	BufferKeyVerification(schema.KeyVerificationRequestV1)
}
