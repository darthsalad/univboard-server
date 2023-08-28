package server

import (
	"net/http"

	"github.com/darthsalad/univboard/pkg/server/auth"
	"github.com/darthsalad/univboard/internal/logger"
	"github.com/darthsalad/univboard/internal/utils"
	"github.com/darthsalad/univboard/pkg/database"
	"github.com/gorilla/mux"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error
type apiDBFunc func(db *database.Database, w http.ResponseWriter, r *http.Request) error

func wrapperDB(f apiDBFunc, db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(db, w, r); err != nil {
			utils.JsonResp(w, http.StatusInternalServerError, map[string]any{
				"error": err.Error(),
			})
		}
	}
}

func wrapper(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			utils.JsonResp(w, http.StatusInternalServerError, map[string]any{
				"error": err.Error(),
			})
		}
	}
}

func validate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr, err := utils.ExtractToken(r)
		if err != nil {
			logger.Logf("err extracting token: %v", err)
			utils.JsonResp(w, http.StatusUnauthorized, map[string]any{
				"error": map[string]any{
					"message":     "Invalid Bearer token",
					"status_code": http.StatusUnauthorized,
				},
			})
			return
		}

		token, err := utils.VerifyToken(tokenStr)
		if err != nil || !token.Valid {
			logger.Logf("err verifying token: %v", err)
			utils.JsonResp(w, http.StatusUnauthorized, map[string]any{
				"error": map[string]any{
					"message":     "Invalid token",
					"status_code": http.StatusUnauthorized,
				},
			})
			return
		}

		ctx := utils.SetTokenPayload(r, token)
		r = r.WithContext(ctx)

		f(w, r)
	}
}

func CreateRoutes(s *mux.Router, db *database.Database) {
	s.HandleFunc("/", wrapper(func(w http.ResponseWriter, r *http.Request) error {
		w.Write([]byte("Welcome to UnivBoard! Check out the API docs for more info."))
		return nil
	})).Methods("GET")

	s.HandleFunc("/test", wrapper(func(w http.ResponseWriter, r *http.Request) error {
		w.Write([]byte("Test"))
		return nil
	})).Methods("GET")

	s.HandleFunc("/auth/register", wrapperDB(auth.RegisterUser, db)).Methods("POST")
	s.HandleFunc("/auth/login", wrapperDB(auth.LoginUser, db)).Methods("POST")
	s.HandleFunc("/auth/logout", validate(wrapper(auth.LogoutUser))).Methods("DELETE")
	s.HandleFunc("/profile", validate(wrapperDB(auth.GetProfile, db))).Methods("GET")
	s.HandleFunc("/profile", validate(wrapperDB(auth.DeleteUser, db))).Methods("DELETE")
}
