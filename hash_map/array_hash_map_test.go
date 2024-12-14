package hashmap

import "testing"

// !Tip: all of my tests are bad practices
// U need to Test one function once not like
// me a big function test all the function of object
func TestArrayMap(t *testing.T) {
	// New a map
	hMap := NewArrayHashMap(100)

	// put key, val paris
	hMap.putPair(3120712, "Rick")
	hMap.putPair(3120711, "Clara")
	hMap.putPair(3120747, "Chris")
	hMap.putPair(3120792, "Jack")
	hMap.putPair(3120767, "Anna")

	// search pair
	resultRick := hMap.getValue(3120712)
	t.Log(resultRick)

	// remove pair
	hMap.remove(3120712)

	// Search nil pair
	resultNil := hMap.getValue(3120712)
	t.Log(resultNil)

	// print pair
	hMap.print()

}
