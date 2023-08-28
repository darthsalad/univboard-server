package database

import (
	"errors"

	"github.com/darthsalad/univboard/internal/logger"
	"github.com/darthsalad/univboard/pkg/models"
)

func (db *Database) Register(user *models.User) error {
	exists, err := db.Query("SELECT * FROM users WHERE email = ?", user.Email)
	if err != nil {
		logger.Logf("err querying: %v", err)
		return err
	}
	defer exists.Close()

	if exists.Next() {
		logger.Logf("err querying: %v", err)
		err = errors.New("err: Email already exists")
		return err
	}

	exists, err = db.Query(
		"SELECT * FROM users WHERE username = ?", user.Username,
	)
	if err != nil {
		logger.Logf("err querying: %v", err)
		return err
	}
	defer exists.Close()

	if exists.Next() {
		logger.Logf("err querying: %v", err)
		err = errors.New("err: Username already exists")
		return err
	}
	
	_, err = db.Exec(
		"INSERT INTO users (username, password, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		user.Username, user.Password, user.Email, user.CreatedAt, user.UpdatedAt,
	)

	if err != nil {
		logger.Logf("err inserting: %v", err)
		err = errors.New("err: Error inserting into database")
		return err
	}

	return nil
}

func (db *Database) Login(user *models.User) (*models.User, error) {
	exists, err := db.Query(
		"SELECT * FROM users WHERE username = ?", user.Username,
	)
	if err != nil {
		logger.Logf("err querying: %v", err)
		return nil, err
	}
	defer exists.Close()

	if !exists.Next() {
		logger.Logf("err querying: %v", err)
		err = errors.New("err: Username does not exist")
		return nil, err
	}

	var hashedPass string
	if err := exists.Scan(&user.ID, &user.Username, &hashedPass, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
		logger.Logf("err scanning: %v", err)
		return nil, err
	}

	validPass := user.ComparePass(hashedPass)
	if !validPass {
		logger.Logf("err comparing: %v", err)
		err = errors.New("err: Invalid password")
		return nil, err
	}

	return user, nil
}

func (db *Database) FetchProfile(username string) (*models.User, error) {
	exists, err := db.Query(
		"SELECT * FROM users WHERE username = ?", username,
	)
	if err != nil {
		logger.Logf("err querying: %v", err)
		return nil, err
	}
	defer exists.Close()

	if !exists.Next() {
		logger.Logf("err querying: %v", err)
		err = errors.New("err: Username does not exist")
		return nil, err
	}

	user := &models.User{}
	var hashedPass string
	if err := exists.Scan(&user.ID, &user.Username, &hashedPass, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
		logger.Logf("err scanning: %v", err)
		return nil, err
	}

	return user, nil
}

func (db *Database) UpdateProfile(user *models.User) (*models.User, error) {
	// update profile picture
	return user, nil
}

func (db *Database) DeleteUser(username string) error {
	_, err := db.Exec("DELETE FROM users WHERE username = ?", username)
	if err != nil {
		logger.Logf("err deleting: %v", err)
		return err
	}

	return nil
}