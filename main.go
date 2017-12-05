package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
	"time"
	//"./lib/ctrl"
	//"./lib/Node"
	"strings"
)

func readFile(path string) string {
	var result string
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		result = fmt.Sprintf("%s%s", result, scanner.Text())
	}
	return result
}

//type Stack struct {
//	stack []string
//}
//
//func (s *Stack) Head() string {
//	return s.stack[0]
//}
//func (s *Stack) Unshift(char string) {
//	s.stack = append([]string{char}, s.stack...)
//}
//func (s *Stack) Pop() string {
//	head := s.stack[0]
//	s.stack = s.stack[1:len(s.stack)]
//	return head
//}
//func (s *Stack) Print() {
//	fmt.Println(s.stack)
//}
//
//type Tree struct {
//	//tree []Node.Node
//	tree []string
//	word string
//	flag bool
//}
//
//func (t *Tree) getFlag() bool {
//	return t.flag
//}
//func (t *Tree) createTag() {
//	t.word = ""
//	t.flag = true
//}
//func (t *Tree) updateTag(char string) {
//	t.word += char
//}
//func (t *Tree) getTag() string {
//	t.flag = false
//	t.tree = append(t.tree, t.word)
//	return t.word
//}
//func (t *Tree) getTags() []string {
//	return t.tree
//}
//
//func fsm(char string, stack *Stack, tree *Tree) {
//	stack.Print()
//	switch {
//	case char == "/" && stack.Head() == "<":
//		fmt.Println("start ending tag")
//	case char == "<":
//		fmt.Println("start tag")
//		stack.Unshift(char)
//		tree.createTag()
//	case char == ">":
//		fmt.Println("end tag")
//		stack.Pop()
//		tree.getTag()
//	case char != "/" && stack.Head() == "<":
//		fmt.Println("tag -> ", char)
//		tree.updateTag(char)
//	}
//
//}

type FSM struct {
	table map[string]map[string]string
	currentState string
}
func (fsm *FSM) Init() {
	fsm.currentState = "init"
	fsm.table = map[string]map[string]string{
		"init": map[string]string{
			"<": "tagNameCreatingStart",
		},
		"tagNameCreatingStart": map[string]string{
			"*": "tagNameCreating",
		},
		"tagNameCreating": map[string]string{
			" ": "tagAttributeNameCreating",
			"*": "tagNameCreating",
		},
		"tagAttributeNameCreating": map[string]string{
			"=": "tagAttributeValueCreatingStart",
			"*": "tagAttributeNameCreating",
		},
		"tagAttributeValueCreatingStart": map[string]string{
			"*": "tagAttributeValueCreating",
		},
		"tagAttributeValueCreating": map[string]string{
			">": "tagCreatingEnd",
			" ": "tagAttributeNameCreating",
			"*": "tagAttributeValueCreating",
		},
	}
	fmt.Println("GetCurrentState", fsm.currentState)
}
func (fsm *FSM) Input(char string) {
	newState, ok := fsm.table[fsm.currentState][char]
	if !ok {
		newState = fsm.table[fsm.currentState]["*"]
	}
	fsm.currentState = newState
}
func (fsm *FSM) GetCurrentState() string {
	return fsm.currentState
}

type Node struct {
	tagName string
	attributes map[string]string
}

func main() {
	start := time.Now()

	fsm := new(FSM)
	fsm.Init()

	//file := readFile("html/PDF-report-8b89e64e1132cbd877b479e517f46db7375370adccade7f085d4492a0315f225.html")
	file := readFile("html/test.html")
	chars := strings.Split(file, "")
	node := new(Node)
	attributeName := ""
	attributeValue := ""

	for n, char := range chars {
		fsm.Input(char)
		switch fsm.GetCurrentState() {
		case "tagNameCreating":
			node.tagName += char
		case "tagAttributeNameCreating":
			attributeName += char
		case "tagCreatingEnd":
			node.attributes = make(map[string]string)
			node.attributes[strings.Trim(attributeName, " ")] = attributeValue
			attributeName = ""
			attributeValue = ""
		case "tagAttributeValueCreating":
			if char == " " {
				node.attributes = make(map[string]string)
				node.attributes[strings.Trim(attributeName, " ")] = attributeValue
				attributeName = ""
				attributeValue = ""
			} else if char != "\"" {
				attributeValue += char
			}
		}
		fmt.Println(n, char, fsm.GetCurrentState(), attributeName, attributeValue)
	}
	fmt.Println(node)

	end := time.Now()
	fmt.Println("\n")
	fmt.Println(end.Sub(start))
}
