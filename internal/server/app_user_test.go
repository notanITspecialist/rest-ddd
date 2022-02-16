package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"rest-ddd/internal/endpoints"
	"rest-ddd/internal/repository"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	"rest-ddd/internal/config"
	mockEndpoints "rest-ddd/internal/mocks/endpoints"
	mockService "rest-ddd/internal/mocks/service"
)

type UserEndpointsSuite struct {
	suite.Suite

	ctrl *gomock.Controller

	mockUserEndpoints *mockEndpoints.MockUserEndpoints
	userEndpoints     endpoints.UserEndpoints

	mockUserService *mockService.MockUserService

	app     *appServer
	mockApp *appServer
}

func TestUserEndpointsSuite(t *testing.T) {
	suite.Run(t, new(UserEndpointsSuite))
}

func (s *UserEndpointsSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())

	s.mockUserEndpoints = mockEndpoints.NewMockUserEndpoints(s.ctrl)

	s.mockUserService = mockService.NewMockUserService(s.ctrl)
	s.userEndpoints = endpoints.NewUserEndpoints(zap.NewNop(), s.mockUserService)

	cfg := config.Server{Port: 1234}
	s.mockApp = newAppServer(zap.NewNop(), cfg, s.mockUserEndpoints, nil)
	s.app = newAppServer(zap.NewNop(), cfg, s.userEndpoints, nil)
}

func (s *UserEndpointsSuite) TestMockGetAllUsersEndpoint() {
	req, _ := http.NewRequest(http.MethodGet, Paths["GetUsers"], nil)
	res := httptest.NewRecorder()

	s.mockUserEndpoints.EXPECT().GetUsers(res, gomock.Any())
	s.mockApp.server.Handler.ServeHTTP(res, req)

	s.Equal(http.StatusOK, res.Code)
}

func (s *UserEndpointsSuite) TestGetAllUsersEndpoint() {
	path := Paths["GetUsers"]
	method := http.MethodGet

	s.Run("GetUsers call success", func() {
		endpointUrl := path + "" // there may be query parameters
		req, _ := http.NewRequest(method, endpointUrl, nil)
		res := httptest.NewRecorder()

		expectedData := []repository.User{
			{"1", "John", "Doe"},
			{"2", "John", "Doe"},
		}

		s.mockUserService.EXPECT().GetAllUsers(gomock.Any()).Return(expectedData, nil)
		s.app.server.Handler.ServeHTTP(res, req)

		var mapBody []repository.User
		err := json.Unmarshal(res.Body.Bytes(), &mapBody)
		s.NoError(err)
		s.Equal(http.StatusOK, res.Code)
		s.Equal(expectedData, mapBody)
	})
}
