package setup

import (
	"context"
	"net/http"
	"runtime"
	"strconv"
	"text/template"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/willbarkoff/donorfide/donorfide-server/database"
	"github.com/willbarkoff/donorfide/donorfide-server/errors"
	"github.com/willbarkoff/donorfide/donorfide-server/util"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

var setupCode string
var srv http.Server
var setupPage = template.Must(template.ParseFiles("./setup/setup.html"))
var setupCompletePage = template.Must(template.ParseFiles("./setup/complete.html"))

type setupPageData struct {
	OS        string
	Arch      string
	GoVers    string
	HasErrors bool
	Errors    []string
	Time      string
}

var db *gorm.DB

// Setup starts an HTTP server used for setting up Donorfide. This is done seperately so as to prevent the client from becoming too large. It's also easier to manage.
// Setup returns control to the caller once setup has been completed.
func Setup(port int, database *gorm.DB) {
	var err error
	db = database
	setupCode, err = util.GenerateRandomString(24)
	if err != nil {
		errors.Fatal(err)
	}

	errors.Logger.Info().Int("Port", port).Msg("Starting setup server")
	errors.Logger.Info().Msg("")
	errors.Logger.Info().Msg("==== USE THIS CODE FOR SETUP ====")
	errors.Logger.Info().Msg("")
	errors.Logger.Info().Msg(setupCode)
	errors.Logger.Info().Msg("")
	errors.Logger.Info().Msg("====      END SETUP CODE     ====")
	errors.Logger.Info().Msg("")

	r := httprouter.New()
	r.GET("/", index)
	r.POST("/setup", setup)

	srv.Handler = r
	srv.Addr = ":" + strconv.Itoa(port)

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		errors.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setupPage.Execute(w, setupPageData{
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		GoVers:    runtime.Version(),
		HasErrors: false,
		Time:      time.Now().Format(time.RFC1123), // use RFC1123 because it's easy to read for non-technical audiences
	})
}

func setup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setupCodeE := r.FormValue("setup-code")
	publishableKey := r.FormValue("stripe-pk")
	secretKey := r.FormValue("stripe-sk")
	orgName := r.FormValue("org-name")
	orgSite := r.FormValue("org-site")
	orgPhone := r.FormValue("org-phone")
	orgEmail := r.FormValue("org-email")
	telemetryOptOut := r.FormValue("telemetry-opt-out")
	adminFName := r.FormValue("admin-fname")
	adminLName := r.FormValue("admin-lname")
	adminEmail := r.FormValue("admin-email")
	adminPassword := r.FormValue("admin-password")
	adminPassword2 := r.FormValue("admin-password2")

	pageErrors := []string{}
	if setupCodeE == "" {
		pageErrors = append(pageErrors, "setup code is required")
	} else if setupCodeE != setupCode {
		pageErrors = append(pageErrors, "setup code is invalid")
	}

	if publishableKey == "" {
		pageErrors = append(pageErrors, "publishable key is required")
	}

	if secretKey == "" {
		pageErrors = append(pageErrors, "secret key is required")
	}

	if orgName == "" {
		pageErrors = append(pageErrors, "organazation name is required")
	}

	if orgSite == "" {
		pageErrors = append(pageErrors, "organization website is required")
	}

	if orgPhone == "" {
		pageErrors = append(pageErrors, "organization phone number is required")
	}

	if orgEmail == "" {
		pageErrors = append(pageErrors, "organization email is required")
	} else if !util.EmailIsValid(orgEmail) {
		pageErrors = append(pageErrors, "organization email is not valid")
	}

	if adminFName == "" {
		pageErrors = append(pageErrors, "admin first name is required")
	}

	if adminLName == "" {
		pageErrors = append(pageErrors, "admin last name is required")
	}

	if adminEmail == "" {
		pageErrors = append(pageErrors, "admin email is required")
	} else if !util.EmailIsValid(adminEmail) {
		pageErrors = append(pageErrors, "admin email is not valid")
	}

	if adminPassword == "" {
		pageErrors = append(pageErrors, "admin password is required")
	} else if !util.PasswordIsValid(adminPassword) {
		pageErrors = append(pageErrors, "admin password does not meet security requirements")
	} else if adminPassword != adminPassword2 {
		pageErrors = append(pageErrors, "admin passwords do not match")
	}

	errors.Logger.Debug().Str("telemetry", telemetryOptOut)

	HasErrors := len(pageErrors) > 0

	if HasErrors {
		setupPage.Execute(w, setupPageData{
			OS:        runtime.GOOS,
			Arch:      runtime.GOARCH,
			GoVers:    runtime.Version(),
			HasErrors: true,
			Errors:    pageErrors,
			Time:      time.Now().Format(time.RFC1123), // use RFC1123 because it's easy to read for non-technical audiences
		})

		return
	}

	db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&database.Prefs{Key: "stripePK", Value: publishableKey}).Error; err != nil {
			return err
		}

		if err := tx.Create(&database.Prefs{Key: "stripeSK", Value: secretKey}).Error; err != nil {
			return err
		}

		if err := tx.Create(&database.Prefs{Key: "orgName", Value: orgName}).Error; err != nil {
			return err
		}

		if err := tx.Create(&database.Prefs{Key: "orgSite", Value: orgSite}).Error; err != nil {
			return err
		}

		if err := tx.Create(&database.Prefs{Key: "orgPhone", Value: orgPhone}).Error; err != nil {
			return err
		}

		if err := tx.Create(&database.Prefs{Key: "orgEmail", Value: orgEmail}).Error; err != nil {
			return err
		}

		if err := tx.Create(&database.Prefs{Key: "telemetryOptOut", Value: telemetryOptOut}).Error; err != nil {
			return err
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		if err := tx.Create(&database.Users{
			FirstName: adminFName,
			LastName:  adminLName,
			Email:     adminEmail,
			Password:  string(hash),
			Level:     1,
		}).Error; err != nil {
			return err
		}

		return nil
	})

	setupCompletePage.Execute(w, setupPageData{
		OS:     runtime.GOOS,
		Arch:   runtime.GOARCH,
		GoVers: runtime.Version(),
		Time:   time.Now().Format(time.RFC1123), // use RFC1123 because it's easy to read for non-technical audiences
	})

	errors.Logger.Info().Msg("Setup is complete!")

	go func() {
		err := srv.Shutdown(context.Background())
		if err != nil {
			errors.FatalMsg(err, "Setup has been completed, but the setup server couldn't shut down. This is usually okay, and the process will terminate; however, upon restart of the process, Donorfide will be set up.")
		}
	}()
}
