package xsync

import (
	"sync"
)

type Slice struct {
	mu   sync.RWMutex
	list []interface{}
}

func New(array []interface{}) *Slice {
	return &Slice{
		list: array,
	}
}

func (s *Slice) Append(args ...interface{}) {
	s.mu.Lock()
	s.list = append(s.list, args...)
	s.mu.Unlock()
}

func (s *Slice) Get(idx int) interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if idx < 0 || idx >= len(s.list) {
		return nil
	}
	return s.list[idx]
}

func (s *Slice) Remove(idx int) interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	if idx < 0 || idx >= len(s.list) {
		return nil
	}
	ret := s.list[idx]
	s.list = append(s.list[:idx], s.list[idx+1:]...)
	return ret
}

func (s *Slice) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.list)
}

func (s *Slice) Clear() {
	s.mu.Lock()
	s.list = nil
	s.mu.Unlock()
}

func (s *Slice) Range(f func(index int, value interface{}) bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for key, value := range s.list {
		if !f(key, value) {
			break
		}
	}
}
