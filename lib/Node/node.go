package Node

type Node struct {
	tagName string
	attrs   map[string]string
	content []Node
}

func (n *Node) GetAttr(attrName string) string {
	return n.attrs[attrName]
}

func (n *Node) GetTagName() string {
	return n.tagName
}
