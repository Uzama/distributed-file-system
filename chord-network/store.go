package main

import (
	"context"
	"errors"
)

type Store interface {
	Get(ctx context.Context, key string) (string, error)
	Put(ctx context.Context, key string) (error)
}

type store struct {
	store map[string]struct{}
}

func NewStore() Store {
	return &store{
		store: make(map[string]struct{}),
	}
}

func (s *store) Get(ctx context.Context, key string) (string, error) {
	_, ok := s.store[key]
	if ok {
		return key, nil
	}
	return "", errors.New("key not available")
}

func (s *store) Put(ctx context.Context, key string) (error) {
	s.store[key] = struct{}{}

	return nil
}
