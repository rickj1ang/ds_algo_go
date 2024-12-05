package heap

import "sort"

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
