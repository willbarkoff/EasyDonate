package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// InitilizeApi initilizes the donorfide API.

var db *gorm.DB

var statusOK = map[string]interface{}{"status": "ok"}
var statusInternalServerError = map[string]interface{}{"status": "error", "error": "internal_server_error"}
var statusBadRequest = map[string]interface{}{"status": "error", "error": "bad_request"}
var statusUnauthorized = map[string]interface{}{"status": "error", "error": "unauthorized"}

func writeJSON(w http.ResponseWriter, status int, thing interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(thing)
	if err != nil {
		panic(err)
	}
}

func SetupAPI(r *mux.Router, database *gorm.DB) {
	db = database

	r.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, statusOK)
	})

	setupAuthEndpoints(r.PathPrefix("/auth").Subrouter())
}
