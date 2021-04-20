package server

import (
	"ecommerce/controllers"
	"ecommerce/middleware"
	"ecommerce/services/config"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

var conf = config.New()

func Run() {
	defer func() {
		rec := recover()
		if rec != nil {
			fmt.Println("Panic recovered in server  Run")
			fmt.Println(rec)
		}
	}()

	router := mux.NewRouter()
	srv := serverConfig(router)

	middleware.Use(router)
	controllers.Router(router)

	// Check what mode the server is running in
	switch conf.Environment {
	case "production":
		getCertificateManager()
	case "testing":
		// Else set the certManager to nil
		conf.CertManager = nil
	case "development":
		// Else set the certManager to nil
		conf.CertManager = nil
	default:
		getCertificateManager()
	}

	// Make a channel for error handling in the Goroutines
	c := make(chan error)

	// Start a Goroutine for both HTTP and HTTPS listeners
	go httpListener(router, c)
	go httpsListener(srv, c)

	// If the go routine returns an error fetch the error from the channel and assign it to err and panic
	// The panic will be picked up in the server recovery section.
	err := <-c
	if err != nil {
		panic(err)
	}
}


// caselessMatcher ensures that the URL resolves regardless of the case entered by the end user
// it converts the request URL Path to lowercase
func caselessMatcher(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.ToLower(r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

