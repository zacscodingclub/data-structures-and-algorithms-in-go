package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

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

func getInput() []rune {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a word to be reversed:")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Failed to convert input to string: %v", err)
	}
	text = strings.Replace(text, "\n", "", -1)
	return []rune(text)
}

func ReverseString() {
	input := getInput()
	ss := NewStringStack(len(input))
	for _, char := range input {
		ss.Push(string(char))
	}

	var b bytes.Buffer
	for !ss.IsEmpty() {
		b.WriteString(ss.Pop())
	}
	fmt.Println(b.String())
}
