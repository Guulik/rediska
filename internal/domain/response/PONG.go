package response

import (
	"fmt"
	"net"
	"rediska/internal/util/resper"
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