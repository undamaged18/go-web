package server

import (
	"crypto/tls"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func serverConfig(mux *mux.Router) *http.Server {
	return &http.Server{
		Addr:              conf.Ports.HTTPS,
		Handler:           http.TimeoutHandler(caselessMatcher(csrfProtection()(mux)), time.Second*5, ""),
		TLSConfig:         tlsConfig(),
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
	}
}

func csrfProtection() func (handler http.Handler) http.Handler {
	return csrf.Protect(
		[]byte(conf.Keys.CSRF),
		csrf.CookieName("csrf"),
		csrf.FieldName("csrf"),
		csrf.HttpOnly(true),
		csrf.Secure(true),
		csrf.SameSite(csrf.SameSiteStrictMode),
		csrf.Path("/"),
	)
}

func tlsConfig() *tls.Config {
	return &tls.Config{
		PreferServerCipherSuites: true,
		MinVersion:               tls.VersionTLS13,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
	}
}