package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Array struct {
	items []int
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

// Projects 3.2
func (a *Array) Median() int {
	a.InsertionSort()
	half := len(a.items) / 2
	return a.items[half-1]
}

func InsertionSortRun() {
	i := []int{77, 99, 11, 13, 44, 12, 12, 55, 00, 77}
	a := Array{items: i}

	a.Display()

	a.InsertionSort()
	fmt.Println("after")
	a.Display()
}

func SelectionSortRun() {
	i := []int{77, 99, 11, 13, 44, 12, 12, 55, 00, 77}
	a := Array{items: i}

	a.Display()

	a.SelectionSort()
	fmt.Println("after")
	a.Display()
}

func random(max int) int {
	return rand.Intn(max)
}

// Experiment 1
func TenThousandSort() {
	var a []int
	for i := 0; i < 10000; i++ {
		a = append(a, random(10000))
	}

	b := Array{items: a}
	c := Array{items: a}
	d := Array{items: a}

	bNow := time.Now()
	b.BubbleSort()
	bDur := time.Since(bNow)
	fmt.Printf("BubbleSort: %v\n", bDur)

	cNow := time.Now()
	c.InsertionSort()
	cDur := time.Since(cNow)
	fmt.Printf("InsertionSort: %v\n", cDur)

	dNow := time.Now()
	d.SelectionSort()
	dDur := time.Since(dNow)
	fmt.Printf("SelectionSort: %v\n", dDur)
}

// Experiment 1
// BubbleSort: 19.644929445s
// InsertionSort: 590.368µs
// SelectionSort: 6.201012026s
func HundredThousandSort() {
	var a []int
	for i := 0; i < 100000; i++ {
		a = append(a, random(100000))
	}

	b := Array{items: a}
	c := Array{items: a}
	d := Array{items: a}

	bNow := time.Now()
	b.BubbleSort()
	bDur := time.Since(bNow)
	fmt.Printf("BubbleSort: %v\n", bDur)

	cNow := time.Now()
	c.InsertionSort()
	cDur := time.Since(cNow)
	fmt.Printf("InsertionSort: %v\n", cDur)

	dNow := time.Now()
	d.SelectionSort()
	dDur := time.Since(dNow)
	fmt.Printf("SelectionSort: %v\n", dDur)
}

// Experiment 2
// BubbleSort: 11.882628593s
// InsertionSort: 205.563µs
// SelectionSort: 6.156577306s
func ReverseHundredThousandSort() {
	var a []int
	for i := 100000; i >= 0; i-- {
		a = append(a, i)
	}

	b := Array{items: a}
	c := Array{items: a}
	d := Array{items: a}

	bNow := time.Now()
	b.BubbleSort()
	bDur := time.Since(bNow)
	fmt.Printf("BubbleSort: %v\n", bDur)

	cNow := time.Now()
	c.InsertionSort()
	cDur := time.Since(cNow)
	fmt.Printf("InsertionSort: %v\n", cDur)

	dNow := time.Now()
	d.SelectionSort()
	dDur := time.Since(dNow)
	fmt.Printf("SelectionSort: %v\n", dDur)
}

// SortedHundredThousandSort
// Experiment 3
// BubbleSort: 4.51892939s
// InsertionSort: 196.94µs
// SelectionSort: 6.198048522s
func SortedHundredThousandSort() {
	var a []int
	for i := 0; i <= 100000; i++ {
		a = append(a, i)
	}

	b := Array{items: a}
	c := Array{items: a}
	d := Array{items: a}

	bNow := time.Now()
	b.BubbleSort()
	bDur := time.Since(bNow)
	fmt.Printf("BubbleSort: %v\n", bDur)

	cNow := time.Now()
	c.InsertionSort()
	cDur := time.Since(cNow)
	fmt.Printf("InsertionSort: %v\n", cDur)

	dNow := time.Now()
	d.SelectionSort()
	dDur := time.Since(dNow)
	fmt.Printf("SelectionSort: %v\n", dDur)
}
