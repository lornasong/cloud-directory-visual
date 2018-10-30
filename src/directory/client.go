package directory

import (
	"github.com/aws/aws-sdk-go/service/clouddirectory"
)

// Client TODO:
type Client interface {
	ListObjectAttributes(*clouddirectory.ListObjectAttributesInput) (*clouddirectory.ListObjectAttributesOutput, error)
}
