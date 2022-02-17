package dependencies

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"rest-ddd/internal/config"
	"rest-ddd/internal/db"
	"rest-ddd/internal/endpoints"
	"rest-ddd/internal/repository"
	"rest-ddd/internal/server"
	"rest-ddd/internal/service"
)

type (
	Dependencies interface {
		AppServer() server.Server
	}

	dependencies struct {
		log    *zap.Logger
		config *config.Config

		appServer server.Server

		psqlClient *db.PostgresqlClient

		userEndpoints endpoints.UserEndpoints

		userService service.UserService

		userRepository repository.UserRepository
	}
)

func NewDependencies() Dependencies {
	return newDependencies()
}

func newDependencies() *dependencies {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	logger := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			zapcore.Lock(os.Stdout),
			zap.NewAtomicLevel(),
		),
	).Named("dependencies")

	return &dependencies{
		log: logger,
	}
}

func (d *dependencies) Config() *config.Config {
	if d.config == nil {
		msg := "Initialize [dependencies.Config]"

		cfg, err := config.NewConfig()
		if err != nil {
			d.log.Panic(msg, zap.Error(err))
		}
		d.log.Info(msg)

		d.config = cfg
	}
	return d.config
}

func (d *dependencies) AppServer() server.Server {
	if d.appServer == nil {
		msg := "Initialize [dependencies.AppServer]"
		d.log.Info(msg)

		appServer, err := server.NewAppServer(
			d.log,
			d.Config().AppServer,
			d.UserEndpoints(),
		)
		if err != nil {
			d.log.Panic(msg, zap.Error(err))
		}
		d.appServer = appServer
	}
	return d.appServer
}

func (d *dependencies) PostgresqlClient() *db.PostgresqlClient {
	if d.psqlClient == nil {
		msg := "initialize [dependencies.PostgresqlClient]"
		psqlClient, err := db.NewPostgresClient(d.Config().Postgresql)
		if err != nil {
			d.log.Panic(msg, zap.Error(err))
		}

		d.psqlClient = psqlClient
	}
	return d.psqlClient
}

func (d *dependencies) UserEndpoints() endpoints.UserEndpoints {
	if d.userEndpoints == nil {
		d.userEndpoints = endpoints.NewUserEndpoints(d.log, d.UserService())
		d.log.Info("Initialize [dependencies.UserEndpoints]")
	}
	return d.userEndpoints
}

func (d *dependencies) UserService() service.UserService {
	if d.userService == nil {
		msg := "Initialize [dependencies.UserService]"
		d.userService = service.NewUserService(d.UserRepository())
		d.log.Info(msg)
	}
	return d.userService
}

func (d *dependencies) UserRepository() repository.UserRepository {
	if d.userRepository == nil {
		msg := "Initialize [dependencies.UserRepository]"
		d.userRepository = repository.NewPGUserRepository(d.log, d.PostgresqlClient())
		d.log.Info(msg)
	}
	return d.userRepository
}
