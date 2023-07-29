package main

import (
	"context"
	"errors"
)

type Store interface {
	Get(ctx context.Context, key string) (string, error)
	Put(ctx context.Context, key string) (error)
	GetKeys(ctx context.Context) ([]string)
	Delete(ctx context.Context, key string)
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

func (s *store) GetKeys(ctx context.Context) ([]string) {

	keys := make([]string, len(s.store))

	i := 0
	for k := range s.store {
		keys[i] = k
		i++
	}

	return keys
}

func (s *store) Delete(ctx context.Context, key string) {

	delete(s.store, key)
}
