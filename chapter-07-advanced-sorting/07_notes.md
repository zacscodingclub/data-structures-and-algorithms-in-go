# Advanced Sorting

## Shellsort

- O(n\*log(n)^2)
- Based on Insertion sort, but improves upon performance
- Works well with medium-sized arrays, up to a few thousand items
- Many experts suggest starting with a shellsort, switch only if performance is too slow in practice
- Sorts by intervals quickly interleaving themselves over time
- common interval sequence: `h = h*3 + 1`
- Only hard requirement is that the diminishing sequence must end with 1, so that the last pass is a fast insertion sort (since it's partially sorted)

## Partitioning

- O(n)
- Underlying mechanism of quicksort
- Dividing data into two groups, one greater and one less than the pivot point
  Algorithm:
- Two indices, one at each end of the array. `leftPtr` starts on the left moves to the right, and `rightPtr` starts on the right and moves left
- Stopping and Swapping:
  - When `leftPtr` encounters an item smaller than the pivot, it keeps going because that value is already on the correct side of the array
  - When `leftPtr` encounters an item larger than the pivot, it stops
  - Similarly, but opposite, for `rightPtr` if value is smaller than pivot
  - When both of these have stopped, we swap!
  - Repeat process until the left and right meet. Once that happens the sort has been completed.

## Quicksort

- O(n\*log(n))
- Most popular sorting algorithm since it is fastest in the majority of situations (Specifically in-memory sorting, sorting data on disk may prefer other algorithms)
- Can be inefficient with smaller arrays and there may be clauses to use insertion sort if below a certain size.
- Another option: quicksort then insertion sort. Since it is partially sorted after initial quicksort, the insertion sort is very effecient.

### Pivot Point

- Should be the key value of an actual data item
- You can pick it more or less at random
- After partition, the pivot will be at the boundary of the left and right subarrays, thus in it's final sorted position.
- Median of Three: Take median value of first, middle, and last elements and use that as the pivot point. Swap the elements on first look so they are already in the correct place and the search can start left+1 and right-1.

## Radix Sort

- O(n\*log(n))
- Disassembles the key into digits and arranges the data items according to their value of digits without needing comparisons.
- Typically implemented in base 2 to take advantage of bit manipulation

### Algorithm

1. Divide all data items into 10 groups based on their 1s digit
2. These 10 groups are reassembled. All keys ending with 0 go first, followed by 1, etc. e.g. sub-sort
3. Divide all data into 10 groups based on their 10s digit. This must be done without changing the order of the previous sort.
4. Again the 10 groups are sub-sorted
5. This process is repeated for the remaining digits.

## Questions

1. ~~a~~ c
2. 40
3. d
4. false
5. O(n\*log(n)), O(n^2)
6. ~~c~~ a
7. pivot
8. ~~a~~ d
9. True
10. ~~b~~ c
11. sorting each partition
12. ~~c?~~ b
13. pivot
14. log base2 (N)
15. True
