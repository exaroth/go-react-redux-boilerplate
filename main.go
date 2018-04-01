package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	config, err := getConfig()
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter()
	initViews(router, config)
	initAPI(router, config)

	http.HandleFunc("/healthz",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
			return
		})
	http.Handle("/", router)
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir(config.StaticDir))))

	addr := fmt.Sprintf("%s:%d", config.Hostname, config.Port)
	fmt.Printf("Starting server on %s", addr)

	http.ListenAndServe(addr, nil)
}
