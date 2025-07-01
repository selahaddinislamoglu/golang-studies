package server

type HTTPServer interface {
	Serve(addr string) error
	Shutdown() error
}
