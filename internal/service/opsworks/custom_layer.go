package opsworks

import (
	"github.com/aws/aws-sdk-go/service/opsworks"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// @SDKResource("aws_opsworks_custom_layer")
func ResourceCustomLayer() *schema.Resource {
	layerType := &opsworksLayerType{
		TypeName:        opsworks.LayerTypeCustom,
		CustomShortName: true,

		// The "custom" layer type has no additional attributes.
		Attributes: map[string]*opsworksLayerTypeAttribute{},
	}

	return layerType.resourceSchema()
}
