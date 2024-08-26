package list

import (
	"fmt"

	"github.com/ckshitij/go-collection/collections/node"
)

type List[T any] struct {
	head *node.Node[T]
	tail *node.Node[T]
	size uint64
}

// Create a new list with the initializer element
func NewList[T any]() List[T] {
	return List[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// insert element in the beginning of the list
func (list *List[T]) Insert(element T) {
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

// Return the size of list
func (list *List[T]) Len() uint64 {
	return list.size
}

func (list *List[T]) Debug() {
	from := list.head
	to := list.tail
	for from != to {
		fmt.Printf("\n %v \n", from.Element())
		from = from.Next()
	}
	fmt.Printf("\n %v \n", from.Element())
}
