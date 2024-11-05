package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"rediska/config"
	"rediska/internal/Storage"
	"rediska/internal/api"
	"rediska/internal/lib/logger/handlers/slogpretty"
)

func main() {
	cfg := config.MustLoad()

	_ = setupLogger("local")

	address := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Failed to bind to port 6379")
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
