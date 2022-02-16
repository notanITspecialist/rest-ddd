package app

import "rest-ddd/internal/dependencies"

type (
	Application interface {
		Run()
	}

	application struct {
		deps dependencies.Dependencies
	}
)

func NewApplications() Application {
	return newApplications(dependencies.NewDependencies())
}

func newApplications(deps dependencies.Dependencies) *application {
	return &application{
		deps: deps,
	}
}

func (app *application) Run() {
	deps := app.deps

	appServer := deps.AppServer()
	appServer.Start()
}
