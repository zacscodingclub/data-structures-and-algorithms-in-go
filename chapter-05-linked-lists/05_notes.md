# Linked Lists

## Simple Linked List

1. Insert
2. Delete
3. Iterate

## Linked List Efficiency

- Insertion and deletion at the beginning are O(1)
- Find, delete, or insert next to a specific item requires searching through O(n)
- Uses only the exact amount of memory needed and can expand to fill all of available memory.

## Abstract Data Types

A way of looking at a data structre focusing on what it does and ignoring how it does its job.

## Sorted Lists

- Insertion and deletion require O(N) worst case with O(N/2) average case.
- Minimum can be deleted in O(1)
- Priority queue may be implemented as a sorted list

## Doubly Linked Lists

Each link knows the next AND the previous links. Practical application is a text editor where input is stored as a linked list. If it is a singly linked and a user attempts to go backwards in text, the editor will have to start at the beginning and find the location that matches the cursor. Doubly Linked list makes it easy to navigate backwards. Could also be used as the ADT for a deque, since it requires insert & delete at both ends.

## Iterators

- A reference, encapsulated in a class that points to a link in an associated list
- Methods allow the user to move the iterator along the list an access the link currently pointed to.

## Questions

1. ~~c~~ b
2. ~~current~~ first
3. d
4. 2
5. 1
6. c
7. `current.next = nil`
8. garbage collection
9. a
10. ~~double-ended (?)~~ empty
11. linked list
12. ~~Twice
13. Doubly Ended List
14. b
15. singly linked list
