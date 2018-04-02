package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var Log logger
var Cfg *MainConfig

func main() {
	config, err := getConfig()
	if err != nil {
		panic(err)
	}

	Log = NewLogger(config)

	router := mux.NewRouter()
	initViews(router, config)
	initAPI(router, config)

	http.Handle("/", router)
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir(config.StaticDir))))

	addr := fmt.Sprintf("%s:%d", config.Hostname, config.Port)
	fmt.Printf("Starting server on %s", addr)
	http.ListenAndServe(addr, nil)
}
