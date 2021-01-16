package main

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/willbarkoff/donorfide/donorfide-server/database"
	"github.com/willbarkoff/donorfide/donorfide-server/errors"
	"github.com/willbarkoff/donorfide/donorfide-server/setup"
)

func main() {
	errors.Logger.Info().Msg("Starting Donorfide")

	port := 8989
	if os.Getenv("DONORFIDE_PORT") != "" {
		var err error
		port, err = strconv.Atoi(os.Getenv("DONORFIDE_PORT"))
		if err != nil {
			errors.FatalMsg(err, "The enviornment variable DONORFIDE_PORT is invalid. Read more at the Donorfide documentation.")
		}
	}

	_ = godotenv.Load()
	// explicity ignore error because not everyone uses a .env file

	dsn := ""
	databaseType := ""

	if os.Getenv("DONORFIDE_DSN") != "" && os.Getenv("DATABASE_URL") != "" {
		errors.Logger.Warn().Msg("The DATABASE_URL and DONORFIDE_DSN enviorment variables are both set. Using DONORFIDE_DSN.")
		dsn = os.Getenv("DONORFIDE_DSN")
	} else if os.Getenv("DATABASE_URL") != "" {
		errors.Logger.Info().Msg("The DATABASE_URL enviorment variable is set, but the DONORFIDE_DSN enviorment varible isn't. It is assumed you are " +
			"deploying on Heroku and want to use DATABASE_URL as the DSN")
		dsn = os.Getenv("DATABASE_URL")
	} else if os.Getenv("DONORFIDE_DSN") == "" {
		errors.FatalMsg(nil, "The DONORFIDE_DSN enviorment variable is unset. Read more at the Donorfide documentation.")
	} else {
		dsn = os.Getenv("DONORFIDE_DSN")
	}

	if os.Getenv("DATABASE_URL") != "" && os.Getenv("DONORFIDE_DATABASE") == "" {
		errors.Logger.Warn().Msg("The DONORFIDE_DATABASE enviorment variable is not set, but since you're using the DATABASE_URL enviorment variable, " +
			"we're assuming you're deploying on Heroku and want to use Postgres. We reccomend setting the DONORFIDE_DATABASE enviorment variable, as this " +
			"behavior may be changed in the future.")
		databaseType = "postgres"
	} else if os.Getenv("DONORFIDE_DATABASE") == "" {
		errors.FatalMsg(nil, "The DONORFIDE_DATABASE enviorment variable is unset. Read more at the Donorfide documentation.")
	} else {
		databaseType = os.Getenv("DONORFIDE_DATABASE")
	}

	initDB(databaseType, dsn)

	err := db.AutoMigrate(database.Prefs{})
	if err != nil {
		errors.FatalMsg(err, "An error occured performing the migration.\nYou may need to fix the database yourself.")
	}
	err = db.AutoMigrate(database.Users{})
	if err != nil {
		errors.FatalMsg(err, "An error occured performing the migration.\nYou may need to fix the database yourself.")
	}

	count := int64(0)
	db.Model(&database.Prefs{}).Count(&count)

	if count <= 0 {
		// Donorfide isn't set up.
		errors.Logger.Info().Msg("Donorfide hasn't been setup yet.")

		setup.Setup(port, db)
	}
}
