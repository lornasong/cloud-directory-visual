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

// DescribeObject TODO:
func (v *Visual) DescribeObject(id string) (*Node, error) {

	out, err := v.dir.ListObjectAttributes(id)
	if err != nil {
		return nil, err
	}

	facetName := ""
	attributes := make([]Attribute, len(out.Attributes))
	for ix, attr := range out.Attributes {
		facetName = *attr.Key.FacetName
		a := Attribute{
			Key:   *attr.Key.Name,
			Value: valueString(attr.Value),
		}
		attributes[ix] = a
	}

	stuff := Node{
		FacetName:  facetName,
		Attributes: &attributes,
	}

	return &stuff, nil
}

func valueString(value *clouddirectory.TypedAttributeValue) string {

	if len(*value.StringValue) > 0 {
		return *value.StringValue
	}
	if len(*value.NumberValue) > 0 {
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