package main

import "fmt"

func main() {
	t := triangle(10)
	fmt.Println(t)
}

func pow(base, exponent int) int {
	if exponent == 1 {
		return base
	}
	exponent--
	return base * pow(base, exponent)
}

func combinationUtil(a, d []int, start, end, index, k int) {
	if index == k {
		for j := 0; j < k; j++ {
			fmt.Printf("%d ", d[j])
		}
		fmt.Println("")
	}
	for i := start; i <= end && end-i+1 >= k-index; i++ {
		d[index] = a[i]
		combinationUtil(a, d, i+1, end, index+1, k)
	}
}

func buildArr(n int) []int {
	var tmp []int
	for i := 1; i < n+1; i++ {
		tmp = append(tmp, i)
	}
	return tmp
}

func combinations(n, k int) {
	arr := buildArr(n)
	data := make([]int, n)
	combinationUtil(arr, data, 0, n-1, 0, k)
}

func sumItems(items []int) int {
	n := 0
	for _, item := range items {
		n += item
	}
	return n
}

func knapsack(target int, items []int) int {
	newTarget := sumItems(items)
	if target == newTarget {
		return target
	} else {
		knapsack(newTarget, items[1:])
	}
	return -1
}

// (n^2 + n)/2
func triangle(n int) int {
	fmt.Printf("Entering %d\n", n)
	if n == 1 {
		fmt.Printf("Returning 1\n")
		return 1
	}
	t := n + triangle(n-1)
	fmt.Printf("Returning %d\n", t)
	return t
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

type params struct {
	n             int
	returnAddress int
}

func NewParams(n, r int) params {
	return params{
		n:             n,
		returnAddress: r,
	}
}

type stack struct {
	maxSize int
	top     int
	items   []params
}

func NewStack(size int) stack {
	items := make([]params, size, size)
	return stack{
		maxSize: size,
		top:     -1,
		items:   items,
	}
}

func (s *stack) display(st string) {
	fmt.Printf(" Top: %d, Items: %v%v\n", s.top, s.items, st)
}

func (s *stack) isEmpty() bool {
	return s.top == -1
}

func (s *stack) push(p params) {
	s.top++
	s.items[s.top] = p
}

func (s *stack) pop() params {
	s.top--
	return s.items[s.top+1]
}

func (s *stack) peek() params {
	return s.items[s.top]
}

type stackTriangleApp struct {
	theNumber, theAnswer, codePart int
	theseParams                    params
	theStack                       stack
}

func (s *stackTriangleApp) run(n int) {
	s.theNumber = n
	s.theStack = NewStack(n)
	s.codePart = 1
	s.triangle()

	fmt.Printf("Triangle=%d\n", s.theAnswer)
}

func (s *stackTriangleApp) triangle() {
	for s.theNumber > 0 {
		newParams := NewParams(s.theNumber, 7)
		s.theStack.push(newParams)
		s.theNumber--
	}
	for !s.theStack.isEmpty() {
		newN := s.theStack.pop()
		s.theAnswer += newN.n
	}
}

func (s *stackTriangleApp) recursiveTriangle() {
	shouldStop := false
	for !shouldStop {
		shouldStop = s.step()
	}
}

func (s *stackTriangleApp) step() bool {
	switch s.codePart {
	case 1:
		s.theseParams = NewParams(s.theNumber, 6)
		s.theStack.push(s.theseParams)
		s.codePart = 2
	case 2:
		s.theseParams = s.theStack.peek()
		if s.theseParams.n == 1 {
			s.theAnswer = 1
			s.codePart = 5
		} else {
			s.codePart = 3
		}
	case 3:
		newParams := NewParams(s.theseParams.n-1, 4)
		s.theStack.push(newParams)
		s.codePart = 2
	case 4:
		s.theseParams = s.theStack.peek()
		s.theAnswer = s.theAnswer + s.theseParams.n
		s.codePart = 5
	case 5:
		s.theseParams = s.theStack.peek()
		s.codePart = s.theseParams.returnAddress
		s.theStack.pop()
	case 6:
		return true
	}
	return false
}
