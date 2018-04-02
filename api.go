package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func initAPI(router *mux.Router) {

	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"development": Config.DevEnv,
			"appName":     Config.AppName,
		}
		err := SendJSONResponse(w, response, defaultJSONHeaders)
		if err != nil {
			SendJSONError(w, err)
		}
	}).Methods("GET")

}
