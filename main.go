package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var Log logger
var Config *MainConfig

func init() {
	var err error
	Config, err = getConfig()
	if err != nil {
		panic(err)
	}
	Log = NewLogger(Config)
}

func main() {

	router := mux.NewRouter()
	initViews(router)
	initAPI(router)

	http.Handle("/", router)
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir(Config.GetStaticDir()))))

	addr := Config.GetAddress()
	fmt.Printf("Starting server on %s", addr)
	http.ListenAndServe(addr, nil)
}
