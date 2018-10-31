package directory

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/clouddirectory"
)

// Directory TODO:
type Directory struct {
	client    Client
	arn       string
	schemaArn string
}

// New TODO:
func New(client Client, arn, schemaArn string) *Directory {
	return &Directory{
		client:    client,
		arn:       arn,
		schemaArn: schemaArn,
	}
}

// ListObjectAttributes TODO:
func (d *Directory) ListObjectAttributes(id string) (*clouddirectory.ListObjectAttributesOutput, error) {
	in := clouddirectory.ListObjectAttributesInput{
		DirectoryArn:     aws.String(d.arn),
		ConsistencyLevel: aws.String(clouddirectory.ConsistencyLevelEventual),
		ObjectReference: &clouddirectory.ObjectReference{
			Selector: aws.String(fmt.Sprintf("$%s", id)),
		},
	}

	out, err := d.client.ListObjectAttributes(&in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListObjectChildren TODO:
func (d *Directory) ListObjectChildren(id string) (*clouddirectory.ListObjectChildrenOutput, error) {
	in := clouddirectory.ListObjectChildrenInput{
		DirectoryArn:     aws.String(d.arn),
		ConsistencyLevel: aws.String(clouddirectory.ConsistencyLevelEventual),
		ObjectReference: &clouddirectory.ObjectReference{
			Selector: aws.String(fmt.Sprintf("$%s", id)),
		},
	}

	out, err := d.client.ListObjectChildren(&in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListObjectParents TODO:
func (d *Directory) ListObjectParents(id string) (*clouddirectory.ListObjectParentsOutput, error) {
	in := clouddirectory.ListObjectParentsInput{
		DirectoryArn:     aws.String(d.arn),
		ConsistencyLevel: aws.String(clouddirectory.ConsistencyLevelEventual),
		ObjectReference: &clouddirectory.ObjectReference{
			Selector: aws.String(fmt.Sprintf("$%s", id)),
		},
	}

	out, err := d.client.ListObjectParents(&in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListIncomingTypedLinks TODO:
func (d *Directory) ListIncomingTypedLinks(id string) (*clouddirectory.ListIncomingTypedLinksOutput, error) {
	in := clouddirectory.ListIncomingTypedLinksInput{
		DirectoryArn:     aws.String(d.arn),
		ConsistencyLevel: aws.String(clouddirectory.ConsistencyLevelEventual),
		ObjectReference: &clouddirectory.ObjectReference{
			Selector: aws.String(fmt.Sprintf("$%s", id)),
		},
	}

	out, err := d.client.ListIncomingTypedLinks(&in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListOutgoingTypedLinks TODO:
func (d *Directory) ListOutgoingTypedLinks(id string) (*clouddirectory.ListOutgoingTypedLinksOutput, error) {
	in := clouddirectory.ListOutgoingTypedLinksInput{
		DirectoryArn:     aws.String(d.arn),
		ConsistencyLevel: aws.String(clouddirectory.ConsistencyLevelEventual),
		ObjectReference: &clouddirectory.ObjectReference{
			Selector: aws.String(fmt.Sprintf("$%s", id)),
		},
	}

	out, err := d.client.ListOutgoingTypedLinks(&in)
	if err != nil {
		return nil, err
	}
	return out, nil
}
