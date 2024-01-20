package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) Push(next int) {
	s.top = &Node{
		Value: next,
		Next:  s.top,
	}
}

func (s *Stack) Pop() *int {
	var ans *int
	top := s.top
	if top == nil {
		return ans
	} else {
		s.top = top.Next
		return &top.Value
	}
}

func isValid(parenth []int) bool {
	var stack Stack
	for _, term := range parenth {
		if term == 0 {
            stack.Push(term)
		} else {
            last := stack.Pop()
            if last == nil {
                return false
            }
		}
	}
	return stack.Pop() == nil
}

func main() {
	x := []int{0, 0, 1, 1}
	fmt.Println(isValid(x))
	x = []int{0, 1, 1}
	fmt.Println(isValid(x))
	x = []int{0, 0, 0, 1}
	fmt.Println(isValid(x))
	x = []int{}
	fmt.Println(isValid(x))
}
