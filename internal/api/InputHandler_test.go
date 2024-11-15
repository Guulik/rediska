package api

import (
	"github.com/tidwall/resp"
	"log/slog"
	"net"
	"reflect"
	"testing"
)

func TestAPI_readInput(t *testing.T) {
	type fields struct {
		log  *slog.Logger
		conn net.Conn
	}
	type args struct {
		conn    net.Conn
		message []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    resp.Value
		wantErr bool
	}{
		{
			name: "ping",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &API{
				log:  tt.fields.log,
				conn: tt.fields.conn,
			}
			got, err := a.readInput(tt.args.conn)
			if (err != nil) != tt.wantErr {
				t.Errorf("readInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readInput() got = %v, want %v", got, tt.want)
			}
		})
	}
}
