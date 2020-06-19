package api

import (
	"fmt"
	"net/http"

	"github.com/exaroth/go-react-redux-boilerplate/pkg/config"
	"github.com/gorilla/mux"
)

func Init(router *mux.Router) {
	cfg := config.Config
	router.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"env": cfg.ServiceEnv,
		}
		fmt.Println(response)

	}).Methods("GET")
}
