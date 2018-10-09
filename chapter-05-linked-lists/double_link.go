package main

import "fmt"

type DoubleLink struct {
	data int
	next *DoubleLink
	prev *DoubleLink
}

func NewDoubleLink(i int) *DoubleLink {
	return &DoubleLink{
		data: i,
		next: nil,
		prev: nil,
	}
}

func (l *DoubleLink) display() {
	fmt.Printf("%d, ", l.data)
}

type DoublyLinkedList struct {
	first *DoubleLink
	last  *DoubleLink
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		first: nil,
		last:  nil,
	}
}

func (d *DoublyLinkedList) isEmpty() bool {
	return d.first == nil
}

func (d *DoublyLinkedList) insertFirst(i int) {
	nl := NewDoubleLink(i)

	if d.isEmpty() {
		d.last = nl
	} else {
		d.first.prev = nl
	}
	nl.next = d.first
	d.first = nl
}

func (d *DoublyLinkedList) insertLast(i int) {
	nl := NewDoubleLink(i)

	if d.isEmpty() {
		d.first = nl
	} else {
		d.last.next = nl
		nl.prev = d.last
	}
	d.last = nl
}

func (d *DoublyLinkedList) deleteFirst() *DoubleLink {
	tmp := d.first
	if d.first.next == nil {
		d.last = nil
	} else {
		d.first.next.prev = nil
	}
	d.first = d.first.next
	return tmp
}

func (d *DoublyLinkedList) deleteLast() *DoubleLink {
	tmp := d.last
	if d.first.next == nil {
		d.first = nil
	} else {
		d.last.prev.next = nil
	}
	d.last = d.last.prev
	return tmp
}

func (d *DoublyLinkedList) insertAfter(key, i int) bool {
	current := d.first

	for current.data != key {
		current = current.next
		if current == nil {
			return false
		}
	}
	nl := NewDoubleLink(i)
	if current == d.last {
		nl.next = nil
		d.last = nl
	} else {
		nl.next = current.next
		current.next.prev = nl
	}
	nl.prev = current
	current.next = nl
	return true
}

func (d *DoublyLinkedList) deleteKey(key int) *DoubleLink {
	current := d.first

	for current.data != key {
		current = current.next
		if current == nil {
			return nil
		}
	}

	if current == d.first {
		d.first = current.next
	} else {
		current.prev.next = current.next
	}

	if current == d.last {
		d.last = current.prev
	} else {
		current.next.prev = current.prev
	}
	return current
}

func (d *DoublyLinkedList) displayForward() {
	fmt.Printf("List (first-->last): ")
	current := d.first

	for current != nil {
		current.display()
		current = current.next
	}
	fmt.Println(")")
}

func (d *DoublyLinkedList) displayBackward() {
	fmt.Printf("List (last-->first): ")
	current := d.last

	for current != nil {
		current.display()
		current = current.prev
	}
	fmt.Println(")")
}
