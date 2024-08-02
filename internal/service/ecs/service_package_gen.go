// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package ecs

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	ecs_sdkv2 "github.com/aws/aws-sdk-go-v2/service/ecs"
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
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceCluster,
			TypeName: "aws_ecs_cluster",
			Name:     "Cluster",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceContainerDefinition,
			TypeName: "aws_ecs_container_definition",
			Name:     "Container Definition",
		},
		{
			Factory:  dataSourceService,
			TypeName: "aws_ecs_service",
			Name:     "Service",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceTaskDefinition,
			TypeName: "aws_ecs_task_definition",
			Name:     "Task Definition",
		},
		{
			Factory:  dataSourceTaskExecution,
			TypeName: "aws_ecs_task_execution",
			Name:     "Task Execution",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAccountSettingDefault,
			TypeName: "aws_ecs_account_setting_default",
			Name:     "Account Setting Default",
		},
		{
			Factory:  resourceCapacityProvider,
			TypeName: "aws_ecs_capacity_provider",
			Name:     "Capacity Provider",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceCluster,
			TypeName: "aws_ecs_cluster",
			Name:     "Cluster",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceClusterCapacityProviders,
			TypeName: "aws_ecs_cluster_capacity_providers",
			Name:     "Cluster Capacity Providers",
		},
		{
			Factory:  resourceService,
			TypeName: "aws_ecs_service",
			Name:     "Service",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceTag,
			TypeName: "aws_ecs_tag",
			Name:     "ECS Resource Tag",
		},
		{
			Factory:  resourceTaskDefinition,
			TypeName: "aws_ecs_task_definition",
			Name:     "Task Definition",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceTaskSet,
			TypeName: "aws_ecs_task_set",
			Name:     "Task Set",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ECS
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*ecs_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return ecs_sdkv2.NewFromConfig(cfg,
		ecs_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}