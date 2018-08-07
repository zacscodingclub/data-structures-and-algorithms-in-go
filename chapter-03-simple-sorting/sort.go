package main

import "fmt"

type Array struct {
	items []int
}

func (a *Array) Display() {
	for i := 0; i < len(a.items); i++ {
		if a.items[i] != 0 {
			fmt.Printf("%v\n", a.items[i])
		}
	}
}

func (a *Array) swap(one, two int) {
	tmp := a.items[one]
	a.items[one] = a.items[two]
	a.items[two] = tmp
}

func (a *Array) BubbleSort() {
	for out := len(a.items) - 1; out > 1; out-- {
		for in := 0; in < out; in++ {
			if a.items[in] > a.items[in+1] {
				a.swap(in, in+1)
			}
		}
	}
}

func (a *Array) InsertionSort() {
	for i := 1; i < len(a.items); i++ {
		k := a.items[i]
		j := i - 1
		for j >= 0 && a.items[j] > k {
			a.items[j+1] = a.items[j]
			j--
		}
		a.items[j+1] = k
	}
}

func InsertionSortRun() {
	i := make([]int, 7, 7)
	a := Array{items: i}

	a.items = append(a.items, 77)
	a.items = append(a.items, 99)
	a.items = append(a.items, 11)
	a.items = append(a.items, 13)
	a.items = append(a.items, 44)
	a.items = append(a.items, 12)
	a.items = append(a.items, 12)
	a.items = append(a.items, 55)
	a.items = append(a.items, 77)
	a.Display()

	a.InsertionSort()
	fmt.Println("after")
	a.Display()
}
