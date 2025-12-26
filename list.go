package collections

import "fmt"

// ListNode represents a node in a linked list.
type ListNode[T any] struct {
	// Value holds the data stored in the node.
	Value T

	// Next is a pointer to the next node in the list.
	Next *ListNode[T]
}

// LinkedList represents a singly linked list data structure.
type LinkedList[T any] struct {
	head     *ListNode[T]
	tail     *ListNode[T]
	size     int
	circular bool
}

// NewLinkedList creates and returns a new empty linked list.
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// NewCircularLinkedList creates and returns a new empty circular linked list.
func NewCircularLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{circular: true}
}

// NewListNode creates and returns a new list node with the given value.
func NewListNode[T any](value T) *ListNode[T] {
	return &ListNode[T]{Value: value}
}

// Append adds a new value to the end of the list.
func (l *LinkedList[T]) Append(value T) {
	newNode := NewListNode(value)

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		if l.circular {
			newNode.Next = newNode
		}
	} else {
		l.tail.Next = newNode
		l.tail = newNode
		if l.circular {
			l.tail.Next = l.head
		}
	}
	l.size++
}

// Prepend adds a new value to the beginning of the list.
func (l *LinkedList[T]) Prepend(value T) {
	newNode := NewListNode(value)

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		if l.circular {
			newNode.Next = newNode
		}
	} else {
		newNode.Next = l.head
		l.head = newNode
		if l.circular {
			l.tail.Next = l.head
		}
	}
	l.size++
}

// InsertAt inserts a value at the specified index.
// Returns false if the index is out of bounds.
func (l *LinkedList[T]) InsertAt(index int, value T) bool {
	if index < 0 || index > l.size {
		return false
	}

	if index == 0 {
		l.Prepend(value)
		return true
	}

	if index == l.size {
		l.Append(value)
		return true
	}

	newNode := NewListNode(value)
	current := l.head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	newNode.Next = current.Next
	current.Next = newNode
	l.size++
	return true
}

// RemoveFirst removes and returns the first element from the list.
// Returns false if the list is empty.
func (l *LinkedList[T]) RemoveFirst() (T, bool) {
	var zero T

	if l.head == nil {
		return zero, false
	}

	value := l.head.Value

	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.Next
		if l.circular {
			l.tail.Next = l.head
		}
	}

	l.size--
	return value, true
}

// RemoveLast removes and returns the last element from the list.
// Returns false if the list is empty.
func (l *LinkedList[T]) RemoveLast() (T, bool) {
	var zero T

	if l.head == nil {
		return zero, false
	}

	value := l.tail.Value

	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		l.size--
		return value, true
	}

	// Find the second-to-last node
	current := l.head
	for current.Next != l.tail {
		current = current.Next
	}

	current.Next = nil
	l.tail = current
	if l.circular {
		l.tail.Next = l.head
	}

	l.size--
	return value, true
}

// RemoveAt removes the element at the specified index.
// Returns the removed value and true if successful, zero value and false otherwise.
func (l *LinkedList[T]) RemoveAt(index int) (T, bool) {
	var zero T

	if index < 0 || index >= l.size {
		return zero, false
	}

	if index == 0 {
		return l.RemoveFirst()
	}

	if index == l.size-1 {
		return l.RemoveLast()
	}

	current := l.head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	value := current.Next.Value
	current.Next = current.Next.Next
	l.size--
	return value, true
}

// Remove removes the first occurrence of the specified value.
// Returns true if an element was removed, false otherwise.
func (l *LinkedList[T]) Remove(value T) bool {
	if l.head == nil {
		return false
	}

	// Check if head needs to be removed
	if any(l.head.Value) == any(value) {
		l.RemoveFirst()
		return true
	}

	// Search for the value
	current := l.head
	maxIterations := l.size
	for i := 0; i < maxIterations && current.Next != nil; i++ {
		if l.circular && current.Next == l.head {
			break
		}

		if any(current.Next.Value) == any(value) {
			if current.Next == l.tail {
				l.tail = current
			}
			current.Next = current.Next.Next
			if l.circular {
				l.tail.Next = l.head
			}
			l.size--
			return true
		}
		current = current.Next
	}

	return false
}

// Get returns the value at the specified index.
// Returns false if the index is out of bounds.
func (l *LinkedList[T]) Get(index int) (T, bool) {
	var zero T

	if index < 0 || index >= l.size {
		return zero, false
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	return current.Value, true
}

// GetFirst returns the first element in the list.
// Returns false if the list is empty.
func (l *LinkedList[T]) GetFirst() (T, bool) {
	var zero T

	if l.head == nil {
		return zero, false
	}

	return l.head.Value, true
}

// GetLast returns the last element in the list.
// Returns false if the list is empty.
func (l *LinkedList[T]) GetLast() (T, bool) {
	var zero T

	if l.tail == nil {
		return zero, false
	}

	return l.tail.Value, true
}

// Contains checks if the list contains the specified value.
func (l *LinkedList[T]) Contains(value T) bool {
	current := l.head
	for i := 0; i < l.size; i++ {
		if any(current.Value) == any(value) {
			return true
		}
		current = current.Next
	}
	return false
}

// IndexOf returns the index of the first occurrence of the specified value.
// Returns -1 if the value is not found.
func (l *LinkedList[T]) IndexOf(value T) int {
	current := l.head
	for i := 0; i < l.size; i++ {
		if any(current.Value) == any(value) {
			return i
		}
		current = current.Next
	}
	return -1
}

// Find returns the first node with the specified value.
// Returns nil if not found.
func (l *LinkedList[T]) Find(value T) *ListNode[T] {
	current := l.head
	for i := 0; i < l.size; i++ {
		if any(current.Value) == any(value) {
			return current
		}
		current = current.Next
	}
	return nil
}

// Len returns the number of elements in the list.
func (l *LinkedList[T]) Len() int {
	return l.size
}

// IsEmpty returns true if the list has no elements.
func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

// Clear removes all elements from the list.
func (l *LinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

// Reverse reverses the order of elements in the list.
func (l *LinkedList[T]) Reverse() {
	if l.head == nil || l.head == l.tail {
		return
	}

	// Break circle if circular
	wasCircular := l.circular
	if wasCircular {
		l.BreakCircle()
	}

	var prev *ListNode[T]
	current := l.head
	l.tail = l.head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	l.head = prev

	// Restore circle if it was circular
	if wasCircular {
		l.MakeCircular()
	}
}

// MakeCircular converts the list to a circular linked list.
func (l *LinkedList[T]) MakeCircular() {
	if l.circular {
		return
	}

	if l.tail != nil {
		l.tail.Next = l.head
	}
	l.circular = true
}

// BreakCircle converts a circular list to a regular linked list.
func (l *LinkedList[T]) BreakCircle() {
	if !l.circular {
		return
	}

	if l.tail != nil {
		l.tail.Next = nil
	}
	l.circular = false
}

// IsCircular returns true if the list is circular.
func (l *LinkedList[T]) IsCircular() bool {
	return l.circular
}

// ToSlice returns all elements as a slice.
func (l *LinkedList[T]) ToSlice() []T {
	result := make([]T, 0, l.size)
	current := l.head

	for i := 0; i < l.size; i++ {
		result = append(result, current.Value)
		current = current.Next
	}

	return result
}

// ForEach applies a function to each element in the list.
func (l *LinkedList[T]) ForEach(fn func(T)) {
	current := l.head
	for i := 0; i < l.size; i++ {
		fn(current.Value)
		current = current.Next
	}
}

// String returns a string representation of the list.
func (l *LinkedList[T]) String() string {
	if l.head == nil {
		return "LinkedList{empty}"
	}

	result := "LinkedList{"
	current := l.head

	for i := 0; i < l.size; i++ {
		result += fmt.Sprintf("%v", current.Value)
		if i < l.size-1 {
			result += " -> "
		}
		current = current.Next
	}

	if l.circular {
		result += " -> (circular)"
	}

	result += "}"
	return result
}
