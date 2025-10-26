package uniqueue

import (
	"sync"
)

// Uniqueue is a thread-safe generic unique queue that enforces uniqueness
// of items. Duplicate items are automatically ignored when added.
// All operations are safe for concurrent access.
type Uniqueue[T comparable] struct {
	mu       sync.RWMutex
	uniqueue *UniqueueUnsafe[T]
}

// NewUniqueue creates and returns a new empty thread-safe unique queue.
func NewUniqueue[T comparable]() *Uniqueue[T] {
	return &Uniqueue[T]{
		uniqueue: NewUniqueueUnsafe[T](),
	}
}

// PushBack adds an item to the end of the queue if it doesn't already exist.
// If the item is already in the queue, this operation does nothing.
// Time complexity: O(1)
func (u *Uniqueue[T]) PushBack(item T) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.uniqueue.PushBack(item)
}

// PopHead removes and returns the first item from the queue.
// Returns the zero value and false if the queue is empty.
// Time complexity: O(1)
func (u *Uniqueue[T]) PopHead() (T, bool) {
	u.mu.Lock()
	defer u.mu.Unlock()

	return u.uniqueue.PopHead()
}

// Size returns the number of unique items in the queue.
// Time complexity: O(1)
func (u *Uniqueue[T]) Size() int {
	u.mu.RLock()
	defer u.mu.RUnlock()

	return u.uniqueue.Size()

}

// Contains checks if an item exists in the queue.
// Time complexity: O(1) due to hash map lookup
func (u *Uniqueue[T]) Contains(item T) bool {
	u.mu.RLock()
	defer u.mu.RUnlock()

	return u.uniqueue.Contains(item)
}

// IsEmpty returns true if the queue is empty, false otherwise.
// Time complexity: O(1)
func (u *Uniqueue[T]) IsEmpty() bool {
	return u.uniqueue.IsEmpty()
}
