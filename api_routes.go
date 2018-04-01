package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func initAPI(router *mux.Router, config *MainConfig) {
	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		err := SendJSONResponse(w, map[string]string{
			"version": config.Version,
		}, defaultJSONHeaders)
		if err != nil {
			SendJSONError(w, err)
		}
	}).Methods("GET")

}
