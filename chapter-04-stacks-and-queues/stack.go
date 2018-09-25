package main

import "fmt"

type Stack struct {
	maxSize int
	items   []int
	top     int
}

func NewStack(m int) *Stack {
	items := make([]int, m, m)

	return &Stack{
		maxSize: m,
		items:   items,
		top:     0,
	}
}

func (s *Stack) Push(i int) {
	s.top++
	s.items[s.top] = i
}

func (s *Stack) Pop() int {
	s.top--
	return s.items[s.top+1]
}

func (s *Stack) Peek() int {
	return s.items[s.top]
}

func (s *Stack) IsEmpty() bool {
	return s.top == 0
}

func (s *Stack) Size() int {
	return s.top + 1
}

func (s *Stack) IsFull() bool {
	return s.top == s.maxSize-1
}

func RunStack() {
	s := NewStack(10)
	s.Push(20)
	s.Push(40)
	s.Push(60)
	s.Push(80)

	for !s.IsEmpty() {
		v := s.Pop()
		fmt.Printf("Current Value: %v \n", v)
	}

	fmt.Println("Fin")
}
