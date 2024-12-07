package heap

import (
	"fmt"
	"sort"
)

// We need to impl the interface of IntSlice(Less, Len, Swap) and Push(), Pop()
// to impl a MinHeap, but this code is a easy way to impl MinHeap with Golang
// but it is a MinHeap can not be a MaxHeap, you can use MinHeap as a MaxHeap
// by make your all input negative and reverse them when output. or you can
// impl the Less(i, j int)int by your self
type MinHeap struct{ sort.IntSlice }

// we do not need to use any type because all we impl is Int for simplity
//
//	but we respect to the interafce
func (m *MinHeap) Push(val any) {
	m.IntSlice = append(m.IntSlice, val.(int))
}

func (m *MinHeap) Pop() any {
	temp := m.IntSlice
	val := temp[len(temp)-1]
	m.IntSlice = temp[:len(temp)-1]
	return val
}

func (m *MinHeap) Top() any {
	return m.IntSlice[0]
}

func Constructor() MinHeap {
	return MinHeap{}
}

// own maxHeap for sake of learning make no sense
// impl it by yourself when you need to use them
type maxHeap struct {
	data []any
}

func NewMaxHeap() maxHeap {
	return maxHeap{
		data: make([]any, 0),
	}
}

/*
				5       value:[5, 3, 7, 2, 6, 9, 4]
			/  \ 			index:[0, 1, 2, 3, 4, 5, 6]
		 3    7
	  / \  / \
	 2  6  9  4
*/
func (m *maxHeap) left(i int) int {
	return 2*i + 1
}
func (m *maxHeap) right(i int) int {
	return 2*i + 2
}
func (m *maxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (m *maxHeap) peek() any {
	return m.data[0]
}

func (m *maxHeap) push(val any) {
	m.data = append(m.data, val)
	m.siftUp(len(m.data) - 1)
}

func (m *maxHeap) siftUp(i int) {
	for true {
		p := m.parent(i)
		if p < 0 || m.data[i].(int) <= m.data[p].(int) {
			break
		}

		m.swap(i, p)
		i = p
	}
}

func (m *maxHeap) swap(i, j int) {
	m.data[i], m.data[j] = m.data[j], m.data[i]
}

func (m *maxHeap) pop() any {
	if m.isEmpty() {
		fmt.Println("error")
		return nil
	}
	m.swap(0, len(m.data)-1)
	val := m.data[len(m.data)-1]
	m.data = m.data[:len(m.data)-1]

	m.siftDown(0)

	return val
}

func (m *maxHeap) siftDown(i int) {
	for true {
		leftSon := m.left(i)
		rightSon := m.right(i)
		maxIndex := i

		if leftSon < m.size() && m.data[leftSon].(int) > m.data[maxIndex].(int) {
			maxIndex = leftSon
		}
		if rightSon < m.size() && m.data[rightSon].(int) > m.data[maxIndex].(int) {
			maxIndex = rightSon
		}
		if maxIndex == i {
			break
		}
		m.swap(i, maxIndex)

		i = maxIndex
	}
}

func (m *maxHeap) size() int {
	return len(m.data)
}

func (m *maxHeap) isEmpty() bool {
	if len(m.data) > 0 {
		return false
	}
	return true
}

func buildMaxHeap(nums []any) *maxHeap {
	h := &maxHeap{data: nums}
	for i := h.parent(len(h.data) - 1); i >= 0; i-- {
		//heapify
		h.siftDown(i)
	}
	return h
}
