package main

import "fmt"

type Queue struct {
	maxSize int
	items   []int
	front   int
	rear    int
	nItems  int
}

func RunQueueExample() {
	q := NewQueue(10)
	q.Insert(10)
	q.Insert(20)
	q.Insert(30)
	q.Insert(40)
	q.Remove()
	q.Remove()
	q.Remove()
	q.Insert(50)
	q.Insert(60)

	for !q.IsEmpty() {
		n := q.Remove()
		fmt.Printf("%v\n", n)
	}

}

func NewQueue(m int) *Queue {
	maxSize := m + 1
	items := make([]int, maxSize, maxSize)

	return &Queue{
		maxSize: maxSize,
		items:   items,
		front:   0,
		rear:    -1,
	}
}

func (q *Queue) Insert(j int) {
	if q.rear == q.maxSize-1 {
		q.rear = -1
	}
	q.rear++
	q.items[q.rear] = j
}

func (q *Queue) Remove() int {
	n := q.items[q.front]
	q.front++
	if q.front == q.maxSize {
		q.front = 0
	}
	return n
}

func (q *Queue) PeekFront() int {
	return q.items[q.front]
}

func (q *Queue) IsEmpty() bool {
	return (q.rear+1 == q.front) || (q.front+q.maxSize-1 == q.rear)
}

func (q *Queue) IsFull() bool {
	return (q.rear+2 == q.front) || (q.front+q.maxSize-2 == q.rear)
}

func (q *Queue) Size() int {
	if q.rear >= q.front {
		return q.rear - q.front + 1
	} else {
		return q.maxSize - q.front + q.rear + 1
	}
}
