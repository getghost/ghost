package routes

import (
	"github.com/getghost/ghost/apps/agent/pkg/api/validation"
	"github.com/getghost/ghost/apps/agent/pkg/logging"
	"github.com/getghost/ghost/apps/agent/pkg/metrics"
	"github.com/getghost/ghost/apps/agent/services/ratelimit"
	"github.com/getghost/ghost/apps/agent/services/vault"
)

type Services struct {
	Logger           logging.Logger
	Metrics          metrics.Metrics
	Vault            *vault.Service
	Ratelimit        ratelimit.Service
	OpenApiValidator validation.OpenAPIValidator
	Sender           Sender
}
