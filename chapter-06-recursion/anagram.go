package main

import "fmt"

type Anagram struct {
	size       int
	count      int
	characters []byte
}

func NewAnagram(s []byte) *Anagram {
	return &Anagram{
		size:       len(s),
		count:      0,
		characters: s,
	}
}

func (a *Anagram) doAnagram(n int) {
	if n == 1 {
		return
	}
	for j := 0; j < n; j++ {
		a.doAnagram(n - 1)
		if n == 2 {
			a.displayWord()
		}
		a.rotate(n)
	}
}

func (a *Anagram) rotate(n int) {
	var j int
	pos := a.size - n
	tmp := a.characters[pos]
	for j = pos + 1; j < a.size; j++ {
		a.characters[j-1] = a.characters[j]
	}
	a.characters[j-1] = tmp
}

func (a *Anagram) displayWord() {
	if a.count < 99 {
		fmt.Printf(" ")
	}
	if a.count < 9 {
		fmt.Printf(" ")
	}

	a.count++
	fmt.Printf("%d ", a.count)
	for j := 0; j < a.size; j++ {
		fmt.Printf("%s", string(a.characters[j]))
	}
	fmt.Println("")
	if a.count%6 == 0 {
		fmt.Println("")
	}

}
