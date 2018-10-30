# Advanced Sorting

## Shellsort

- O(n\*log(n)^2)
- Based on Insertion sort, but improves upon performance
- Works well with medium-sized arrays, up to a few thousand items
- Many experts suggest starting with a shellsort, switch only if performance is too slow in practice
- Sorts by intervals quickly interleaving themselves over time
- common interval sequence: `h = h*3 + 1`
- Only hard requirement is that the diminishing sequence must end with 1, so that the last pass is a fast insertion sort (since it's partially sorted)

## Quicksort

- O(n\*log(n))

### Partitioning

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

## Radix Sort
