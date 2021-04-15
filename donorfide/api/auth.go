package api

import (
	"errors"
	"github.com/willbarkoff/donorfide/donorfide/database"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/wader/gormstore/v2"
	"github.com/willbarkoff/donorfide/donorfide/logging"
)

var store *gormstore.Store
var quitStore chan struct{}

var malformedSession = errors.New("auth: malformed session")

func setupAuthEndpoints(r *mux.Router) {
	openSessionsTable()

	r.HandleFunc("/me", me).Methods(GET)
	r.HandleFunc("/login", login).Methods(POST)
	r.HandleFunc("/logout", logout).Methods(POST)

	logging.Logger.Info().Msg("Loaded authentication endpoints.")
}

func openSessionsTable() {
	store = gormstore.New(db, []byte("secret"))
	// db cleanup every hour
	// close quit channel to stop cleanup
	quitStore = make(chan struct{})
	go store.PeriodicCleanup(1*time.Hour, quitStore)
}

func me(w http.ResponseWriter, r *http.Request) {
	id, err := getUserId(r, w)

	if err != nil {
		writeJSON(w, http.StatusInternalServerError, statusInternalServerError)
		logging.Logger.Err(err).Msg("Getting user ID in auth/me")
	}

	if id == -1 {
		writeJSON(w, http.StatusUnauthorized, statusUnauthorized)
		return
	}

	user := database.User{}
	db.First(&user, id)

	writeJSON(w, http.StatusOK, user)
}

func login(w http.ResponseWriter, r *http.Request) {
	ok := paramsOk(w, r, "email", "password")
	if !ok {
		return
	}

	user := database.User{}
	db.First(&user, "email = ?", r.FormValue("email"))

	if (user == database.User{}) {
		writeJSON(w, http.StatusUnauthorized, statusInvalidLogin)
		return
	}

	hash := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(hash, []byte(r.FormValue("password")))
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, statusInvalidLogin)
		return
	}

	updateUser(1, w, r)
}

func logout(w http.ResponseWriter, r *http.Request) {
	updateUser(-1, w, r)
}

func updateUser(newId int, w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "id")
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, statusInternalServerError)
		logging.Logger.Err(err).Msg("Getting session in updateUser")
		return
	}

	session.Values["id"] = newId

	err = session.Save(r, w)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, statusInternalServerError)
		logging.Logger.Err(err).Msg("Writing session in updateUser")
		return
	}

	writeJSON(w, http.StatusOK, statusOK)
}

func getUserId(r *http.Request, w http.ResponseWriter) (int, error) {
	session, err := store.Get(r, "id")

	if err != nil {
		return -1, err
	}

	id, ok := session.Values["id"].(int)

	if !ok || id < 1 {
		session.Values["id"] = -1
		err := session.Save(r, w)
		if err != nil {
			return -1, err
		}
		return -1, nil
	}

	return id, nil
}
