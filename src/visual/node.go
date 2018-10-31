package visual

import (
	"fmt"
)

// Node holds information about a node and nodes it has an immediate relationship with
type Node struct {
	ID                 string
	FacetName          string
	Attributes         []*Attribute
	Parent             []*RelatedNode
	Children           []*RelatedNode
	OutgoingTypedLinks []*LinkedNode
	IncomingTypedLinks []*LinkedNode
}

func (n *Node) String() string {
	return fmt.Sprintf("{ID: %s , FacetName: %s}", n.ID, n.FacetName)
}

// FullString returns full information about a node
func (n *Node) FullString() string {
	desc := "ID: %s, FacetName: %s\nAttributes: %s\nParent: %v\nChildren: %v\nIncoming Typed Links: %v\nOutgoingTypedLinks: %v"
	return fmt.Sprintf(desc, n.ID, n.FacetName, n.Attributes, n.Parent, n.Children, n.IncomingTypedLinks, n.OutgoingTypedLinks)
}

// Attribute is a key value pair of information
type Attribute struct {
	Key   string
	Value string
}

func (a *Attribute) String() string {
	return fmt.Sprintf("{%s: %s}", a.Key, a.Value)
}

// RelatedNode holds a node and its linkname in relationship to another node
type RelatedNode struct {
	Node     *Node
	Linkname string
}

func (n *RelatedNode) String() string {
	return fmt.Sprintf("{Linkname: %s, Node: %s}", n.Linkname, n.Node)
}

// AttachedNode holds a node and its relationship to another node
type AttachedNode struct {
	Node         *Node
	Relationship string
}

func (n *AttachedNode) String() string {
	return fmt.Sprintf("{Relationship: %s, Node: %s}", n.Relationship, n.Node)
}

// LinkedNode holds a node and the attributes tied to the relationship with another node
type LinkedNode struct {
	Node       *Node
	Attributes []*Attribute
}

func (n *LinkedNode) String() string {
	return fmt.Sprintf("{Attributes: %v, Node: %s}", n.Attributes, n.Node)
}
