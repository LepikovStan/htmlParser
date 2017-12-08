package stack

import (
	"fmt"
	"../node"
)

type Stack struct {
	stack []*node.Node
}

func (s *Stack) Init() *Stack {
	s.stack = []*node.Node{}
	return s
}

func (s *Stack) Head() *node.Node {
	if len(s.stack) > 0 {
		return s.stack[0]
	} else {
		return nil
	}
}
func (s *Stack) Unshift(nd *node.Node) {
	s.stack = append([]*node.Node{nd}, s.stack...)
}
func (s *Stack) Pop() *node.Node {
	head := s.stack[0]
	s.stack = s.stack[1:len(s.stack)]
	return head
}
func (s *Stack) Print() {
	fmt.Println("stack ->", s.stack)
}

func Init() *Stack {
	return new(Stack).Init()
}
