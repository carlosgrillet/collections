package collections

import (
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	list := NewLinkedList[int]()
	if list == nil {
		t.Fatal("NewLinkedList() returned nil")
	}
	if !list.IsEmpty() {
		t.Error("New list should be empty")
	}
	if list.Len() != 0 {
		t.Errorf("Expected length 0, got %d", list.Len())
	}
	if list.IsCircular() {
		t.Error("New list should not be circular")
	}
}

func TestNewCircularLinkedList(t *testing.T) {
	list := NewCircularLinkedList[int]()
	if list == nil {
		t.Fatal("NewCircularLinkedList() returned nil")
	}
	if !list.IsCircular() {
		t.Error("New circular list should be circular")
	}
}

func TestAppend(t *testing.T) {
	list := NewLinkedList[int]()

	list.Append(1)
	if list.IsEmpty() {
		t.Error("List should not be empty after append")
	}
	if list.Len() != 1 {
		t.Errorf("Expected length 1, got %d", list.Len())
	}

	list.Append(2)
	list.Append(3)
	if list.Len() != 3 {
		t.Errorf("Expected length 3, got %d", list.Len())
	}

	// Check order
	expected := []int{1, 2, 3}
	slice := list.ToSlice()
	for i, val := range expected {
		if slice[i] != val {
			t.Errorf("Expected slice[%d] = %d, got %d", i, val, slice[i])
		}
	}
}

func TestPrepend(t *testing.T) {
	list := NewLinkedList[int]()

	list.Prepend(1)
	list.Prepend(2)
	list.Prepend(3)

	if list.Len() != 3 {
		t.Errorf("Expected length 3, got %d", list.Len())
	}

	// Check order (should be reversed)
	expected := []int{3, 2, 1}
	slice := list.ToSlice()
	for i, val := range expected {
		if slice[i] != val {
			t.Errorf("Expected slice[%d] = %d, got %d", i, val, slice[i])
		}
	}
}

func TestInsertAt(t *testing.T) {
	list := NewLinkedList[int]()

	// Insert at beginning
	if !list.InsertAt(0, 1) {
		t.Error("InsertAt should succeed")
	}

	// Insert at end
	if !list.InsertAt(1, 3) {
		t.Error("InsertAt should succeed")
	}

	// Insert in middle
	if !list.InsertAt(1, 2) {
		t.Error("InsertAt should succeed")
	}

	expected := []int{1, 2, 3}
	slice := list.ToSlice()
	for i, val := range expected {
		if slice[i] != val {
			t.Errorf("Expected slice[%d] = %d, got %d", i, val, slice[i])
		}
	}

	// Test out of bounds
	if list.InsertAt(-1, 99) {
		t.Error("InsertAt with negative index should fail")
	}
	if list.InsertAt(10, 99) {
		t.Error("InsertAt with out of bounds index should fail")
	}
}

func TestRemoveFirst(t *testing.T) {
	list := NewLinkedList[int]()

	// Remove from empty list
	if val, ok := list.RemoveFirst(); ok {
		t.Errorf("RemoveFirst on empty list should fail, got %v", val)
	}

	list.Append(1)
	list.Append(2)
	list.Append(3)

	val, ok := list.RemoveFirst()
	if !ok || val != 1 {
		t.Errorf("Expected to remove 1, got %v", val)
	}
	if list.Len() != 2 {
		t.Errorf("Expected length 2, got %d", list.Len())
	}

	first, _ := list.GetFirst()
	if first != 2 {
		t.Errorf("Expected first element to be 2, got %d", first)
	}
}

func TestRemoveLast(t *testing.T) {
	list := NewLinkedList[int]()

	// Remove from empty list
	if val, ok := list.RemoveLast(); ok {
		t.Errorf("RemoveLast on empty list should fail, got %v", val)
	}

	list.Append(1)
	list.Append(2)
	list.Append(3)

	val, ok := list.RemoveLast()
	if !ok || val != 3 {
		t.Errorf("Expected to remove 3, got %v", val)
	}
	if list.Len() != 2 {
		t.Errorf("Expected length 2, got %d", list.Len())
	}

	last, _ := list.GetLast()
	if last != 2 {
		t.Errorf("Expected last element to be 2, got %d", last)
	}
}

func TestRemoveAt(t *testing.T) {
	list := NewLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	// Remove from middle
	val, ok := list.RemoveAt(2)
	if !ok || val != 3 {
		t.Errorf("Expected to remove 3, got %v", val)
	}

	// Remove first
	val, ok = list.RemoveAt(0)
	if !ok || val != 1 {
		t.Errorf("Expected to remove 1, got %v", val)
	}

	// Remove last
	val, ok = list.RemoveAt(list.Len() - 1)
	if !ok || val != 5 {
		t.Errorf("Expected to remove 5, got %v", val)
	}

	// Test out of bounds
	if _, ok := list.RemoveAt(10); ok {
		t.Error("RemoveAt with out of bounds index should fail")
	}
}

func TestRemove(t *testing.T) {
	list := NewLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(2)
	list.Append(4)

	// Remove existing value
	if !list.Remove(2) {
		t.Error("Remove should succeed")
	}
	if list.Len() != 4 {
		t.Errorf("Expected length 4, got %d", list.Len())
	}

	// Verify first occurrence was removed
	slice := list.ToSlice()
	expected := []int{1, 3, 2, 4}
	for i, val := range expected {
		if slice[i] != val {
			t.Errorf("Expected slice[%d] = %d, got %d", i, val, slice[i])
		}
	}

	// Remove non-existing value
	if list.Remove(99) {
		t.Error("Remove should fail for non-existing value")
	}
}

func TestGet(t *testing.T) {
	list := NewLinkedList[string]()
	list.Append("first")
	list.Append("second")
	list.Append("third")

	val, ok := list.Get(1)
	if !ok || val != "second" {
		t.Errorf("Expected 'second', got %v", val)
	}

	// Test out of bounds
	if _, ok := list.Get(10); ok {
		t.Error("Get with out of bounds index should fail")
	}
	if _, ok := list.Get(-1); ok {
		t.Error("Get with negative index should fail")
	}
}

func TestGetFirst(t *testing.T) {
	list := NewLinkedList[int]()

	// Empty list
	if _, ok := list.GetFirst(); ok {
		t.Error("GetFirst on empty list should fail")
	}

	list.Append(1)
	list.Append(2)

	val, ok := list.GetFirst()
	if !ok || val != 1 {
		t.Errorf("Expected 1, got %v", val)
	}
}

func TestGetLast(t *testing.T) {
	list := NewLinkedList[int]()

	// Empty list
	if _, ok := list.GetLast(); ok {
		t.Error("GetLast on empty list should fail")
	}

	list.Append(1)
	list.Append(2)

	val, ok := list.GetLast()
	if !ok || val != 2 {
		t.Errorf("Expected 2, got %v", val)
	}
}

func TestContains(t *testing.T) {
	list := NewLinkedList[int]()

	if list.Contains(1) {
		t.Error("Empty list should not contain any element")
	}

	list.Append(1)
	list.Append(2)
	list.Append(3)

	if !list.Contains(2) {
		t.Error("List should contain 2")
	}
	if list.Contains(5) {
		t.Error("List should not contain 5")
	}
}

func TestIndexOf(t *testing.T) {
	list := NewLinkedList[int]()
	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Append(20)

	// Find first occurrence
	if index := list.IndexOf(20); index != 1 {
		t.Errorf("Expected index 1, got %d", index)
	}

	// Not found
	if index := list.IndexOf(99); index != -1 {
		t.Errorf("Expected index -1, got %d", index)
	}
}

func TestFind(t *testing.T) {
	list := NewLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	node := list.Find(2)
	if node == nil {
		t.Fatal("Find should return node")
	}
	if node.Value != 2 {
		t.Errorf("Expected node value 2, got %d", node.Value)
	}

	// Not found
	if node := list.Find(99); node != nil {
		t.Error("Find should return nil for non-existing value")
	}
}

func TestLen(t *testing.T) {
	list := NewLinkedList[int]()

	if list.Len() != 0 {
		t.Errorf("Expected length 0, got %d", list.Len())
	}

	for i := 1; i <= 5; i++ {
		list.Append(i)
		if list.Len() != i {
			t.Errorf("Expected length %d, got %d", i, list.Len())
		}
	}
}

func TestIsEmpty(t *testing.T) {
	list := NewLinkedList[int]()

	if !list.IsEmpty() {
		t.Error("New list should be empty")
	}

	list.Append(1)
	if list.IsEmpty() {
		t.Error("List with elements should not be empty")
	}

	list.Clear()
	if !list.IsEmpty() {
		t.Error("Cleared list should be empty")
	}
}

func TestClear(t *testing.T) {
	list := NewLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.Clear()

	if !list.IsEmpty() {
		t.Error("List should be empty after Clear()")
	}
	if list.Len() != 0 {
		t.Errorf("Expected length 0, got %d", list.Len())
	}
}

func TestReverse(t *testing.T) {
	list := NewLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	list.Reverse()

	expected := []int{4, 3, 2, 1}
	slice := list.ToSlice()

	for i, val := range expected {
		if slice[i] != val {
			t.Errorf("Expected slice[%d] = %d, got %d", i, val, slice[i])
		}
	}

	// Check first and last
	first, _ := list.GetFirst()
	last, _ := list.GetLast()
	if first != 4 {
		t.Errorf("Expected first element 4, got %d", first)
	}
	if last != 1 {
		t.Errorf("Expected last element 1, got %d", last)
	}
}

func TestMakeCircular(t *testing.T) {
	list := NewLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	if list.IsCircular() {
		t.Error("List should not be circular initially")
	}

	list.MakeCircular()

	if !list.IsCircular() {
		t.Error("List should be circular after MakeCircular()")
	}

	// Verify tail points to head
	if list.tail.Next != list.head {
		t.Error("Tail should point to head in circular list")
	}
}

func TestBreakCircle(t *testing.T) {
	list := NewCircularLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	if !list.IsCircular() {
		t.Error("List should be circular initially")
	}

	list.BreakCircle()

	if list.IsCircular() {
		t.Error("List should not be circular after BreakCircle()")
	}

	// Verify tail does not point to head
	if list.tail.Next != nil {
		t.Error("Tail should point to nil in non-circular list")
	}
}

func TestCircularListAppend(t *testing.T) {
	list := NewCircularLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	if list.Len() != 3 {
		t.Errorf("Expected length 3, got %d", list.Len())
	}

	// Verify circular structure
	if list.tail.Next != list.head {
		t.Error("Tail should point to head in circular list")
	}
}

func TestCircularListPrepend(t *testing.T) {
	list := NewCircularLinkedList[int]()
	list.Prepend(1)
	list.Prepend(2)
	list.Prepend(3)

	if list.Len() != 3 {
		t.Errorf("Expected length 3, got %d", list.Len())
	}

	// Verify circular structure
	if list.tail.Next != list.head {
		t.Error("Tail should point to head in circular list")
	}

	// Check order
	expected := []int{3, 2, 1}
	slice := list.ToSlice()
	for i, val := range expected {
		if slice[i] != val {
			t.Errorf("Expected slice[%d] = %d, got %d", i, val, slice[i])
		}
	}
}

func TestCircularListRemove(t *testing.T) {
	list := NewCircularLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.Remove(2)

	if list.Len() != 2 {
		t.Errorf("Expected length 2, got %d", list.Len())
	}

	// Verify circular structure maintained
	if list.tail.Next != list.head {
		t.Error("Tail should still point to head after removal")
	}
}

func TestCircularListReverse(t *testing.T) {
	list := NewCircularLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.Reverse()

	expected := []int{3, 2, 1}
	slice := list.ToSlice()
	for i, val := range expected {
		if slice[i] != val {
			t.Errorf("Expected slice[%d] = %d, got %d", i, val, slice[i])
		}
	}

	// Verify still circular
	if !list.IsCircular() {
		t.Error("List should still be circular after reverse")
	}
	if list.tail.Next != list.head {
		t.Error("Tail should point to head after reverse")
	}
}

func TestToSlice(t *testing.T) {
	list := NewLinkedList[int]()

	// Empty list
	slice := list.ToSlice()
	if len(slice) != 0 {
		t.Errorf("Expected empty slice, got length %d", len(slice))
	}

	// With elements
	list.Append(1)
	list.Append(2)
	list.Append(3)

	slice = list.ToSlice()
	expected := []int{1, 2, 3}

	for i, val := range expected {
		if slice[i] != val {
			t.Errorf("Expected slice[%d] = %d, got %d", i, val, slice[i])
		}
	}
}

func TestForEach(t *testing.T) {
	list := NewLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	sum := 0
	list.ForEach(func(val int) {
		sum += val
	})

	if sum != 6 {
		t.Errorf("Expected sum 6, got %d", sum)
	}
}

func TestString(t *testing.T) {
	list := NewLinkedList[int]()

	// Empty list
	str := list.String()
	if str != "LinkedList{empty}" {
		t.Errorf("Expected 'LinkedList{empty}', got '%s'", str)
	}

	// With elements
	list.Append(1)
	list.Append(2)
	list.Append(3)

	str = list.String()
	if str == "" {
		t.Error("String() should return non-empty string")
	}

	t.Logf("List string representation: %s", str)
}

func TestCircularListString(t *testing.T) {
	list := NewCircularLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	str := list.String()
	if str == "" {
		t.Error("String() should return non-empty string")
	}

	// Should indicate it's circular
	t.Logf("Circular list string representation: %s", str)
}

func TestLinkedListWithStrings(t *testing.T) {
	list := NewLinkedList[string]()
	list.Append("hello")
	list.Append("world")
	list.Append("test")

	if !list.Contains("world") {
		t.Error("List should contain 'world'")
	}

	slice := list.ToSlice()
	expected := []string{"hello", "world", "test"}

	for i, val := range expected {
		if slice[i] != val {
			t.Errorf("Expected slice[%d] = %s, got %s", i, val, slice[i])
		}
	}
}

func TestLinkedListWithStructs(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	list := NewLinkedList[Person]()
	list.Append(Person{Name: "Alice", Age: 30})
	list.Append(Person{Name: "Bob", Age: 25})

	if list.Len() != 2 {
		t.Errorf("Expected length 2, got %d", list.Len())
	}

	first, _ := list.GetFirst()
	if first.Name != "Alice" {
		t.Errorf("Expected first person to be Alice, got %s", first.Name)
	}
}

func TestSingleElementOperations(t *testing.T) {
	list := NewLinkedList[int]()
	list.Append(42)

	// Test operations on single element list
	if val, _ := list.GetFirst(); val != 42 {
		t.Error("GetFirst failed on single element list")
	}
	if val, _ := list.GetLast(); val != 42 {
		t.Error("GetLast failed on single element list")
	}

	list.Reverse()
	if val, _ := list.GetFirst(); val != 42 {
		t.Error("Reverse failed on single element list")
	}

	val, ok := list.RemoveFirst()
	if !ok || val != 42 {
		t.Error("RemoveFirst failed on single element list")
	}
	if !list.IsEmpty() {
		t.Error("List should be empty after removing single element")
	}
}
