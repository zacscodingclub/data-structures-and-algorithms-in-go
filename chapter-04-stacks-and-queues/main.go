package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	RunQueueExample()
}

func getInput(s string) []rune {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(s)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Failed to convert input to string: %v", err)
	}
	text = strings.Replace(text, "\n", "", -1)
	return []rune(text)
}
