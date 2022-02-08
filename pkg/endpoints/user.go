package endpoints

import (
	"go.uber.org/zap"
	"net/http"
)

type (
	UserEndpoints interface {
		GetUsers(http.ResponseWriter, *http.Request)
		CreateUser(http.ResponseWriter, *http.Request)
	}

	userHandler struct {
		log *zap.Logger
	}
)

func NewUserEndpoints(
	log *zap.Logger,
) UserEndpoints {
	return newUserEndpoints(log)
}

func newUserEndpoints(
	log *zap.Logger,
) *userHandler {
	return &userHandler{
		log: log.With(zap.String("module", "user_endpoints")),
	}
}

func (h *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Get users call")
}

func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Create user call")
}
