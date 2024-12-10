package api

import (
	"bytes"
	"fmt"
	"github.com/tidwall/resp"
	"io"
	"log/slog"
	"net"
	"rediska/internal/domain/response"
)

type API struct {
	log    *slog.Logger
	conn   net.Conn
	router *Router
}

func New(
	log *slog.Logger,
) *API {
	return &API{log: log, router: NewRouter()}
}

func (a *API) setConn(conn net.Conn) {
	a.conn = conn
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
			return
		}

		command := v.Array()[0].String()
		fmt.Println("command: ", command)
		args := a.convertRespValuesToAnyArray(v.Array()[1:])

		handler, ok := a.router.routes[command]
		if !ok {
			fmt.Printf("Unknown command: %s\n", command)
			return
		}
		err = handler(args)
		if err != nil {
			response.Error(a.conn, err)
		}
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
			return resp.NullValue(), fmt.Errorf("connection closed by client: %w", err)
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
