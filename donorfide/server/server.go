package server

import (
	"github.com/willbarkoff/donorfide/donorfide/util"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/willbarkoff/donorfide/donorfide/api"
	"gorm.io/gorm"
)

func SetupRoutes(static http.FileSystem, db *gorm.DB, f util.Flags) *mux.Router {
	r := mux.NewRouter()

	apiSubroute := r.PathPrefix("/api").Subrouter()

	api.SetupAPI(apiSubroute, db, f)

	r.PathPrefix("/").Handler(http.FileServer(static))

	return r
}
