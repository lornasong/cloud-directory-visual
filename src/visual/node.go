package visual

import (
	"fmt"
)

// Node holds information about a node and its immediate relationships
type Node struct {
	ID                 string
	FacetName          string
	Attributes         []*Attribute
	Parent             []*RelatedNode
	Children           []*RelatedNode
	OutgoingTypedLinks []*LinkedNode
	IncomingTypedLinks []*LinkedNode
	AttachedPolicies   []*AttachedNode
	AttachedObjects    []*AttachedNode
}

func (n *Node) String() string {
	return fmt.Sprintf("{ID: %s , FacetName: %s}", n.ID, n.FacetName)
}

// FullString returns full information about a node
func (n *Node) FullString() string {
	desc := "ID: %s, FacetName: %s\nAttributes: %s\n" +
		"Parent: %v\nChildren: %v\n" +
		"Incoming Typed Links: %v\nOutgoingTypedLinks: %v\n" +
		"Attached Policies: %v\nAttached Objects: %v"
	return fmt.Sprintf(desc, n.ID, n.FacetName, n.Attributes, n.Parent, n.Children,
		n.IncomingTypedLinks, n.OutgoingTypedLinks, n.AttachedPolicies, n.AttachedObjects)
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

// AttachedNode holds a node that is attached usually via policy attachments
type AttachedNode struct {
	Node *Node
}

func (n *AttachedNode) String() string {
	return fmt.Sprintf("{Node: %s}", n.Node)
}

// LinkedNode holds a node and the attributes tied to the relationship with another node
type LinkedNode struct {
	Node       *Node
	Attributes []*Attribute
}

func (n *LinkedNode) String() string {
	return fmt.Sprintf("{Attributes: %v, Node: %s}", n.Attributes, n.Node)
}
