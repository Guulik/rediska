package api

import (
	"github.com/stretchr/testify/require"
	"log/slog"
	"net"
	"testing"
)

func TestAPI_readInput(t *testing.T) {
	type ApiFields struct {
		log  *slog.Logger
		conn net.Conn
	}
	tests := []struct {
		name        string
		apiFields   ApiFields
		message     []byte
		wantCommand string
	}{
		{
			name: "ping",

			wantCommand: "ping",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &API{
				log:  tt.apiFields.log,
				conn: tt.apiFields.conn,
			}
			got, err := a.readInput(tt.apiFields.conn)

			require.NoError(t, err)
			command := got.Array()[0].String()
			require.Equal(t, tt.wantCommand, command)
		})
	}
}
