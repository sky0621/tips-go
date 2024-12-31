package server

import (
	"log"
	"time"
)

type Server struct {
	param Param
}

func (s *Server) Start() error {
	if s.param.logger != nil {
		s.param.logger.Println("server started")
	}
	return nil
}

type Param struct {
	host    string
	port    int
	timeout time.Duration
	logger  *log.Logger
}

func NewParamBuilder(host string, port int) *Param {
	return &Param{host: host, port: port}
}

func (p *Param) Timeout(timeout time.Duration) *Param {
	p.timeout = timeout
	return p
}

func (p *Param) Logger(logger *log.Logger) *Param {
	p.logger = logger
	return p
}

func (p *Param) Build() *Server {
	svr := &Server{param: *p}
	return svr
}
