package go_set

import (
	"fmt"
)

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](items ...T) Set[T] {
	s := Set[T]{}
	s.AddAll(items...)
	return s
}

func (s *Set[T]) Add(item T) bool {
	if s.Contains(item) {
		return false
	}
	(*s)[item] = struct{}{}
	return true
}

func (s *Set[T]) AddAll(items ...T) bool {
	changed := false
	for _, item := range items {
		if s.Add(item) {
			changed = true
		}
	}
	return changed
}

func (s *Set[T]) Clear() {
	*s = map[T]struct{}{}
}

func (s *Set[T]) Contains(item T) bool {
	_, ok := (*s)[item]
	return ok
}

func (s *Set[T]) ContainsAll(items ...T) bool {
	for _, item := range items {
		if !s.Contains(item) {
			return false
		}
	}
	return true
}

func (s *Set[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Set[T]) Remove(item T) bool {
	if s.Contains(item) {
		delete(*s, item)
		return true
	}
	return false
}

func (s *Set[T]) RemoveAll(items ...T) bool {
	changed := false
	for _, item := range items {
		if s.Remove(item) {
			changed = true
		}
	}
	return changed
}

func (s *Set[T]) retainAll(items ...T) bool {
	retainSet := Set[T]{}
	retainSet.AddAll(items...)

	changed := false
	for item := range *s {
		if !retainSet.Contains(item) {
			_ = s.Remove(item)
			changed = true
		}
	}
	return changed
}

func (s *Set[T]) Size() int {
	return len(*s)
}

func (s *Set[T]) ToSlice() []T {
	slice := make([]T, len(*s))
	index := 0
	for item := range *s {
		slice[index] = item
		index++
	}
	return slice
}

func (s *Set[T]) ToString() string {
	return fmt.Sprintf("%v", s.ToSlice())
}
