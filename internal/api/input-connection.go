package api

import (
	"bytes"
	"fmt"
	"github.com/tidwall/resp"
	"io"
	"log/slog"
	"net"
	"rediska/internal/util/resper"
)

func (a *API) HandleInput(conn net.Conn) {
	const op = "api.HandleInput"
	log := a.log.With(slog.String("op", op))

	a.setConn(conn)
	for {
		v, err := a.readInput()
		log.Debug("resp value and err",
			slog.Any("value", resper.RespValuesToAny(v.Array())),
			slog.Any("error", err))

		if err != nil {
			log.Error("FATAL Err = ", err)
			a.conn.Close()
			return
		}

		command := v.Array()[0].String()
		fmt.Println("command: ", command)
		args := resper.RespValuesToAny(v.Array()[1:])

		handler, ok := a.router.routes[command]
		if !ok {
			fmt.Printf("Unknown command: %service\n", command)
			return
		}
		handler(args)
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
		slog.Any("readerValue: ", resper.RespValuesToAny(v.Array())))

	return v, nil
}
