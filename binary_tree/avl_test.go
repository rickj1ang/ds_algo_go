package binarytree

import "testing"

var aT *avlTree = &avlTree{
	root: nil,
}

func TestAVLInsert(t *testing.T) {
	aT.insert(8)
	aT.insert(123)
	aT.insert(34)
	aT.insert(6)
	aT.insert(87)
	aT.insert(13)
	aT.insert(90)
	aT.insert(34)
	res := aT.avlBFS()
	t.Log(res)

}
func TestAVLRemove(t *testing.T) {
	aT.remove(123)
	res := aT.avlBFS()
	t.Log(res)

}
