package heap

import (
	"container/heap"
	"testing"
)

func TestMaxHeap(t *testing.T) {
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
