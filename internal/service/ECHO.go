package service

import (
	"bytes"
	"rediska/internal/util/resper"
)

func (s *CommandsService) ECHO(phrase string) (bytes.Buffer, error) {
	log := s.log.With("op", "Service.ECHO")

	buf, err := resper.EncodeSimpleString(phrase)
	if err != nil {
		log.Error("failed to encode:", err)
		return bytes.Buffer{}, err
	}
	return buf, nil
}
