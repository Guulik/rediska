package main

import (
	"fmt"
	"net"
	"os"
	"rediska/internal/Storage"
	"rediska/internal/api"
)

func main() {

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
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
