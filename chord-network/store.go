package main

import (
	"context"
	"errors"
	"sync"
)

type Store interface {
	Get(ctx context.Context, key string) (string, error)
	Put(ctx context.Context, key string) (error)
}

type store struct {
	store map[string]struct{}
	mutex    *sync.RWMutex
}

func NewStore() Store {
	return &store{
		store: make(map[string]struct{}),
	}
}

func (s *store) Get(ctx context.Context, key string) (string, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	_, ok := s.store[key]
	if ok {
		return key, nil
	}
	return "", errors.New("key not existing")
}

func (s *store) Put(ctx context.Context, key string) (error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	s.store[key] = struct{}{}

	return nil
}
