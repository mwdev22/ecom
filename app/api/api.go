package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mwdev22/ecom/app/routes/auth"
)

type Server struct {
	addr string
	db   *sql.DB
}

func NewServer(addr string, db *sql.DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}

func (s *Server) Run() error {
	router := mux.NewRouter()
	// prefix, because if api changes to new version we can change it
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	authHandler := auth.NewHandler()
	authHandler.RegisterRoutes(subrouter)

	return http.ListenAndServe(s.addr, router)
}
