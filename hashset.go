package set

import (
	"fmt"
	"strings"
)

// A hash set implementation, using a valueless map.
// Yes, it is just a glorified wrapper.
type HashSet[T comparable] struct {
	elements map[T]struct{}
}

// NewHashSet creates a new instance of a hash set for type T.
func NewHashSet[T comparable]() *HashSet[T] {
	return &HashSet[T]{elements: make(map[T]struct{})}
}

// Add inserts an element of type T to the set.
func (set *HashSet[T]) Add(newElements ...T) bool {
	if len(newElements) < 1 {
		return false
	}
	//set.elements[element] = struct{}{}
	for n := 0; n < len(newElements); n++ {
		set.elements[newElements[n]] = struct{}{}
	}
	return true
}

// Remove deletes an element of type T from the set.
func (set *HashSet[T]) Remove(element T) {
	delete(set.elements, element)
}

// Contains checks if an element of type T is in the set.
func (set *HashSet[T]) Has(element T) bool {
	_, exists := set.elements[element]
	return exists
}

// Size returns the number of elements in the set.
func (set *HashSet[T]) Size() int {
	return len(set.elements)
}

// so it can implement the stringer interface for println
func (set *HashSet[T]) String() string {
	var sb strings.Builder
	sb.WriteString("{")

	first := true
	for element := range set.elements {
		if !first {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%v", element))
		first = false
	}

	sb.WriteString("}")
	return sb.String()
}

// Performs a copy.
func (set *HashSet[T]) Clone() *HashSet[T] {
	keys := make([]T, 0, len(set.elements))
	for k := range set.elements {
		keys = append(keys, k)
	}
	var newSet *HashSet[T] = NewHashSet[T]()
	newSet.Add(keys...)
	return newSet
}
