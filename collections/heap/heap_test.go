package heap

import (
	"testing"
)

func TestHeap_Push_Pop_MinHeap(t *testing.T) {
	// Min-Heap comparison function: returns true if a < b
	minHeap := NewHeap[int](func(a, b int) bool {
		return a < b
	})

	// Test pushing elements into the heap
	minHeap.Push(10)
	minHeap.Push(5)
	minHeap.Push(15)
	minHeap.Push(1)

	if minHeap.Size() != 4 {
		t.Errorf("Expected heap size 4, but got %d", minHeap.Size())
	}

	// Test popping elements from the heap
	expectedPopOrder := []int{1, 5, 10, 15}
	for i, expected := range expectedPopOrder {
		val := minHeap.Pop()
		if val != expected {
			t.Errorf("TestHeap_Push_Pop_MinHeap: expected %d at position %d, got %d", expected, i, val)
		}
	}

	if minHeap.Size() != 0 {
		t.Errorf("Expected heap size 0 after popping all elements, but got %d", minHeap.Size())
	}
}

func TestHeap_Push_Pop_MaxHeap(t *testing.T) {
	// Max-Heap comparison function: returns true if a > b
	maxHeap := NewHeap[int](func(a, b int) bool {
		return a > b
	})

	// Test pushing elements into the heap
	maxHeap.Push(10)
	maxHeap.Push(5)
	maxHeap.Push(15)
	maxHeap.Push(1)

	if maxHeap.Size() != 4 {
		t.Errorf("Expected heap size 4, but got %d", maxHeap.Size())
	}

	// Test popping elements from the heap
	expectedPopOrder := []int{15, 10, 5, 1}
	for i, expected := range expectedPopOrder {
		val := maxHeap.Pop()
		if val != expected {
			t.Errorf("TestHeap_Push_Pop_MaxHeap: expected %d at position %d, got %d", expected, i, val)
		}
	}

	if maxHeap.Size() != 0 {
		t.Errorf("Expected heap size 0 after popping all elements, but got %d", maxHeap.Size())
	}
}

func TestHeap_BuildHeap_MinHeap(t *testing.T) {
	// Min-Heap comparison function: returns true if a < b
	heap := NewHeap[int](func(a, b int) bool {
		return a < b
	}, 10, 15, 5, 1)

	expectedPopOrder := []int{1, 5, 10, 15}
	for i, expected := range expectedPopOrder {
		val := heap.Pop()
		if val != expected {
			t.Errorf("TestHeap_BuildHeap_MinHeap: expected %d at position %d, got %d", expected, i, val)
		}
	}
}

func TestHeap_BuildHeap_MaxHeap(t *testing.T) {
	// Max-Heap comparison function: returns true if a > b
	heap := NewHeap[int](func(a, b int) bool {
		return a > b
	}, 10, 15, 5, 1)

	expectedPopOrder := []int{15, 10, 5, 1}
	for i, expected := range expectedPopOrder {
		val := heap.Pop()
		if val != expected {
			t.Errorf("TestHeap_BuildHeap_MaxHeap: expected %d at position %d, got %d", expected, i, val)
		}
	}
}
