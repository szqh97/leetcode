package main

import "fmt"

type Stack struct {
	dlist  []byte
	length int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(c byte) {
	s.dlist = append(s.dlist[:s.length], c)
	s.length = s.length + 1

}
func (s *Stack) Pop() byte {
	c := s.dlist[s.length-1]
	s.length = s.length - 1
	return c
}
func (s *Stack) Len() int {
	return s.length
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 == 1 {
		return false
	}
	stack := NewStack()
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack.Push(s[i])
		} else if stack.Len() > 0 {
			c := stack.Pop()
			switch c {
			case '(':
				if s[i] != ')' {
					return false
				}
			case '[':
				if s[i] != ']' {
					return false
				}
			case '{':
				if s[i] != '}' {
					return false
				}
			}
		} else {
			return false
		}
	}
	return stack.Len() == 0
}
func main() {
	s := "(("
	s = "(("
	fmt.Println(isValid(s))
}
