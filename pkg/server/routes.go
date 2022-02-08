package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (s *appServer) initRoutes(r *mux.Router) {
	r.HandleFunc("/api/v1/user", s.userEndpoints.GetUsers).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/user", s.userEndpoints.CreateUser).Methods(http.MethodPost)
}
