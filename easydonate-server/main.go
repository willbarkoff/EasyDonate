package main

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/willbarkoff/easydonate/easydonate-server/database"
	"github.com/willbarkoff/easydonate/easydonate-server/errors"
	"github.com/willbarkoff/easydonate/easydonate-server/setup"
)

func main() {
	errors.Logger.Info().Msg("Starting EasyDonate")

	port := 8989
	if os.Getenv("EASYDONATE_PORT") != "" {
		var err error
		port, err = strconv.Atoi(os.Getenv("EASYDONATE_PORT"))
		if err != nil {
			errors.FatalMsg(err, "The enviornment variable EASYDONATE_PORT is invalid. Read more at the EasyDonate documentation.")
		}
	}

	_ = godotenv.Load()
	// explicity ignore error because not everyone uses a .env file

	if os.Getenv("EASYDONATE_DSN") == "" {
		errors.FatalMsg(nil, "The EASYDONATE_DSN enviorment variable is unset. Read more at the EasyDonate documentation.")
	}

	if os.Getenv("EASYDONATE_DATABASE") == "" {
		errors.FatalMsg(nil, "The EASYDONATE_DATABASE enviorment variable is unset. Read more at the EasyDonate documentation.")
	}

	initDB(os.Getenv("EASYDONATE_DATABASE"), os.Getenv("EASYDONATE_DSN"))

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
		// EasyDonate isn't set up.
		errors.Logger.Info().Msg("EasyDonate hasn't been setup yet.")
		setup.Setup(port, db)
	}
}
