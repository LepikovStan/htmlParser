package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	//"github.com/LepikovStan/htmlParser/lib/node"
	//"github.com/LepikovStan/htmlParser/lib/stack"
	//"github.com/LepikovStan/htmlParser/lib/fsm"
	"./lib/node"
	"./lib/stack"
	"./lib/fsm"
	"strings"
)

var simpleTags = map[string]bool{
	"meta": true,
	"img": true,
	"link": true,
	"br": true,
	"hr": true,
	"input": true,
	"area": true,
	"param": true,
	"col": true,
	"base": true,
}

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

func pr(n *node.Node) {
	if n.GetTagName() != "" {
		fmt.Println("node ->", n.GetTagName(), n.GetAttributes(), n.GetParent().GetTagName())
	} else {
		fmt.Println("node ->", "text", n.GetText(), n.GetParent().GetTagName())
	}
	for _, in := range n.GetContent() {
		pr(in)
	}
}

func main() {
	start := time.Now()

	finsm := fsm.Create()

	//file := readFile("html/PDF-report-8b89e64e1132cbd877b479e517f46db7375370adccade7f085d4492a0315f225.html")
	file := readFile("html/test2.html")
	chars := strings.Split(file, "")
	stck := stack.Init()

	tagNode := node.CreateNewNode()
	textNode := node.CreateNewNode()
	top := tagNode

	var nodeToClose *node.Node

	for _, char := range chars {
		finsm.Input(char)

		switch finsm.GetCurrentState() {
		//case "tagClosingStart":
		//case "tagClosingTagName":
		case "tagNameCreatingStart":
			fmt.Println(stck.Head(), textNode)
			stackNode := stck.Head()
			if stackNode != nil && textNode.HasText() {
				textNode.SetParent(stackNode)
				stackNode.SetContent(textNode)
				textNode = node.CreateNewNode()
			}
		case "textCreating":
			textNode.AddText(char)
		case "tagClosingEnd":
			stackNode := stck.Head()
			for stackNode == nodeToClose && stackNode != nil {
				stackNode = nil
				stck.Pop()
			}
			nodeToClose = stck.Head()
		case "tagNameCreating":
			tagNode.CreateTagName(char)
		case "tagAttributeNameCreating":
			if char == " " {
				tagNode.SetAttribute()
			} else {
				tagNode.CreateAttributeName(char)
			}
		case "tagAttributeValueCreating":
			if char == " " {
				tagNode.SetAttribute()
			} else if char != "\"" {
				tagNode.CreateAttributeValue(char)
			}
		case "tagCreatingEnd":
			tagNode.SetAttribute()
			parentNode := stck.Head()
			if parentNode != nil {
				tagNode.SetParent(parentNode)
				parentNode.SetContent(tagNode)
			}

			if _, ok := simpleTags[tagNode.GetTagName()]; !ok {
				stck.Unshift(tagNode)
			}
			nodeToClose = tagNode
			tagNode = node.CreateNewNode()
		}

		fmt.Println(char, finsm.GetCurrentState())
	}
	stck.Print()
	pr(top)

	end := time.Now()
	fmt.Println("\n")
	fmt.Println(end.Sub(start))
}
