package response

import (
	"bytes"
	"fmt"
	"rediska/internal/util/resper"
)

func CreateError(err error) bytes.Buffer {
	buf, e := resper.EncodeSimpleError(err)
	if e != nil {
		fmt.Println("failed to encode error:", err)
		return bytes.Buffer{}
	}
	return buf
}
