package list

import (
	"errors"
	"fmt"

	"github.com/ckshitij/go-collection/collections/node"
)

type List[T any] struct {
	head *node.Node[T]
	tail *node.Node[T]
	size uint
}

var (
	ErrInvalidPosition = errors.New("invalid position, please check the list size")
)

// Create a new list with the initializer element
func NewList[T any]() List[T] {
	return List[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// insert element in the beginning of the list
func (list *List[T]) Push(element T) {
	newNode := node.NewNode(element)
	if list.head != nil {
		newNode.SetNext(list.head)
		list.head.SetPrevious(newNode)
		list.head = newNode
	} else {
		list.head = newNode
		list.tail = newNode
	}
	list.size++
}

// insert element in the beginning of the list
func (list *List[T]) Append(element T) {
	newNode := node.NewNode(element)
	if list.tail != nil {
		newNode.SetPrevious(list.tail)
		list.tail.SetNext(newNode)
		list.tail = newNode
	} else {
		list.head = newNode
		list.tail = newNode
	}
	list.size++
}

// Remove element from the beginning of the list
func (list *List[T]) PopFront() {
	if list.head == nil {
		return
	}
	list.head = list.head.Next()
	if list.head == nil {
		list.tail = nil
	} else {
		list.head.SetPrevious(nil)
	}
	list.size--
}

// Remove element from the end of the list
func (list *List[T]) PopBack() {
	if list.head == nil {
		return
	}
	list.tail = list.tail.Previous()
	if list.tail == nil {
		list.head = nil
	} else {
		list.tail.SetNext(nil)
	}
	list.size--
}

// Get element from the beginning of the list
func (list *List[T]) Front() *node.Node[T] {
	return list.head
}

// Get element from the back of the list
func (list *List[T]) Back() *node.Node[T] {
	return list.tail
}

// Return the size of list
func (list *List[T]) Len() uint {
	return list.size
}

// InsertAtPosition inserts a node at a specific position in the list
func (dll *List[T]) InsertAtPosition(data T, position int) error {
	newNode := node.NewNode(data)

	if position < 0 {
		return fmt.Errorf("position must be non-negative")
	}

	if position == 0 {
		dll.Push(data)
		return nil
	}

	current := dll.head
	for i := 0; i < position-1; i++ {
		if current == nil {
			return fmt.Errorf("position out of bounds")
		}
		current = current.Next()
	}

	if current == nil || current.Next() == nil {
		dll.Append(data)
	} else {
		newNode.SetNext(current.Next())
		newNode.SetPrevious(current)
		current.Next().SetPrevious(newNode)
		current.SetNext(newNode)
	}
	dll.size++
	return nil
}

func (list *List[T]) IterateForward() {
	itr := list.head
	fmt.Println()
	i := 1
	for itr != nil {
		fmt.Printf(" %d : %v ", i, itr.Element())
		i++
		itr = itr.Next()
	}
	fmt.Println()
}

func (list *List[T]) IterateBackward() {
	itr := list.tail
	fmt.Println()
	i := list.size
	for itr != nil {
		fmt.Printf(" %d : %v ", i, itr.Element())
		i--
		itr = itr.Previous()
	}
	fmt.Println()
}
