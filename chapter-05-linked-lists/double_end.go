package main

import "fmt"

type DoubleEnd struct {
	first *DLink
	last  *DLink
}

func NewDoubleEndList() *DoubleEnd {
	return &DoubleEnd{
		first: nil,
		last:  nil,
	}
}

func (d *DoubleEnd) isEmpty() bool {
	return d.first == nil
}

func (d *DoubleEnd) insertFirst(dd int) {
	nl := NewDLink(dd)

	if d.isEmpty() {
		d.last = nl
	}
	nl.next = d.first
	d.first = nl
}

func (d *DoubleEnd) insertLast(dd int) {
	nl := NewDLink(dd)

	if d.isEmpty() {
		d.first = nl
	} else {
		d.last.next = nl
	}
	d.last = nl
}

func (d *DoubleEnd) deleteFirst() int {
	tmp := d.first.data
	if d.first.next == nil {
		d.last = nil
	}
	d.first = d.first.next
	return tmp
}

func (d *DoubleEnd) displayList() {
	fmt.Printf("List (first-->last: ")
	current := d.first
	for current != nil {
		current.display()
		current = current.next
	}
	fmt.Println(" )")
}
