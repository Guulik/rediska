package service

import "log/slog"

type CommandsService struct {
	log      *slog.Logger
	provider Provider
	saver    Saver
}

type Provider interface {
	Get(key string) (string, bool)
}
type Saver interface {
	Set(key string, value string)
}

func New(
	log *slog.Logger,
	provider Provider,
	saver Saver,
) *CommandsService {
	return &CommandsService{
		log:      log,
		provider: provider,
		saver:    saver,
	}
}
