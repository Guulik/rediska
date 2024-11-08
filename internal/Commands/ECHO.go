package Commands

import (
	"fmt"
	"net"
	"rediska/internal/util/RESP"
)

func ECHO(conn net.Conn, phrase string) {
	buf, err := RESP.EncodeSimpleString(phrase)
	if err != nil {
		fmt.Println("failed to encode:", err)
	}

	_, err = conn.Write(buf.Bytes())
	if err != nil {
		fmt.Println("failed to write response to client")
	}
}
