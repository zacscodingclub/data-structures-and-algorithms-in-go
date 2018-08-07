# Simple Sorting

- Since sorting is so important and potentially time-consuming, it has been the subject of extensive research in computer science. This chapter will discuss some of the slower, more naive algorithms because they are easier to understand and can actually be better for specific use cases.

1.  Bubble Sort

    - O(n^2)
    - Notriously slow, but simplest to consider.
    - Until we have reached the end of the list, compare each element to the one to it's right and if left element is larger, swap the elements. This can visually appear like a bubble moving up to the surface.
    - Invariants: items on the right of the outer loop will always be sorted

2.  Selection Sort

    - O(n^2)
    - For large datasets it can be significantly faster than other n^2 algorithms.
    - Minimizes swaps but still has high number of comparisons.
    - Pass through all elements and select the smallest. Swap that smallest with the element on the left hand side of the array. Increment one spot to the right, find the smallest element, then swap those elements.
    - Invariants: items on the left of the outer loop will always be sorted

3.  Insertion Sort
    - O(n^2)
    - About twice as fast as bubble and slightly faster than selection for most datasets. If data is partially sorted, then it can be really fast.
    - Invariants: At the end of each pass, following the insertion of the item from tmp, the data items with smaller indices than outer are partially sorted.

## Questions

1.  ~~c~~ d
2.  comparing, swapping
3.  false
4.  a
5.  false
6.  b
7.  ~~true~~ false
8.  ~~2~~ 3
9.  items <= outer loop var will always be sorted
10. tmp
11. ~~a~~ d
12. copies
13. b
14. see notes above
15. ~~d~~ b
