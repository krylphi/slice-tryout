package repo

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"sync"
)

type Store struct {
	mtx  sync.RWMutex
	repo map[string]string
}

const shortTempl = "https://slice.us/%s"
const (
	initialStoreLen = 10
)

var ErrNoValue = errors.New("no value")

func NewStore() *Store {
	return &Store{
		mtx:  sync.RWMutex{},
		repo: make(map[string]string, initialStoreLen),
	}
}

func (s *Store) Save(_ context.Context, value string) (string, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	key := shorten(value)
	s.repo[key] = value
	return key, nil
}

func (s *Store) Load(_ context.Context, key string) (string, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	value, ok := s.repo[key]
	if !ok {
		return "", ErrNoValue
	}
	return value, nil
}

func shorten(value string) string {
	shrt := base64.StdEncoding.EncodeToString([]byte(value))[:7]
	return fmt.Sprintf(shortTempl, shrt)
}
