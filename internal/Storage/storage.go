package Storage

import "sync"

type Storage struct {
	data map[string]string
	mu   sync.RWMutex
}

var storage *Storage

// New creates and return new storage.
func New() *Storage {
	return &Storage{
		data: make(map[string]string),
	}
}

// Init initialize global storage
func Init() {
	storage = New()
}

// GetInstance returns pointer to local storage
func GetInstance() *Storage {
	return storage
}

// Get gets value by key
func (s *Storage) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok := s.data[key]
	return value, ok
}

// Set sets key by value
func (s *Storage) Set(key string, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}
