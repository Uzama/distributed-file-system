package chordnetwork

import (
	"context"
	"sync"
)

type Store struct {
	KeyValue map[string]string
	mutex    *sync.RWMutex
}

func (s *Store) Ping(ctx context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func (s *Store) Get(ctx context.Context, request interface{}) (interface{}, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	return nil, nil
}

func (s *Store) Put(ctx context.Context, request interface{}) (interface{}, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	return nil, nil
}