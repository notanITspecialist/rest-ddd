package server

type Server interface {
	Start()
	Stop() error
}
