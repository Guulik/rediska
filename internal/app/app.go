package app

import (
	"log/slog"
	"rediska/config"
	"rediska/internal/app/server"
)

type App struct {
	RedisServer *server.Server
}

func New(
	log *slog.Logger,
	cfg *config.Config,
) *App {
	srv := server.New(log, cfg)

	return &App{RedisServer: srv}
}
