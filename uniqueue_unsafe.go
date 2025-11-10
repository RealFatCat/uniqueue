package uniqueue

// UniqueueUnsafe is a non-thread-safe generic unique queue that enforces
// uniqueness of items. Duplicate items are automatically ignored when added.
// This type should only be used from a single goroutine.
type UniqueueUnsafe[T comparable] struct {
	queue *Queue[T]
	seen  map[T]struct{}
}

// NewUniqueueUnsafe creates and returns a new empty unique queue.
// This type is not thread-safe and should only be used from one goroutine.
func NewUniqueueUnsafe[T comparable]() *UniqueueUnsafe[T] {
	return &UniqueueUnsafe[T]{
		queue: NewQueue[T](),
		seen:  make(map[T]struct{}),
	}
}

// PushBack adds an item to the end of the queue if it doesn't already exist.
// If the item is already in the queue, this operation does nothing.
// Time complexity: O(1)
func (u *UniqueueUnsafe[T]) PushBack(item T) {
	if _, ok := u.seen[item]; ok {
		return
	}
	u.seen[item] = struct{}{}
	u.queue.PushBack(item)
}

// PopHead removes and returns the first item from the queue.
// Returns the zero value and false if the queue is empty.
// Time complexity: O(1)
func (u *UniqueueUnsafe[T]) PopHead() (T, bool) {
	item, ok := u.queue.PopHead()
	if ok {
		delete(u.seen, item)
	}
	return item, ok
}

// Size returns the number of unique items in the queue.
// Time complexity: O(1)
func (u *UniqueueUnsafe[T]) Size() int {
	return u.queue.Size()
}

// Contains checks if an item exists in the queue.
// Time complexity: O(1) due to hash map lookup
func (u *UniqueueUnsafe[T]) Contains(item T) bool {
	_, ok := u.seen[item]
	return ok
}

// IsEmpty returns true if the queue is empty, false otherwise.
// Time complexity: O(1)
func (u *UniqueueUnsafe[T]) IsEmpty() bool {
	return u.Size() == 0
}
