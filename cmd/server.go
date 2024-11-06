package main

import (
	"log/slog"
	"os"
	"rediska/config"
	"rediska/internal/app"
	"rediska/internal/lib/logger/handlers/slogpretty"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger("local")

	a := app.New(log, cfg)

	a.RedisServer.MustRun()
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = setupPrettySlog()
	case "prod":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
