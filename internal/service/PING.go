package service

import (
	"bytes"
	"rediska/internal/lib/logger/sl"
	"rediska/internal/util/resper"
)

func (s *CommandsService) PING() (bytes.Buffer, error) {
	log := s.log.With("op", "Service.PING")

	buf, err := resper.EncodeSimpleString("PONG")
	if err != nil {
		log.Error("failed to encode:", sl.Err(err))
		return bytes.Buffer{}, err
	}

	return buf, nil
}
