package service

import (
	"bytes"
	"rediska/internal/Storage"
	"rediska/internal/util/resper"
	"time"
)

type SetOptionFunc func(*SetArgs)

func WithTTL(ttl time.Duration) SetOptionFunc {
	return func(s *SetArgs) {
		s.ttl = ttl
	}
}

type SetArgs struct {
	key   string
	value string
	GET   bool
	ttl   time.Duration
}

func (s *CommandsService) SET(key string, value string, opts ...SetOptionFunc) (bytes.Buffer, error) {
	log := s.log.With("op", "Service.SET")

	args := &SetArgs{
		key:   key,
		value: value,
	}

	for _, opt := range opts {
		opt(args)
	}

	storage := Storage.GetInstance()
	// we do not handle error because setting to map is trivial...
	storage.Set(args.key, args.value)

	buf, err := resper.EncodeSimpleString("OK")
	if err != nil {
		log.Error("failed to encode: ", err)
		return bytes.Buffer{}, err
	}
	return buf, nil
}
