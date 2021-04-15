package api

import (
	"github.com/gorilla/mux"
	"github.com/willbarkoff/donorfide/donorfide/database"
	"github.com/willbarkoff/donorfide/donorfide/logging"
	"net/http"
)

type orgContext struct {
	StripePK string `json:"stripe_pk"`
	Name     string `json:"name"`
	Site     string `json:"site"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Imprint  string `json:"imprint"`
}

func setupContextEndpoints(r *mux.Router) {
	openSessionsTable()

	r.HandleFunc("/org", contextOrg).Methods(GET)

	logging.Logger.Info().Msg("Loaded context endpoints.")
}

func contextOrg(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, okWithData(orgContext{
		StripePK: database.GetPref(db, "stripePK"),
		Name:     database.GetPref(db, "orgName"),
		Site:     database.GetPref(db, "orgSite"),
		Phone:    database.GetPref(db, "orgPhone"),
		Email:    database.GetPref(db, "orgEmail"),
		Imprint:  database.GetPref(db, "orgImprint"),
	}))
}
