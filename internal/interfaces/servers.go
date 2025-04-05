package interfaces

type HTTPServer interface {
	Start()
	Notify() <-chan error
	Shutdown() error
}
