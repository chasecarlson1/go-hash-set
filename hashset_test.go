package set

import (
	"testing"
)

func TestReferenceSet(t *testing.T) {
	type PtrStruct struct {
		Name string
	}

	set := NewHashSet[*PtrStruct]()
	set.Add(
		&PtrStruct{Name: "first"},
		&PtrStruct{Name: "second"},
		&PtrStruct{Name: "third"},
		&PtrStruct{Name: "fourth"},
	)

	t.Run("test copying and editing one", func(t *testing.T) {
		copy := set.Clone()
		copy.Add(&PtrStruct{Name: "newone"})
		t.Log("original =", set)
		t.Log("copy 	=", copy)
	})
}

func TestClone(t *testing.T) {
	set := NewHashSet[float32]()
	set.Add(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	t.Run("testing cloning to make a copy", func(t *testing.T) {
		var clone *HashSet[float32] = set.Clone()
		clone.Add(11)
		clone.Remove(2)
		t.Log("original set is", set)
		t.Log("cloned and edited set is", clone)
		if isDiff := clone != set; isDiff {
			t.Log("Correctly found to be different!")
		} else {
			t.Error("Found them to be the same! mission failed.")
		}
	})
}

func TestHas(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(9, 21, 32, 9191, 31, 99)
	t.Run("when the requested key does exist", func(t *testing.T) {
		if exists := set.Has(21); exists {
			t.Log("Has(21) found existing!")
		} else {
			t.Error("failed to find existing value of 21 on set", set)
		}
	})
	t.Run("when the requested key does NOT exist", func(t *testing.T) {
		if exists := set.Has(4); !exists {
			t.Log("Has(4) was not found (correctly)!")
		} else {
			t.Error("failed to find existing value of 4 on set", set)
		}
	})
}
