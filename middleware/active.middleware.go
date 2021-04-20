package middleware

import (
	"ecommerce/services/config"
	"github.com/gorilla/mux"
)

var conf = config.New()

func Use(router *mux.Router) {
	router.Use(accessLogger)
	router.Use(errorLogger)
	router.Use(headers)
	router.Use(session)
}