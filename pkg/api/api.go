package api

import (
	"github.com/gorilla/mux"
)

func Init(router *mux.Router) {
	router.HandleFunc("/example", Example).Methods("GET")
}
