package keys

import (
	"context"
	"os"
	"testing"
	"time"

	ghost "github.com/getghost/ghost-go"
	"github.com/getghost/ghost-go/models/operations"
	"github.com/getghost/ghost/apps/agent/pkg/uid"
	"github.com/getghost/ghost/apps/agent/pkg/util"
	"github.com/stretchr/testify/require"
)

func TestUpdateRatelimits(t *testing.T) {
	// Step 1 --------------------------------------------------------------------
	// Setup the sdk, create an API and a key
	// ---------------------------------------------------------------------------

	ctx := context.Background()
	rootKey := os.Getenv("INTEGRATION_TEST_ROOT_KEY")
	if rootKey == "" {
		t.Skip("INTEGRATION_TEST_ROOT_KEY is not set")
	}
	baseURL := os.Getenv("GHOST_BASE_URL")

	options := []ghost.SDKOption{
		ghost.WithSecurity(rootKey),
	}

	if baseURL != "" {
		options = append(options, ghost.WithServerURL(baseURL))
	}
	sdk := ghost.New(options...)

	api, err := sdk.Apis.CreateAPI(ctx, operations.CreateAPIRequestBody{
		Name: uid.New("testapi"),
	})
	require.NoError(t, err)

	key, err := sdk.Keys.CreateKey(ctx, operations.CreateKeyRequestBody{
		APIID: api.Object.APIID,
		Ratelimit: &operations.Ratelimit{
			Async:    util.Pointer(true),
			Limit:    100,
			Duration: util.Pointer(time.Minute.Milliseconds()),
		},
	})
	require.NoError(t, err)

	// Step 2 --------------------------------------------------------------------
	// Update the ratelimit
	// ---------------------------------------------------------------------------

	_, err = sdk.Keys.UpdateKey(ctx, operations.UpdateKeyRequestBody{
		KeyID: key.Object.KeyID,
		Ratelimit: &operations.UpdateKeyRatelimit{
			Async:    util.Pointer(true),
			Limit:    60,
			Duration: util.Pointer(2 * time.Minute.Milliseconds()),
		},
	})
	require.NoError(t, err)

	// Step 4 --------------------------------------------------------------------
	// Ensure the ratelimit is updated
	// ---------------------------------------------------------------------------

	retrievedKey, err := sdk.Keys.GetKey(ctx, operations.GetKeyRequest{
		KeyID: key.Object.KeyID,
	})
	require.NoError(t, err)

	require.Equal(t, true, retrievedKey.Key.Ratelimit.Async)
	require.Equal(t, int64(60), retrievedKey.Key.Ratelimit.Limit)
	require.Equal(t, int64(120000), retrievedKey.Key.Ratelimit.Duration)

}
