package database

import (
	"errors"

	"github.com/darthsalad/univboard/internal/logger"
	"github.com/darthsalad/univboard/pkg/models"
)

func (db *Database) AddDevice(device *models.Device) error {
	exists, err := db.Query(
		"SELECT * FROM devices WHERE user_id = ? AND name = ?",
		device.UserID, device.Name,
	)
	if err != nil {
		// logger.Logf("err querying: %v", err)
		return err
	}
	defer exists.Close()

	if exists.Next() {
		logger.Logf("err querying: %v", err)
		err = errors.New("err: Device already exists")
		return err
	}

	_, err = db.Exec(
		"INSERT INTO devices (user_id, name, os, os_version, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		device.UserID, device.Name, device.OS, device.OSVersion, device.CreatedAt, device.UpdatedAt,
	)

	if err != nil {
		logger.Logf("err inserting: %v", err)
		err = errors.New("err: Error inserting into database")
		return err
	}

	return nil
}

func (db *Database) GetDevices(userId string) ([]models.Device, error) {
	results, err := db.Query("SELECT * FROM devices WHERE user_id = ?", userId)
	if err != nil {
		logger.Logf("err querying: %v", err)
		return nil, err
	}

	devices := []models.Device{}
	if err := results.Scan(&devices); err != nil {
		logger.Logf("err scanning: %v", err)
		return nil, err
	}

	return devices, nil
}
