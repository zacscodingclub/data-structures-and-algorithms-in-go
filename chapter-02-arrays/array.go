package main

import "fmt"

type Array struct {
	values []int
	length int
}

func NewArray(size int) *Array {
	a := make([]int, size, size)
	arr := Array{values: a}
	return &arr
}

func (a *Array) Insert(value int) {
	a.values[a.Size()] = value
	a.length += 1
}

func (a *Array) GetMax() int {
	var j int
	max := -1
	for j = 0; j < a.Size(); j++ {
		if a.values[j] > max {
			max = a.values[j]
		}
	}
	return max
}

func (a *Array) NoDups() []int {
	filtered := []int{}
	found := map[int]int{}

	for i := 0; i < a.Size(); i++ {
		if _, ok := found[a.values[i]]; !ok {
			found[a.values[i]] = 1
			filtered = append(filtered, a.values[i])
		}
	}

	return filtered
}

func (a *Array) RemoveMax() {
	var j, maxIndex int
	max := -1
	for j = 0; j < a.Size(); j++ {
		if a.values[j] > max {
			max = a.values[j]
			maxIndex = j
		}
	}
	if max != -1 {
		for k := maxIndex; k < a.Size(); k++ {
			a.values[k] = a.values[k+1]
		}
		a.length -= 1
	}

}

func (a *Array) Find(key int) bool {
	var j int
	for j = 0; j < a.Size(); j++ {
		if a.values[j] == key {
			break
		}
	}
	if j == a.Size() {
		return false
	} else {
		return true
	}
}

// Delete removes an element and returns true if it was found
func (a *Array) Delete(key int) bool {
	var j int
	for j = 0; j < a.Size(); j++ {
		if a.values[j] == key {
			break
		}
	}
	if j == a.Size() {
		return false
	} else {
		for k := j; k < a.Size(); k++ {
			a.values[k] = a.values[k+1]
		}
		a.length -= 1

		return true
	}
}

// Display prints out each value on a new line
func (a *Array) Display() {
	for i := 0; i < a.Size(); i++ {
		fmt.Printf("%v\n", a.values[i])
	}
}

// Size is a getter for length attribute
func (a *Array) Size() int {
	return a.length
}

func ArrayRun() {
	a := NewArray(100)

	a.Insert(77)
	a.Insert(99)
	a.Insert(44)
	a.Insert(12)
	a.Insert(12)
	a.Insert(55)
	a.Insert(77)
	a.Display()

	searchKeyOne := 36
	if a.Find(searchKeyOne) {
		fmt.Printf("Found: %v\n", searchKeyOne)
	} else {
		fmt.Printf("Didn't find: %v\n", searchKeyOne)
	}
	searchKeyTwo := 12
	if a.Find(searchKeyTwo) {
		fmt.Printf("Found: %v\n", searchKeyTwo)
	} else {
		fmt.Printf("Didn't find: %v\n", searchKeyTwo)
	}

	a.Display()

	fmt.Printf("%v\n", a.NoDups())
}
