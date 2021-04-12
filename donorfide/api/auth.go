package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func setupAuthEndpoints(r *mux.Router) {
	r.HandleFunc("/whoami", whoami).Methods("GET")
}

func whoami(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusUnauthorized, statusUnauthorized)
}
