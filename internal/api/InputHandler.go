package api

import (
	"bytes"
	"fmt"
	"github.com/tidwall/resp"
	"io"
	"log/slog"
	"net"
	"os"
	"rediska/internal/Commands"
	"strings"
	"time"
)

type API struct {
	log  *slog.Logger
	conn net.Conn
}

func New(
	log *slog.Logger,
) *API {
	return &API{log: log}
}

func (a *API) setConn(conn net.Conn) {
	a.conn = conn
}

func (a *API) execute(command string, args []any) {
	switch command {
	case "PING":
		fmt.Println("ponging...")
		Commands.PING(a.conn)
	case "ECHO":
		fmt.Println("echoing...")
		phrase := args[0].(string)
		Commands.ECHO(a.conn, phrase)
	case "SET":
		fmt.Println("setting...")
		key := args[0].(string)
		value := args[1].(string)
		ttlType := args[2].(string)
		expire := args[3].(string)

		var (
			ttl time.Duration
			err error
		)
		if strings.EqualFold(ttlType, "px") {
			ttl, err = time.ParseDuration(expire + "ms")
			if err != nil {
				fmt.Println("failed to parse duration")
			}
			Commands.SET(a.conn, key, value, Commands.WithTTL(ttl))
			break
		}
		Commands.SET(a.conn, key, value)
	case "GET":
		fmt.Println("getting...")
		key := args[0].(string)
		Commands.GET(a.conn, key)
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}

func (a *API) HandleInput(conn net.Conn) {
	const op = "api.HandleInput"
	log := a.log.With(slog.String("op", op))

	a.setConn(conn)
	for {
		v, err := a.readInput()
		log.Debug("resp value and err",
			slog.Any("value", a.convertRespValuesToAnyArray(v.Array())),
			slog.Any("error", err))

		if err != nil {
			log.Error("FATAL Err = ", err)
			a.conn.Close()
			os.Exit(1)
		}

		command := v.Array()[0].String()
		fmt.Println("command: ", command)

		args := a.convertRespValuesToAnyArray(v.Array()[1:])
		a.execute(command, args)
	}
}

func (a *API) readInput() (resp.Value, error) {
	const op = "api.readInput"
	log := a.log.With(slog.String("op", op))

	buf := make([]byte, 128)
	n, err := a.conn.Read(buf)
	if err != nil {
		if err == io.EOF {
			log.Warn("connection closed by client")
			return resp.NullValue(), err
		}
		log.Error("failed to read bytes", err)
		return resp.NullValue(), fmt.Errorf("failed to read bytes: %w", err)
	}
	if n == 0 {
		log.Warn("no data received")
		return resp.NullValue(), fmt.Errorf("no data to read")
	}

	var v resp.Value
	rd := resp.NewReader(bytes.NewReader(buf[:n]))
	v, _, err = rd.ReadValue()
	if err != nil {
		log.Error("failed to parse RESP value", err)
		return resp.NullValue(), fmt.Errorf("failed to parse RESP value: %w", err)
	}

	log.Debug("buffer and RESP value debug",
		slog.Any("received bytes:", buf[:n]),
		slog.Any("readerValue: ", a.convertRespValuesToAnyArray(v.Array())))

	return v, nil
}

func (a *API) convertRespValuesToAnyArray(values []resp.Value) []any {
	var result []any
	for _, v := range values {
		switch v.Type() {
		case resp.BulkString:
			result = append(result, v.String())
		case resp.SimpleString:
			result = append(result, v.String())
		case resp.Integer:
			result = append(result, v.Integer())
		case resp.Error:
			result = append(result, v.Error())
		case resp.Array:
			result = append(result, v.Array())
		default:
			result = append(result, nil)
		}
	}
	return result
}
