package app

type (
	Application interface {
		Run()
	}

	application struct {
		// there will be dependencies
	}
)

func NewApplications() Application {
	return newApplications()
}

func newApplications() *application {
	return &application{}
}

func (application) Run() {
	// there will be dependencies initialization
}
