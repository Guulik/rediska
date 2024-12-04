package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"rediska/config"
	"rediska/internal/app/server"
	"rediska/internal/lib/logger"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger("local")

	redisServer := server.New(log, cfg)

	go redisServer.MustRun()
	log.Info("Redis is Started :)", slog.String("address", fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)))

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = logger.SetupPrettySlog()
	case "prod":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
