package Commands

import (
	"fmt"
	"net"
	"rediska/internal/util/resper"
)

func ECHO(conn net.Conn, phrase string) {
	buf, err := resper.EncodeSimpleString(phrase)
	if err != nil {
		fmt.Println("failed to encode:", err)
	}

	_, err = conn.Write(buf.Bytes())
	if err != nil {
		fmt.Println("failed to write response to client")
	}
}
