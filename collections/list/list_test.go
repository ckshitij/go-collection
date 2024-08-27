package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_Push(t *testing.T) {
	list := NewList[int]()
	list.Push(10)
	list.Push(20)

	assert.Equal(t, uint(2), list.Len(), "The list size should be 2")
	assert.Equal(t, 20, list.Front().Element(), "The first element should be 20")
	assert.Equal(t, 10, list.Back().Element(), "The last element should be 10")
}

func TestList_Append(t *testing.T) {
	list := NewList[int]()
	list.Append(10)
	list.Append(20)

	assert.Equal(t, uint(2), list.Len(), "The list size should be 2")
	assert.Equal(t, 10, list.Front().Element(), "The first element should be 10")
	assert.Equal(t, 20, list.Back().Element(), "The last element should be 20")
}

func TestList_PopFront(t *testing.T) {
	list := NewList[int]()
	list.Push(10)
	list.Push(20)
	list.PopFront()

	assert.Equal(t, uint(1), list.Len(), "The list size should be 1 after popping the front")
	assert.Equal(t, 10, list.Front().Element(), "The first element should now be 10")
}

func TestList_PopBack(t *testing.T) {
	list := NewList[int]()
	list.Push(10)
	list.Push(20)
	list.PopBack()

	assert.Equal(t, uint(1), list.Len(), "The list size should be 1 after popping the back")
	assert.Equal(t, 20, list.Front().Element(), "The first element should still be 20")
}

func TestList_InsertAtPosition(t *testing.T) {
	list := NewList[int]()
	list.Push(10)
	list.Push(20)

	err := list.InsertAtPosition(15, 1)
	assert.Nil(t, err, "Inserting at position 1 should not return an error")
	assert.Equal(t, uint(3), list.Len(), "The list size should be 3 after insertion")
	assert.Equal(t, 20, list.Front().Element(), "The first element should be 20")
	assert.Equal(t, 15, list.Front().Next().Element(), "The element at position 1 should be 15")
	assert.Equal(t, 10, list.Back().Element(), "The last element should be 10")

	// Test inserting at an invalid position
	err = list.InsertAtPosition(25, 10)
	assert.NotNil(t, err, "Inserting at an out-of-bounds position should return an error")
}

func TestList_Len(t *testing.T) {
	list := NewList[int]()
	list.Push(10)
	list.Push(20)

	assert.Equal(t, uint(2), list.Len(), "The list size should be 2")
	list.PopFront()
	assert.Equal(t, uint(1), list.Len(), "The list size should be 1 after popping the front")
	list.PopBack()
	assert.Equal(t, uint(0), list.Len(), "The list size should be 0 after popping the back")
}

func TestList_IterateForward(t *testing.T) {
	list := NewList[int]()
	list.Push(10)
	list.Push(20)
	list.Append(30)

	expected := []int{20, 10, 30}
	var result []int

	itr := list.Front()
	for itr != nil {
		result = append(result, itr.Element())
		itr = itr.Next()
	}

	assert.Equal(t, expected, result, "The list should iterate forward correctly")
}

func TestList_IterateBackward(t *testing.T) {
	list := NewList[int]()
	list.Push(10)
	list.Push(20)
	list.Append(30)

	expected := []int{30, 10, 20}
	var result []int

	itr := list.Back()
	for itr != nil {
		result = append(result, itr.Element())
		itr = itr.Previous()
	}

	assert.Equal(t, expected, result, "The list should iterate backward correctly")
}
