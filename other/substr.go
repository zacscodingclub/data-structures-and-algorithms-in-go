package main

import (
	"fmt"
)

func RunSubstr() {
	haystack := "averylongtestphrase"
	needle := "test"

	found := substr(haystack, needle)
	fmt.Println(found)
}

func substr(haystack, needle string) int {
	for i, c := range haystack {
		for _, m := range needle {
			if c == m {
				return i
			}
		}
	}
	return -1
}
