package main

import "fmt"

func main() {
	RunTowers()
}

// (n^2 + n)/2
func triangle(n int) int {
	fmt.Printf("Entering %d\n", n)
	if n == 1 {
		fmt.Printf("Returning 1\n")
		return 1
	}
	t := n + triangle(n-1)
	fmt.Printf("Returning %d\n", t)
	return t
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
