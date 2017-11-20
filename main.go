package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
	"time"
	//"github.com/LepikovStan/htmlParser/lib/stack"
	"github.com/LepikovStan/htmlParser/lib/fsm"
	"github.com/LepikovStan/htmlParser/lib/stack"
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

type Controller struct {
	fsm *fsm.Fsm
	stack *stack.Stack
}
func (c *Controller) Init(fsm *fsm.Fsm, stack *stack.Stack) {
	c.fsm = fsm
	c.stack = stack
}
func (c *Controller) Input(char string) {
	switch {
	case char == "<":
		c.stack.Unshift(char)
		c.fsm.SetCurrentState("waiting")
	case char == ">" && c.stack.Head() == "<":
		c.stack.Pop()
		c.stack.Unshift("")
		c.fsm.SetCurrentState("waiting")
	case char != "/" && c.stack.Head() == "<":
		c.fsm.SetCurrentState("tagParsing")
	case char == "/" && c.stack.Head() == "<":
		c.stack.Pop()
		c.stack.Unshift("")
		c.fsm.SetCurrentState("waiting")
	}
}

func main() {
	start := time.Now()

	file := readFile("html/test.html")
	chars := strings.Split(file, "")

	stk := stack.Init()
	fsm := fsm.Init()
	ctrl := Controller{}
	ctrl.Init(fsm, stk)

	//fmt.Println(fsm.GetCurrentState())
	//file := readFile("html/PDF-report-8b89e64e1132cbd877b479e517f46db7375370adccade7f085d4492a0315f225.html")

	//stack := Stack{}
	//tree := Tree{}
	//
	//chars := strings.Split(file, "")
	//count := 0
	for _, char := range chars {
		ctrl.Input(char)
		stk.Print()
		fmt.Println(char, fsm.GetCurrentState())
	}
	//fmt.Println("end", count, tree.getTags()[0])

	end := time.Now()
	fmt.Println("\n")
	fmt.Println(end.Sub(start))
}
