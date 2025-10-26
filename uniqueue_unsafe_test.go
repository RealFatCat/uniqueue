package uniqueue

import (
	"testing"
)

func TestNewUniqueueUnsafe(t *testing.T) {
	u := NewUniqueueUnsafe[int]()
	if u == nil {
		t.Fatal("NewUniqueueUnsafe returned nil")
	}
	if u.Size() != 0 {
		t.Errorf("Expected size 0, got %d", u.Size())
	}
}

func TestUniqueueUnsafe_PushBack(t *testing.T) {
	u := NewUniqueueUnsafe[string]()

	u.PushBack("first")
	if u.Size() != 1 {
		t.Errorf("Expected size 1, got %d", u.Size())
	}

	u.PushBack("second")
	if u.Size() != 2 {
		t.Errorf("Expected size 2, got %d", u.Size())
	}
}

func TestUniqueueUnsafe_PushBack_Uniqueness(t *testing.T) {
	u := NewUniqueueUnsafe[int]()

	// Push same item twice
	u.PushBack(42)
	u.PushBack(42)

	if u.Size() != 1 {
		t.Errorf("Expected size 1 for duplicate items, got %d", u.Size())
	}
}

func TestUniqueueUnsafe_PushBack_MultipleDuplicates(t *testing.T) {
	u := NewUniqueueUnsafe[string]()

	// Push same item multiple times
	u.PushBack("duplicate")
	u.PushBack("duplicate")
	u.PushBack("duplicate")
	u.PushBack("duplicate")

	if u.Size() != 1 {
		t.Errorf("Expected size 1 for multiple duplicates, got %d", u.Size())
	}
}

func TestUniqueueUnsafe_PopHead(t *testing.T) {
	t.Run("empty uniqueue", func(t *testing.T) {
		u := NewUniqueueUnsafe[int]()
		val, ok := u.PopHead()
		if ok {
			t.Error("Expected ok=false for empty uniqueue")
		}
		var zero int
		if val != zero {
			t.Errorf("Expected zero value, got %v", val)
		}
		if u.Size() != 0 {
			t.Errorf("Expected size 0, got %d", u.Size())
		}
	})

	t.Run("single element", func(t *testing.T) {
		u := NewUniqueueUnsafe[string]()
		u.PushBack("only")
		val, ok := u.PopHead()
		if !ok {
			t.Error("Expected ok=true")
		}
		if val != "only" {
			t.Errorf("Expected 'only', got %s", val)
		}
		if u.Size() != 0 {
			t.Errorf("Expected size 0, got %d", u.Size())
		}
	})

	t.Run("multiple elements", func(t *testing.T) {
		u := NewUniqueueUnsafe[int]()
		u.PushBack(1)
		u.PushBack(2)
		u.PushBack(3)

		val1, ok1 := u.PopHead()
		if !ok1 || val1 != 1 {
			t.Errorf("Expected (1, true), got (%d, %v)", val1, ok1)
		}
		if u.Size() != 2 {
			t.Errorf("Expected size 2, got %d", u.Size())
		}

		val2, ok2 := u.PopHead()
		if !ok2 || val2 != 2 {
			t.Errorf("Expected (2, true), got (%d, %v)", val2, ok2)
		}
		if u.Size() != 1 {
			t.Errorf("Expected size 1, got %d", u.Size())
		}

		val3, ok3 := u.PopHead()
		if !ok3 || val3 != 3 {
			t.Errorf("Expected (3, true), got (%d, %v)", val3, ok3)
		}
		if u.Size() != 0 {
			t.Errorf("Expected size 0, got %d", u.Size())
		}
	})
}

func TestUniqueueUnsafe_PopHead_CanReAdd(t *testing.T) {
	u := NewUniqueueUnsafe[string]()

	// Add and remove
	u.PushBack("item")
	val, ok := u.PopHead()
	if !ok || val != "item" {
		t.Errorf("Expected ('item', true), got (%s, %v)", val, ok)
	}

	// Can add same item again after removing
	u.PushBack("item")
	if u.Size() != 1 {
		t.Errorf("Expected size 1 after re-adding, got %d", u.Size())
	}
	if !u.Contains("item") {
		t.Error("Expected Contains to return true for re-added item")
	}
}

func TestUniqueueUnsafe_Contains(t *testing.T) {
	t.Run("empty uniqueue", func(t *testing.T) {
		u := NewUniqueueUnsafe[int]()
		if u.Contains(42) {
			t.Error("Expected Contains to return false for empty uniqueue")
		}
	})

	t.Run("single element", func(t *testing.T) {
		u := NewUniqueueUnsafe[string]()
		u.PushBack("test")
		if !u.Contains("test") {
			t.Error("Expected Contains to return true")
		}
		if u.Contains("other") {
			t.Error("Expected Contains to return false")
		}
	})

	t.Run("multiple elements", func(t *testing.T) {
		u := NewUniqueueUnsafe[int]()
		u.PushBack(1)
		u.PushBack(2)
		u.PushBack(3)

		if !u.Contains(1) {
			t.Error("Expected Contains(1) to return true")
		}
		if !u.Contains(2) {
			t.Error("Expected Contains(2) to return true")
		}
		if !u.Contains(3) {
			t.Error("Expected Contains(3) to return true")
		}
		if u.Contains(99) {
			t.Error("Expected Contains(99) to return false")
		}
	})

	t.Run("after removal", func(t *testing.T) {
		u := NewUniqueueUnsafe[string]()
		u.PushBack("keep")
		u.PushBack("remove")

		if !u.Contains("keep") {
			t.Error("Expected Contains to return true before removal")
		}

		u.PopHead() // removes "keep"

		if u.Contains("keep") {
			t.Error("Expected Contains to return false after removal")
		}
		if !u.Contains("remove") {
			t.Error("Expected Contains to return true for remaining element")
		}
	})
}

func TestUniqueueUnsafe_Size(t *testing.T) {
	u := NewUniqueueUnsafe[int]()

	if u.Size() != 0 {
		t.Errorf("Expected size 0, got %d", u.Size())
	}

	u.PushBack(1)
	if u.Size() != 1 {
		t.Errorf("Expected size 1, got %d", u.Size())
	}

	u.PushBack(2)
	u.PushBack(3)
	if u.Size() != 3 {
		t.Errorf("Expected size 3, got %d", u.Size())
	}

	u.PopHead()
	if u.Size() != 2 {
		t.Errorf("Expected size 2, got %d", u.Size())
	}

	u.PopHead()
	u.PopHead()
	if u.Size() != 0 {
		t.Errorf("Expected size 0, got %d", u.Size())
	}
}

func TestUniqueueUnsafe_IsEmpty(t *testing.T) {
	u := NewUniqueueUnsafe[int]()

	if !u.IsEmpty() {
		t.Error("Expected IsEmpty to return true for empty queue")
	}

	u.PushBack(1)
	if u.IsEmpty() {
		t.Error("Expected IsEmpty to return false for non-empty queue")
	}

	u.PopHead()
	if !u.IsEmpty() {
		t.Error("Expected IsEmpty to return true after popping all items")
	}
}

func TestUniqueueUnsafe_UniquenessConstraint(t *testing.T) {
	u := NewUniqueueUnsafe[int]()

	// Add unique items
	u.PushBack(1)
	u.PushBack(2)
	u.PushBack(3)

	if u.Size() != 3 {
		t.Errorf("Expected size 3, got %d", u.Size())
	}

	// Try to add duplicates
	u.PushBack(1)
	u.PushBack(2)
	u.PushBack(3)

	if u.Size() != 3 {
		t.Errorf("Expected size 3 after duplicates, got %d", u.Size())
	}

	// Mix of unique and duplicate
	u.PushBack(4)
	u.PushBack(1) // duplicate
	u.PushBack(5)
	u.PushBack(2) // duplicate

	if u.Size() != 5 {
		t.Errorf("Expected size 5, got %d", u.Size())
	}
}

func TestUniqueueUnsafe_FIFOOrder(t *testing.T) {
	u := NewUniqueueUnsafe[string]()
	items := []string{"first", "second", "third"}

	for _, item := range items {
		u.PushBack(item)
	}

	// Try to add duplicates - should be ignored
	u.PushBack("first")
	u.PushBack("second")

	// Pop in FIFO order
	val1, _ := u.PopHead()
	if val1 != "first" {
		t.Errorf("Expected 'first', got %s", val1)
	}

	val2, _ := u.PopHead()
	if val2 != "second" {
		t.Errorf("Expected 'second', got %s", val2)
	}

	val3, _ := u.PopHead()
	if val3 != "third" {
		t.Errorf("Expected 'third', got %s", val3)
	}
}

func TestUniqueueUnsafe_GenericTypes(t *testing.T) {
	t.Run("string type", func(t *testing.T) {
		u := NewUniqueueUnsafe[string]()
		u.PushBack("hello")
		u.PushBack("world")
		u.PushBack("hello") // duplicate

		if u.Size() != 2 {
			t.Errorf("Expected size 2, got %d", u.Size())
		}
		if !u.Contains("hello") {
			t.Error("Expected Contains to work with strings")
		}
		val, _ := u.PopHead()
		if val != "hello" {
			t.Errorf("Expected 'hello', got %s", val)
		}
	})

	t.Run("int type", func(t *testing.T) {
		u := NewUniqueueUnsafe[int]()
		u.PushBack(42)
		u.PushBack(100)
		u.PushBack(42) // duplicate

		if u.Size() != 2 {
			t.Errorf("Expected size 2, got %d", u.Size())
		}
		if !u.Contains(42) {
			t.Error("Expected Contains to work with ints")
		}
		val, _ := u.PopHead()
		if val != 42 {
			t.Errorf("Expected 42, got %d", val)
		}
	})

	t.Run("bool type", func(t *testing.T) {
		u := NewUniqueueUnsafe[bool]()
		u.PushBack(true)
		u.PushBack(false)
		u.PushBack(true) // duplicate

		if u.Size() != 2 {
			t.Errorf("Expected size 2, got %d", u.Size())
		}
		if !u.Contains(true) {
			t.Error("Expected Contains to work with bools")
		}
		val, _ := u.PopHead()
		if val != true {
			t.Errorf("Expected true, got %v", val)
		}
	})
}

func TestUniqueueUnsafe_ComplexWorkflow(t *testing.T) {
	u := NewUniqueueUnsafe[int]()

	// Add some items
	u.PushBack(1)
	u.PushBack(2)
	u.PushBack(3)

	// Try to add duplicates (should be ignored)
	u.PushBack(1)
	u.PushBack(2)

	// Add new unique items
	u.PushBack(4)
	u.PushBack(5)

	if u.Size() != 5 {
		t.Errorf("Expected size 5, got %d", u.Size())
	}

	// Remove items
	val1, _ := u.PopHead() // 1
	if val1 != 1 {
		t.Errorf("Expected 1, got %d", val1)
	}

	val2, _ := u.PopHead() // 2
	if val2 != 2 {
		t.Errorf("Expected 2, got %d", val2)
	}

	if u.Size() != 3 {
		t.Errorf("Expected size 3, got %d", u.Size())
	}

	// After removing, we can add them again
	u.PushBack(1)
	u.PushBack(2)

	if u.Size() != 5 {
		t.Errorf("Expected size 5 after re-adding, got %d", u.Size())
	}
}
