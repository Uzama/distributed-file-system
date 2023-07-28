package main

import (
	"context"
	"errors"
	"sync"
)

type Store struct {
	store map[string]struct{}
	mutex    *sync.RWMutex
}

func (s *Store) Get(ctx context.Context, key string) (string, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	_, ok := s.store[key]
	if ok {
		return key, nil
	}
	return "", errors.New("key not existing")
}

func (s *Store) Put(ctx context.Context, key string) (error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	s.store[key] = struct{}{}

	return nil
}
