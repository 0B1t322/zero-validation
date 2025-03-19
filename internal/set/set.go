package set

import "maps"

type Set[T comparable] interface {
	Add(value T) bool
	Contains(value T) bool
	Values() []T
	AddMany(value ...T)
	Clean()
}

type simpleSet[T comparable] struct {
	values    []T
	valuesMap map[T]struct{}
}

func New[T comparable]() Set[T] {
	return &simpleSet[T]{
		values:    make([]T, 0),
		valuesMap: make(map[T]struct{}),
	}
}

func (s *simpleSet[T]) Clean() {
	s.values = s.values[:0]
	maps.DeleteFunc(s.valuesMap, func(_ T, _ struct{}) bool {
		return true
	})
}

func (s *simpleSet[T]) Add(value T) bool {
	if _, isFind := s.valuesMap[value]; isFind {
		return false
	}

	s.valuesMap[value] = struct{}{}
	s.values = append(s.values, value)
	return true
}

func (s *simpleSet[T]) Contains(value T) bool {
	_, isFind := s.valuesMap[value]
	return isFind
}

func (s *simpleSet[T]) Values() []T {
	return s.values
}

func (s *simpleSet[T]) AddMany(value ...T) {
	for _, v := range value {
		s.Add(v)
	}
}
