package visual

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/clouddirectory"
	"github.com/lornasong/aws-cloud-directory-visual/src/directory"
)

// Visual TODO:
type Visual struct {
	dir *directory.Directory
}

// New TODO:
func New(d *directory.Directory) *Visual {
	return &Visual{
		dir: d,
	}
}

// GenerateProfile TODO:
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

// Describe TODO:
func (v *Visual) Describe(id string) (*Node, error) {

	// if selector, clear off $
	if len(id) > 0 && id[0] == '$' {
		id = id[1:]
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

// FindRelationships TODO:
// return identifier of all relationships
func (v *Visual) FindRelationships(node *Node) (*Node, error) {

	cnodes, err := v.FindChildren(node.ID)
	if err != nil {
		return nil, err
	}
	node.Children = cnodes

	pnode, err := v.FindParents(node.ID)
	if err != nil {
		return nil, err
	}
	node.Parent = pnode

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

// FindParents TODO:
func (v *Visual) FindParents(id string) (*RelatedNode, error) {

	// need to catch NotNodeException:

	ps, err := v.dir.ListObjectParents(id)
	if err != nil {
		return nil, err
	}

	if len(ps.Parents) > 1 {
		fmt.Println("ERROR unexpected")
		// FIXME: need error handling
		return nil, nil
	}

	for pid, linkname := range ps.Parents {

		pnode, err := v.Describe(pid)
		if err != nil {
			return nil, err
		}

		return &RelatedNode{
			Node:     pnode,
			Linkname: *linkname,
		}, nil
	}
	// FIXME: need error handling? how should this be handled?
	return nil, nil
}

// FindChildren TODO:
func (v *Visual) FindChildren(id string) ([]*RelatedNode, error) {

	// need to catch NotNodeException:

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

// FindIncomingTypedLinks TODO:
func (v *Visual) FindIncomingTypedLinks(id string) ([]*LinkedNode, error) {

	// need to catch NotNodeException:

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

// FindOutgoingTypedLinks TODO:
func (v *Visual) FindOutgoingTypedLinks(id string) ([]*LinkedNode, error) {

	// need to catch NotNodeException:

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
