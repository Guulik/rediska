package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"rediska/config"
	"rediska/internal/Storage"
	"rediska/internal/api"
)

type Server struct {
	log *slog.Logger

	cfg *config.Config
}

func New(
	log *slog.Logger,
	cfg *config.Config,
) *Server {

	_ = context.Background()

	//TODO: init api layer
	//TODO: init service layer
	//TODO: init repo layer

	srv := &Server{log: log, cfg: cfg}

	return srv
}

func (s *Server) MustRun() {
	if err := s.run(); err != nil {
		panic("failed to run server")
	}
}

func (s *Server) run() error {
	address := fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port)
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	Storage.Init()

	for {
		con, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		go api.ReadInput(con)
	}
}

func (s *Server) Stop() {
	//TODO: implement me
}
