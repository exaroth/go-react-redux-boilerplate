package api

import (
	"encoding/json"
	"net/http"
)

// SendJ will write data in a JSON format into the response.
func SendJ(w http.ResponseWriter, data interface{}) error {
	return json.NewEncoder(w).Encode(data)
}

func SendEJ(w http.ResponseWriter, m string, s int, meta interface{}) error {
	w.WriteHeader(s)
	return SendJ(w, map[string]interface{}{"error": m, "meta": meta})
}
