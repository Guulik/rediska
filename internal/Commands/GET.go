package Commands

import (
	"fmt"
	"net"
	"rediska/internal/Storage"
	"rediska/internal/util/resper"
)

func GET(conn net.Conn, key string) {
	op := " SET "

	storage := Storage.GetInstance()
	if value, exists := storage.Get(key); exists {

		buf, err := resper.EncodeBulkString(value)
		if err != nil {
			fmt.Println("op:", op, "failed to encode: ", err)
		}

		_, err = conn.Write(buf.Bytes())
		if err != nil {
			fmt.Println("op:", op, "failed to write response to client")
		}
	}

}
