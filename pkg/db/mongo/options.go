package mongo

import "time"

type Option func(s *Mongo)

func WithPort(port string) Option {
	return func(s *Mongo) {
		s.Port = port
	}
}

func WithURL(url string) Option {
	return func(s *Mongo) {
		s.URL = url
	}
}

func WithConnectTimeout(connectTimeout time.Duration) Option {
	return func(s *Mongo) {
		s.ConnectTimeout = connectTimeout
	}
}

func WithMaxPoolSize(maxPoolSize uint64) Option {
	return func(s *Mongo) {
		s.MaxPoolSize = maxPoolSize
	}
}
