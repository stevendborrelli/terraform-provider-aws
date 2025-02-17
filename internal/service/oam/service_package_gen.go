// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package oam

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceLink,
			TypeName: "aws_oam_link",
		},
		{
			Factory:  DataSourceLinks,
			TypeName: "aws_oam_links",
		},
		{
			Factory:  DataSourceSink,
			TypeName: "aws_oam_sink",
		},
		{
			Factory:  DataSourceSinks,
			TypeName: "aws_oam_sinks",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceLink,
			TypeName: "aws_oam_link",
		},
		{
			Factory:  ResourceSink,
			TypeName: "aws_oam_sink",
		},
		{
			Factory:  ResourceSinkPolicy,
			TypeName: "aws_oam_sink_policy",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ObservabilityAccessManager
}

var ServicePackage = &servicePackage{}
