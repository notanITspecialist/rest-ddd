package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
	"gopkg.in/validator.v2"

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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string][]string{
			"errors": {"Error while try to read body " + err.Error()},
		})
		return
	}

	var rBody service.UserCreateData
	err = json.Unmarshal(body, &rBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string][]string{
			"errors": {"Error while try to serialize body " + err.Error()},
		})
		return
	}
	if err = validator.Validate(rBody); err != nil {
		json.NewEncoder(w).Encode(map[string][]string{
			"errors": {"Error while try to serialize body " + err.Error()},
		})
		return
	}

	err = h.userService.CreateUser(r.Context(), rBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
