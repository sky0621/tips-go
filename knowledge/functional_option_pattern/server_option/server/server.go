package server

import (
	"log"
	"time"
)

type Server struct {
	host    string
	port    int
	timeout time.Duration
	logger  *log.Logger
}

func (s *Server) Start() error {
	return nil
}

func New(host string, port int, options ...Option) *Server {
	svr := &Server{host: host, port: port}
	for _, opt := range options {
		err := opt(svr)
		if err != nil {
			log.Fatal(err)
		}
	}
	return svr
}

type Option func(*Server) error

func WithTimeout(timeout time.Duration) Option {
	return func(svr *Server) error {
		svr.timeout = timeout
		return nil
	}
}

func WithLogger(logger *log.Logger) Option {
	return func(svr *Server) error {
		svr.logger = logger
		return nil
	}
}
