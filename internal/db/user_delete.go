package db

import (
	loger "github.com/dredfort42/tools/logprinter"
)

// UserDelete deletes a user profile and all associated devices from the database
func UserDelete(email string) (err error) {
	if !UserExistsCheck(email) {
		return
	}

	query := `
		DELETE FROM ` + DB.TableUsers + ` 
		WHERE email = $1;
	`

	_, err = DB.Database.Exec(query, email)
	if err != nil {
		loger.Error("Failed to delete user from the database", err)
		return
	}

	if !UserDevicesExistsCheck(email) {
		return
	}

	query = `
		DELETE FROM ` + DB.TableDevices + ` 
		WHERE email = $1;
	`

	_, err = DB.Database.Exec(query, email)
	if err != nil {
		loger.Error("Failed to delete devices from the database", err)
	}

	return
}
