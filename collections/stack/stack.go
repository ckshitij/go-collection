package stack

import "github.com/ckshitij/go-collection/collections/list"

type Stack[T any] struct {
	head list.List[T]
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		head: list.NewList[T](),
	}
}

func (st *Stack[T]) Push(value T) {
	st.head.Push(value)
}

func (st *Stack[T]) IsEmpty() bool {
	return st.head.Len() < 1
}

func (st *Stack[T]) Pop() {
	st.head.PopFront()
}

func (st *Stack[T]) Top() T {
	var empty T
	if !st.IsEmpty() {
		return st.head.Front().Element()
	}
	return empty
}
