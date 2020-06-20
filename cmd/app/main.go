package main

import (
	"log"
	"net/http"

	"github.com/exaroth/go-react-redux-boilerplate/pkg/api"
	"github.com/exaroth/go-react-redux-boilerplate/pkg/config"
	"github.com/exaroth/go-react-redux-boilerplate/pkg/views"
	"github.com/gorilla/mux"
)

func main() {
	cfg := config.Config

	router := mux.NewRouter()

	views.Init(router)
	api.Init(router.PathPrefix("/api").Subrouter())

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir(cfg.StaticDir))),
	)

	addr := cfg.GetServiceAddr()
	log.Printf("Starting server on %s", addr)

	http.Handle("/", router)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}
