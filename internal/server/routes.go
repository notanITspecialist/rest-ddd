package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

var Paths = map[string]string{
	"GetUsers":   "/api/v1/user",
	"CreateUser": "/api/v1/user",
}

func (s *appServer) initRoutes(r *mux.Router) {
	r.HandleFunc(Paths["GetUsers"], s.userEndpoints.GetUsers).Methods(http.MethodGet)
	r.HandleFunc(Paths["CreateUser"], s.userEndpoints.CreateUser).Methods(http.MethodPost)
}
