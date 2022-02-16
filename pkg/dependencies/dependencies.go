package dependencies

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"rest-ddd/pkg/config"
	"rest-ddd/pkg/endpoints"
	"rest-ddd/pkg/server"
)

type (
	Dependencies interface {
		AppServer() server.Server

		UserEndpoints() endpoints.UserEndpoints
	}

	dependencies struct {
		log    *zap.Logger
		config *config.Config

		appServer server.Server

		userEndpoints endpoints.UserEndpoints
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

func (d *dependencies) UserEndpoints() endpoints.UserEndpoints {
	if d.userEndpoints == nil {
		d.userEndpoints = endpoints.NewUserEndpoints(d.log)
		d.log.Info("initialize user endpoints dependency")
	}
	return d.userEndpoints
}
