package collections

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue[int]()
	if q == nil {
		t.Fatal("NewQueue() returned nil")
	}
	if !q.IsEmpty() {
		t.Error("New queue should be empty")
	}
	if q.Len() != 0 {
		t.Errorf("Expected length 0, got %d", q.Len())
	}
}

func TestNewBoundedQueue(t *testing.T) {
	q := NewBoundedQueue[int](5)
	if q == nil {
		t.Fatal("NewBoundedQueue() returned nil")
	}
	if !q.IsEmpty() {
		t.Error("New bounded queue should be empty")
	}
	if q.IsFull() {
		t.Error("New bounded queue should not be full")
	}
}

func TestEnqueue(t *testing.T) {
	q := NewQueue[int]()

	if !q.Enqueue(1) {
		t.Error("Enqueue should return true for unbounded queue")
	}
	if q.IsEmpty() {
		t.Error("Queue should not be empty after enqueue")
	}
	if q.Len() != 1 {
		t.Errorf("Expected length 1, got %d", q.Len())
	}

	q.Enqueue(2)
	q.Enqueue(3)
	if q.Len() != 3 {
		t.Errorf("Expected length 3, got %d", q.Len())
	}
}

func TestEnqueueBounded(t *testing.T) {
	q := NewBoundedQueue[int](2)

	if !q.Enqueue(1) {
		t.Error("First enqueue should succeed")
	}
	if !q.Enqueue(2) {
		t.Error("Second enqueue should succeed")
	}
	if q.Enqueue(3) {
		t.Error("Third enqueue should fail when queue is full")
	}
	if !q.IsFull() {
		t.Error("Queue should be full")
	}
	if q.Len() != 2 {
		t.Errorf("Expected length 2, got %d", q.Len())
	}
}

func TestNext(t *testing.T) {
	q := NewQueue[string]()

	// Test empty queue
	if val, ok := q.Next(); ok {
		t.Errorf("Next on empty queue should return false, got value: %v", val)
	}

	// Test FIFO order
	q.Enqueue("first")
	q.Enqueue("second")
	q.Enqueue("third")

	val, ok := q.Next()
	if !ok || val != "first" {
		t.Errorf("Expected 'first', got %v", val)
	}
	if q.Len() != 2 {
		t.Errorf("Expected length 2, got %d", q.Len())
	}

	val, ok = q.Next()
	if !ok || val != "second" {
		t.Errorf("Expected 'second', got %v", val)
	}

	val, ok = q.Next()
	if !ok || val != "third" {
		t.Errorf("Expected 'third', got %v", val)
	}

	if !q.IsEmpty() {
		t.Error("Queue should be empty after removing all elements")
	}
}

func TestPeek(t *testing.T) {
	q := NewQueue[int]()

	// Test empty queue
	if val, ok := q.Peek(); ok {
		t.Errorf("Peek on empty queue should return false, got value: %v", val)
	}

	q.Enqueue(10)
	q.Enqueue(20)

	val, ok := q.Peek()
	if !ok || val != 10 {
		t.Errorf("Expected 10, got %v", val)
	}
	if q.Len() != 2 {
		t.Error("Peek should not remove element")
	}

	// Peek again to ensure it doesn't change
	val, ok = q.Peek()
	if !ok || val != 10 {
		t.Errorf("Expected 10 again, got %v", val)
	}
}

func TestPeekLast(t *testing.T) {
	q := NewQueue[int]()

	// Test empty queue
	if val, ok := q.PeekLast(); ok {
		t.Errorf("PeekLast on empty queue should return false, got value: %v", val)
	}

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	val, ok := q.PeekLast()
	if !ok || val != 30 {
		t.Errorf("Expected 30, got %v", val)
	}
	if q.Len() != 3 {
		t.Error("PeekLast should not remove element")
	}
}

func TestQueueClear(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	q.Clear()

	if !q.IsEmpty() {
		t.Error("Queue should be empty after Clear()")
	}
	if q.Len() != 0 {
		t.Errorf("Expected length 0, got %d", q.Len())
	}
}

func TestQueueLen(t *testing.T) {
	q := NewQueue[int]()

	if q.Len() != 0 {
		t.Errorf("Expected length 0, got %d", q.Len())
	}

	q.Enqueue(1)
	if q.Len() != 1 {
		t.Errorf("Expected length 1, got %d", q.Len())
	}

	q.Enqueue(2)
	q.Enqueue(3)
	if q.Len() != 3 {
		t.Errorf("Expected length 3, got %d", q.Len())
	}

	q.Next()
	if q.Len() != 2 {
		t.Errorf("Expected length 2, got %d", q.Len())
	}
}

func TestQueueIsEmpty(t *testing.T) {
	q := NewQueue[int]()

	if !q.IsEmpty() {
		t.Error("New queue should be empty")
	}

	q.Enqueue(1)
	if q.IsEmpty() {
		t.Error("Queue with elements should not be empty")
	}

	q.Next()
	if !q.IsEmpty() {
		t.Error("Queue should be empty after removing all elements")
	}
}

func TestIsFull(t *testing.T) {
	// Unbounded queue
	q1 := NewQueue[int]()
	q1.Enqueue(1)
	q1.Enqueue(2)
	if q1.IsFull() {
		t.Error("Unbounded queue should never be full")
	}

	// Bounded queue
	q2 := NewBoundedQueue[int](2)
	if q2.IsFull() {
		t.Error("Empty bounded queue should not be full")
	}

	q2.Enqueue(1)
	if q2.IsFull() {
		t.Error("Bounded queue with 1/2 elements should not be full")
	}

	q2.Enqueue(2)
	if !q2.IsFull() {
		t.Error("Bounded queue at capacity should be full")
	}
}

func TestQueueContains(t *testing.T) {
	q := NewQueue[int]()

	if q.Contains(1) {
		t.Error("Empty queue should not contain any element")
	}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if !q.Contains(2) {
		t.Error("Queue should contain 2")
	}
	if q.Contains(5) {
		t.Error("Queue should not contain 5")
	}
}

func TestQueueToSlice(t *testing.T) {
	q := NewQueue[int]()

	// Test empty queue
	slice := q.ToSlice()
	if len(slice) != 0 {
		t.Errorf("Expected empty slice, got length %d", len(slice))
	}

	// Test with elements
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	slice = q.ToSlice()
	if len(slice) != 3 {
		t.Errorf("Expected slice length 3, got %d", len(slice))
	}

	expected := []int{1, 2, 3}
	for i, val := range expected {
		if slice[i] != val {
			t.Errorf("Expected slice[%d] = %d, got %d", i, val, slice[i])
		}
	}

	// Ensure it's a copy
	slice[0] = 999
	if val, _ := q.Peek(); val != 1 {
		t.Error("Modifying slice should not affect queue")
	}
}

func TestEnqueueAll(t *testing.T) {
	q := NewQueue[int]()

	elements := []int{1, 2, 3, 4, 5}
	count := q.EnqueueAll(elements)

	if count != 5 {
		t.Errorf("Expected 5 elements enqueued, got %d", count)
	}
	if q.Len() != 5 {
		t.Errorf("Expected queue length 5, got %d", q.Len())
	}
}

func TestEnqueueAllBounded(t *testing.T) {
	q := NewBoundedQueue[int](3)

	elements := []int{1, 2, 3, 4, 5}
	count := q.EnqueueAll(elements)

	if count != 3 {
		t.Errorf("Expected 3 elements enqueued, got %d", count)
	}
	if q.Len() != 3 {
		t.Errorf("Expected queue length 3, got %d", q.Len())
	}
	if !q.IsFull() {
		t.Error("Queue should be full")
	}
}

func TestDequeueN(t *testing.T) {
	q := NewQueue[int]()

	// Test empty queue
	if elements, ok := q.DequeueN(3); ok {
		t.Errorf("DequeueN on empty queue should return false, got: %v", elements)
	}

	// Add elements
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	// Dequeue 3 elements
	elements, ok := q.DequeueN(3)
	if !ok {
		t.Error("DequeueN should succeed")
	}
	if len(elements) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(elements))
	}

	expected := []int{1, 2, 3}
	for i, val := range expected {
		if elements[i] != val {
			t.Errorf("Expected elements[%d] = %d, got %d", i, val, elements[i])
		}
	}

	if q.Len() != 2 {
		t.Errorf("Expected queue length 2, got %d", q.Len())
	}
}

func TestDequeueNMoreThanAvailable(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)

	// Request more than available
	elements, ok := q.DequeueN(5)
	if !ok {
		t.Error("DequeueN should succeed when queue has elements")
	}
	if len(elements) != 2 {
		t.Errorf("Expected 2 elements, got %d", len(elements))
	}
	if !q.IsEmpty() {
		t.Error("Queue should be empty")
	}
}

func TestClone(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	clone := q.Clone()

	if clone.Len() != q.Len() {
		t.Errorf("Clone should have same length as original")
	}

	// Verify contents
	if val, _ := clone.Peek(); val != 1 {
		t.Error("Clone should have same first element")
	}

	// Modify clone and ensure original is unaffected
	clone.Next()
	if q.Len() != 3 {
		t.Error("Modifying clone should not affect original")
	}
	if clone.Len() != 2 {
		t.Errorf("Expected clone length 2, got %d", clone.Len())
	}
}

func TestCloneBounded(t *testing.T) {
	q := NewBoundedQueue[int](5)
	q.Enqueue(1)
	q.Enqueue(2)

	clone := q.Clone()

	// Ensure capacity is preserved
	clone.Enqueue(3)
	clone.Enqueue(4)
	clone.Enqueue(5)

	if !clone.IsFull() {
		t.Error("Clone should have same capacity as original")
	}
}

func TestQueueForEach(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	sum := 0
	q.ForEach(func(val int) {
		sum += val
	})

	if sum != 6 {
		t.Errorf("Expected sum 6, got %d", sum)
	}
}

func TestFilter(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	// Filter even numbers
	filtered := q.Filter(func(val int) bool {
		return val%2 == 0
	})

	if filtered.Len() != 2 {
		t.Errorf("Expected 2 even numbers, got %d", filtered.Len())
	}

	// Verify filtered values
	if val, _ := filtered.Next(); val != 2 {
		t.Errorf("Expected 2, got %d", val)
	}
	if val, _ := filtered.Next(); val != 4 {
		t.Errorf("Expected 4, got %d", val)
	}

	// Ensure original is unmodified
	if q.Len() != 5 {
		t.Error("Filter should not modify original queue")
	}
}

func TestQueueString(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)

	str := q.String()
	if str == "" {
		t.Error("String() should return non-empty string")
	}

	// Just verify it doesn't panic and returns something
	t.Logf("Queue string representation: %s", str)
}

func TestQueueWithStrings(t *testing.T) {
	q := NewQueue[string]()
	q.Enqueue("hello")
	q.Enqueue("world")

	if val, _ := q.Next(); val != "hello" {
		t.Errorf("Expected 'hello', got %v", val)
	}
	if val, _ := q.Next(); val != "world" {
		t.Errorf("Expected 'world', got %v", val)
	}
}

func TestQueueWithStructs(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	q := NewQueue[Person]()
	q.Enqueue(Person{Name: "Alice", Age: 30})
	q.Enqueue(Person{Name: "Bob", Age: 25})

	person, ok := q.Next()
	if !ok || person.Name != "Alice" {
		t.Errorf("Expected Alice, got %v", person)
	}
}
