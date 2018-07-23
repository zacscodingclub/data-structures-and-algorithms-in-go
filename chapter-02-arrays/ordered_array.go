package main

import "fmt"

type OrderedArray struct {
	values []int
	length int
}

func NewOrderedArray(size int) *OrderedArray {
	a := make([]int, size, size)
	arr := OrderedArray{values: a}
	return &arr
}

func MergeOrderedArrays(one, two *OrderedArray) *OrderedArray {
	a := make([]int, one.Size()+two.Size(), one.Size()+two.Size())
	n := OrderedArray{values: a}

	for one.Size() > 0 && two.Size() > 0 {
		o, t := one.values[0], two.values[0]
		if o < t {
			n.Insert(o)
			one.Delete(o)
		} else {
			n.Insert(t)
			two.Delete(t)
		}
	}

	for one.Size() > 0 {
		t := one.values[0]
		n.Insert(t)
		one.Delete(t)
	}

	for two.Size() > 0 {
		t := two.values[0]
		n.Insert(t)
		two.Delete(t)
	}

	return &n
}

func (a *OrderedArray) bsearch(key int) int {
	if a.Size() == 0 {
		return 0
	}
	lower := 0
	upper := a.Size() - 1
	var current int
	found := false
	for !found {
		current = (lower + upper) / 2
		if a.values[current] == key {
			break
		} else if lower > upper {
			current = a.Size()
			break
		} else {
			if a.values[current] < key {
				lower = current + 1
			} else {
				upper = current - 1
			}
		}
	}
	return current
}

func (a *OrderedArray) Insert(key int) {
	j := a.bsearch(key)
	if j == a.Size() {
		for j = 0; j < a.Size(); j++ {
			if a.values[j] > key {
				break
			}
		}
	}
	for k := a.Size(); k > j; k-- {
		a.values[k] = a.values[k-1]
	}
	a.values[j] = key
	a.length += 1
}

func (a *OrderedArray) Find(key int) int {
	f := a.bsearch(key)
	if f == a.Size() {
		return -1
	}
	return f
}

func (a *OrderedArray) Delete(key int) bool {
	f := a.bsearch(key)
	if f == a.Size() {
		return false
	}
	for k := f; k < a.Size(); k++ {
		a.values[k] = a.values[k+1]
	}
	a.length -= 1
	return true
}

func (a *OrderedArray) Display() {
	for i := 0; i < a.Size(); i++ {
		fmt.Printf("%v\n", a.values[i])
	}
}

func (a *OrderedArray) Size() int {
	return a.length
}

func OrderedArrayRun() {
	a := NewOrderedArray(100)
	b := NewOrderedArray(100)

	a.Insert(99)
	a.Insert(55)
	a.Insert(44)
	a.Insert(77)
	a.Insert(12)
	a.Display()

	b.Insert(1)
	b.Insert(2)
	b.Insert(98)
	b.Insert(3)
	b.Display()

	searchKeyOne := 36
	f := a.Find(searchKeyOne)
	if f != -1 {
		fmt.Printf("Found: %v\n", searchKeyOne)
	} else {
		fmt.Printf("Didn't find: %v\n", searchKeyOne)
	}
	searchKeyTwo := 12
	f2 := a.bsearch(searchKeyTwo)
	if f2 != -1 {
		fmt.Printf("Found: %v\n", searchKeyTwo)
	} else {
		fmt.Printf("Didn't find: %v\n", searchKeyTwo)
	}

	c := MergeOrderedArrays(a, b)
	c.Display()
}
