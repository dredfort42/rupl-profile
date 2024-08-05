package db

import (
	"database/sql"

	cfg "github.com/dredfort42/tools/configreader"
	_ "github.com/lib/pq"
)

// Database is the database struct
type Database struct {
	Database     *sql.DB
	TableUsers   string
	TableDevices string
}

var DB Database

// DatabaseInit initializes the database
func DatabaseInit() {
	DB.TableUsers = cfg.Config["db.table.profile.users"]
	if DB.TableUsers == "" {
		panic("db.table.profile.users is empty")
	}

	DB.TableDevices = cfg.Config["db.table.profile.devices"]
	if DB.TableDevices == "" {
		panic("db.table.profile.devices is empty")
	}

	databaseConnect()
	tablesCheck()
}
