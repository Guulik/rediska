package response

import (
	"fmt"
	"net"
	"rediska/internal/util/resper"
)

func Error(conn net.Conn, err error) {
	buf, e := resper.EncodeSimpleError(err)
	if e != nil {
		fmt.Println("failed to encode:", err)
	}

	_, e = conn.Write(buf.Bytes())
	if e != nil {
		fmt.Println("failed to write PONG to client")
	}
}
