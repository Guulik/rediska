package service

import "log/slog"

type CommandsService struct {
	log *slog.Logger
}

func New(log *slog.Logger) *CommandsService {
	return &CommandsService{log: log}
}
