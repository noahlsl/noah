package mapx

import (
	"errors"
	"sync"
)

var (
	ErrValueNil = errors.New("the value nil")
)

type SafeMap[T1, T2 any] struct {
	sync.RWMutex
	m map[*T1]T2
}

func NewSafeMap[T1, T2 any]() *SafeMap[T1, T2] {
	return &SafeMap[T1, T2]{
		m: make(map[*T1]T2),
	}
}

func (s *SafeMap[T1, T2]) Set(key T1, value T2) {
	s.Lock()
	defer s.Unlock()
	s.m[&key] = value
}

func (s *SafeMap[T1, T2]) Get(key T1) (T2, error) {
	s.RLock()
	defer s.RUnlock()
	if v, ok := s.m[&key]; ok {
		return v, nil
	}

	return *new(T2), ErrValueNil
}

func (s *SafeMap[T1, T2]) Exist(key T1) bool {
	s.RLock()
	defer s.RUnlock()
	if _, ok := s.m[&key]; ok {
		return true
	}

	return false
}
