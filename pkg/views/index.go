package views

import (
	"net/http"
)

var IndexView http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	templateData := map[string]interface{}{
		"test": 1,
	}
	err := RenderTemplate(w, "index", templateData)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
