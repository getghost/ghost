package uid_test

import (
	"testing"

	"github.com/getghost/ghost/apps/agent/pkg/uid"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	ids := map[string]bool{}
	for range 1000 {
		id := uid.New("")
		require.Positive(t, len(id))
		_, ok := ids[id]
		require.False(t, ok, "generated id must be unique")
		ids[id] = true
	}
}

func TestNewWithPrefix(t *testing.T) {
	prefixes := []uid.Prefix{
		uid.NodePrefix,
	}

	ids := map[string]bool{}
	for _, prefix := range prefixes {
		for range 1000 {
			id := uid.New(string(prefix))
			require.Positive(t, len(id))
			_, ok := ids[id]
			require.False(t, ok, "generated id must be unique")
			ids[id] = true
		}
	}
}