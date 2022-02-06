package server

import (
	"fmt"
	"net"
	"net/http"

	"go.uber.org/zap"

	"github.com/gorilla/mux"
)

type appServer struct {
	log      *zap.Logger
	listener net.Listener
	server   *http.Server
}

func NewAppServer(
	log *zap.Logger,
) (Server, error) {
	port := "8000"
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return nil, fmt.Errorf("cannot listen app port: %w", port)
	}
	server := newAppServer(log, listener)
	return server, nil
}

func newAppServer(
	log *zap.Logger,
	listener net.Listener,
) *appServer {
	router := mux.NewRouter()

	server := &appServer{
		log:      log,
		listener: listener,
		server: &http.Server{
			Handler: router,
		},
	}

	return server
}

func (s *appServer) Start() {
	s.log.Info("Start app server :8000")

	err := s.server.Serve(s.listener)
	if err != nil {
		s.log.Panic("Error while serve app server", zap.Error(err))
	}
}
