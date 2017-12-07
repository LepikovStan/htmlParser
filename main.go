package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
	"time"
	//"./lib/ctrl"
	//"./lib/Node"
	"./lib/stack"
	"strings"
)

func readFile(path string) string {
	var result string
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		result = fmt.Sprintf("%s%s", result, strings.Trim(scanner.Text(), " "))
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
	fsm.currentState = "waiting"
	fsm.table = map[string]map[string]string{
		"waiting": map[string]string{
			"<": "tagNameCreatingStart",
		},
		"tagNameCreatingStart": map[string]string{
			"/": "tagClosingStart",
			"*": "tagNameCreating",
		},
		"tagNameCreating": map[string]string{
			" ": "tagAttributeNameCreating",
			">": "tagCreatingEnd",
			"/": "tagCreatingEnd",
			"*": "tagNameCreating",
		},
		"tagAttributeNameCreating": map[string]string{
			">": "tagCreatingEnd",
			"=": "tagAttributeValueCreatingStart",
			"*": "tagAttributeNameCreating",
		},
		"tagAttributeValueCreatingStart": map[string]string{
			"\"": "tagAttributeValueCreatingStart",
			"*": "tagAttributeValueCreating",
		},
		"tagAttributeValueCreating": map[string]string{
			">": "tagCreatingEnd",
			" ": "tagAttributeNameCreating",
			"*": "tagAttributeValueCreating",
		},
		"tagCreatingEnd": map[string]string{
			"<": "tagNameCreatingStart",
			">": "tagClosingEnd",
		},
		"tagClosingStart": map[string]string{
			"*": "tagClosingTagName",
		},
		"tagClosingTagName": map[string]string{
			">": "tagClosingEnd",
			"*": "tagClosingTagName",
		},
		"tagClosingEnd": map[string]string{
			"<": "tagNameCreatingStart",
			//"*": "tagClosingTagName",
		},
	}
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
	attrs map[string]string
	content []*Node
}
func (n *Node) Init() {
	n.attrs = make(map[string]string)
	n.content = make([]*Node, 1)
}
func (n *Node) GetTagName() string {
	return n.tagName
}
func (n *Node) GetAttribute(attrName string) string {
	if attr, ok := n.attrs[attrName]; ok {
		return attr
	} else {
		return ""
	}
}
func (n *Node) SetAttribute(attrName string, attrValue string) {
	n.attrs[attrName] = attrValue
}

func main() {
	start := time.Now()

	fsm := new(FSM)
	fsm.Init()

	//file := readFile("html/PDF-report-8b89e64e1132cbd877b479e517f46db7375370adccade7f085d4492a0315f225.html")
	file := readFile("html/test.html")
	chars := strings.Split(file, "")
	stck := stack.Init()

	node := new(Node)
	node.Init()

	nodes := []*Node{}
	attributeName := ""
	attributeValue := ""
	tagNameToClose := ""

	for n, char := range chars {
		fsm.Input(char)

		switch fsm.GetCurrentState() {
		case "tagClosingStart":
			tagNameToClose = ""
		case "tagClosingTagName":
			if char != "/" {
				tagNameToClose += char
			}
		case "tagClosingEnd":
			stackTag := stck.Head()
			for stackTag == tagNameToClose {
				stck.Pop()
				stackTag = stck.Head()
			}
			tagNameToClose = ""
		case "tagNameCreating":
			node.tagName += char
			tagNameToClose += char
		case "tagAttributeNameCreating":
			if char == " " {
				if attributeName != "" && attributeValue != "" {
					node.SetAttribute(strings.Trim(attributeName, " "), attributeValue)
					attributeName = ""
					attributeValue = ""
				}
			} else {
				attributeName += char
			}
		case "tagAttributeValueCreating":
			if char == " " {
				node.SetAttribute(strings.Trim(attributeName, " "), attributeValue)
				attributeName = ""
				attributeValue = ""
			} else if char != "\"" {
				attributeValue += char
			}
		case "tagCreatingEnd":
			node.SetAttribute(strings.Trim(attributeName, " "), attributeValue)
			attributeName = ""
			attributeValue = ""
			stck.Unshift(node.tagName)
			nodes = append(nodes, node)
			node = new(Node)
			node.Init()
			attributeName = ""
			attributeValue = ""
		}

		fmt.Println(n, char, fsm.GetCurrentState(), attributeName, attributeValue)
	}
	for _, node := range nodes {
		fmt.Println("node", *node)
	}
	stck.Print()

	end := time.Now()
	fmt.Println("\n")
	fmt.Println(end.Sub(start))
}
