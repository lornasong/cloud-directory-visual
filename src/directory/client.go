package directory

import (
	"github.com/aws/aws-sdk-go/service/clouddirectory"
)

// Client is the interface for AWS Cloud Directory
type Client interface {
	// obj
	ListObjectAttributes(*clouddirectory.ListObjectAttributesInput) (*clouddirectory.ListObjectAttributesOutput, error)

	// relationships
	ListObjectChildren(*clouddirectory.ListObjectChildrenInput) (*clouddirectory.ListObjectChildrenOutput, error)
	ListObjectParents(*clouddirectory.ListObjectParentsInput) (*clouddirectory.ListObjectParentsOutput, error)
	ListIncomingTypedLinks(*clouddirectory.ListIncomingTypedLinksInput) (*clouddirectory.ListIncomingTypedLinksOutput, error)
	ListOutgoingTypedLinks(*clouddirectory.ListOutgoingTypedLinksInput) (*clouddirectory.ListOutgoingTypedLinksOutput, error)
}
