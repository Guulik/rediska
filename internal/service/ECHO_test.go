package service

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io"
	api_mocks "rediska/internal/api/mocks"
	"rediska/internal/util/resper"
	"testing"
)

func TestServiceECHO_Happy(t *testing.T) {

	tests := []struct {
		name   string
		phrase string
		want   string
	}{
		{
			name:   "hi",
			phrase: "Helloi",
			want:   "$6\r\nHelloi\r\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checker := api_mocks.NewChecker(t)
			checker.On("ECHO", tt.phrase).Return(resper.EncodeBulkString(tt.phrase))

			got, err := checker.ECHO(tt.phrase)
			require.NoError(t, err)

			reader := bytes.NewReader(got.Bytes())
			data, err := io.ReadAll(reader)
			require.NoError(t, err)
			require.Equal(t, tt.want, string(data))
		})
	}
}
