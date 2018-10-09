package main

func main() {
	dl := NewDoublyLinkedList()
	dl.insertFirst(22)
	dl.insertFirst(44)
	dl.insertFirst(66)

	dl.insertLast(11)
	dl.insertLast(33)
	dl.insertLast(55)

	dl.displayForward()
	dl.displayBackward()

	dl.deleteFirst()
	dl.deleteLast()
	dl.deleteKey(11)

	dl.displayForward()
	dl.displayBackward()

	dl.insertAfter(22, 77)
	dl.insertAfter(33, 88)

	dl.displayForward()
}
