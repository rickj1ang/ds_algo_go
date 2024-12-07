package heap

import (
	"container/heap"
	"testing"
)

func TestMinHeap(t *testing.T) {
	minH := Constructor()
	heap.Push(&minH, 4)
	heap.Push(&minH, 2)
	heap.Push(&minH, 7)
	heap.Push(&minH, 1)
	heap.Push(&minH, 5)
	if min := heap.Pop(&minH); min != 1 {
		t.Errorf("Expect: 1; Pop out: %d", min)
	}

	if min := heap.Pop(&minH); min != 2 {
		t.Errorf("Expect: 2; Pop out: %d", min)
	}

	if top := minH.Top(); top != 4 {
		t.Errorf("Expect: 4; Pop out: %d", top)
	}
}

func TestMaxHeap(t *testing.T) {
	// new push pop peek
	maxH := NewMaxHeap()
	maxH.push(6)
	maxH.push(8)
	maxH.push(5)
	maxH.push(1)
	maxH.push(3)
	maxH.push(7)
	if maxH.peek() != 8 {
		t.Errorf("Peek:: expect: 8; actual: %d", maxH.peek())
	}
	if out := maxH.pop(); out != 8 {
		t.Errorf("Pop:: expect: 8; actual: %d", out)
	}
	if maxH.peek() != 7 {
		t.Errorf("Peek:: expect: 7; actual: %d", maxH.peek())
	}
}

func TestBuildMaxHeap(t *testing.T) {
	// This is a debug method simply print something out, fiugure out
	// which step program stuck into
	t.Log("Start test build max heap")
	h := buildMaxHeap([]any{1, 8, 5, 4, 10, 6, 9})
	t.Log("Build the heap successfully")
	if h.peek().(int) != 10 {
		t.Errorf("Peek:: expect: 8; actual: %d", h.peek().(int))
	}

}

func TestTopK(t *testing.T) {
	nums := []int{1, 5, 3, 6, 10, 23, 456, 56, 23}
	minH := topKHeap(nums, 3)
	// check quikly by see the output from log, must use test -v
	t.Log(minH.IntSlice)
}
