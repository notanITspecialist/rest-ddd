package server

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"rest-ddd/pkg/config"
	"rest-ddd/pkg/endpoints"
)

type appServer struct {
	log      *zap.Logger
	listener net.Listener
	server   *http.Server
	cfg      config.Server

	userEndpoints endpoints.UserEndpoints
}

func NewAppServer(
	log *zap.Logger,
	cfg config.Server,
	userEndpoints endpoints.UserEndpoints,
) (Server, error) {
	port := cfg.Port
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return nil, fmt.Errorf("cannot listen app port: %w", port)
	}
	server := newAppServer(
		log,
		cfg,
		listener,
		userEndpoints,
	)
	return server, nil
}

func newAppServer(
	log *zap.Logger,
	cfg config.Server,
	listener net.Listener,
	userEndpoints endpoints.UserEndpoints,
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

func (s *appServer) Start() {
	s.log.Info(fmt.Sprintf("Start app server :%v", s.cfg.Port))

	err := s.server.Serve(s.listener)
	if err != nil {
		s.log.Panic("Error while serve app server", zap.Error(err))
	}
}
