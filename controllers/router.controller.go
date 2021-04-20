package controllers

import (
	"ecommerce/middleware"
	"ecommerce/services/config"
	"github.com/gorilla/mux"
	"net/http"
)

func Router(router *mux.Router) {

	if conf.Environment != "production" {
		router.HandleFunc("/test-get", TestGet).Methods(http.MethodGet)
		router.HandleFunc("/test-post", TestPost).Methods(http.MethodPost)
	}

	router.HandleFunc(config.Routes("index"), index).Methods(http.MethodGet)

	router.HandleFunc(config.Routes("account:login"), login).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc(config.Routes("account:register"), register).Methods(http.MethodGet, http.MethodPost)

	router.PathPrefix("/assets/").Handler(middleware.Assets(http.StripPrefix("/assets/", http.FileServer(http.Dir(conf.Paths.Root+"/assets"))))).Methods(http.MethodGet)
	router.PathPrefix("/images/").Handler(middleware.Assets(http.StripPrefix("/images/", http.FileServer(http.Dir(conf.Paths.Root+"/images"))))).Methods(http.MethodGet)

	router.PathPrefix("/favicon.ico").Handler(http.FileServer(http.Dir(conf.Paths.Root))).Methods(http.MethodGet)
	router.PathPrefix(config.Assets("logo")).Handler(http.FileServer(http.Dir(conf.Paths.Root))).Methods(http.MethodGet)
	router.PathPrefix("/robots.txt").Handler(http.FileServer(http.Dir(conf.Paths.Root))).Methods(http.MethodGet)
}

func TestGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET successful"))
}
func TestPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("POST successful"))
}
