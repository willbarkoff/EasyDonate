package main

import (
	"github.com/willbarkoff/donorfide/donorfide/logging"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB(databaseType, DSN string) {
	var err error
	gormConfig := &gorm.Config{}
	gormConfig.Logger = logging.GormLogger{}

	const databaseError = "An error occurred connecting to the database"

	if databaseType == "sqlite" {
		db, err = gorm.Open(sqlite.Open(DSN), gormConfig)
		logging.Logger.Warn().Str("Database Type", "SQLite").Msg("SQLite should not be used in production environments.")
		if err != nil {
			logging.FatalMsg(err, databaseError)
		}
	} else if databaseType == "mysql" {
		db, err = gorm.Open(mysql.Open(DSN), gormConfig)
		if err != nil {
			logging.FatalMsg(err, databaseError)
		}
	} else if databaseType == "postgres" {
		db, err = gorm.Open(postgres.Open(DSN), gormConfig)
		if err != nil {
			logging.FatalMsg(err, databaseError)
		}
	} else if databaseType == "mssql" {
		db, err = gorm.Open(sqlserver.Open(DSN), gormConfig)
		if err != nil {
			logging.FatalMsg(err, databaseError)
		}
	} else {
		logging.FatalMsg(nil, "The database type selected isn't supported.")
	}

	sqlDB, err := db.DB()
	if err != nil {
		logging.FatalMsg(err, databaseError)
	}

	err = sqlDB.Ping()
	if err != nil {
		logging.FatalMsg(err, databaseError)
	}
}
