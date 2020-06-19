package main

import (
	"fmt"
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
		http.StripPrefix("/static/", http.FileServer(http.Dir(cfg.StaticDir))))

	addr := cfg.GetServiceAddr()
	fmt.Printf("Starting server on %s", addr)
	http.ListenAndServe(addr, nil)
}
