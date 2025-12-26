<div align="center">
  <br> 
  <h1> collections </h1>

  <p>  A package with generic data structures for Go applications.</p>
  <p>  (Queues, linked lists, binary trees and more)</p>
</div>


<!-- Badges -->
<p align="center">
  <a href="https://github.com/carlosgrillet/collections/actions/workflows/go.yml">
    <img src="https://github.com/carlosgrillet/collections/actions/workflows/go.yml/badge.svg?branch=main"
      alt="Github Actions (workflows)" />
  </a>

  <a href="https://github.com/carlosgrillet/collections">
    <img src="https://img.shields.io/github/go-mod/go-version/carlosgrillet/collections.svg"
      alt="Go module version" />
  </a>

  <a href="https://pkg.go.dev/github.com/carlosgrillet/collections">
    <img src="https://img.shields.io/badge/godoc-reference-blue.svg"
      alt="GoDoc reference" />
  </a>

  <a href="https://goreportcard.com/report/github.com/carlosgrillet/collections">
    <img src="https://goreportcard.com/badge/github.com/carlosgrillet/collections"
      alt="Go Report Card" />
  </a>

  <a href="https://codecov.io/gh/carlosgrillet/collections">
    <img src="https://codecov.io/gh/carlosgrillet/collections/branch/main/graph/badge.svg"
      alt="codecov" />
  </a>

  <a href="https://opensource.org/licenses/MIT">
    <img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License: MIT" />
  </a>
</p>

## What is this?

Collections is a Go package that gives you useful data structures. These are ready-to-use building blocks for organizing and storing data in your programs.

## What's included?

- **Queue** - First in, first out (like a line at a store)
- **Linked List** - Items connected in a chain (can be circular too)
- **Binary Tree** - Items organized in a tree shape

## How to install

```bash
go get github.com/carlosgrillet/collections
```

## How to use

### Queue

A queue works like a line at a store - the first person in line gets served first.

```go
package main

import "github.com/carlosgrillet/collections"

func main() {
    // Create a new queue
    q := collections.NewQueue[int]()

    // Add items to the queue
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    // Remove and get the first item
    first, ok := q.Next()  // Returns 1

    // Look at the next item without removing it
    next, ok := q.Peek()   // Returns 2

    // Get the queue size
    size := q.Len()        // Returns 2
}
```

**Queue features:**
- `Enqueue(item)` - Add an item to the back
- `Next()` - Remove and return the first item
- `Peek()` - Look at the first item without removing it
- `PeekLast()` - Look at the last item
- `Clear()` - Remove all items
- `Len()` - Get the number of items
- `IsEmpty()` - Check if the queue is empty
- `IsFull()` - Check if a bounded queue is full
- `Contains(item)` - Check if an item exists
- `ToSlice()` - Convert to a slice
- `EnqueueAll(items)` - Add multiple items at once
- `DequeueN(n)` - Remove and return n items
- `Clone()` - Make a copy of the queue
- `ForEach(fn)` - Run a function on each item
- `Filter(fn)` - Create a new queue with only matching items

**Bounded Queue:**

You can create a queue with a size limit:

```go
q := collections.NewBoundedQueue[int](5)  // Max 5 items
q.Enqueue(1)  // Returns true
q.Enqueue(2)  // Returns true
// ... add 3 more items
q.Enqueue(6)  // Returns false (queue is full)
```

### Linked List

A linked list is like a chain where each item points to the next one.

```go
package main

import "github.com/carlosgrillet/collections"

func main() {
    // Create a new linked list
    list := collections.NewLinkedList[string]()

    // Add items
    list.Append("first")    // Add to the end
    list.Append("second")
    list.Prepend("zero")    // Add to the beginning

    // Get items
    first, ok := list.GetFirst()  // Returns "zero"
    last, ok := list.GetLast()    // Returns "second"
    item, ok := list.Get(1)       // Returns "first"

    // Remove items
    removed, ok := list.RemoveFirst()  // Removes "zero"
    list.Remove("second")              // Removes "second"

    // Check what's in the list
    exists := list.Contains("first")  // Returns true
    index := list.IndexOf("first")    // Returns 0
}
```

**Linked List features:**
- `Append(item)` - Add to the end
- `Prepend(item)` - Add to the beginning
- `InsertAt(index, item)` - Add at a specific position
- `Get(index)` - Get item at position
- `GetFirst()` - Get the first item
- `GetLast()` - Get the last item
- `Remove(item)` - Remove first occurrence of item
- `RemoveFirst()` - Remove the first item
- `RemoveLast()` - Remove the last item
- `RemoveAt(index)` - Remove item at position
- `Contains(item)` - Check if item exists
- `IndexOf(item)` - Find the position of an item
- `Find(item)` - Find the node containing the item
- `Len()` - Get the number of items
- `IsEmpty()` - Check if the list is empty
- `Clear()` - Remove all items
- `Reverse()` - Flip the order of all items
- `ToSlice()` - Convert to a slice
- `ForEach(fn)` - Run a function on each item

**Circular Linked List:**

A circular list is like a regular list, but the last item points back to the first one (like a circle).

```go
// Create a circular list
list := collections.NewCircularLinkedList[int]()
list.Append(1)
list.Append(2)
list.Append(3)

// Convert a regular list to circular
regularList := collections.NewLinkedList[int]()
regularList.Append(1)
regularList.MakeCircular()

// Convert back to a regular list
regularList.BreakCircle()

// Check if a list is circular
isCircular := regularList.IsCircular()  // Returns false
```

### Binary Tree

A tree is like an upside-down family tree where each item can have a left and right child.

```go
package main

import "github.com/carlosgrillet/collections"

func main() {
    // Create a new tree
    tree := collections.NewTree[int]()

    // Add items (adds level by level, left to right)
    tree.Insert(1)
    tree.Insert(2)
    tree.Insert(3)
    tree.Insert(4)

    // Tree looks like:
    //       1
    //      / \
    //     2   3
    //    /
    //   4

    // Get information about the tree
    depth := tree.MaxDepth()       // Returns 3
    minDepth := tree.MinDepth()    // Returns 2
    size := tree.Size()            // Returns 4
    leaves := tree.CountLeaves()   // Returns 2 (nodes 3 and 4)

    // Search for items
    exists := tree.Contains(3)     // Returns true
    node, found := tree.Search(2)  // Returns the node with value 2

    // Get items in different orders
    inOrder := tree.InOrder()        // [4, 2, 1, 3]
    preOrder := tree.PreOrder()      // [1, 2, 4, 3]
    postOrder := tree.PostOrder()    // [4, 2, 3, 1]
    levelOrder := tree.LevelOrder()  // [1, 2, 3, 4]
}
```

**Binary Tree features:**
- `Insert(item)` - Add an item to the tree
- `Contains(item)` - Check if an item exists
- `Search(item)` - Find the node with the item
- `MaxDepth()` - Get the maximum depth (height)
- `MinDepth()` - Get the minimum depth to a leaf
- `Size()` - Get the number of items
- `CountLeaves()` - Count leaf nodes (nodes with no children)
- `IsEmpty()` - Check if the tree is empty
- `Clear()` - Remove all items
- `InOrder()` - Get items in order: Left, Root, Right
- `PreOrder()` - Get items in order: Root, Left, Right
- `PostOrder()` - Get items in order: Left, Right, Root
- `LevelOrder()` - Get items level by level (breadth-first)
- `String()` - Get a text view of the tree

## Using generic types

All data structures work with any type you want:

```go
// With numbers
queueInt := collections.NewQueue[int]()

// With text
queueString := collections.NewQueue[string]()

// With your own types
type Person struct {
    Name string
    Age  int
}
listPeople := collections.NewLinkedList[Person]()
listPeople.Append(Person{Name: "Alice", Age: 30})
```

## Running tests

```bash
go test
```

For more details:

```bash
go test -v
```

## License

MIT License - see LICENSE file for details.
