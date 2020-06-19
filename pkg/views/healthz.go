package views

import "net/http"

var Healthz = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
	return
}
