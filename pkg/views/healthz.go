package views

import "net/http"

// Healthz will always return response OK with code 200,
// this view is usually used for checking health of the app.
var Healthz = func(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("OK"))
	if err != nil {
		w.WriteHeader(500)
	}
	return
}
