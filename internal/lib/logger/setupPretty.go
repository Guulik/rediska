package logger

import (
	"log/slog"
	"os"
	"rediska/internal/lib/logger/handlers/slogpretty"
)

func SetupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
