package hashmap

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	// usage of hashmap in Golang, pretty easy
	hmap := make(map[int]string)

	// a ID name map
	hmap[3120111] = "Rick" //my name
	hmap[3120127] = "Clara"
	hmap[3120274] = "Chris"
	hmap[3120872] = "Jack"

	// Get my name use my StudentID
	name := hmap[3120111]
	t.Log(name)

	// bye Clara
	delete(hmap, 3120127)

	_, isExist := hmap[3120127]
	t.Log(isExist)

	// use for ... range ... to iteration a map
	for key, value := range hmap {
		stuID := fmt.Sprintf("Student: %s --> ID: %d\n", value, key)
		t.Log(stuID)
	}
}
