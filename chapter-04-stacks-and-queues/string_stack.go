package main

type StringStack struct {
	maxSize int
	top     int
	items   []string
}

func NewStringStack(m int) *StringStack {
	items := make([]string, m, m)

	return &StringStack{
		maxSize: m,
		top:     0,
		items:   items,
	}
}

func (ss *StringStack) Push(s string) {
	ss.items[ss.top] = s
	ss.top += 1
}

func (ss *StringStack) Pop() string {
	ss.top -= 1
	return ss.items[ss.top]
}

func (ss *StringStack) Peek() string {
	return ss.items[ss.top]
}

func (ss *StringStack) IsEmpty() bool {
	return ss.top == 0
}

func (ss *StringStack) IsFull() bool {
	return ss.top == ss.maxSize-1
}
