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
