package db

import (
	"database/sql"

	cfg "github.com/dredfort42/tools/configreader"
	_ "github.com/lib/pq"
)

// Database is the database struct
type Database struct {
	database     *sql.DB
	tableUsers   string
	tableDevices string
}

var db Database

// DatabaseInit initializes the database
func DatabaseInit() {
	db.tableUsers = cfg.Config["db.table.profile.users"]
	if db.tableUsers == "" {
		panic("db.table.profile.users is empty")
	}

	db.tableDevices = cfg.Config["db.table.profile.devices"]
	if db.tableDevices == "" {
		panic("db.table.profile.devices is empty")
	}

	databaseConnect()
	tablesCheck()
}
