package api

import (
	"github.com/stretchr/testify/require"
	"net"
	"rediska/internal/lib/logger"
	"strings"
	"testing"
)

func TestAPI_readInput(t *testing.T) {
	tests := []struct {
		name                string
		message             string
		wantReceivedMessage string
		wantCommand         string
		wantArgs            []any
	}{
		{
			name:                "ping",
			message:             "*1\r\n$4\r\nPING\r\n",
			wantReceivedMessage: "PING",
			wantCommand:         "PING",
			wantArgs:            nil,
		},
		{
			name:                "set",
			message:             "*3\r\n$3\r\nSET\r\n$3\r\ncar\r\n$3\r\n911\r\n",
			wantReceivedMessage: "SET car 911",
			wantCommand:         "SET",
			wantArgs:            []any{"car", "911"},
		},
		{
			name:                "get",
			message:             "*2\r\n$3\r\nGET\r\n$3\r\ncar\r\n",
			wantReceivedMessage: "GET car",
			wantCommand:         "GET",
			wantArgs:            []any{"car"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serverConn, clientConn := net.Pipe()
			defer serverConn.Close()
			defer clientConn.Close()

			a := &API{
				log:  logger.SetupPrettySlog(),
				conn: serverConn,
			}

			go func() {
				_, err := clientConn.Write([]byte(tt.message))
				require.NoError(t, err)
			}()

			value, err := a.readInput()
			require.NoError(t, err)

			message := strings.Trim(value.String(), "[]")
			command := value.Array()[0].String()
			respArgs := value.Array()[1:]
			args, convertErr := a.convertRespValuesToAnyArray(respArgs)
			require.NoError(t, convertErr)

			require.Equal(t, tt.wantReceivedMessage, message)
			require.Equal(t, tt.wantCommand, command)
			require.Equal(t, tt.wantArgs, args)
		})
	}
}
