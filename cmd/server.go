package main

import (
	"fmt"
	"log"
	"net"
	"rediska/config"
	"rediska/internal/Storage"
	"rediska/internal/api"
)

func main() {
	cfg := config.MustLoad()

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
