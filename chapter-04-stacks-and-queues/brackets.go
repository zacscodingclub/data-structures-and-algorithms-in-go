package main

import "fmt"

func CheckBrackets() {
	input := getInput("Enter some code so I can check if the brackets are correct:")
	ss := NewStringStack(len(input))
	failed := false
	for i, char := range input {
		schar := string(char)
		if failed {
			break
		}
		fmt.Printf("Evaluating: %v\n", schar)
		switch schar {
		case "{", "[", "(":
			ss.Push(schar)
			break
		case "}", "]", ")":
			if !ss.IsEmpty() {
				fmt.Println("It was empty")
				fmt.Printf("Error: %v at %v\n", schar, i)
				failed = true
			} else {
				failed = checkRightBracket(ss, schar, i)
				if failed {
					fmt.Printf("Error: %v at %v\n", schar, i)
				}
			}
			break
		default:
			break
		}
	}
	if !ss.IsEmpty() {
		fmt.Println("Not balanced.  Missing a right bracket.")
	} else {
		fmt.Println("Code is balanced!")
	}
}

func checkRightBracket(ss *StringStack, schar string, i int) bool {
	c := ss.Pop()
	if (schar == "}" && c != "{") ||
		(schar == "]" && c != "[") ||
		(schar == ")" && c != "(") {
		return true
	}
	return false
}
