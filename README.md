# go-hash-set
## A hash set implementation in Golang.

Yes, it is just a glorified wrapper for map[T]U. At least it has no external dependencies.
It can contain any value satisfying the comparable interface.

## functions:
### NewHashSet[T comparable]() *HashSet[T]
### (set *HashSet[T]) Add(newElements ...T) bool
### (set *HashSet[T]) Add(newElements ...T) bool
### (set *HashSet[T]) Remove(element T)
### (set *HashSet[T]) Has(element T) bool
### (set *HashSet[T]) Size() int
### (set *HashSet[T]) String() string
### (set *HashSet[T]) Clone() *HashSet[T]

## example...
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
set := NewHashSet[int]()
set.Add(2, 3, 1, 9, 11, 9213209)
fmt.Println(set)		// prints out the above values in the set
set.Remove(3, 99)		// 99 does nothing, but does not error
                        //should now contain 2, 1, 9, 11, 9213209.
// Clone() returns a pointer to a new deep copy if you want to edit one and its values
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
*Note: Has() would not work easily with structs and compound types unless you pass the exact same reference to it.*