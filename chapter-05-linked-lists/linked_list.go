package main

import "fmt"

type LinkedList struct {
	first *Link
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		first: nil,
	}
}

func (ll *LinkedList) display() {
	fmt.Printf("List (first-->last: ")
	current := ll.first
	for current != nil {
		current.display()
		current = current.next
	}
	fmt.Println(" )")
}

func (ll *LinkedList) getFirst() *Link {
	return ll.first
}

func (ll *LinkedList) setFirst(l *Link) {
	ll.find = l
}
func (ll *LinkedList) isEmpty() bool {
	return ll.first == nil
}

func (ll *LinkedList) insertFirst(i int, f float64) {
	link := NewLink(i, f)
	link.next = ll.first
	ll.first = link
}

func (ll *LinkedList) deleteFirst() *Link {
	tmp := ll.first
	ll.first = ll.first.next
	return tmp
}

func (ll *LinkedList) find(k int) *Link {
	current := ll.first
	for current.iData != k {
		if current.next == nil {
			return nil
		}
		current = current.next
	}
	return current
}

func (ll *LinkedList) delete(k int) *Link {
	current := ll.first
	previous := ll.first
	for current.iData != k {
		if current.next == nil {
			return nil
		}
		previous = current
		current = current.next
	}

	if current == ll.first {
		ll.first = ll.first.next
	} else {
		previous.next = current.next
	}
	return current
}
