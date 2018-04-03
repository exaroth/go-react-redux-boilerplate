package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func initViews(router *mux.Router) {

	defaultHeaders := map[string]string{
		"Content-Type": "text/html",
	}

	getContext := func() map[string]interface{} {
		return map[string]interface{}{
			"developmentEnv": Config.DevEnv,
			"appName":        Config.AppName,
		}
	}

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := RenderTemplate(w, "index", getContext(), defaultHeaders)
		if err != nil {
			Log.Error("", err, r)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	})

	router.HandleFunc("/healthz",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
			return
		})
}
