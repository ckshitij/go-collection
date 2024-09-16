package heap

type Heap[T any] struct {
	table   []T
	compare Comparable[T]
}

type Comparable[T any] func(a T, b T) bool

func NewHeap[T any](compFunc Comparable[T], elements ...T) Heap[T] {
	table := []T{}
	table = append(table, elements...)
	hp := Heap[T]{
		table:   table,
		compare: compFunc,
	}
	hp.buildHeap()
	return hp
}

func (hp *Heap[T]) InsertElement(data T) {
	hp.table = append(hp.table, data)
	hp.heapify(hp.Size() - 1)
}

func (hp *Heap[T]) left(index int) int {
	return (2 * index) + 1
}

func (hp *Heap[T]) right(index int) int {
	return (2 * index) + 2
}

// func (hp *Heap[T]) parent(index int) int {
// 	return (index - 1) / 2
// }

func (hp *Heap[T]) Pop() T {
	var val T
	if hp.Size() <= 0 {
		return val
	}
	val = hp.table[0]
	hp.swap(0, hp.Size()-1)
	hp.table = hp.table[:hp.Size()-1]
	hp.heapify(0)
	return val
}

func (hp *Heap[T]) isValidInd(ind int) bool {
	return ind >= 0 && ind < hp.Size()
}

func (hp *Heap[T]) swap(a, b int) {
	hp.table[a], hp.table[b] = hp.table[b], hp.table[a]
}

func (hp *Heap[T]) heapify(rootInd int) {
	left := hp.left(rootInd)
	right := hp.right(rootInd)
	cmpInd := rootInd

	// Compare the left child
	if hp.isValidInd(left) && hp.compare(hp.table[left], hp.table[cmpInd]) {
		cmpInd = left
	}

	// Compare the right child
	if hp.isValidInd(right) && hp.compare(hp.table[right], hp.table[cmpInd]) {
		cmpInd = right
	}

	// If the root is not the largest (or smallest, depending on the comparison), swap and continue heapifying
	if cmpInd != rootInd {
		hp.swap(rootInd, cmpInd)
		hp.heapify(cmpInd)
	}
}

func (hp *Heap[T]) buildHeap() {
	totalElement := hp.Size()
	for i := totalElement / 2; i >= 0; i-- {
		hp.heapify(i)
	}
}

func (hp *Heap[T]) Size() int {
	return len(hp.table)
}
