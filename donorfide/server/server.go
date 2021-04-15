package server

import (
	"github.com/willbarkoff/donorfide/donorfide/server/spa"
	"github.com/willbarkoff/donorfide/donorfide/util"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/willbarkoff/donorfide/donorfide/api"
	"gorm.io/gorm"
)

func SetupRoutes(static http.FileSystem, db *gorm.DB, f util.Flags) *mux.Router {
	r := mux.NewRouter()
	spaHandler := spa.Handler{
		IndexPath:  "index.html",
		FileSystem: static,
	}

	apiSubroute := r.PathPrefix("/api").Subrouter()

	api.SetupAPI(apiSubroute, db, f)

	r.PathPrefix("/").Handler(spaHandler)

	return r
}
