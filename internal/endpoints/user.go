package endpoints

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"rest-ddd/internal/service"
)

type (
	UserEndpoints interface {
		GetUsers(http.ResponseWriter, *http.Request)
		CreateUser(http.ResponseWriter, *http.Request)
	}

	userHandler struct {
		log         *zap.Logger
		userService service.UserService
	}
)

func NewUserEndpoints(
	log *zap.Logger,
	userService service.UserService,
) UserEndpoints {
	return newUserEndpoints(log, userService)
}

func newUserEndpoints(
	log *zap.Logger,
	userService service.UserService,
) *userHandler {
	return &userHandler{
		log:         log.With(zap.String("module", "user_endpoints")),
		userService: userService,
	}
}

func (h *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Get users call")

	users, err := h.userService.GetAllUsers(r.Context())
	if err != nil {
		h.log.Error("Error while call [service.UserService.GetAllUsers]", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		h.log.Error("Error while encode [endpoints.UserEndpoints.GetUsers]", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
