package server

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/darthsalad/univboard/pkg/database"
)

type Server struct {
	Router *mux.Router
	DB     *database.Database
	server *http.Server
}

func CreateServer(db *database.Database) *Server {
	server := &Server{
		Router: mux.NewRouter(),
		DB:     db,
		server: &http.Server{},
	}
	CreateRoutes(server.Router, db)
	server.Router.Use(middleware)
	return server
}

func (s *Server) Start(address string) error {
	s.server.Addr = address
	s.server.Handler = s.Router
	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	s.DB.Close()
	return s.server.Shutdown(ctx)
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") 
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		next.ServeHTTP(w, r)
	})
}
	