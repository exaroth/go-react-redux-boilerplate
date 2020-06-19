package views

import (
	"github.com/gorilla/mux"
)

func Init(router *mux.Router) {

	router.HandleFunc("/", IndexView)
	router.HandleFunc("/healthz", Healthz)
}
