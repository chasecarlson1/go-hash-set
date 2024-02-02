package set

// A hash set implementation, using a valueless map.
// Yes, it is just a glorified wrapper. At least it has no external dependencies.
// It can contain any value satisfying the comparable interface.
//
// example...
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// set := NewHashSet[int]()
// set.Add(2, 3, 1, 9, 11, 9213209)
// fmt.Println(set)			- prints out the above values in the set
// set.Remove(3, 99)		- 99 does nothing, but does not error
// should now contain 2, 1, 9, 11, 9213209.
// Clone() returns a pointer to a new deep copy if you want to edit one and its values

import (
	"fmt"
	"strings"
)

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

// Remove deletes an element(s) of type T from the set.
func (set *HashSet[T]) Remove(element T) {
	delete(set.elements, element)
}

// Contains checks if an element of type T is in the set.
// does not work with compound types yet.
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
