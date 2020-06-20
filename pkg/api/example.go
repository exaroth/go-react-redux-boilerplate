package api

import (
	"encoding/json"
	"net/http"

	"github.com/exaroth/go-react-redux-boilerplate/pkg/config"
	"github.com/exaroth/go-react-redux-boilerplate/pkg/logger"
	"github.com/pkg/errors"
)

var Example http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {

	cfg := config.Config
	response := map[string]interface{}{
		"env": cfg.ServiceEnv,
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		logger.Logger.Error(errors.Wrap(err, "Error parsing JSON response"))
		w.WriteHeader(500)
	}

}
