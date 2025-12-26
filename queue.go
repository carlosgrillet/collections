package collections

import "fmt"

// Queue is a data structure used for enqueueing elements.
// The queue follows the FIFO (First-In-First-Out) method.
type Queue[T any] struct {
	elements []T
	len      int
	capacity int // 0 means unbounded
}

// NewQueue creates and returns a new empty queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// NewBoundedQueue creates a new queue with a maximum capacity.
func NewBoundedQueue[T any](capacity int) *Queue[T] {
	return &Queue[T]{capacity: capacity}
}

// Enqueue adds a new element to the queue.
// Returns false if the queue is at capacity (for bounded queues).
func (q *Queue[T]) Enqueue(element T) bool {
	if q.capacity > 0 && q.len >= q.capacity {
		return false
	}
	q.elements = append(q.elements, element)
	q.len += 1
	return true
}

// Next returns and removes the first element from the queue.
// Returns false if there are no elements in the queue.
func (q *Queue[T]) Next() (T, bool) {
	var zero T
	if q.len == 0 {
		return zero, false
	}
	next := q.elements[0]
	q.elements = q.elements[1:]
	q.len -= 1
	return next, true
}

// Peek returns the next element in the queue without removing it.
// Returns false if there is no next element.
func (q *Queue[T]) Peek() (T, bool) {
	var zero T
	if q.len == 0 {
		return zero, false
	}
	return q.elements[0], true
}

// Clear removes all elements from the queue.
func (q *Queue[T]) Clear() {
	q.elements = []T{}
	q.len = 0
}

// Len returns the current length of the queue.
func (q *Queue[T]) Len() int {
	return q.len
}

// IsEmpty returns true if the queue has no elements.
func (q *Queue[T]) IsEmpty() bool {
	return q.len == 0
}

// IsFull returns true if the queue is at capacity (for bounded queues).
func (q *Queue[T]) IsFull() bool {
	return q.capacity > 0 && q.len >= q.capacity
}

// Contains checks if an element exists in the queue.
func (q *Queue[T]) Contains(element T) bool {
	for _, e := range q.elements {
		if any(e) == any(element) {
			return true
		}
	}
	return false
}

// ToSlice returns a copy of all elements as a slice.
func (q *Queue[T]) ToSlice() []T {
	result := make([]T, q.len)
	copy(result, q.elements)
	return result
}

// EnqueueAll adds multiple elements to the queue.
// Returns the number of elements successfully enqueued.
func (q *Queue[T]) EnqueueAll(elements []T) int {
	count := 0
	for _, element := range elements {
		if q.Enqueue(element) {
			count++
		} else {
			break
		}
	}
	return count
}

// DequeueN removes and returns up to n elements from the queue.
// Returns the elements and true if at least one element was dequeued.
func (q *Queue[T]) DequeueN(n int) ([]T, bool) {
	if q.len == 0 {
		return nil, false
	}

	count := min(n, q.len)

	result := make([]T, count)
	for i := range count {
		result[i] = q.elements[i]
	}

	q.elements = q.elements[count:]
	q.len -= count
	return result, true
}

// PeekLast returns the last element in the queue without removing it.
// Returns false if there is no last element.
func (q *Queue[T]) PeekLast() (T, bool) {
	var zero T
	if q.len == 0 {
		return zero, false
	}
	return q.elements[q.len-1], true
}

// Clone creates a deep copy of the queue.
func (q *Queue[T]) Clone() *Queue[T] {
	newQueue := &Queue[T]{
		elements: make([]T, q.len),
		len:      q.len,
		capacity: q.capacity,
	}
	copy(newQueue.elements, q.elements)
	return newQueue
}

// ForEach applies a function to each element in the queue.
func (q *Queue[T]) ForEach(fn func(T)) {
	for _, element := range q.elements {
		fn(element)
	}
}

// Filter returns a new queue containing only elements that match the predicate.
func (q *Queue[T]) Filter(fn func(T) bool) *Queue[T] {
	newQueue := &Queue[T]{capacity: q.capacity}
	for _, element := range q.elements {
		if fn(element) {
			newQueue.Enqueue(element)
		}
	}
	return newQueue
}

// String returns a string representation of the queue for debugging.
func (q *Queue[T]) String() string {
	return fmt.Sprintf("Queue{len: %d, capacity: %d, elements: %v}", q.len, q.capacity, q.elements)
}
