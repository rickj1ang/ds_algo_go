package search

import "fmt"

func hashSearch() {
	// hashSearch is a kind of search method based on hashmap
	// I can not talk more aboht it
	hMap := make(map[int]string)

	hMap[123] = "Rick"
	name := hMap[123]

	fmt.Print(name)
}
