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
		opt(svr)
	}
	return svr
}

type Option func(*Server)

func WithTimeout(timeout time.Duration) Option {
	return func(svr *Server) {
		svr.timeout = timeout
	}
}

func WithLogger(logger *log.Logger) Option {
	return func(svr *Server) {
		svr.logger = logger
	}
}
