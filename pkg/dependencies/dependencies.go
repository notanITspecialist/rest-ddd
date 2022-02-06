package dependencies

type (
	Dependencies interface {
	}

	dependencies struct {
	}
)

func NewDependencies() Dependencies {
	return newDependencies()
}

func newDependencies() *dependencies {
	return &dependencies{}
}
