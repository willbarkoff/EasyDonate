package api

import (
	_ "embed"
	"encoding/json"
	"github.com/willbarkoff/donorfide/donorfide/util"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/willbarkoff/donorfide/donorfide/logging"
	"gorm.io/gorm"
)

//go:embed api-tester.html
var apiTester string

var db *gorm.DB

var flags util.Flags

var statusOK = map[string]interface{}{"status": "ok"}
var statusInternalServerError = map[string]interface{}{"status": "error", "error": "internal_server_error"}
var statusMissingParams = map[string]interface{}{"status": "error", "error": "bad_request"}
var statusInvalidParams = map[string]interface{}{"status": "error", "error": "missing_params"}
var statusUnauthorized = map[string]interface{}{"status": "error", "error": "unauthorized"}
var statusLoggedOut = map[string]interface{}{"status": "error", "error": "logged_out"}
var statusInvalidLogin = map[string]interface{}{"status": "error", "error": "invalid_login"}
var statusUserMissing = map[string]interface{}{"status": "error", "error": "user_missing"}
var statusNotFound = map[string]interface{}{"status": "error", "error": "not_found"}

func okWithData(data interface{}) map[string]interface{} {
	return map[string]interface{}{"status": "ok", "data": data}
}

const (
	GET  = "GET"
	POST = "POST"
)

func writeJSON(w http.ResponseWriter, status int, thing interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(thing)
	if err != nil {
		logging.Logger.Err(err).Msg("Encoding JSON")
	}
}

func SetupAPI(r *mux.Router, database *gorm.DB, f util.Flags) {
	db = database
	flags = f

	if f.APITester {
		logging.Logger.Info().Msg("The API tester is enabled. For more information, visit https://donorfide.org/docs/dev/api-tester")
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "text/html")
			_, _ = w.Write([]byte(apiTester))
		})
	}

	r.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, statusOK)
	})

	setupAuthEndpoints(r.PathPrefix("/auth").Subrouter())
	setupDonationEndpoints(r.PathPrefix("/donate").Subrouter())
	setupContextEndpoints(r.PathPrefix("/context").Subrouter())

	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusNotFound, statusNotFound)
	})
}

func paramsOk(w http.ResponseWriter, r *http.Request, params ...string) bool {
	for _, v := range params {
		if r.FormValue(v) == "" {
			writeJSON(w, http.StatusBadRequest, statusMissingParams)
			return false
		}
	}
	return true
}
