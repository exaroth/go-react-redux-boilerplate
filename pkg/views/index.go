package views

import (
	"net/http"

	"github.com/exaroth/go-react-redux-boilerplate/pkg/logger"
)

var IndexView http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	templateData := map[string]interface{}{
		"test": 1,
	}
	err := RenderTemplate(w, "index", templateData)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	logger := logger.Logger
	logger.Error()
}
