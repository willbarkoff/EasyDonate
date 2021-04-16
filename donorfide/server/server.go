package server

import (
	_ "embed"
	"github.com/willbarkoff/donorfide/donorfide/logging"
	"github.com/willbarkoff/donorfide/donorfide/server/spa"
	"github.com/willbarkoff/donorfide/donorfide/util"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/willbarkoff/donorfide/donorfide/api"
	"gorm.io/gorm"
)

//go:embed apple-developer-merchantid-domain-association
var merchantIDassoc string

func SetupRoutes(static http.FileSystem, db *gorm.DB, f util.Flags) *mux.Router {
	r := mux.NewRouter()
	spaHandler := spa.Handler{
		IndexPath:  "index.html",
		FileSystem: static,
	}

	apiSubroute := r.PathPrefix("/api").Subrouter()

	api.SetupAPI(apiSubroute, db, f)

	r.HandleFunc("/.well-known/apple-developer-merchantid-domain-association", func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte(merchantIDassoc))
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			logging.Logger.Err(err).Msg("writing merchant id assciation")
		}
	})

	r.PathPrefix("/").Handler(spaHandler)

	return r
}
