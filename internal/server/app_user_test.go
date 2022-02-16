package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	"rest-ddd/internal/config"
	mockEndpoints "rest-ddd/internal/mocks/endpoints"
)

type UserEndpointsSuite struct {
	suite.Suite

	ctrl *gomock.Controller

	mockUserEndpoints *mockEndpoints.MockUserEndpoints

	app *appServer
}

func TestUserEndpointsSuite(t *testing.T) {
	suite.Run(t, new(UserEndpointsSuite))
}

func (s *UserEndpointsSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())

	s.mockUserEndpoints = mockEndpoints.NewMockUserEndpoints(s.ctrl)

	cfg := config.Server{Port: 1234}
	s.app = newAppServer(zap.NewNop(), cfg, s.mockUserEndpoints, nil)
}

func (s *UserEndpointsSuite) TestGetAllUsersEndpoint() {
	req, _ := http.NewRequest(http.MethodGet, Paths["GetUsers"], nil)
	res := httptest.NewRecorder()

	s.mockUserEndpoints.EXPECT().GetUsers(res, gomock.Any())
	s.app.server.Handler.ServeHTTP(res, req)

	s.Equal(http.StatusOK, res.Code)
}
