package dependencies

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"rest-ddd/pkg/server"
)

type (
	Dependencies interface {
		AppServer() server.Server
	}

	dependencies struct {
		log *zap.Logger

		appServer server.Server
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

func (d *dependencies) AppServer() server.Server {
	if d.appServer == nil {
		msg := "Initialize [dependencies.AppServer]"
		d.log.Info(msg)

		appServer, err := server.NewAppServer(d.log)
		if err != nil {
			d.log.Panic(msg, zap.Error(err))
		}
		d.appServer = appServer
	}
	return d.appServer
}
