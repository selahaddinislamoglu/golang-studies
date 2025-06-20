package server

type Server interface {
	Serve(addr string) error
	Shutdown() error
}
