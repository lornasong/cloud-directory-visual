package visual

import (
	"bytes"
	"fmt"
)

// Node TODO:
type Node struct {
	FacetName  string
	Attributes *[]Attribute
	// TypedLinks
	Parent   *Node
	Children []*Node
}

// Attribute TODO:
type Attribute struct {
	Key   string
	Value string
}

func (a *Attribute) String() string {
	return fmt.Sprintf("{%s: %s}", a.Key, a.Value)
}

// String TODO:
func (n *Node) String() string {
	return fmt.Sprintf("FacetName: %s\nAttributes: %s\n", n.FacetName, n.AttributesString())
}

// AttributesString TODO:
func (n *Node) AttributesString() string {
	var buffer bytes.Buffer
	for _, a := range *n.Attributes {
		buffer.WriteString(a.String())
	}
	return buffer.String()
}
