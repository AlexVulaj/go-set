package go_set

import (
	"fmt"
)

// Set A collection that contains no duplicate elements.
type Set[T comparable] map[T]struct{}

// NewSet Creates a new set of type T with the given items if present.
func NewSet[T comparable](items ...T) Set[T] {
	s := Set[T]{}
	s.AddAll(items...)
	return s
}

// Add Adds the specified item to this set if it is not already present. Returns whether the set was modified.
func (s *Set[T]) Add(item T) bool {
	if s.Contains(item) {
		return false
	}
	(*s)[item] = struct{}{}
	return true
}

// AddAll Adds all the given items to this set if they're not already present. Returns whether the set was modified.
func (s *Set[T]) AddAll(items ...T) bool {
	changed := false
	for _, item := range items {
		if s.Add(item) {
			changed = true
		}
	}
	return changed
}

// Clear Removes all items from this set
func (s *Set[T]) Clear() {
	*s = map[T]struct{}{}
}

// Contains Returns whether this set contains the specified item
func (s *Set[T]) Contains(item T) bool {
	_, ok := (*s)[item]
	return ok
}

// ContainsAll Returns whether this set contains all the specified items or not
func (s *Set[T]) ContainsAll(items ...T) bool {
	for _, item := range items {
		if !s.Contains(item) {
			return false
		}
	}
	return true
}

// IsEmpty Returns true if the set contains no items
func (s *Set[T]) IsEmpty() bool {
	return len(*s) == 0
}

// Remove Removes the specified item from this set if it is present. Returns whether the set was modified.
func (s *Set[T]) Remove(item T) bool {
	if s.Contains(item) {
		delete(*s, item)
		return true
	}
	return false
}

// RemoveAll Removes all the specified items from this set if they are present. Returns whether the set was modified.
func (s *Set[T]) RemoveAll(items ...T) bool {
	changed := false
	for _, item := range items {
		if s.Remove(item) {
			changed = true
		}
	}
	return changed
}

// RetainAll Retains only the items in this set that are specified. Returns whether the set was modified.
func (s *Set[T]) RetainAll(items ...T) bool {
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

// Size Returns the number of items in this set
func (s *Set[T]) Size() int {
	return len(*s)
}

// ToSlice Returns a Slice containing all the items in this Set
func (s *Set[T]) ToSlice() []T {
	slice := make([]T, len(*s))
	index := 0
	for item := range *s {
		slice[index] = item
		index++
	}
	return slice
}

// ToString Returns a string representation of this Set
func (s *Set[T]) ToString() string {
	return fmt.Sprintf("%v", s.ToSlice())
}
