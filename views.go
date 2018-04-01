package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func initViews(router *mux.Router, config *MainConfig) {

	defaultHeaders := map[string]string{
		"Content-Type": "text/html",
	}

	getContext := func() map[string]interface{} {
		return map[string]interface{}{
			"development": config.DevEnv,
			"appName":     config.AppName,
		}
	}

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl := config.GetTemplate("index.tpl")
		err := tpl.Execute(w, getContext())
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		SetHeaders(w, defaultHeaders)
	})
}
