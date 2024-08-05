package db

import (
	loger "github.com/dredfort42/tools/logprinter"
)

// DeviceDelete deletes a device from the database
func DeviceDelete(email string, deviceUUID string) (err error) {
	query := `
		DELETE FROM ` + DB.TableDevices + ` 
		WHERE email = $1 AND device_uuid = $2;
	`

	_, err = DB.Database.Exec(query, email, deviceUUID)
	if err != nil {
		loger.Error("Failed to delete device from the database", err)
	}

	return
}
