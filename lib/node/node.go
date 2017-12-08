package node

import "strings"

type Node struct {
	tagName        string
	attributeName  string
	attributeValue string
	attrs          map[string]string
	content        []*Node
	parent *Node
}

func (n *Node) init() {
	n.attrs = make(map[string]string)
	n.content = []*Node{}
	n.parent = new(Node)
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
func (n *Node) GetAttributes() map[string]string {
	return n.attrs
}
func (n *Node) CreateTagName(char string) {
	n.tagName += char
}
func (n *Node) CreateAttributeName(char string) {
	n.attributeName += char
}
func (n *Node) CreateAttributeValue(char string) {
	n.attributeValue += char
}
func (n *Node) SetAttribute() {
	if n.attributeName != "" && n.attributeValue != "" {
		n.attrs[strings.Trim(n.attributeName, " ")] = n.attributeValue
		n.attributeName = ""
		n.attributeValue = ""
	}
}
func (n *Node) SetContent(nd *Node) {
	n.content = append(n.content, nd)
}
func (n *Node) GetContent() []*Node {
	return n.content
}
func (n *Node) SetParent(nd *Node) {
	n.parent = nd
}
func (n *Node) GetParent() *Node {
	return n.parent
}

func CreateNewNode() *Node {
	nd := new(Node)
	nd.init()
	return nd
}
