package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/darthsalad/univboard/internal/logger"
	"github.com/darthsalad/univboard/internal/utils"
	"github.com/darthsalad/univboard/pkg/database"
	"github.com/darthsalad/univboard/pkg/models"
)

func RegisterUser(db *database.Database, w http.ResponseWriter, r *http.Request) error {
	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logger.Fatalf("err decoding: %v", err)
		return err
	}

	hashedPass, err := user.HashPassword(user.Password)
	if err != nil {
		logger.Fatalf("err hashing: %v", err)
		return err
	}

	user = *models.NewUser(user.Username, user.Email, hashedPass)

	if err = db.Register(&user); err != nil {
		logger.Logf("err registering: %v", err)
		utils.JsonResp(w, http.StatusInternalServerError, map[string]any{
			"error": map[string]any{
				"message":     err.Error(),
				"status_code": http.StatusInternalServerError,
			},
		})
		return nil
	}

	err = utils.JsonResp(w, http.StatusCreated, map[string]any{
		"message": "Successfully created account",
		"user": map[string]string{
			"username":   user.Username,
			"email":      user.Email,
			"created_at": user.CreatedAt,
		},
	})
	if err != nil {
		logger.Logf("err responding: %v", err)
		return err
	}

	return nil
}

func LoginUser(db *database.Database, w http.ResponseWriter, r *http.Request) error {
	user := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		logger.Fatalf("err decoding: %v", err)
		return err
	}

	userNew, err := db.Login(&user)
	if err != nil {
		logger.Logf("err logging in: %v", err)
		utils.JsonResp(w, http.StatusInternalServerError, map[string]any{
			"error": map[string]any{
				"message":     err.Error(),
				"status_code": http.StatusInternalServerError,
			},
		})
		return nil
	}

	token, err := utils.CreateToken(userNew)
	if err != nil {
		logger.Fatalf("err creating token: %v", err)
		return err
	}

	http.SetCookie(w, &http.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 24 * 15),
			HttpOnly: true,
			Secure:   true,
			Path:     "/",
			SameSite: http.SameSiteNoneMode,
		},
	)

	err = utils.JsonResp(w, http.StatusOK, map[string]any{
		"message": "Successfully logged in",
		"user": map[string]string{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
		"token": token,
	})
	if err != nil {
		logger.Logf("err responding: %v", err)
		return err
	}

	return nil
}

func GetProfile(db *database.Database, w http.ResponseWriter, r *http.Request) error {
	username := r.Context().Value("username").(string)
	userNew, err := db.FetchProfile(username)
	if err != nil {
		logger.Logf("err getting profile: %v", err)
		utils.JsonResp(w, http.StatusInternalServerError, map[string]any{
			"error": map[string]any{
				"message":     err.Error(),
				"status_code": http.StatusInternalServerError,
			},
		})
		return err
	}

	err = utils.JsonResp(w, http.StatusOK, map[string]any{
		"message": "Successfully retrieved profile",
		"user": map[string]string{
			"id":       userNew.ID,
			"username": userNew.Username,
			"email":    userNew.Email,
			"created_at": userNew.CreatedAt,
			"modified_at": userNew.UpdatedAt,
		},
	})
	if err != nil {
		logger.Logf("err responding: %v", err)
		return err
	}

	return nil
}

func UpdateProfile(db *database.Database, w http.ResponseWriter, r *http.Request) error {
	// update profile picture
	
	// username := r.Context().Value("username").(string)

	// user := models.User{}
	// user.Username = username
	// if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
	// 	logger.Fatalf("err decoding: %v", err)
	// 	return err
	// }

	// userNew, err := db.UpdateProfile(&user)

	return nil
}

func LogoutUser(w http.ResponseWriter, r *http.Request) error {
	http.SetCookie(w, &http.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Now(),
			HttpOnly: true,
			Secure:   true,
			Path:     "/",
			SameSite: http.SameSiteNoneMode,
		},
	)

	err := utils.JsonResp(w, http.StatusOK, map[string]any{
		"message": "Successfully logged out",
	})
	if err != nil {
		logger.Logf("err responding: %v", err)
		return err
	}

	return nil
}

func DeleteUser (db *database.Database, w http.ResponseWriter, r *http.Request) error {
	username := r.Context().Value("username").(string)

	if err := db.DeleteUser(username); err != nil {
		utils.JsonResp(w, http.StatusInternalServerError, map[string]any{
			"error": map[string]any{
				"message":     err.Error(),
				"status_code": http.StatusInternalServerError,
			},
		})
		return nil
	}

	err := utils.JsonResp(w, http.StatusOK, map[string]any{
		"message": "Successfully deleted account",
	})
	if err != nil {
		logger.Logf("err responding: %v", err)
		return err
	}

	return nil
}