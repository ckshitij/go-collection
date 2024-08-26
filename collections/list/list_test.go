package list

import "testing"

func TestNewList(t *testing.T) {
	newList := NewList[int]()
	if newList.head != nil && newList.tail != nil {
		t.Errorf("Expected nil but found different %v", newList)
	}

	if newList.Len() != 0 {
		t.Errorf("Expected length to be 0 but, got %v", newList.Len())
	}
}

func TestListInsert(t *testing.T) {
	newList := NewList[int]()
	newList.Insert(10)
	if newList.head.Element() != 10 {
		t.Errorf("Expected 10 but, got %v", newList.head)
	}

	newList.Insert(20)
	if newList.head.Element() != 20 {
		t.Errorf("Expected 20 but, got %v", newList.head)
	}

	newList.Append(30)
	newList.Append(40)
	if newList.Len() != 4 {
		t.Errorf("Expected length to be 4 but, got %v", newList.Len())
	}
}
