package service

import (
	"bytes"
	"errors"
	"rediska/internal/lib/logger/sl"
	"rediska/internal/util/resper"
)

var ErrorNotFound = errors.New("value by provided key is not found")

func (s *CommandsService) GET(key string) (bytes.Buffer, error) {
	log := s.log.With("op", "Service.GET")

	var (
		buf bytes.Buffer
		err error
	)
	value, exists := s.provider.Get(key)
	if !exists {
		return bytes.Buffer{}, ErrorNotFound
	}

	buf, err = resper.EncodeBulkString(value)
	if err != nil {
		log.Error("failed to encode:", sl.Err(err))
		return bytes.Buffer{}, err
	}

	return buf, nil
}
