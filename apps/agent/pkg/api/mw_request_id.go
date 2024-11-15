package api

import (
	"net/http"

	"github.com/getghost/ghost/apps/agent/pkg/api/ctxutil"
	"github.com/getghost/ghost/apps/agent/pkg/uid"
)

func withRequestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = ctxutil.SetRequestId(ctx, uid.New(uid.Request()))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
