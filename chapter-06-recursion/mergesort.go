package main

import "fmt"

func RunMergesort() {
	a := []int{23, 47, 81, 95}
	b := []int{7, 14, 39, 55, 62, 74}

	c := mergeIterate(a, b)
	display(c)
}

func RunRecursiveMergeSort() {
	d := NewDArray(20)

	d.insert(64)
	d.insert(21)
	d.insert(33)
	d.insert(70)
	d.insert(12)
	d.insert(85)
	d.insert(44)
	d.insert(3)
	d.insert(99)
	d.insert(0)
	d.insert(109)
	d.insert(36)
	d.insert(100)
	d.insert(2)
	d.insert(108)
	d.insert(111)
	d.insert(88)
	d.insert(55)
	d.insert(77)
	d.insert(69)

	d.display()
	d.mergeSort()
	d.display()
}

type dArray struct {
	size   int
	max    int
	values []int
}

func NewDArray(size int) *dArray {
	v := make([]int, size)
	return &dArray{
		size:   0,
		max:    size,
		values: v,
	}
}

func (d *dArray) insert(i int) {
	d.values[d.size] = i
	d.size++
}

func (d *dArray) display() {
	for _, n := range d.values {
		fmt.Printf("%v ", n)
	}
	fmt.Println()
}

func (d *dArray) mergeSort() {
	space := make([]int, d.max)
	d.recMergeSort(space, 0, d.size-1)
}

func (d *dArray) recMergeSort(space []int, lower, upper int) {
	if lower < upper {
		mid := (lower + upper) / 2
		d.recMergeSort(space, lower, mid)
		d.recMergeSort(space, mid+1, upper)
		d.merge(space, lower, mid+1, upper)
	}
}

func (d *dArray) merge(space []int, low, high, up int) {
	func() {
		j := 0
		mid := high - 1

		for low <= mid && high <= up {
			fmt.Printf("Comparing %d to %d\n", d.values[low], d.values[high])
			if d.values[low] < d.values[high] {
				fmt.Printf("%d is less than %d\n", d.values[low], d.values[high])
				fmt.Printf("Adding %d at %d\n", d.values[low], j)
				space[j] = d.values[low]
				j++
				low++
				fmt.Println(space)
				fmt.Printf("j up to: %d, low: %d\n", j, low)
			} else {
				fmt.Printf("%d is greater than than %d\n", d.values[low], d.values[high])
				fmt.Printf("Adding %d at %d\n", d.values[high], j)
				space[j] = d.values[high]
				j++
				high++
				fmt.Println(space)
				fmt.Printf("j up to: %d, low: %d\n", j, low)
			}
		}

		for low <= mid {
			space[j] = d.values[low]
			j++
			low++
		}

		for high <= up {
			space[j] = d.values[high]
			j++
			high++
		}
	}()
	copy(d.values, space)
}

func mergeIterate(a, b []int) []int {
	lenA := len(a)
	lenB := len(b)
	var aDex, bDex int
	var c []int

	for aDex < lenA && bDex < lenB {
		if a[aDex] < b[bDex] {
			c = append(c, a[aDex])
			aDex++
		} else {
			c = append(c, b[bDex])
			bDex++
		}
	}
	for aDex < lenA {
		c = append(c, a[aDex])
		aDex++
	}
	for bDex < lenB {
		c = append(c, b[bDex])
		bDex++

	}
	return c
}

func display(c []int) {
	for _, el := range c {
		fmt.Printf("%d ", el)
	}
	fmt.Println("")
}
