package middleware

import (
	"context"
	"crypto/rand"
	"ecommerce/services/config"
	errors2 "ecommerce/services/utils/errors"
	"fmt"
	"net/http"
)

var httpHeaders = config.Headers()
func headers(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		csp, ctx := setCSP(ctx)

		w.Header().Set("X-Content-Type-Options", "nosniff")
		if httpHeaders.XSSProtection.Enabled {
			w.Header().Set("X-XSS-Protection", httpHeaders.XSSProtection.Value)
		}

		if httpHeaders.ReferrerPolicy.Enabled {
			w.Header().Set("Referrer-Policy", httpHeaders.ReferrerPolicy.Value)
		}
		if httpHeaders.HSTS.Enabled {
			w.Header().Set("Strict-Transport-Security", httpHeaders.HSTS.Value)
		}
		if httpHeaders.CSP.Enabled {
			w.Header().Set("Content-Security-Policy", csp)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func setCSP(ctx context.Context) (string, context.Context) {
	scriptNonce := nonce()
	styleNonce := nonce()

	ctx = context.WithValue(ctx, "scriptNonce", scriptNonce)
	ctx = context.WithValue(ctx, "styleNonce", styleNonce)

	csp := fmt.Sprintf(httpHeaders.CSP.Value, styleNonce, scriptNonce)
	return csp, ctx
}

// Generate the CSP values
func nonce() string {
	key := make([]byte, 20) // Generate a cryptographically random string with a length of 20 characters
	_, err := rand.Read(key)
	if err != nil {
		fn, line := errors2.FuncTrace()
		errors2.Panic(http.StatusInternalServerError, fn, line, err)
	}
	return fmt.Sprintf("%x", key)
}