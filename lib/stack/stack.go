package stack

import "fmt"

type Stack struct {
	stack []string
}

func (s *Stack) Init() *Stack {
	s.stack = []string{}
	return s
}

func (s *Stack) Head() string {
	return s.stack[0]
}
func (s *Stack) Unshift(char string) {
	s.stack = append([]string{char}, s.stack...)
}
func (s *Stack) Pop() string {
	head := s.stack[0]
	s.stack = s.stack[1:len(s.stack)]
	return head
}
func (s *Stack) Print() {
	fmt.Print(s.stack)
}

func Init() *Stack {
	return new(Stack).Init()
}
