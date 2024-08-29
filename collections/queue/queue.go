package queue

import (
	"errors"

	"github.com/ckshitij/go-collection/collections/list"
)

type Queue[T any] struct {
	head list.List[T]
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		head: list.NewList[T](),
	}
}

func (st *Queue[T]) Enqueue(value T) {
	st.head.Push(value)
}

func (st *Queue[T]) IsEmpty() bool {
	return st.head.Len() < 1
}

func (st *Queue[T]) Dequeue() error {
	if st.IsEmpty() {
		return errors.New("invalid operation: empty Queue")
	}
	st.head.PopBack()
	return nil
}

func (st *Queue[T]) Front() T {
	var empty T
	if !st.IsEmpty() {
		return st.head.Back().Element()
	}
	return empty
}

func (st *Queue[T]) Back() T {
	var empty T
	if !st.IsEmpty() {
		return st.head.Front().Element()
	}
	return empty
}
