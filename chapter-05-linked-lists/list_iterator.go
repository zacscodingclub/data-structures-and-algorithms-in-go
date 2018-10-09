package main

type ListIterator struct {
	current  *DLink
	previous *DLink
	list     *LinkedList
}

func NewListIterator(l *LinkedList) *ListIterator {
	return &ListIterator{
		current:  nil,
		previous: nil,
		list:     l,
	}
}

func (l *ListIterator) reset() {
	l.current = l.list.getFirst()
	l.previous = nil
}

func (l *ListIterator) atEnd() bool {
	return l.current.next == nil
}

func (l *ListIterator) nextLink() {
	l.previous = l.current
	l.current = l.current.next
}

func (l *ListIterator) getCurrent() *DLink {
	return l.current
}

func (l *ListIterator) insertAfter(i int) {
	nl := NewDLink(i)

	if l.list.isEmpty() {
		l.list.setFirst(nl)
		l.current = nl
	} else {
		nl.next = l.current.next
		l.current.next = nl
		l.nextLink()
	}
}

func (l *ListIterator) insertBefore(i int) {
	nl := NewDLink(i)

	if l.previous == nil {
		nl.next = l.list.getFirst()
		l.list.setFirst(nl)
		l.reset()
	} else {
		nl.next = l.previous.next
		l.previous.next = nl
		l.current = nl
	}
}

func (l *ListIterator) deleteCurrent() int {
	v := l.current.data
	if l.previous == nil {
		l.list.setFirst(l.current.next)
		l.reset()
	} else {
		l.previous.next = l.current.next
		if l.atEnd() {
			l.reset()
		} else {
			l.current = l.current.next
		}
	}
	return v
}
