package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func initViews(router *mux.Router, c *MainConfig) {

	defaultHeaders := map[string]string{
		"Content-Type": "text/html",
	}

	getContext := func() map[string]interface{} {
		return map[string]interface{}{
			"developmentEnv": c.DevEnv,
			"appName":        c.AppName,
		}
	}

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := RenderTemplate(w, c.GetTemplate("index"), getContext(), defaultHeaders)
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
