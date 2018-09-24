package main

import "fmt"

type PriorityQueue struct {
	maxSize int
	items   []int
	nItems  int
}

func RunPriorityQueueExample() {
	q := NewPriorityQueue(10)
	q.Insert(30)
	q.Insert(50)
	q.Insert(10)
	q.Insert(40)
	q.Insert(20)

	for !q.IsEmpty() {
		i := q.Remove()
		fmt.Printf("Number removed from queue: %v\n", i)
	}

}

func NewPriorityQueue(m int) *PriorityQueue {
	maxSize := m + 1
	items := make([]int, maxSize, maxSize)

	return &PriorityQueue{
		maxSize: maxSize,
		items:   items,
		nItems:  0,
	}
}

func (q *PriorityQueue) Insert(item int) {
	var j int
	if q.nItems == 0 {
		q.items[q.nItems] = item
		q.nItems++
	} else {
		for j = q.nItems - 1; j >= 0; j-- {
			if item > q.items[j] {
				q.items[j+1] = q.items[j]
			} else {
				break
			}
		}
		q.items[j+1] = item
		q.nItems++
	}
}

func (q *PriorityQueue) Remove() int {
	q.nItems--
	return q.items[q.nItems]
}

func (q *PriorityQueue) PeekMin() int {
	return q.items[q.nItems-1]
}

func (q *PriorityQueue) IsEmpty() bool {
	return q.nItems == 0
}

func (q *PriorityQueue) IsFull() bool {
	return q.nItems == q.maxSize
}
