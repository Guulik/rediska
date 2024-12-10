package resper

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
)

func TestEncodeBulkString_Happy(t *testing.T) {
	tests := []struct {
		name   string
		string string
		want   string
	}{
		{
			name:   "1 bukva",
			string: "j",
			want:   "$1\r\nj\r\n",
		},
		{
			name:   "with space",
			string: "unga bunga",
			want:   "$10\r\nunga bunga\r\n",
		},
		{
			name:   "mnoga",
			string: "Саша, мы прислали новый промокод, чтобы в череде сложных дел появилось время отдохнуть. Просто сделайте заказ от 1 290₽ и пицца Карбонара 25 см будет стоить всего 1₽.\nПромокод: J1DGE\nДействует до 19 декабря\n\nЧтобы получить пиццу по специальной цене, нужно:\n— собрать заказ от 1290₽ на dodopizza.ru, в мобильном приложении (vk.cc/99vs23) или в ресторане\n— добавить в корзину пиццу Карбонара 25 см\n— ввести или назвать промокод J1DGE. А если вы делаете заказ в ресторане на кассе, назовите свой номер телефона\n\nИ новая цена появится в корзине. Получите ваш заказ и наслаждайтесь!\n\nПолучите ваш заказ и наслаждайтесь!\nАкция действует в Москве, полный список городов доступен по ссылке https://vk.cc/cEyxKO\nНе суммируется с комбо и другими акциями, не действует на продукт с кастомизацией. Акция действительна пока товар есть в наличии.",
			want:   "$1451\r\nСаша, мы прислали новый промокод, чтобы в череде сложных дел появилось время отдохнуть. Просто сделайте заказ от 1 290₽ и пицца Карбонара 25 см будет стоить всего 1₽.\nПромокод: J1DGE\nДействует до 19 декабря\n\nЧтобы получить пиццу по специальной цене, нужно:\n— собрать заказ от 1290₽ на dodopizza.ru, в мобильном приложении (vk.cc/99vs23) или в ресторане\n— добавить в корзину пиццу Карбонара 25 см\n— ввести или назвать промокод J1DGE. А если вы делаете заказ в ресторане на кассе, назовите свой номер телефона\n\nИ новая цена появится в корзине. Получите ваш заказ и наслаждайтесь!\n\nПолучите ваш заказ и наслаждайтесь!\nАкция действует в Москве, полный список городов доступен по ссылке https://vk.cc/cEyxKO\nНе суммируется с комбо и другими акциями, не действует на продукт с кастомизацией. Акция действительна пока товар есть в наличии.\r\n",
		},
		{
			name:   "Pusto",
			string: "",
			want:   "$-1\r\n",
		},
		{
			name:   "space",
			string: " ",
			want:   "$1\r\n \r\n",
		},
		{
			name:   "slash n",
			string: "\r\n",
			want:   "$2\r\n\r\n\r\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncodeBulkString(tt.string)
			require.NoError(t, err)

			reader := bytes.NewReader(got.Bytes())
			data, err := io.ReadAll(reader)
			require.NoError(t, err)
			require.Equal(t, tt.want, string(data))
		})
	}
}

func TestEncodeSimpleError_Happy(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "chto-to",
			err:  errors.New("mya"),
			want: "-mya\r\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncodeSimpleError(tt.err)
			require.NoError(t, err)

			reader := bytes.NewReader(got.Bytes())
			data, err := io.ReadAll(reader)
			require.NoError(t, err)
			require.Equal(t, tt.want, string(data))
		})
	}
}

func TestEncodeSimpleError_Bad(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		wantErr error
	}{
		{
			name:    "nulik",
			err:     nil,
			wantErr: errors.New("error should not be null"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncodeSimpleError(tt.err)
			require.Error(t, err)
			require.EqualError(t, err, tt.wantErr.Error())
			require.Equal(t, bytes.Buffer{}, got)
		})
	}
}

func TestEncodeSimpleString_Happy(t *testing.T) {
	tests := []struct {
		name    string
		message string
		want    string
	}{
		{
			name:    "ok",
			message: "OK",
			want:    "+OK\r\n",
		},
		{
			name:    "pong",
			message: "PONG",
			want:    "+PONG\r\n",
		},
		{
			name:    "empty",
			message: "",
			want:    "+\r\n",
		},
		{
			name:    "space",
			message: " ",
			want:    "+ \r\n",
		},
		{
			name:    "mnogo bukv but equal limit",
			message: strings.Repeat("C", SimpleString_CharLimit),
			want:    "+" + strings.Repeat("C", SimpleString_CharLimit) + "\r\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncodeSimpleString(tt.message)
			require.NoError(t, err)

			reader := bytes.NewReader(got.Bytes())
			data, err := io.ReadAll(reader)
			require.NoError(t, err)
			require.Equal(t, tt.want, string(data))
		})
	}
}
func TestEncodeSimpleString_Bad(t *testing.T) {
	tests := []struct {
		name    string
		message string
		wantErr error
	}{
		{
			name:    "crlf",
			message: "\r\n",
			wantErr: errors.New("simple string should not contain \"\\r\\n\""),
		},
		{
			name:    "mnogo bukv",
			message: strings.Repeat("A", SimpleString_CharLimit+1),
			wantErr: fmt.Errorf("simple string should be less than %d characters", SimpleString_CharLimit),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncodeSimpleString(tt.message)
			require.Error(t, err)
			require.EqualError(t, err, tt.wantErr.Error())
			require.Equal(t, bytes.Buffer{}, got)
		})
	}
}
