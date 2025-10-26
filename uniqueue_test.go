package uniqueue

import (
	"sync"
	"testing"
)

func TestNewUniqueue(t *testing.T) {
	u := NewUniqueue[int]()
	if u == nil {
		t.Fatal("NewUniqueue returned nil")
	}
	if u.Size() != 0 {
		t.Errorf("Expected size 0, got %d", u.Size())
	}
}

func TestUniqueue_PushBack(t *testing.T) {
	u := NewUniqueue[string]()

	u.PushBack("first")
	if u.Size() != 1 {
		t.Errorf("Expected size 1, got %d", u.Size())
	}

	u.PushBack("second")
	if u.Size() != 2 {
		t.Errorf("Expected size 2, got %d", u.Size())
	}
}

func TestUniqueue_PopHead(t *testing.T) {
	t.Run("empty uniqueue", func(t *testing.T) {
		u := NewUniqueue[int]()
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
		u := NewUniqueue[string]()
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
		u := NewUniqueue[int]()
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

func TestUniqueue_Contains(t *testing.T) {
	t.Run("empty uniqueue", func(t *testing.T) {
		u := NewUniqueue[int]()
		if u.Contains(42) {
			t.Error("Expected Contains to return false for empty uniqueue")
		}
	})

	t.Run("single element", func(t *testing.T) {
		u := NewUniqueue[string]()
		u.PushBack("test")
		if !u.Contains("test") {
			t.Error("Expected Contains to return true")
		}
		if u.Contains("other") {
			t.Error("Expected Contains to return false")
		}
	})

	t.Run("multiple elements", func(t *testing.T) {
		u := NewUniqueue[int]()
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
}

func TestUniqueue_Size(t *testing.T) {
	u := NewUniqueue[int]()

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

func TestUniqueue_Concurrency(t *testing.T) {
	u := NewUniqueue[int]()
	const numGoroutines = 10
	const itemsPerGoroutine = 100

	var wg sync.WaitGroup

	// Push items concurrently
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			for j := 0; j < itemsPerGoroutine; j++ {
				u.PushBack(start*itemsPerGoroutine + j)
			}
		}(i)
	}

	wg.Wait()
	expectedSize := numGoroutines * itemsPerGoroutine
	if u.Size() != expectedSize {
		t.Errorf("Expected size %d, got %d", expectedSize, u.Size())
	}

	// Pop items concurrently
	popped := make(chan int, expectedSize)
	wg = sync.WaitGroup{}

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				val, ok := u.PopHead()
				if !ok {
					break
				}
				popped <- val
			}
		}()
	}

	wg.Wait()
	close(popped)

	count := 0
	for range popped {
		count++
	}

	if count != expectedSize {
		t.Errorf("Expected %d items popped, got %d", expectedSize, count)
	}

	if u.Size() != 0 {
		t.Errorf("Expected size 0 after popping, got %d", u.Size())
	}
}

func TestUniqueue_IsEmpty(t *testing.T) {
	u := NewUniqueue[int]()

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

func TestUniqueue_GenericTypes(t *testing.T) {
	t.Run("string type", func(t *testing.T) {
		u := NewUniqueue[string]()
		u.PushBack("hello")
		u.PushBack("world")
		if !u.Contains("hello") {
			t.Error("Expected Contains to work with strings")
		}
		val, _ := u.PopHead()
		if val != "hello" {
			t.Errorf("Expected 'hello', got %s", val)
		}
	})

	t.Run("int type", func(t *testing.T) {
		u := NewUniqueue[int]()
		u.PushBack(42)
		u.PushBack(100)
		if !u.Contains(42) {
			t.Error("Expected Contains to work with ints")
		}
		val, _ := u.PopHead()
		if val != 42 {
			t.Errorf("Expected 42, got %d", val)
		}
	})
}
