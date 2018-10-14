package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type intCmp struct{}

func (c *intCmp) Equal(i, j interface{}) bool {
	return i.(int) == j.(int)
}

func (c *intCmp) Less(i, j interface{}) bool {
	return i.(int) < j.(int)
}

func TestHeapify(t *testing.T) {
	expected := []interface{}{45, 32, 10, 20, 2, 5}
	heap := Heapify([]interface{}{10, 2, 5, 20, 32, 45}, new(intCmp))
	assert.Equal(t, expected, heap.Slice)
}

func TestPop(t *testing.T) {
	expected := []interface{}{32, 20, 10, 5, 2}
	heap := &Heap{
		comp:  new(intCmp),
		Slice: []interface{}{45, 32, 10, 20, 2, 5},
	}
	root := heap.Pop()
	assert.Equal(t, 45, root)
	assert.Equal(t, expected, heap.Slice)
	assert.Len(t, heap.Slice, 5)
}

func TestPopEmpty(t *testing.T) {
	heap := NewHeap(new(intCmp))
	assert.Equal(t, nil, heap.Pop())
	assert.Len(t, heap.Slice, 0)
}

func TestPopSingle(t *testing.T) {
	heap := &Heap{
		comp:  new(intCmp),
		Slice: []interface{}{10},
	}
	assert.Equal(t, 10, heap.Pop())
	assert.Len(t, heap.Slice, 0)
}

func TestPush(t *testing.T) {
	expected := []interface{}{51, 32, 45, 20, 2, 5, 10}
	heap := &Heap{
		comp:  new(intCmp),
		Slice: []interface{}{45, 32, 10, 20, 2, 5},
	}
	heap.Push(51)
	assert.Equal(t, expected, heap.Slice)
}

func TestSort(t *testing.T) {
	expected := []interface{}{2, 5, 10, 20, 32, 45}
	heap := &Heap{
		comp:  new(intCmp),
		Slice: []interface{}{45, 32, 10, 20, 2, 5},
	}
	heap.Sort()
	assert.Equal(t, expected, heap.Slice)
}
