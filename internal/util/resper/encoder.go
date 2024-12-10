package resper

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/tidwall/resp"
	"strings"
)

var (
	SimpleString_CharLimit = 50
)

func EncodeSimpleString(s string) (bytes.Buffer, error) {
	var buf bytes.Buffer

	if strings.ContainsAny(s, "\r\n") {
		return bytes.Buffer{}, errors.New("simple string should not contain \"\\r\\n\"")
	}
	if len(s) > SimpleString_CharLimit {
		return bytes.Buffer{}, fmt.Errorf("simple string should be less than %d characters", SimpleString_CharLimit)
	}

	wr := resp.NewWriter(&buf)
	err := wr.WriteSimpleString(s)

	if err != nil {
		fmt.Println("failed to encode string with resper")
		return bytes.Buffer{}, fmt.Errorf("%s: %w", "failed to encode string with resper", err)
	}

	return buf, nil
}

func EncodeBulkString(s string) (bytes.Buffer, error) {
	var (
		buf bytes.Buffer
		err error
	)
	wr := resp.NewWriter(&buf)
	if s != "" {
		err = wr.WriteString(s)
	} else {
		err = wr.WriteNull()
	}

	if err != nil {
		fmt.Println("failed to encode string with resper")
		return bytes.Buffer{}, fmt.Errorf("%s: %w", "failed to encode string with resper", err)
	}

	return buf, nil
}

func EncodeSimpleError(error error) (bytes.Buffer, error) {
	var buf bytes.Buffer

	if error == nil {
		return bytes.Buffer{}, errors.New("error should not be null")
	}

	wr := resp.NewWriter(&buf)

	err := wr.WriteError(error)
	if err != nil {
		fmt.Println("failed to encode error with resper")
		return bytes.Buffer{}, fmt.Errorf("%s: %w", "failed to encode string with resper", err)
	}

	return buf, nil
}
