package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mwdev22/ecom/app/routes/auth"
	"gorm.io/gorm"
)

type Server struct {
	addr string
	db   *gorm.DB
}

func NewServer(addr string, db *gorm.DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}

func (s *Server) Run() error {
	router := mux.NewRouter()
	// prefix, because if api changes to new version we can change it
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	authStore := auth.NewStore(s.db)
	authHandler := auth.NewHandler(authStore)
	authHandler.RegisterRoutes(subrouter)

	return http.ListenAndServe(s.addr, router)
}
