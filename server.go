package main

import (
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
	return
}
func initRoutes(config *MainConfig) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl := config.GetTemplate("index.tpl")
		err := tpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

}
