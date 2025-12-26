package collections

import "fmt"

// Node represents a node in a binary tree.
type Node[T any] struct {
	// Value holds the data that the node is going to carry.
	Value T

	// Left is a pointer to the left child node.
	Left *Node[T]

	// Right is a pointer to the right child node.
	Right *Node[T]
}

// Tree represents a binary tree data structure.
type Tree[T any] struct {
	root *Node[T]
	size int
}

// NewTree creates and returns a new empty binary tree.
func NewTree[T any]() *Tree[T] {
	return &Tree[T]{}
}

// NewNode creates and returns a new node with the given value.
func NewNode[T any](value T) *Node[T] {
	return &Node[T]{Value: value}
}

// Insert adds a new value to the tree.
// For a binary search tree, use InsertBST instead.
// This method inserts nodes level by level from left to right.
func (t *Tree[T]) Insert(value T) {
	newNode := NewNode(value)
	if t.root == nil {
		t.root = newNode
		t.size++
		return
	}

	queue := NewQueue[*Node[T]]()
	queue.Enqueue(t.root)

	for !queue.IsEmpty() {
		current, _ := queue.Next()
		if current.Left == nil {
			current.Left = newNode
			t.size++
			return
		}
		queue.Enqueue(current.Left)

		if current.Right == nil {
			current.Right = newNode
			t.size++
			return
		}
		queue.Enqueue(current.Right)
	}
}

// String returns a string representation of the tree using BFS (level-order traversal).
func (t *Tree[T]) String() string {
	if t.root == nil {
		return "Tree{empty}\n"
	}
	output := ""
	queue := NewQueue[*Node[T]]()
	queue.Enqueue(t.root)

	for !queue.IsEmpty() {
		current_level := []T{}
		level_size := queue.Len()

		for range level_size {
			current_node, _ := queue.Next()
			current_level = append(current_level, current_node.Value)
			if current_node.Left != nil {
				queue.Enqueue(current_node.Left)
			}
			if current_node.Right != nil {
				queue.Enqueue(current_node.Right)
			}
		}
		output += fmt.Sprintf("%v\n", current_level)
	}
	return output
}

// MaxDepth returns the maximum depth of the tree using DFS algorithm.
func (t *Tree[T]) MaxDepth() int {
	return dfs(t.root)
}

func dfs[T any](root *Node[T]) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	return 1 + max(left, right)
}

// Contains checks if a value exists in the tree using BFS.
func (t *Tree[T]) Contains(value T) bool {
	if t.root == nil {
		return false
	}

	queue := NewQueue[*Node[T]]()
	queue.Enqueue(t.root)

	for !queue.IsEmpty() {
		current, _ := queue.Next()
		if any(current.Value) == any(value) {
			return true
		}
		if current.Left != nil {
			queue.Enqueue(current.Left)
		}
		if current.Right != nil {
			queue.Enqueue(current.Right)
		}
	}
	return false
}

// Search finds and returns the first node with the given value using BFS.
// Returns the node and true if found, nil and false otherwise.
func (t *Tree[T]) Search(value T) (*Node[T], bool) {
	if t.root == nil {
		return nil, false
	}

	queue := NewQueue[*Node[T]]()
	queue.Enqueue(t.root)

	for !queue.IsEmpty() {
		current, _ := queue.Next()
		if any(current.Value) == any(value) {
			return current, true
		}
		if current.Left != nil {
			queue.Enqueue(current.Left)
		}
		if current.Right != nil {
			queue.Enqueue(current.Right)
		}
	}
	return nil, false
}

// IsEmpty returns true if the tree has no nodes.
func (t *Tree[T]) IsEmpty() bool {
	return t.root == nil
}

// Size returns the number of nodes in the tree.
func (t *Tree[T]) Size() int {
	return t.size
}

// Clear removes all nodes from the tree.
func (t *Tree[T]) Clear() {
	t.root = nil
	t.size = 0
}

// InOrder performs an in-order traversal (Left, Root, Right) and returns the values.
func (t *Tree[T]) InOrder() []T {
	result := []T{}
	inOrderHelper(t.root, &result)
	return result
}

func inOrderHelper[T any](node *Node[T], result *[]T) {
	if node == nil {
		return
	}
	inOrderHelper(node.Left, result)
	*result = append(*result, node.Value)
	inOrderHelper(node.Right, result)
}

// PreOrder performs a pre-order traversal (Root, Left, Right) and returns the values.
func (t *Tree[T]) PreOrder() []T {
	result := []T{}
	preOrderHelper(t.root, &result)
	return result
}

func preOrderHelper[T any](node *Node[T], result *[]T) {
	if node == nil {
		return
	}
	*result = append(*result, node.Value)
	preOrderHelper(node.Left, result)
	preOrderHelper(node.Right, result)
}

// PostOrder performs a post-order traversal (Left, Right, Root) and returns the values.
func (t *Tree[T]) PostOrder() []T {
	result := []T{}
	postOrderHelper(t.root, &result)
	return result
}

func postOrderHelper[T any](node *Node[T], result *[]T) {
	if node == nil {
		return
	}
	postOrderHelper(node.Left, result)
	postOrderHelper(node.Right, result)
	*result = append(*result, node.Value)
}

// LevelOrder performs a level-order (BFS) traversal and returns the values.
func (t *Tree[T]) LevelOrder() []T {
	result := []T{}
	if t.root == nil {
		return result
	}

	queue := NewQueue[*Node[T]]()
	queue.Enqueue(t.root)

	for !queue.IsEmpty() {
		current, _ := queue.Next()
		result = append(result, current.Value)
		if current.Left != nil {
			queue.Enqueue(current.Left)
		}
		if current.Right != nil {
			queue.Enqueue(current.Right)
		}
	}
	return result
}

// MinDepth returns the minimum depth of the tree (shortest path from root to leaf).
func (t *Tree[T]) MinDepth() int {
	return minDepthHelper(t.root)
}

func minDepthHelper[T any](node *Node[T]) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		return 1
	}
	if node.Left == nil {
		return 1 + minDepthHelper(node.Right)
	}
	if node.Right == nil {
		return 1 + minDepthHelper(node.Left)
	}
	return 1 + min(minDepthHelper(node.Left), minDepthHelper(node.Right))
}

// CountLeaves returns the number of leaf nodes in the tree.
func (t *Tree[T]) CountLeaves() int {
	return countLeavesHelper(t.root)
}

func countLeavesHelper[T any](node *Node[T]) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		return 1
	}
	return countLeavesHelper(node.Left) + countLeavesHelper(node.Right)
}
