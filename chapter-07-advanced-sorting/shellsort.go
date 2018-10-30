package main

import (
	"fmt"
	"math/rand"
)

func RunShellSort() {
	max := 100
	ss := NewArray(max)

	for i := 0; i < max; i++ {
		ss.insert(rand.Intn(max * 10))
	}

	ss.shellSort()
	ss.display()
}

type array struct {
	items []int
	size  int
}

func NewArray(max int) *array {
	items := make([]int, max, max)
	return &array{
		items: items,
		size:  0,
	}
}

func (a *array) insert(v int) {
	a.items[a.size] = v
	a.size++
}

func (a *array) display() {
	fmt.Printf("items: ")
	for _, el := range a.items {
		fmt.Printf("%d ", el)
	}
	fmt.Println("")
}

func (a *array) shellSort() {
	var inner, outer, temp int

	h := 1
	for h <= a.size/3 {
		h = h*3 + 1
	}

	for h > 0 {
		for outer = h; outer < a.size; outer++ {
			temp = a.items[outer]
			inner = outer

			for inner > h-1 && a.items[inner-h] >= temp {
				a.items[inner] = a.items[inner-h]
				inner -= h
			}
			a.items[inner] = temp
		}
		h = (h - 1) / 3
	}
}
