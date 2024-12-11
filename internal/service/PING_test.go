package service

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io"
	api_mocks "rediska/internal/service/mocks"
	"rediska/internal/util/resper"
	"testing"
)

func TestServicePING(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "hi",
			want: "+OK\r\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checker := api_mocks.NewChecker(t)
			checker.On("PING").Return(resper.EncodeSimpleString("OK"))

			got, err := checker.PING()
			require.NoError(t, err)

			reader := bytes.NewReader(got.Bytes())
			data, err := io.ReadAll(reader)
			require.NoError(t, err)
			require.Equal(t, tt.want, string(data))
		})
	}
}
