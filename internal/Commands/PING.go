package Commands

import (
	"net"
	"rediska/internal/domain/response"
)

func PING(conn net.Conn) {
	response.PONG(conn)
}
