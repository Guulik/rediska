package service

import (
	"bytes"
	"rediska/internal/lib/logger/sl"
	"rediska/internal/util/resper"
)

func (s *CommandsService) ECHO(phrase string) (bytes.Buffer, error) {
	log := s.log.With("op", "Service.ECHO")

	buf, err := resper.EncodeBulkString(phrase)
	if err != nil {
		log.Error("failed to encode:", sl.Err(err))
		return bytes.Buffer{}, err
	}
	return buf, nil
}
