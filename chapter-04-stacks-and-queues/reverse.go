package main

import (
	"bytes"
	"fmt"
)

func ReverseString() {
	input := getInput("Enter a word to be reversed:")
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
