package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"rediska/config"
	"rediska/internal/api"
	"rediska/internal/service"
	"rediska/internal/storage"
)

type Server struct {
	log *slog.Logger
	cfg *config.Config
	api *api.API
}

func New(
	log *slog.Logger,
	cfg *config.Config,
) *Server {

	//TODO: use context
	_ = context.Background()

	repo := storage.New()
	serv := service.New(log, repo)
	API := api.New(log, serv, serv, serv)
	API.RegisterCommands()

	srv := &Server{log: log, cfg: cfg, api: API}

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

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		s.api.HandleInput(conn)
	}
}

func (s *Server) Stop() {
	//TODO: implement me
}
