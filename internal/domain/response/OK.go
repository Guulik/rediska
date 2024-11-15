package response

import (
	"fmt"
	"net"
	"rediska/internal/util/resper"
)

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
