package main

type SortedList struct {
	*LinkedList
}

func NewSortedList() *SortedList {
	return &SortedList{
		LinkedList: NewLinkedList(),
	}
}

func (sl *SortedList) insert(i int, f float64) {
	var prev *Link
	nl := NewLink(i, f)
	current := sl.first

	for current != nil && f > current.dData {
		prev = current
		current = current.next
	}
	if prev == nil {
		sl.first = nl
	} else {
		prev.next = nl
	}
	nl.next = current
}
