# Recursion

Programming technique in which a method calls itself. It is often not inherently more efficient, but that it simplifies the problem conceptually.

## Characteristics of Recursive Methods

1. It calls itself.
2. When it calls itself, it does so to solve a smaller problem.
3. There is a "base case" that is simple enough that the routine can solve it, and return, without calling itself.

## Divide and Conquer Algorithms

Separate your problem into two separate problems and solve them independently. The solution to each smaller problem is the same: divide it into two even smaller problems and solve them. Continue this process until you get to the base case, which can be solved easily with no further division into halves.

## Mergesort

- N\*log(N)
- requires an additional array in memory, equal in size to the one being sorted
- Number of copies is N\*log2(N)
- Number of comparisons is always somewhat less than the number of copies (Max = N-1, Min = N/2)

## Elminiating Recursion

In many cases, removing recursion allows you to implement more effecient methods. However, many divide and conquer algorithms work very well in recursive routines.

While recursion may be easy to conceptualize, it is often inefficient. In these cases we can transform the recursive method into a non-recursive approach.
