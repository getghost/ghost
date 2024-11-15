package v1Liveness_test

import (
	"testing"

	v1Liveness "github.com/getghost/ghost/apps/agent/pkg/api/routes/v1_liveness"
	"github.com/getghost/ghost/apps/agent/pkg/api/testutil"
	"github.com/getghost/ghost/apps/agent/pkg/openapi"
	"github.com/stretchr/testify/require"
)

func TestLiveness(t *testing.T) {

	h := testutil.NewHarness(t)
	route := h.SetupRoute(v1Liveness.New)
	res := testutil.CallRoute[any, openapi.V1LivenessResponseBody](t, route, nil, nil)

	require.Equal(t, 200, res.Status)
}
