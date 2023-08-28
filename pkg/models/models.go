package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Device struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	OS        string `json:"os"`
	OSVersion string `json:"os_version"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Push struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	DeviceID  string `json:"device_id"`
	Type      string `json:"type"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewUser(username, email, password string) *User {
	return &User{
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
}

func NewDevice(userID, name, os, osVersion string) *Device {
	return &Device{
		UserID:    userID,
		Name:      name,
		OS:        os,
		OSVersion: osVersion,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
}

func NewPush(userID, deviceID, pushType, content string) *Push {
	return &Push{
		UserID:    userID,
		DeviceID:  deviceID,
		Type:      pushType,
		Content:   content,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
}

func (u *User) HashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPass), nil
}

func (u *User) ComparePass(hashedPass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(u.Password)) == nil
}
