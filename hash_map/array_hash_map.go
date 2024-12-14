package hashmap

import "fmt"

type pair struct {
	key int
	val string
}

type arrayHashMap struct {
	buckets []*pair
	capcity int
}

func NewArrayHashMap(capcity int) *arrayHashMap {
	buckets := make([]*pair, capcity)
	return &arrayHashMap{
		buckets: buckets,
		capcity: capcity,
	}
}

// our hash(key) = key very easy hash function for simplicity
func (a *arrayHashMap) hashFunc(key int) int {
	return key
}

// hash(key) % capcity
func (a *arrayHashMap) getIndex(key int) int {
	index := a.hashFunc(key) % a.capcity
	return index
}

func (a *arrayHashMap) getValue(key int) string {
	index := a.getIndex(key)
	pair := a.buckets[index]

	if pair == nil {
		return "Not Found"
	}

	return pair.val
}

func (a *arrayHashMap) putPair(key int, val string) {
	pair := &pair{
		key: key,
		val: val,
	}

	index := a.getIndex(pair.key)
	a.buckets[index] = pair
}

func (a *arrayHashMap) remove(key int) {
	index := a.getIndex(key)
	a.buckets[index] = nil
}

// for future extend capcity need, maybe we need len()
func (a *arrayHashMap) len() int {
	return len(a.buckets)
}

func (a *arrayHashMap) print() {
	for _, pair := range a.buckets {
		if pair != nil {
			fmt.Println(pair.key, "->", pair.val)
		}
	}
}
