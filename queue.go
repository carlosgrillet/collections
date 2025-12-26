package collections

type Queue[T any] struct {
	elements []T
	len int
}

// Generare a new empty queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Enqueue a new element into the queue
func (q *Queue[T]) Enqueue(element T) {
	q.elements = append(q.elements, element)
	q.len += 1
}

// Next removes and returns the first element from the queue
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

// Peek return what is the next element in the queue without removing it
func (q *Queue[T]) Peek() (T, bool) {
	var zero T
	if q.len == 0 {
		return zero, false
	}
	return q.elements[0], true
}

// Clear remove all elements from the queue
func (q *Queue[T]) Clear() {
	q.elements = []T{}
	q.len = 0
}

// Len return the current lenght of the queue
func (q *Queue[T]) Len() int {
	return q.len
}

// IsEmpty return true if the queue has no elements
func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}
