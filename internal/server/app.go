package server

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"rest-ddd/internal/config"
	"rest-ddd/internal/endpoints"
)

var Paths = map[string]string{
	"GetUsers":   "/api/v1/user",
	"CreateUser": "/api/v1/user",
}

type (
	appServer struct {
		log      *zap.Logger
		listener net.Listener
		server   *http.Server
		cfg      config.Server

		userEndpoints endpoints.UserEndpoints
	}

	Server interface {
		Start()
	}
)

func NewAppServer(
	log *zap.Logger,
	cfg config.Server,
	userEndpoints endpoints.UserEndpoints,
) (Server, error) {
	port := cfg.Port
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return nil, fmt.Errorf("cannot listen app port: %v", port)
	}
	server := newAppServer(
		log,
		cfg,
		userEndpoints,
		listener,
	)
	return server, nil
}

func newAppServer(
	log *zap.Logger,
	cfg config.Server,
	userEndpoints endpoints.UserEndpoints,
	listener net.Listener,
) *appServer {
	router := mux.NewRouter()

	server := &appServer{
		log:      log,
		cfg:      cfg,
		listener: listener,
		server:   &http.Server{Handler: router},

		userEndpoints: userEndpoints,
	}
	server.initRoutes(router)

	return server
}

func (s *appServer) initRoutes(r *mux.Router) {
	r.HandleFunc(Paths["GetUsers"], s.userEndpoints.GetUsers).Methods(http.MethodGet)
	r.HandleFunc(Paths["CreateUser"], s.userEndpoints.CreateUser).Methods(http.MethodPost)
}

func (s *appServer) Start() {
	s.log.Info(fmt.Sprintf("Start app server :%v", s.cfg.Port))

	err := s.server.Serve(s.listener)
	if err != nil {
		s.log.Panic("Error while serve app server", zap.Error(err))
	}
}
