package api

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/exaroth/go-react-redux-boilerplate/pkg/config"
	"github.com/exaroth/go-react-redux-boilerplate/pkg/logger"
	"github.com/gorilla/mux"
)

func Init(router *mux.Router) {
	cfg := config.Config
	router.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {

		response := map[string]interface{}{
			"env": cfg.ServiceEnv,
		}
		err := SendJ(w, response)
		if err != nil {
			logger.Logger.Error(errors.Wrap(err, "Error parsing JSON response"))
			SendEJ(w, "Internal Error", 500, nil)
		}

	}).Methods("GET")
}
