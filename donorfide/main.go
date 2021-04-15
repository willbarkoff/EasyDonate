package main

import (
	"context"
	"embed"
	"github.com/willbarkoff/donorfide/donorfide/api"
	"github.com/willbarkoff/donorfide/donorfide/util"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/willbarkoff/donorfide/donorfide/database"
	"github.com/willbarkoff/donorfide/donorfide/logging"
	"github.com/willbarkoff/donorfide/donorfide/server"
	"github.com/willbarkoff/donorfide/donorfide/setup"
)

//go:embed client/dist
var static embed.FS

func main() {
	flags := util.ParseFlags()

	logging.Logger.Info().Msg("Starting Donorfide")

	port := 8989
	if os.Getenv("DONORFIDE_PORT") != "" {
		var err error
		port, err = strconv.Atoi(os.Getenv("DONORFIDE_PORT"))
		if err != nil {
			logging.FatalMsg(err, "The enviornment variable DONORFIDE_PORT is invalid. Read more at the Donorfide documentation.")
		}
	}

	_ = godotenv.Load()
	// explicity ignore error because not everyone uses a .env file

	dsn := ""
	databaseType := ""

	if os.Getenv("DONORFIDE_DSN") != "" && os.Getenv("DATABASE_URL") != "" {
		logging.Logger.Warn().Msg("The DATABASE_URL and DONORFIDE_DSN enviorment variables are both set. Using DONORFIDE_DSN.")
		dsn = os.Getenv("DONORFIDE_DSN")
	} else if os.Getenv("DATABASE_URL") != "" {
		logging.Logger.Info().Msg("The DATABASE_URL enviorment variable is set, but the DONORFIDE_DSN enviorment varible isn't. It is assumed you are " +
			"deploying on Heroku and want to use DATABASE_URL as the DSN")
		dsn = os.Getenv("DATABASE_URL")
	} else if os.Getenv("DONORFIDE_DSN") == "" {
		logging.FatalMsg(nil, "The DONORFIDE_DSN enviorment variable is unset. Read more at the Donorfide documentation.")
	} else {
		dsn = os.Getenv("DONORFIDE_DSN")
	}

	if os.Getenv("DATABASE_URL") != "" && os.Getenv("DONORFIDE_DATABASE") == "" {
		logging.Logger.Warn().Msg("The DONORFIDE_DATABASE enviorment variable is not set, but since you're using the DATABASE_URL enviorment variable, " +
			"we're assuming you're deploying on Heroku and want to use Postgres. We reccomend setting the DONORFIDE_DATABASE enviorment variable, as this " +
			"behavior may be changed in the future.")
		databaseType = "postgres"
	} else if os.Getenv("DONORFIDE_DATABASE") == "" {
		logging.FatalMsg(nil, "The DONORFIDE_DATABASE enviorment variable is unset. Read more at the Donorfide documentation.")
	} else {
		databaseType = os.Getenv("DONORFIDE_DATABASE")
	}

	initDB(databaseType, dsn)

	err := db.AutoMigrate(database.Pref{})
	if err != nil {
		logging.FatalMsg(err, "An error occurred performing the migration.\nYou may need to fix the database yourself.")
	}
	err = db.AutoMigrate(database.User{})
	if err != nil {
		logging.FatalMsg(err, "An error occurred performing the migration.\nYou may need to fix the database yourself.")
	}
	err = db.AutoMigrate(database.Donation{})
	if err != nil {
		logging.FatalMsg(err, "An error occurred performing the migration.\nYou may need to fix the database yourself.")
	}

	count := int64(0)
	db.Model(&database.Pref{}).Count(&count)

	if count <= 0 {
		// Donorfide isn't set up.
		logging.Logger.Info().Msg("Donorfide hasn't been setup yet.")

		setup.Setup(port, db)
	}

	siteFilesFS, err := fs.Sub(static, "client/dist")
	if err != nil {
		logging.FatalMsg(err, "An error occurred while setting up the static site filesystem.")
	}

	siteFiles := http.FS(siteFilesFS)
	if flags.ClientDebug {
		logging.Logger.Info().Msg("Donorfide is running in client debug mode. For more information, visit https://donorfide.org/docs/dev/client-debug")
		siteFiles = http.Dir("client/dist")
	}

	r := server.SetupRoutes(siteFiles, db, flags)

	srv := http.Server{}
	srv.Handler = r
	srv.Addr = ":" + strconv.Itoa(port)
	logging.Logger.Info().Int("Port", port).Msg("Starting Donorfide server")

	defer func() {
		_ = recover()
		cleanup()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			logging.Logger.Info().Msg("Interrupt received. Shutting down server...")
			err := srv.Shutdown(context.Background())
			if err != nil {
				logging.FatalMsg(err, "Couldn't gracefully shutdown server.")
			}
		}
	}()

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logging.Fatal(err)
	}
}

func cleanup() {
	api.CleanupStripe()
}
