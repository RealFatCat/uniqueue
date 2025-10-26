# uniqueue

A Go package providing thread-safe and non-thread-safe implementations of a unique queue data structure.

## Features

- **Generic implementation** - Works with any comparable type
- **Thread-safe variant** - Thread-safe `Uniqueue` for concurrent access
- **Efficient** - O(1) push and pop operations
- **Uniqueness enforcement** - Automatically prevents duplicate items
- **FIFO ordering** - Items are retrieved in first-in-first-out order

## Installation

```bash
go get github.com/realfatcat/uniqueue
```

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/realfatcat/uniqueue"
)

func main() {
    // Create a new thread-safe unique queue
    q := uniqueue.NewUniqueue[string]()
    
    // Add items
    q.PushBack("item1")
    q.PushBack("item2")
    q.PushBack("item1") // Won't be added (duplicate)
    
    fmt.Println(q.Size()) // Output: 2
    
    // Retrieve items
    if val, ok := q.PopHead(); ok {
        fmt.Println(val) // Output: "item1"
    }
}
```

### Thread-Safe Uniqueue

Use `Uniqueue` when you need thread-safe operations:

```go
q := uniqueue.NewUniqueue[int]()

// Concurrent pushes
go func() {
    q.PushBack(1)
    q.PushBack(2)
}()

go func() {
    q.PushBack(3)
}()

// Thread-safe reads
if q.Contains(2) {
    fmt.Println("Found 2")
}
```

### Unsafe Uniqueue (Single Goroutine Only)

For single-goroutine scenarios, use `UniqueueUnsafe`:

```go
q := uniqueue.NewUniqueueUnsafe[string]()

q.PushBack("a")
q.PushBack("b")
q.PushBack("a") // Ignored

fmt.Println(q.Size()) // 2
```

### Low-Level Queue

For a basic queue without uniqueness constraints:

```go
q := uniqueue.NewQueue[int]()
q.PushBack(1)
q.PushBack(2)

val, ok := q.PopHead()
if ok {
    fmt.Println(val) // 1
}
```

## API Reference

### Uniqueue (Thread-Safe)

- `NewUniqueue[T comparable]() *Uniqueue[T]` - Creates a new thread-safe unique queue
- `PushBack(item T)` - Adds an item to the queue (ignores duplicates)
- `PopHead() (T, bool)` - Removes and returns the first item
- `Contains(item T) bool` - Checks if an item exists in the queue
- `Size() int` - Returns the number of items

### UniqueueUnsafe

- `NewUniqueueUnsafe[T comparable]() *UniqueueUnsafe[T]` - Creates a new unique queue (not thread-safe)
- Same methods as `Uniqueue`

### Queue (Basic)

- `NewQueue[T comparable]() *Queue[T]` - Creates a new queue
- `PushBack(item T)` - Adds an item to the end
- `PopHead() (T, bool)` - Removes and returns the first item
- `Contains(item T) bool` - Checks if an item exists
- `Size() int` - Returns the number of items

## Performance

- Push: O(1)
- Pop: O(1)
- Contains: O(n) where n is queue length
- Space: O(n)

## License

MIT
