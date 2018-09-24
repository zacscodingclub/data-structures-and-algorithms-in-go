# Stacks and Queues

These are different types of structures that excell at being used as programmer tools, limited access (FIFO, LIFO), and higher level of abstractions.

## Stack

Last in, first out
Stacks allow access to only the last item inserted. Most microprocessors use a stack-based architecture. When a method is called, its return address and arguments are pushed onto a stack, and when it returns, they're popped off.

- Postal Analogy
  - You read through mail once per week and as you receive new letters, you push each onto the stack. When processing you start at the last inserted and work your way in reverse to the first element inserted.
  - The danger is if the stack gets too deep you may not ever make it back to the zeroth positon.

### Actions

1. `push`
2. `pop`
3. `peek`

## Queue

First in, first out
Mimics real life experiences of people waiting in line, airplanes waiting to take off, or data packets waiting to be transimitted over the internet.

### Deqeue

Double-ended queue where you can insert/remove from either side. More versatile data structure than stack or queue and is sometimes used in container class libraries to serve both purposes. However it is not used as often as stacks and queues.

### Actions

1. `insert` (put, add, enque)
2. `remove` (delete, get, deque)

## Priority Queue

More specialized data structure than stack or queue. Like an ordinary queue, it has a front and a rear, and items are removed from the front. However, items are ordered by key-value so that th eitem with the lowest key (or possibly highest key) is always at the front. Items are inserted into the proper position to maintain the order. Typically implemented as a heap to provide quick insertion.

### Mail Sorting Analogy

Each time you receive a piece of mail, you file it into a pile of pending letters according to it's urgency. More important/urgent letters on top. When you have time to take care of your mail, you start at the top and pop off each element.
