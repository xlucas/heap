package heap

import (
	"math"
)

// Comparator is the interface to be implemented by heap value comparators.
type Comparator interface {
	// Equal returns whether i is equal to j.
	Equal(i, j interface{}) bool

	// Less returns whether is is lower than j.
	Less(i, j interface{}) bool
}

// parentIndex returns the index of the parent of the child located at the
// specified index.
func parentIndex(childIndex int) int {
	return int(math.Floor(float64(childIndex-1) / 2))
}

// leftChildIndex returns the index of the left child of the parent located at
// the given index.
func leftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

// Heap represents a simple datastructure that can contain entries for which the
// following rule is always true: a parent's value is at least at large as any
// of its direct children. Values stored in the heap should be of the same type
// and comparable with the given comparator.
type Heap struct {
	comp  Comparator
	Slice []interface{}
}

// NewHeap creates a new empty Heap.
func NewHeap(c Comparator) *Heap {
	return &Heap{
		comp: c,
	}
}

// Heapify creates a new Heap given a slice and a value comparator.
func Heapify(slice []interface{}, c Comparator) *Heap {
	cp := make([]interface{}, len(slice))
	copy(cp, slice)

	heap := &Heap{
		comp:  c,
		Slice: cp,
	}

	end := len(slice) - 1

	for start := parentIndex(end); start >= 0; start-- {
		heap.siftDown(start, end)
	}

	return heap
}

// Pop removes the root from the Heap and returns it.
func (h *Heap) Pop() interface{} {
	if len(h.Slice) < 1 {
		return nil
	}

	end := len(h.Slice) - 1
	root := h.Slice[0]

	if len(h.Slice) == 1 {
		h.Slice = nil
		return root
	}

	h.swapItems(0, end)
	h.Slice = h.Slice[:end]
	h.siftDown(0, len(h.Slice)-1)

	return root
}

// Push inserts a new value into the Heap.
func (h *Heap) Push(val interface{}) {
	h.Slice = append(h.Slice, val)
	h.siftUp(len(h.Slice) - 1)
}

// RepairDown is used to sift down starting at index i.
func (h *Heap) RepairDown(i int) {
	h.siftDown(i, len(h.Slice)-1)
}

// RepairUp is used to sift up starting at index i.
func (h *Heap) RepairUp(i int) {
	h.siftUp(i)
}

// Sort is used to run the Heapsort algorithm in-place. The Heap's slice will
// contain sorted values in ascending order.
func (h *Heap) Sort() {
	for end := len(h.Slice) - 1; end >= 0; end-- {
		h.swapItems(0, end)
		h.siftDown(0, end-1)
	}
}

// siftDown is used to sift down starting at index start until the index end.
func (h *Heap) siftDown(start, end int) {
	root := start

	for leftChildIndex(root) <= end {
		child := leftChildIndex(root)
		swap := root

		if h.comp.Less(h.Slice[swap], h.Slice[child]) {
			swap = child
		}
		if child+1 <= end && h.comp.Less(h.Slice[swap], h.Slice[child+1]) {
			swap = child + 1
		}
		if swap == root {
			return
		}

		h.swapItems(root, swap)
		root = swap
	}
}

// siftUp is used to sift up starting at index start until index 0.
func (h *Heap) siftUp(start int) {
	root := start

	for parentIndex(root) >= 0 {
		parent := parentIndex(root)
		swap := root

		if h.comp.Less(h.Slice[parent], h.Slice[swap]) {
			swap = parent
		}
		if swap == root {
			return
		}

		h.swapItems(root, swap)
		root = swap
	}
}

// swapItems execute an in-place swap of the two items located at index i and j.
func (h *Heap) swapItems(i, j int) {
	h.Slice[i], h.Slice[j] = h.Slice[j], h.Slice[i]
}
