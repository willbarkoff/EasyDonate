package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/willbarkoff/donorfide/donorfide/database"
	"github.com/willbarkoff/donorfide/donorfide/logging"
	"gorm.io/gorm"
)

func setupAdminEndpoints(r *mux.Router) {
	r.HandleFunc("/updateSetting", updateSetting).Methods(POST)
	logging.Logger.Info().Msg("Loaded administration endpoints.")
}

func updateSetting(w http.ResponseWriter, r *http.Request) {
	if !paramsOk(w, r, "key") {
		return
	}

	id, err := getUserId(r, w)

	if err != nil {
		writeJSON(w, http.StatusInternalServerError, statusInternalServerError)
		logging.Logger.Err(err).Msg("Getting user ID")
		return
	}

	if id <= 0 {
		writeJSON(w, http.StatusUnauthorized, statusLoggedOut)
		return
	}

	user := database.GetUserInfo(db, id)
	if user.Level < 2 {
		writeJSON(w, http.StatusUnauthorized, statusUnauthorized)
		return
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		tx.Delete(&database.Pref{}, "`key` = ?", r.FormValue("key"))
		if r.FormValue("value") != "" {
			tx.Create(&database.Pref{
				Key:   r.FormValue("key"),
				Value: r.FormValue("value"),
			})
		}
		return nil
	})

	if err != nil {
		writeJSON(w, http.StatusInternalServerError, statusInternalServerError)
		logging.Logger.Err(err).Msg("Updating setting")
		return
	}

	writeJSON(w, http.StatusOK, statusOK)
}
