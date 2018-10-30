package directory

import (
	"github.com/aws/aws-sdk-go/service/clouddirectory"
)

// Client TODO:
type Client interface {
	// obj
	ListObjectAttributes(*clouddirectory.ListObjectAttributesInput) (*clouddirectory.ListObjectAttributesOutput, error)

	// relationships
	ListObjectChildren(*clouddirectory.ListObjectChildrenInput) (*clouddirectory.ListObjectChildrenOutput, error)
	ListIncomingTypedLinks(*clouddirectory.ListIncomingTypedLinksInput) (*clouddirectory.ListIncomingTypedLinksOutput, error)
	ListOutgoingTypedLinks(*clouddirectory.ListOutgoingTypedLinksInput) (*clouddirectory.ListOutgoingTypedLinksOutput, error)
	ListObjectParents(*clouddirectory.ListObjectParentsInput) (*clouddirectory.ListObjectParentsOutput, error)
}
