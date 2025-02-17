// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package storagegateway

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
			Factory:  DataSourceLocalDisk,
			TypeName: "aws_storagegateway_local_disk",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceCache,
			TypeName: "aws_storagegateway_cache",
		},
		{
			Factory:  ResourceCachediSCSIVolume,
			TypeName: "aws_storagegateway_cached_iscsi_volume",
		},
		{
			Factory:  ResourceFileSystemAssociation,
			TypeName: "aws_storagegateway_file_system_association",
		},
		{
			Factory:  ResourceGateway,
			TypeName: "aws_storagegateway_gateway",
		},
		{
			Factory:  ResourceNFSFileShare,
			TypeName: "aws_storagegateway_nfs_file_share",
		},
		{
			Factory:  ResourceSMBFileShare,
			TypeName: "aws_storagegateway_smb_file_share",
		},
		{
			Factory:  ResourceStorediSCSIVolume,
			TypeName: "aws_storagegateway_stored_iscsi_volume",
		},
		{
			Factory:  ResourceTapePool,
			TypeName: "aws_storagegateway_tape_pool",
		},
		{
			Factory:  ResourceUploadBuffer,
			TypeName: "aws_storagegateway_upload_buffer",
		},
		{
			Factory:  ResourceWorkingStorage,
			TypeName: "aws_storagegateway_working_storage",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.StorageGateway
}

var ServicePackage = &servicePackage{}
