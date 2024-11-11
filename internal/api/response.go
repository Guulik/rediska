package api

import (
	"fmt"
	"rediska/internal/util/resper"

	"net"
)

func PONG(conn net.Conn) {
	buf, err := resper.EncodeSimpleString("PONG")
	if err != nil {
		fmt.Println("failed to encode:", err)
	}

	_, err = conn.Write(buf.Bytes())
	if err != nil {
		fmt.Println("failed to write PONG to client")
	}
}

func OK(conn net.Conn) {
	op := "OK response"

	buf, err := resper.EncodeSimpleString("OK")
	if err != nil {
		fmt.Println("op:", op, "failed to encode: ", err)
	}

	_, err = conn.Write(buf.Bytes())
	if err != nil {
		fmt.Println("op:", op, "failed to write response to client")
	}
}
