package Commands

import (
	"net"
	"rediska/internal/api"
)

func PING(conn net.Conn) {
	api.PONG(conn)
}
