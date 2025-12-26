package collections

import (
	"testing"
)

func TestNewTree(t *testing.T) {
	tree := NewTree[int]()
	if tree == nil {
		t.Fatal("NewTree() returned nil")
	}
	if !tree.IsEmpty() {
		t.Error("New tree should be empty")
	}
	if tree.Size() != 0 {
		t.Errorf("Expected size 0, got %d", tree.Size())
	}
}

func TestNewNode(t *testing.T) {
	node := NewNode(42)
	if node == nil {
		t.Fatal("NewNode() returned nil")
	}
	if node.Value != 42 {
		t.Errorf("Expected value 42, got %d", node.Value)
	}
	if node.Left != nil || node.Right != nil {
		t.Error("New node should have nil children")
	}
}

func TestInsert(t *testing.T) {
	tree := NewTree[int]()

	tree.Insert(1)
	if tree.IsEmpty() {
		t.Error("Tree should not be empty after insert")
	}
	if tree.Size() != 1 {
		t.Errorf("Expected size 1, got %d", tree.Size())
	}

	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	if tree.Size() != 4 {
		t.Errorf("Expected size 4, got %d", tree.Size())
	}
}

func TestInsertOrder(t *testing.T) {
	tree := NewTree[int]()
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)

	// Level order should be: 1, 2, 3, 4, 5
	// Tree structure:
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5
	levelOrder := tree.LevelOrder()
	expected := []int{1, 2, 3, 4, 5}

	if len(levelOrder) != len(expected) {
		t.Fatalf("Expected %d nodes, got %d", len(expected), len(levelOrder))
	}

	for i, val := range expected {
		if levelOrder[i] != val {
			t.Errorf("Expected levelOrder[%d] = %d, got %d", i, val, levelOrder[i])
		}
	}
}

func TestIsEmpty(t *testing.T) {
	tree := NewTree[int]()

	if !tree.IsEmpty() {
		t.Error("New tree should be empty")
	}

	tree.Insert(1)
	if tree.IsEmpty() {
		t.Error("Tree with elements should not be empty")
	}

	tree.Clear()
	if !tree.IsEmpty() {
		t.Error("Cleared tree should be empty")
	}
}

func TestSize(t *testing.T) {
	tree := NewTree[int]()

	if tree.Size() != 0 {
		t.Errorf("Expected size 0, got %d", tree.Size())
	}

	for i := 1; i <= 10; i++ {
		tree.Insert(i)
		if tree.Size() != i {
			t.Errorf("Expected size %d, got %d", i, tree.Size())
		}
	}
}

func TestClear(t *testing.T) {
	tree := NewTree[int]()
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)

	tree.Clear()

	if !tree.IsEmpty() {
		t.Error("Tree should be empty after Clear()")
	}
	if tree.Size() != 0 {
		t.Errorf("Expected size 0, got %d", tree.Size())
	}
}

func TestContains(t *testing.T) {
	tree := NewTree[int]()

	if tree.Contains(1) {
		t.Error("Empty tree should not contain any element")
	}

	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)

	if !tree.Contains(2) {
		t.Error("Tree should contain 2")
	}
	if tree.Contains(5) {
		t.Error("Tree should not contain 5")
	}
}

func TestSearch(t *testing.T) {
	tree := NewTree[int]()

	// Search in empty tree
	if node, found := tree.Search(1); found {
		t.Errorf("Search in empty tree should return false, got node: %v", node)
	}

	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(30)

	// Search for existing value
	node, found := tree.Search(20)
	if !found {
		t.Error("Should find value 20")
	}
	if node.Value != 20 {
		t.Errorf("Expected node value 20, got %d", node.Value)
	}

	// Search for non-existing value
	if _, found := tree.Search(99); found {
		t.Error("Should not find value 99")
	}
}

func TestMaxDepth(t *testing.T) {
	tree := NewTree[int]()

	// Empty tree has depth 0
	if depth := tree.MaxDepth(); depth != 0 {
		t.Errorf("Expected depth 0 for empty tree, got %d", depth)
	}

	// Single node has depth 1
	tree.Insert(1)
	if depth := tree.MaxDepth(); depth != 1 {
		t.Errorf("Expected depth 1, got %d", depth)
	}

	// Tree structure:
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)

	if depth := tree.MaxDepth(); depth != 3 {
		t.Errorf("Expected depth 3, got %d", depth)
	}
}

func TestMinDepth(t *testing.T) {
	tree := NewTree[int]()

	// Empty tree has min depth 0
	if depth := tree.MinDepth(); depth != 0 {
		t.Errorf("Expected min depth 0 for empty tree, got %d", depth)
	}

	// Single node has min depth 1
	tree.Insert(1)
	if depth := tree.MinDepth(); depth != 1 {
		t.Errorf("Expected min depth 1, got %d", depth)
	}

	// Tree structure:
	//       1
	//      / \
	//     2   3
	tree.Insert(2)
	tree.Insert(3)

	if depth := tree.MinDepth(); depth != 2 {
		t.Errorf("Expected min depth 2, got %d", depth)
	}
}

func TestInOrder(t *testing.T) {
	tree := NewTree[int]()

	// Empty tree
	inOrder := tree.InOrder()
	if len(inOrder) != 0 {
		t.Errorf("Expected empty slice, got %v", inOrder)
	}

	// Tree structure:
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)

	inOrder = tree.InOrder()
	expected := []int{4, 2, 5, 1, 3}

	if len(inOrder) != len(expected) {
		t.Fatalf("Expected %d elements, got %d", len(expected), len(inOrder))
	}

	for i, val := range expected {
		if inOrder[i] != val {
			t.Errorf("Expected inOrder[%d] = %d, got %d", i, val, inOrder[i])
		}
	}
}

func TestPreOrder(t *testing.T) {
	tree := NewTree[int]()

	// Empty tree
	preOrder := tree.PreOrder()
	if len(preOrder) != 0 {
		t.Errorf("Expected empty slice, got %v", preOrder)
	}

	// Tree structure:
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)

	preOrder = tree.PreOrder()
	expected := []int{1, 2, 4, 5, 3}

	if len(preOrder) != len(expected) {
		t.Fatalf("Expected %d elements, got %d", len(expected), len(preOrder))
	}

	for i, val := range expected {
		if preOrder[i] != val {
			t.Errorf("Expected preOrder[%d] = %d, got %d", i, val, preOrder[i])
		}
	}
}

func TestPostOrder(t *testing.T) {
	tree := NewTree[int]()

	// Empty tree
	postOrder := tree.PostOrder()
	if len(postOrder) != 0 {
		t.Errorf("Expected empty slice, got %v", postOrder)
	}

	// Tree structure:
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)

	postOrder = tree.PostOrder()
	expected := []int{4, 5, 2, 3, 1}

	if len(postOrder) != len(expected) {
		t.Fatalf("Expected %d elements, got %d", len(expected), len(postOrder))
	}

	for i, val := range expected {
		if postOrder[i] != val {
			t.Errorf("Expected postOrder[%d] = %d, got %d", i, val, postOrder[i])
		}
	}
}

func TestLevelOrder(t *testing.T) {
	tree := NewTree[int]()

	// Empty tree
	levelOrder := tree.LevelOrder()
	if len(levelOrder) != 0 {
		t.Errorf("Expected empty slice, got %v", levelOrder)
	}

	// Tree structure:
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)

	levelOrder = tree.LevelOrder()
	expected := []int{1, 2, 3, 4, 5}

	if len(levelOrder) != len(expected) {
		t.Fatalf("Expected %d elements, got %d", len(expected), len(levelOrder))
	}

	for i, val := range expected {
		if levelOrder[i] != val {
			t.Errorf("Expected levelOrder[%d] = %d, got %d", i, val, levelOrder[i])
		}
	}
}

func TestCountLeaves(t *testing.T) {
	tree := NewTree[int]()

	// Empty tree has 0 leaves
	if leaves := tree.CountLeaves(); leaves != 0 {
		t.Errorf("Expected 0 leaves for empty tree, got %d", leaves)
	}

	// Single node has 1 leaf
	tree.Insert(1)
	if leaves := tree.CountLeaves(); leaves != 1 {
		t.Errorf("Expected 1 leaf, got %d", leaves)
	}

	// Tree structure:
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5
	// Leaves: 3, 4, 5 = 3 leaves
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)

	if leaves := tree.CountLeaves(); leaves != 3 {
		t.Errorf("Expected 3 leaves, got %d", leaves)
	}
}

func TestString(t *testing.T) {
	tree := NewTree[int]()

	// Empty tree
	str := tree.String()
	if str != "Tree{empty}\n" {
		t.Errorf("Expected empty tree string, got %s", str)
	}

	// Tree with elements
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)

	str = tree.String()
	if str == "" {
		t.Error("String() should return non-empty string for non-empty tree")
	}

	t.Logf("Tree string representation:\n%s", str)
}

func TestTreeWithStrings(t *testing.T) {
	tree := NewTree[string]()
	tree.Insert("hello")
	tree.Insert("world")
	tree.Insert("test")

	if !tree.Contains("world") {
		t.Error("Tree should contain 'world'")
	}

	levelOrder := tree.LevelOrder()
	expected := []string{"hello", "world", "test"}

	for i, val := range expected {
		if levelOrder[i] != val {
			t.Errorf("Expected levelOrder[%d] = %s, got %s", i, val, levelOrder[i])
		}
	}
}

func TestTreeWithStructs(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	tree := NewTree[Person]()
	tree.Insert(Person{Name: "Alice", Age: 30})
	tree.Insert(Person{Name: "Bob", Age: 25})

	if tree.Size() != 2 {
		t.Errorf("Expected size 2, got %d", tree.Size())
	}

	levelOrder := tree.LevelOrder()
	if levelOrder[0].Name != "Alice" {
		t.Errorf("Expected first person to be Alice, got %s", levelOrder[0].Name)
	}
}

func TestManualTreeConstruction(t *testing.T) {
	// Manually construct a tree for testing
	//       1
	//      / \
	//     2   3
	//    /
	//   4
	tree := NewTree[int]()
	tree.root = NewNode(1)
	tree.root.Left = NewNode(2)
	tree.root.Right = NewNode(3)
	tree.root.Left.Left = NewNode(4)
	tree.size = 4

	if tree.MaxDepth() != 3 {
		t.Errorf("Expected max depth 3, got %d", tree.MaxDepth())
	}

	if tree.MinDepth() != 2 {
		t.Errorf("Expected min depth 2, got %d", tree.MinDepth())
	}

	if tree.CountLeaves() != 2 {
		t.Errorf("Expected 2 leaves (3 and 4), got %d", tree.CountLeaves())
	}
}
