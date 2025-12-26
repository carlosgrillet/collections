package collections

import "fmt"

type Node[T any] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

type Tree[T any] struct {
	root *Node[T]
}

// Generate a new empty binary tree
func NewTree[T any]() *Tree[T] {
	return &Tree[T]{}
}

// Print all the nodes of a binary tree using BFS algorithm.
func (t *Tree[T]) String() string {
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

// MaxDepth return the maximum depth of the tree using DFS algorithm
func (t *Tree[T]) MaxDepth() int {
	return DFS(t.root)	
}

func DFS[T any](root *Node[T]) int {
	if root == nil {
		return 0
	}
	left := DFS(root.Left)
	right := DFS(root.Right)
	return 1 + max(left, right)
}
