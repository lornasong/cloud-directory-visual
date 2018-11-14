package directory

import (
	"github.com/aws/aws-sdk-go/service/clouddirectory"
)

// Client is the interface for AWS Cloud Directory
type Client interface {
	// obj
	ListObjectAttributes(*clouddirectory.ListObjectAttributesInput) (*clouddirectory.ListObjectAttributesOutput, error)
	GetObjectInformation(*clouddirectory.GetObjectInformationInput) (*clouddirectory.GetObjectInformationOutput, error)

	// relationships
	ListObjectChildren(*clouddirectory.ListObjectChildrenInput) (*clouddirectory.ListObjectChildrenOutput, error)
	ListObjectParents(*clouddirectory.ListObjectParentsInput) (*clouddirectory.ListObjectParentsOutput, error)

	ListIncomingTypedLinks(*clouddirectory.ListIncomingTypedLinksInput) (*clouddirectory.ListIncomingTypedLinksOutput, error)
	ListOutgoingTypedLinks(*clouddirectory.ListOutgoingTypedLinksInput) (*clouddirectory.ListOutgoingTypedLinksOutput, error)

	ListPolicyAttachments(*clouddirectory.ListPolicyAttachmentsInput) (*clouddirectory.ListPolicyAttachmentsOutput, error)
	ListObjectPolicies(*clouddirectory.ListObjectPoliciesInput) (*clouddirectory.ListObjectPoliciesOutput, error)
}
