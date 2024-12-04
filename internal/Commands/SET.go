package Commands

import (
	"net"
	"rediska/internal/Storage"
	"rediska/internal/domain/response"
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

func SET(conn net.Conn, key string, value string, opts ...SetOptionFunc) {
	args := &SetArgs{
		key:   key,
		value: value,
	}

	for _, opt := range opts {
		opt(args)
	}

	storage := Storage.GetInstance()
	storage.Set(args.key, args.value)

	response.OK(conn)
}
