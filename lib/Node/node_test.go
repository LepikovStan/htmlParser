package Node

import (
	"testing"
)

func TestGetAttr(t *testing.T) {
	n := Node{
		tagName: "div",
		attrs: map[string]string{
			"class": "test",
		},
	}

	if n.GetTagName() != "div" {
		t.Fatalf("Wrong tag name")
	}

	if n.GetAttr("class") != "test" {
		t.Fatalf("Wrong tag attribute class")
	}
}
