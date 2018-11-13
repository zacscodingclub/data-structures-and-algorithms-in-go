package main

import (
	"fmt"
	"math/rand"
)

func RunQuicksort() {
	max := 100
	qs := NewQarray(max)

	for i := 0; i < max; i++ {
		qs.Insert(rand.Intn(max * 9))
	}
	fmt.Println("Before:")
	qs.Display()
	qs.QuickSort()
	fmt.Println("After:")
	qs.Display()
}

type qarray struct {
	items []int
	size  int
}

func NewQarray(max int) *qarray {
	items := make([]int, max, max)
	return &qarray{
		items: items,
		size:  0,
	}
}

func (q *qarray) Insert(i int) {
	q.items[q.size] = i
	q.size++
}

func (q *qarray) Pop() int {
	tmp := q.items[q.size]
	q.size--
	return tmp
}

func (q *qarray) Display() {
	fmt.Printf("items: ")
	for _, el := range q.items {
		fmt.Printf("%d ", el)
	}
	fmt.Println("")
}

func (q *qarray) QuickSort() {
	q.recQuickSort(0, q.size-1)
}

func (q *qarray) recQuickSort(left, right int) {
	if right-left <= 0 {
		return
	}
	if left < right {
		partition := q.partitionIt(left, right)

		q.recQuickSort(left, partition-1)
		q.recQuickSort(partition+1, right)
	}
}

func (q *qarray) partitionIt(left, right int) int {
	pivot := q.items[right]
	i := left - 1
	for j := left; j < right; j++ {
		if q.items[j] <= pivot {
			i++
			q.swap(i, j)
		}
	}
	q.swap(i+1, right)
	return i + 1
}

func (q *qarray) swap(one, two int) {
	tmp := q.items[one]
	q.items[one] = q.items[two]
	q.items[two] = tmp
}
