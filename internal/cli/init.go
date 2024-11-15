package cli

import (
	"fmt"
	"log"
	"net"
)

type CliClient struct {
	conn net.Conn
}

var cli *CliClient

func InitCobra(address string) {
	cli = New(address)
	RootCmd.AddCommand(pingCmd, getCmd, setCmd)
}

func New(
	address string,
) *CliClient {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to connect to server: %w", err))
	}
	defer conn.Close()

	return &CliClient{conn: conn}
}
