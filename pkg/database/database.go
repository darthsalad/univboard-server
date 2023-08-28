package database

import (
	"database/sql"

	"github.com/darthsalad/univboard/internal/logger"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	*sql.DB
}

func (db *Database) Init() error {
	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS users (
			id CHAR(36) NOT NULL DEFAULT (UUID()) PRIMARY KEY,
			username VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE,
			created_at DATETIME NOT NULL DEFAULT (NOW()),
			updated_at DATETIME NOT NULL DEFAULT (NOW())
		)`,
	)
	if err != nil {
		logger.Logf("err creating table: %v", err)
		return err
	}

	return nil
}

func Connect(uriString string) (*Database, error) {
	db, err := sql.Open("mysql", uriString)
	if err != nil {
			logger.Fatalf("failed to connect: %v", err)
			return nil, err
	}

	if err := db.Ping(); err != nil {
			logger.Fatalf("failed to ping: %v", err)
			return nil, err
	}

	logger.Logln("Successfully connected to PlanetScale DB!")

	return &Database{db}, nil
}