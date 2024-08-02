// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package datasync

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	datasync_sdkv2 "github.com/aws/aws-sdk-go-v2/service/datasync"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
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
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAgent,
			TypeName: "aws_datasync_agent",
			Name:     "Agent",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceLocationAzureBlob,
			TypeName: "aws_datasync_location_azure_blob",
			Name:     "Location Microsoft Azure Blob Storage",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceLocationEFS,
			TypeName: "aws_datasync_location_efs",
			Name:     "Location EFS",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceLocationFSxLustreFileSystem,
			TypeName: "aws_datasync_location_fsx_lustre_file_system",
			Name:     "Location FSx for Lustre File System",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceLocationFSxONTAPFileSystem,
			TypeName: "aws_datasync_location_fsx_ontap_file_system",
			Name:     "Location FSx for NetApp ONTAP File System",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceLocationFSxOpenZFSFileSystem,
			TypeName: "aws_datasync_location_fsx_openzfs_file_system",
			Name:     "Location FSx for OpenZFS File System",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceLocationFSxWindowsFileSystem,
			TypeName: "aws_datasync_location_fsx_windows_file_system",
			Name:     "Location FSx for Windows File Server File System",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceLocationHDFS,
			TypeName: "aws_datasync_location_hdfs",
			Name:     "Location HDFS",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceLocationNFS,
			TypeName: "aws_datasync_location_nfs",
			Name:     "Location NFS",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceLocationObjectStorage,
			TypeName: "aws_datasync_location_object_storage",
			Name:     "Location Object Storage",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceLocationS3,
			TypeName: "aws_datasync_location_s3",
			Name:     "Location S3",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceLocationSMB,
			TypeName: "aws_datasync_location_smb",
			Name:     "Location SMB",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceTask,
			TypeName: "aws_datasync_task",
			Name:     "Task",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.DataSync
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*datasync_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return datasync_sdkv2.NewFromConfig(cfg,
		datasync_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}