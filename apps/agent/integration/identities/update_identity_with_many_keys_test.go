package identities

import (
	"context"
	"encoding/json"
	"os"
	"sync"
	"testing"
	"time"

	ghost "github.com/getghost/ghost-go"
	"github.com/getghost/ghost-go/models/components"
	"github.com/getghost/ghost-go/models/operations"
	"github.com/getghost/ghost/apps/agent/pkg/uid"
	"github.com/stretchr/testify/require"
)

func TestUpdateAutomaticallyCreatedIdentityWithManyKeys(t *testing.T) {
	// Step 1 --------------------------------------------------------------------
	// Setup the sdk, create an API and keys
	// ---------------------------------------------------------------------------

	ctx := context.Background()
	rootKey := os.Getenv("INTEGRATION_TEST_ROOT_KEY")
	require.NotEmpty(t, rootKey, "INTEGRATION_TEST_ROOT_KEY must be set")
	baseURL := os.Getenv("GHOST_BASE_URL")
	require.NotEmpty(t, baseURL, "GHOST_BASE_URL must be set")

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

	t.Cleanup(func() {
		_, err = sdk.Apis.DeleteAPI(ctx, operations.DeleteAPIRequestBody{
			APIID: api.Object.APIID,
		})
		require.NoError(t, err)
	})

	externalID := uid.New("testuser")

	keys := make([]*operations.CreateKeyResponseBody, 1000)
	concurrency := make(chan struct{}, 32)
	wg := sync.WaitGroup{}
	for index := range keys {
		wg.Add(1)
		go func(i int) {

			concurrency <- struct{}{}

			key, createErr := sdk.Keys.CreateKey(ctx, operations.CreateKeyRequestBody{
				APIID:   api.Object.APIID,
				OwnerID: ghost.String(externalID),
			})
			require.NoError(t, createErr)

			keys[i] = key.Object

			<-concurrency
			wg.Done()
		}(index)

	}

	wg.Wait()

	// Step 2 --------------------------------------------------------------------
	// Update the identity with ratelimits
	// ---------------------------------------------------------------------------

	// Create fake ratelimits
	ratelimits := make([]operations.UpdateIdentityRatelimits, 50)
	for i := range ratelimits {
		ratelimits[i] = operations.UpdateIdentityRatelimits{
			Name:     uid.New("ratelimit"),
			Limit:    100,
			Duration: time.Second.Milliseconds(),
		}
	}
	// Add a default ratelimit cause it's somewhat of a magic value'
	ratelimits = append(ratelimits, operations.UpdateIdentityRatelimits{
		Name:     "default",
		Limit:    1000,
		Duration: time.Second.Milliseconds(),
	})

	_, err = sdk.Identities.UpdateIdentity(ctx, operations.UpdateIdentityRequestBody{
		ExternalID: ghost.String(externalID),
		Meta: map[string]any{
			"hello": "world",
		},
		Ratelimits: ratelimits,
	})
	require.NoError(t, err)

	// Step 3 --------------------------------------------------------------------
	// Verify the keys to see if they are updated
	// ---------------------------------------------------------------------------

	for _, key := range keys {
		verifyRes, err := sdk.Keys.VerifyKey(ctx, components.V1KeysVerifyKeyRequest{
			APIID: ghost.String(api.Object.APIID),
			Key:   key.Key,
		})
		require.NoError(t, err)

		require.True(t, verifyRes.V1KeysVerifyKeyResponse.Valid)
		require.NotNil(t, verifyRes.V1KeysVerifyKeyResponse.Identity)
		require.Equal(t, externalID, verifyRes.V1KeysVerifyKeyResponse.Identity.ExternalID)

		meta, err := json.Marshal(verifyRes.V1KeysVerifyKeyResponse.Identity.Meta)
		require.NoError(t, err)
		require.JSONEq(t, `{"hello":"world"}`, string(meta))
	}
}
