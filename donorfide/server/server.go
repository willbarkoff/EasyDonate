package server

import (
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/willbarkoff/donorfide/donorfide/api"
	"gorm.io/gorm"
)

func SetupRoutes(static fs.FS, db *gorm.DB) *mux.Router {
	r := mux.NewRouter()
	apiSubroute := r.PathPrefix("/api").Subrouter()
	api.SetupAPI(apiSubroute, db)

	r.Path("/").Handler(http.FileServer(http.FS(static)))

	return r
}
