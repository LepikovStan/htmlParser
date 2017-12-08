package fsm

import (
	"fmt"
	"../node"
)

type FSM struct {
	table        map[string]map[string]string
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
			"/": "tagCreatingEnd",
			"=": "tagAttributeValueCreatingStart",
			"*": "tagAttributeNameCreating",
		},
		"tagAttributeValueCreatingStart": map[string]string{
			"\"": "tagAttributeValueCreatingStart",
			"*":  "tagAttributeValueCreating",
		},
		"tagAttributeValueCreating": map[string]string{
			">": "tagCreatingEnd",
			"\"": "tagAttributeValueCreatingEnd",
			" ": "tagAttributeNameCreating",
			"*": "tagAttributeValueCreating",
		},
		"tagAttributeValueCreatingEnd": map[string]string{
			" ": "tagAttributeNameCreating",
			">": "tagCreatingEnd",
			"/": "tagCreatingEnd",
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

func pr(n *node.Node) {
	fmt.Println("node ->", n.GetTagName(), n.GetAttributes())
	for _, in := range n.GetContent() {
		pr(in)
	}
}

func Create() *FSM {
	fsm := new(FSM)
	fsm.Init()
	return fsm
}
