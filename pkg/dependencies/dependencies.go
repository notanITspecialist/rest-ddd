package dependencies

import "rest-ddd/pkg/server"

type (
	Dependencies interface {
		AppServer() server.Server
	}

	dependencies struct {
		appServer server.Server
	}
)

func NewDependencies() Dependencies {
	return newDependencies()
}

func newDependencies() *dependencies {
	return &dependencies{}
}

func (d *dependencies) AppServer() server.Server {
	if d.appServer == nil {
		// initialization server.NewAppServer
	}
	return d.appServer
}
