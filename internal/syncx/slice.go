package syncx

import (
	"iter"
	"sync"
)

type Slice[T any] struct {
	_ noCopy

	s  []T
	mu sync.RWMutex
}

func (s *Slice[T]) Append(v T) {
	s.mu.Lock()
	s.s = append(s.s, v)
	s.mu.Unlock()
}

func (s *Slice[T]) Elements() iter.Seq2[int, T] {
	s.mu.RLock()
	return func(yield func(int, T) bool) {
		defer s.mu.RUnlock()

		for idx, x := range s.s {
			if !yield(idx, x) {
				return
			}
		}
	}
}
