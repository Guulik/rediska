package service

import "log/slog"

type CommandsService struct {
	log     *slog.Logger
	storage RedisStorage
}

type RedisStorage interface {
	Get(key string) (string, bool)
	Set(key string, value string)
}

func New(
	log *slog.Logger,
	storage RedisStorage,
) *CommandsService {
	return &CommandsService{
		log:     log,
		storage: storage,
	}
}
