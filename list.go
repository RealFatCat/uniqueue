// Package uniqueue provides thread-safe and non-thread-safe implementations
// of a unique queue data structure with generic type support.
package uniqueue

// Queue is a generic doubly-linked list queue that supports FIFO operations.
// It allows duplicate items and provides O(1) push/pop operations.
type Queue[T comparable] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

// NewQueue creates and returns a new empty queue.
func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{}
}

type node[T comparable] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

// PushBack adds an item to the end of the queue.
// Time complexity: O(1)
func (q *Queue[T]) PushBack(item T) {
	node := &node[T]{value: item}
	if q.head == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		node.prev = q.tail
		q.tail = node
	}
	q.length++
}

// PopHead removes and returns the first item from the queue.
// Returns the zero value and false if the queue is empty.
// Time complexity: O(1)
func (q *Queue[T]) PopHead() (T, bool) {
	var result T
	if q.head == nil {
		return result, false
	}

	node := q.head
	q.head = node.next
	if q.head != nil {
		q.head.prev = nil
	} else {
		q.tail = nil
	}
	q.length--
	return node.value, true
}

// Contains checks if an item exists in the queue.
// Time complexity: O(n) where n is the queue length
func (q *Queue[T]) Contains(item T) bool {
	for node := q.head; node != nil; node = node.next {
		if node.value == item {
			return true
		}
	}
	return false
}

// Size returns the number of items in the queue.
// Time complexity: O(1)
func (q *Queue[T]) Size() int {
	return q.length
}
