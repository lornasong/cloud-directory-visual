package directory

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/clouddirectory"
)

// Directory holds everything needed to access AWS Cloud Directory
type Directory struct {
	client    Client
	arn       string
	schemaArn string
}

// New returns a new Directory
func New(client Client, arn, schemaArn string) *Directory {
	return &Directory{
		client:    client,
		arn:       arn,
		schemaArn: schemaArn,
	}
}

// GetObjectInformation returns a Cloud Directory's object's attributes
func (d *Directory) GetObjectInformation(ref string) (*clouddirectory.GetObjectInformationOutput, error) {
	id := ref
	if !isPath(ref) {
		id = fmt.Sprintf("$%s", id)
	}

	in := clouddirectory.GetObjectInformationInput{
		DirectoryArn:     aws.String(d.arn),
		ConsistencyLevel: aws.String(clouddirectory.ConsistencyLevelEventual),
		ObjectReference: &clouddirectory.ObjectReference{
			Selector: aws.String(id),
		},
	}

	out, err := d.client.GetObjectInformation(&in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListObjectAttributes returns a Cloud Directory's object's attributes
func (d *Directory) ListObjectAttributes(ref string) (*clouddirectory.ListObjectAttributesOutput, error) {
	id := ref
	if !isPath(ref) {
		id = fmt.Sprintf("$%s", id)
	}

	in := clouddirectory.ListObjectAttributesInput{
		DirectoryArn:     aws.String(d.arn),
		ConsistencyLevel: aws.String(clouddirectory.ConsistencyLevelEventual),
		ObjectReference: &clouddirectory.ObjectReference{
			Selector: aws.String(id),
		},
	}

	out, err := d.client.ListObjectAttributes(&in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListObjectChildren returns a list of Cloud Directory's object's children
func (d *Directory) ListObjectChildren(ref string) (*clouddirectory.ListObjectChildrenOutput, error) {
	id := ref
	if !isPath(ref) {
		id = fmt.Sprintf("$%s", id)
	}

	in := clouddirectory.ListObjectChildrenInput{
		DirectoryArn:     aws.String(d.arn),
		ConsistencyLevel: aws.String(clouddirectory.ConsistencyLevelEventual),
		ObjectReference: &clouddirectory.ObjectReference{
			Selector: aws.String(id),
		},
	}

	out, err := d.client.ListObjectChildren(&in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListObjectParents returns a list of Cloud Directory's object's parents
func (d *Directory) ListObjectParents(ref string) (*clouddirectory.ListObjectParentsOutput, error) {
	id := ref
	if !isPath(ref) {
		id = fmt.Sprintf("$%s", id)
	}

	in := clouddirectory.ListObjectParentsInput{
		DirectoryArn:     aws.String(d.arn),
		ConsistencyLevel: aws.String(clouddirectory.ConsistencyLevelEventual),
		ObjectReference: &clouddirectory.ObjectReference{
			Selector: aws.String(id),
		},
	}

	out, err := d.client.ListObjectParents(&in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListIncomingTypedLinks returns a list of Cloud Directory's object's incoming typed links
func (d *Directory) ListIncomingTypedLinks(ref string) (*clouddirectory.ListIncomingTypedLinksOutput, error) {
	id := ref
	if !isPath(ref) {
		id = fmt.Sprintf("$%s", id)
	}

	in := clouddirectory.ListIncomingTypedLinksInput{
		DirectoryArn:     aws.String(d.arn),
		ConsistencyLevel: aws.String(clouddirectory.ConsistencyLevelEventual),
		ObjectReference: &clouddirectory.ObjectReference{
			Selector: aws.String(id),
		},
	}

	out, err := d.client.ListIncomingTypedLinks(&in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListOutgoingTypedLinks returns a list of Cloud Directory's object's outgoing typed links
func (d *Directory) ListOutgoingTypedLinks(ref string) (*clouddirectory.ListOutgoingTypedLinksOutput, error) {
	id := ref
	if !isPath(ref) {
		id = fmt.Sprintf("$%s", id)
	}

	in := clouddirectory.ListOutgoingTypedLinksInput{
		DirectoryArn:     aws.String(d.arn),
		ConsistencyLevel: aws.String(clouddirectory.ConsistencyLevelEventual),
		ObjectReference: &clouddirectory.ObjectReference{
			Selector: aws.String(id),
		},
	}

	out, err := d.client.ListOutgoingTypedLinks(&in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListPolicyAttachments returns a CloudDirectory's policy's list of object ids it's attached to.
func (d *Directory) ListPolicyAttachments(ref string) (*clouddirectory.ListPolicyAttachmentsOutput, error) {
	id := ref
	if !isPath(ref) {
		id = fmt.Sprintf("$%s", id)
	}

	in := clouddirectory.ListPolicyAttachmentsInput{
		DirectoryArn:     aws.String(d.arn),
		ConsistencyLevel: aws.String(clouddirectory.ConsistencyLevelEventual),
		PolicyReference: &clouddirectory.ObjectReference{
			Selector: aws.String(id),
		},
	}

	out, err := d.client.ListPolicyAttachments(&in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListObjectPolicies returns a CloudDirectory's object's list of attached policies ids.
func (d *Directory) ListObjectPolicies(ref string) (*clouddirectory.ListObjectPoliciesOutput, error) {
	id := ref
	if !isPath(ref) {
		id = fmt.Sprintf("$%s", id)
	}

	in := clouddirectory.ListObjectPoliciesInput{
		DirectoryArn:     aws.String(d.arn),
		ConsistencyLevel: aws.String(clouddirectory.ConsistencyLevelEventual),
		ObjectReference: &clouddirectory.ObjectReference{
			Selector: aws.String(id),
		},
	}

	out, err := d.client.ListObjectPolicies(&in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func isPath(ref string) bool {
	return len(ref) > 0 && ref[0] == '/'
}
