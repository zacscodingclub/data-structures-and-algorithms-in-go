package main

import "fmt"

type Array struct {
	items []int16
}

func (a *Array) Display() {
	for i := 0; i < len(a.items); i++ {
		fmt.Printf("%v\n", a.items[i])
	}
}

// may be more performant to write this in-line(?)
// might be a Java only thing
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

func (a *Array) SelectionSort() {
	for out := 0; out < len(a.items)-1; out++ {
		min := out

		for in := out + 1; in < len(a.items); in++ {
			if a.items[in] < a.items[min] {
				min = in
			}
		}

		a.swap(out, min)
	}
}

func (a *Array) InsertionSort() {
	for out := 1; out < len(a.items); out++ {
		tmp := a.items[out]
		in := out

		for in > 0 && a.items[in-1] >= tmp {
			a.items[in] = a.items[in-1]
			in--
		}
		a.items[in] = tmp
	}
}

func InsertionSortRun() {
	i := []int16{77, 99, 11, 13, 44, 12, 12, 55, 00, 77}
	a := Array{items: i}

	a.Display()

	a.InsertionSort()
	fmt.Println("after")
	a.Display()
}

func SelectionSortRun() {
	i := []int16{77, 99, 11, 13, 44, 12, 12, 55, 00, 77}
	a := Array{items: i}

	a.Display()

	a.SelectionSort()
	fmt.Println("after")
	a.Display()
}
