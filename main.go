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
	fmt.Println("node ->", n.GetTagName(), n.GetAttributes())
	for _, in := range n.GetContent() {
		pr(in)
	}
}

func main() {
	start := time.Now()

	finsm := fsm.Create()

	//file := readFile("html/PDF-report-8b89e64e1132cbd877b479e517f46db7375370adccade7f085d4492a0315f225.html")
	file := readFile("html/test.html")
	chars := strings.Split(file, "")
	stck := stack.Init()

	nd := node.CreateNewNode()
	top := nd

	tagNameToClose := ""

	for n, char := range chars {
		finsm.Input(char)

		switch finsm.GetCurrentState() {
		case "tagClosingStart":
			tagNameToClose = ""
		case "tagClosingTagName":
			if char != "/" {
				tagNameToClose += char
			}
		case "tagClosingEnd":
			stackNode := *stck.Head()
			for stackNode.GetTagName() == tagNameToClose {
				stackNode = node.Node{}
				stck.Pop()
			}
			tagNameToClose = ""
		case "tagNameCreating":
			nd.CreateTagName(char)
			tagNameToClose += char
		case "tagAttributeNameCreating":
			if char == " " {
				nd.SetAttribute()
			} else {
				nd.CreateAttributeName(char)
			}
		case "tagAttributeValueCreating":
			if char == " " {
				nd.SetAttribute()
			} else if char != "\"" {
				nd.CreateAttributeValue(char)
			}
		case "tagCreatingEnd":
			nd.SetAttribute()
			parentNode := stck.Head()
			if parentNode != nil {
				parentNode.SetContent(nd)
			}
			stck.Unshift(nd)
			nd = node.CreateNewNode()
		}

		fmt.Println(n, char, finsm.GetCurrentState())
	}
	stck.Print()
	pr(top)

	end := time.Now()
	fmt.Println("\n")
	fmt.Println(end.Sub(start))
}
