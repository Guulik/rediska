package service

import (
	"bytes"
	"rediska/internal/lib/logger/sl"
	"rediska/internal/util/resper"
)

func (s *CommandsService) GET(key string) (bytes.Buffer, error) {
	log := s.log.With("op", "Service.GET")

	var (
		buf bytes.Buffer
		err error
	)
	if value, exists := s.storage.Get(key); exists {

		buf, err = resper.EncodeBulkString(value)
		if err != nil {
			log.Error("failed to encode:", sl.Err(err))
			return bytes.Buffer{}, err
		}
	}
	return buf, nil
}
