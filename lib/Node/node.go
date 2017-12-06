package node

type Node struct {
	tagName string
	attrs   map[string]string
	content []Node
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
