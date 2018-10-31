package visual

import (
	"bytes"
	"fmt"
)

// Node TODO: stand alone
type Node struct {
	ID                 string
	FacetName          string
	Attributes         []*Attribute
	Parent             *RelatedNode
	Children           []*RelatedNode
	OutgoingTypedLinks []*LinkedNode
	IncomingTypedLinks []*LinkedNode
}

// RelatedNode TODO:
type RelatedNode struct {
	Node     *Node
	Linkname string
}

func (n *RelatedNode) String() string {
	return fmt.Sprintf("{Linkname: %s, Node: %s}", n.Linkname, n.Node)
}

// AttachedNode TODO:
type AttachedNode struct {
	Node         *Node
	Relationship string
}

func (n *AttachedNode) String() string {
	return fmt.Sprintf("{Relationship: %s, Node: %s}", n.Relationship, n.Node)
}

// LinkedNode TODO:
type LinkedNode struct {
	Node       *Node
	Attributes []*Attribute
}

func (n *LinkedNode) String() string {
	return fmt.Sprintf("{Attributes: %v, Node: %s}", n.Attributes, n.Node)
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
	return fmt.Sprintf("{ID: %s , FacetName: %s}", n.ID, n.FacetName)
}

// FullString TODO:
func (n *Node) FullString() string {
	desc := "ID: %s, FacetName: %s\nAttributes: %s\nParent: %v\nChildren: %v\nIncoming Typed Links: %v\nOutgoingTypedLinks: %v"
	return fmt.Sprintf(desc, n.ID, n.FacetName, n.Attributes, n.Parent, n.Children, n.IncomingTypedLinks, n.OutgoingTypedLinks)
}

// summary string

// AttributesString TODO:
func (n *Node) AttributesString() string {
	var buffer bytes.Buffer
	for _, a := range n.Attributes {
		buffer.WriteString(a.String())
	}
	return buffer.String()
}
