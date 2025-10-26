package uniqueue

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue[int]()
	if q == nil {
		t.Fatal("NewQueue returned nil")
	}
	if q.Size() != 0 {
		t.Errorf("Expected length 0, got %d", q.Size())
	}
}

func TestQueue_PushBack(t *testing.T) {
	q := NewQueue[string]()

	q.PushBack("first")
	if q.Size() != 1 {
		t.Errorf("Expected size 1, got %d", q.Size())
	}

	q.PushBack("second")
	if q.Size() != 2 {
		t.Errorf("Expected size 2, got %d", q.Size())
	}

	q.PushBack("third")
	if q.Size() != 3 {
		t.Errorf("Expected size 3, got %d", q.Size())
	}
}

func TestQueue_PopHead(t *testing.T) {
	t.Run("empty queue", func(t *testing.T) {
		q := NewQueue[int]()
		val, ok := q.PopHead()
		if ok {
			t.Error("Expected ok=false for empty queue")
		}
		var zero int
		if val != zero {
			t.Errorf("Expected zero value, got %v", val)
		}
		if q.Size() != 0 {
			t.Errorf("Expected size 0, got %d", q.Size())
		}
	})

	t.Run("single element", func(t *testing.T) {
		q := NewQueue[string]()
		q.PushBack("only")
		val, ok := q.PopHead()
		if !ok {
			t.Error("Expected ok=true")
		}
		if val != "only" {
			t.Errorf("Expected 'only', got %s", val)
		}
		if q.Size() != 0 {
			t.Errorf("Expected size 0, got %d", q.Size())
		}
	})

	t.Run("multiple elements", func(t *testing.T) {
		q := NewQueue[int]()
		q.PushBack(1)
		q.PushBack(2)
		q.PushBack(3)

		val1, ok1 := q.PopHead()
		if !ok1 || val1 != 1 {
			t.Errorf("Expected (1, true), got (%d, %v)", val1, ok1)
		}
		if q.Size() != 2 {
			t.Errorf("Expected size 2, got %d", q.Size())
		}

		val2, ok2 := q.PopHead()
		if !ok2 || val2 != 2 {
			t.Errorf("Expected (2, true), got (%d, %v)", val2, ok2)
		}
		if q.Size() != 1 {
			t.Errorf("Expected size 1, got %d", q.Size())
		}

		val3, ok3 := q.PopHead()
		if !ok3 || val3 != 3 {
			t.Errorf("Expected (3, true), got (%d, %v)", val3, ok3)
		}
		if q.Size() != 0 {
			t.Errorf("Expected size 0, got %d", q.Size())
		}

		// Pop from empty queue
		_, ok4 := q.PopHead()
		if ok4 {
			t.Error("Expected ok=false for empty queue")
		}
	})

	t.Run("FIFO order", func(t *testing.T) {
		q := NewQueue[string]()
		items := []string{"first", "second", "third", "fourth"}
		for _, item := range items {
			q.PushBack(item)
		}

		for _, expected := range items {
			val, ok := q.PopHead()
			if !ok {
				t.Errorf("Expected ok=true for %s", expected)
			}
			if val != expected {
				t.Errorf("Expected %s, got %s", expected, val)
			}
		}
	})
}

func TestQueue_Contains(t *testing.T) {
	t.Run("empty queue", func(t *testing.T) {
		q := NewQueue[int]()
		if q.Contains(42) {
			t.Error("Expected Contains to return false for empty queue")
		}
	})

	t.Run("single element", func(t *testing.T) {
		q := NewQueue[string]()
		q.PushBack("test")
		if !q.Contains("test") {
			t.Error("Expected Contains to return true")
		}
		if q.Contains("other") {
			t.Error("Expected Contains to return false")
		}
	})

	t.Run("multiple elements", func(t *testing.T) {
		q := NewQueue[int]()
		q.PushBack(1)
		q.PushBack(2)
		q.PushBack(3)

		if !q.Contains(1) {
			t.Error("Expected Contains(1) to return true")
		}
		if !q.Contains(2) {
			t.Error("Expected Contains(2) to return true")
		}
		if !q.Contains(3) {
			t.Error("Expected Contains(3) to return true")
		}
		if q.Contains(99) {
			t.Error("Expected Contains(99) to return false")
		}
	})

	t.Run("after removal", func(t *testing.T) {
		q := NewQueue[string]()
		q.PushBack("keep")
		q.PushBack("remove")
		q.PushBack("also_keep")

		if !q.Contains("keep") {
			t.Error("Expected Contains to return true before removal")
		}

		q.PopHead() // removes "keep"

		if q.Contains("keep") {
			t.Error("Expected Contains to return false after removal")
		}
		if !q.Contains("remove") {
			t.Error("Expected Contains to return true for remaining element")
		}
		if !q.Contains("also_keep") {
			t.Error("Expected Contains to return true for remaining element")
		}
	})
}

func TestQueue_Size(t *testing.T) {
	q := NewQueue[int]()

	if q.Size() != 0 {
		t.Errorf("Expected size 0, got %d", q.Size())
	}

	q.PushBack(1)
	if q.Size() != 1 {
		t.Errorf("Expected size 1, got %d", q.Size())
	}

	q.PushBack(2)
	q.PushBack(3)
	if q.Size() != 3 {
		t.Errorf("Expected size 3, got %d", q.Size())
	}

	q.PopHead()
	if q.Size() != 2 {
		t.Errorf("Expected size 2, got %d", q.Size())
	}

	q.PopHead()
	q.PopHead()
	if q.Size() != 0 {
		t.Errorf("Expected size 0, got %d", q.Size())
	}
}

func TestQueue_GenericTypes(t *testing.T) {
	t.Run("string type", func(t *testing.T) {
		q := NewQueue[string]()
		q.PushBack("hello")
		q.PushBack("world")
		if !q.Contains("hello") {
			t.Error("Expected Contains to work with strings")
		}
		val, _ := q.PopHead()
		if val != "hello" {
			t.Errorf("Expected 'hello', got %s", val)
		}
	})

	t.Run("int type", func(t *testing.T) {
		q := NewQueue[int]()
		q.PushBack(42)
		q.PushBack(100)
		if !q.Contains(42) {
			t.Error("Expected Contains to work with ints")
		}
		val, _ := q.PopHead()
		if val != 42 {
			t.Errorf("Expected 42, got %d", val)
		}
	})

	t.Run("bool type", func(t *testing.T) {
		q := NewQueue[bool]()
		q.PushBack(true)
		q.PushBack(false)
		if !q.Contains(true) {
			t.Error("Expected Contains to work with bools")
		}
		val, _ := q.PopHead()
		if val != true {
			t.Errorf("Expected true, got %v", val)
		}
	})
}
