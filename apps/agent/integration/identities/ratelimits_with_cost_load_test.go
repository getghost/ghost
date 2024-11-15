package identities_test

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	ghost "github.com/getghost/ghost-go"
	"github.com/getghost/ghost-go/models/components"
	"github.com/getghost/ghost-go/models/operations"
	attack "github.com/getghost/ghost/apps/agent/pkg/testutil"
	"github.com/getghost/ghost/apps/agent/pkg/uid"
	"github.com/getghost/ghost/apps/agent/pkg/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestIdentityRatelimitsWithCost0Accuracy(t *testing.T) {
	// Step 1 --------------------------------------------------------------------
	// Setup the sdk, create an API
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

	for _, tc := range []struct {
		rate         attack.Rate
		testDuration time.Duration
	}{
		{
			rate:         attack.Rate{Freq: 100, Per: time.Second},
			testDuration: 1 * time.Minute,
		},
		{
			rate:         attack.Rate{Freq: 100, Per: time.Second},
			testDuration: 5 * time.Minute,
		},
		{
			rate:         attack.Rate{Freq: 100, Per: time.Second},
			testDuration: 30 * time.Minute,
		},
	} {
		t.Run(fmt.Sprintf("[%s] over %s", tc.rate.String(), tc.testDuration), func(t *testing.T) {
			api, err := sdk.Apis.CreateAPI(ctx, operations.CreateAPIRequestBody{
				Name: uid.New("testapi"),
			})
			require.NoError(t, err)

			// Step 2 --------------------------------------------------------------------
			// Create the identity with ratelimits
			// ---------------------------------------------------------------------------

			ratelimit := operations.Ratelimits{
				Name:     "ratelimit-a",
				Limit:    600,
				Duration: time.Minute.Milliseconds(),
			}

			externalID := uuid.NewString()
			_, err = sdk.Identities.CreateIdentity(ctx, operations.CreateIdentityRequestBody{
				ExternalID: externalID,
				Ratelimits: []operations.Ratelimits{
					ratelimit,
				},
			})
			require.NoError(t, err)

			// Step 3 --------------------------------------------------------------------
			// Create key for this identity
			// ---------------------------------------------------------------------------

			key, err := sdk.Keys.CreateKey(ctx, operations.CreateKeyRequestBody{
				APIID:      api.Object.APIID,
				ExternalID: util.Pointer(externalID),
			})
			require.NoError(t, err)

			// Step 5 --------------------------------------------------------------------
			// Test ratelimits
			// ---------------------------------------------------------------------------

			total := 0
			passed := 0
			withCost := 0
			errors := 0

			results := attack.Attack(t, tc.rate, tc.testDuration, func() bool {

				cost := int64(0)
				if rand.Intn(100) == 0 {
					withCost++
					cost = 1
				}

				res, err := sdk.Keys.VerifyKey(context.Background(), components.V1KeysVerifyKeyRequest{
					APIID: ghost.String(api.Object.APIID),
					Key:   key.Object.Key,
					Ratelimits: []components.Ratelimits{
						{Name: ratelimit.Name,
							Cost: util.Pointer(cost),
						},
					},
				})
				if err != nil {
					errors++
					return false
				}

				return res.V1KeysVerifyKeyResponse.Valid

			})

			for valid := range results {
				total++
				if valid {
					passed++
				}

			}

			// Step 6 --------------------------------------------------------------------
			// Assert ratelimits worked
			// ---------------------------------------------------------------------------

			t.Logf("Total: %d, Passed: %d, withCost=1: %d", total, passed, withCost)

			// check requests::api is not exceeded
		})

	}
}
