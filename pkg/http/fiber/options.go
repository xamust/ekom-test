package fiber

import (
	"net"
	"time"
)

type Option func(*Server)

func WithPort(port string) Option {
	return func(s *Server) {
		s.address = net.JoinHostPort("", port)
	}
}

func WithPrefork(prefork bool) Option {
	return func(s *Server) {
		s.prefork = prefork
	}
}

func WithReadTimeout(readTimeout time.Duration) Option {
	return func(s *Server) {
		s.readTimeout = readTimeout
	}
}

func WithWriteTimeout(writeTimeout time.Duration) Option {
	return func(s *Server) {
		s.writeTimeout = writeTimeout
	}
}

func WithShutdownTimeout(shutdownTimeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = shutdownTimeout
	}
}
