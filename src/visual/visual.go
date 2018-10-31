package visual

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/clouddirectory"
	"github.com/lornasong/aws-cloud-directory-visual/src/directory"
)

// Visual handles everything needed to build a visualization
type Visual struct {
	dir *directory.Directory
}

// New returns a new Visual
func New(d *directory.Directory) *Visual {
	return &Visual{
		dir: d,
	}
}

// GenerateProfile returns a node fully hydrated with details about itself and
// the nodes it has relationships with
func (v *Visual) GenerateProfile(id string) (*Node, error) {
	node, err := v.Describe(id)
	if err != nil {
		return nil, err
	}

	node, err = v.FindRelationships(node)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Describe returns a node with basic information about itself
func (v *Visual) Describe(id string) (*Node, error) {

	if isSelector(id) {
		id = selectorToID(id)
	}

	out, err := v.dir.ListObjectAttributes(id)
	if err != nil {
		return nil, err
	}

	facetName := ""
	attributes := make([]*Attribute, len(out.Attributes))
	for ix, attr := range out.Attributes {
		// note, may need to filter by schema arn TODO:
		facetName = *attr.Key.FacetName
		a := Attribute{
			Key:   *attr.Key.Name,
			Value: valueString(attr.Value),
		}
		attributes[ix] = &a
	}

	return &Node{
		ID:         id,
		FacetName:  facetName,
		Attributes: attributes,
	}, nil
}

// FindRelationships returns a node with information about the nodes it has
// a direct relationship with: parent-child, typed-link.
func (v *Visual) FindRelationships(node *Node) (*Node, error) {
	cnodes, err := v.FindChildren(node.ID)
	if err != nil {
		return nil, err
	}
	node.Children = cnodes

	pnodes, err := v.FindParents(node.ID)
	if err != nil {
		return nil, err
	}
	node.Parent = pnodes

	innodes, err := v.FindIncomingTypedLinks(node.ID)
	if err != nil {
		return nil, err
	}
	node.IncomingTypedLinks = innodes

	outnodes, err := v.FindOutgoingTypedLinks(node.ID)
	if err != nil {
		return nil, err
	}
	node.OutgoingTypedLinks = outnodes

	return node, nil
}

// FindParents returns for the id of a node, its parent nodes
func (v *Visual) FindParents(id string) ([]*RelatedNode, error) {
	ps, err := v.dir.ListObjectParents(id)
	if err != nil {
		return nil, err
	}

	ix := 0
	pnodes := make([]*RelatedNode, len(ps.Parents))
	for pid, linkname := range ps.Parents {

		pnode, err := v.Describe(pid)
		if err != nil {
			return nil, err
		}

		pnodes[ix] = &RelatedNode{
			Node:     pnode,
			Linkname: *linkname,
		}
		ix++
	}
	return pnodes, nil
}

// FindChildren returns for the id of a node, its children nodes
func (v *Visual) FindChildren(id string) ([]*RelatedNode, error) {
	cs, err := v.dir.ListObjectChildren(id)
	if err != nil {
		return nil, err
	}

	cnodes := make([]*RelatedNode, len(cs.Children))
	ix := 0
	for linkname, cid := range cs.Children {

		cnode, err := v.Describe(*cid)
		if err != nil {
			return nil, err
		}

		cnodes[ix] = &RelatedNode{
			Node:     cnode,
			Linkname: linkname,
		}
		ix++
	}
	return cnodes, nil
}

// FindIncomingTypedLinks returns for the id of a node, its incoming typed-link nodes
func (v *Visual) FindIncomingTypedLinks(id string) ([]*LinkedNode, error) {
	ins, err := v.dir.ListIncomingTypedLinks(id)
	if err != nil {
		return nil, err
	}

	innodes := make([]*LinkedNode, len(ins.LinkSpecifiers))
	ix := 0
	for _, link := range ins.LinkSpecifiers {
		// only care about source. the id is the target (incoming)
		innode, err := v.Describe(*link.SourceObjectReference.Selector)
		if err != nil {
			return nil, err
		}

		attrs := make([]*Attribute, len(link.IdentityAttributeValues))
		for ixv, val := range link.IdentityAttributeValues {
			attrs[ixv] = &Attribute{
				Key:   *val.AttributeName,
				Value: valueString(val.Value),
			}
		}
		innodes[ix] = &LinkedNode{
			Node:       innode,
			Attributes: attrs,
		}
		ix++
	}
	return innodes, nil
}

// FindOutgoingTypedLinks returns for the id of a node, its outgoing typed-link nodes
func (v *Visual) FindOutgoingTypedLinks(id string) ([]*LinkedNode, error) {
	outs, err := v.dir.ListOutgoingTypedLinks(id)
	if err != nil {
		return nil, err
	}

	outnodes := make([]*LinkedNode, len(outs.TypedLinkSpecifiers))
	ix := 0
	for _, link := range outs.TypedLinkSpecifiers {
		// only care about target. the id is the source (outgoing)
		outnode, err := v.Describe(*link.TargetObjectReference.Selector)
		if err != nil {
			return nil, err
		}

		attrs := make([]*Attribute, len(link.IdentityAttributeValues))
		for ixv, val := range link.IdentityAttributeValues {
			attrs[ixv] = &Attribute{
				Key:   *val.AttributeName,
				Value: valueString(val.Value),
			}
		}
		outnodes[ix] = &LinkedNode{
			Node:       outnode,
			Attributes: attrs,
		}
		ix++
	}
	return outnodes, nil
}

// isSelector checks if an id is a selector (id with '$' preprended)
func isSelector(id string) bool {
	return len(id) > 0 && id[0] == '$'
}

// selectorToID converts a selector to an id (removes prepended '$')
func selectorToID(selector string) string {
	if len(selector) > 1 {
		return selector[1:]
	}
	return ""
}

func valueString(value *clouddirectory.TypedAttributeValue) string {

	if value.StringValue != nil {
		return *value.StringValue
	}
	if value.NumberValue != nil {
		return *value.NumberValue
	}
	if len(value.BinaryValue) > 0 {
		return string(value.BinaryValue)
	}
	if value.DatetimeValue != nil {
		return value.DatetimeValue.String()
	}
	if value.BooleanValue != nil {
		return fmt.Sprintf("%t", *value.BooleanValue)
	}
	return ""
}
