// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package glue

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	glue_sdkv2 "github.com/aws/aws-sdk-go-v2/service/glue"
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
			Factory:  DataSourceCatalogTable,
			TypeName: "aws_glue_catalog_table",
		},
		{
			Factory:  DataSourceConnection,
			TypeName: "aws_glue_connection",
		},
		{
			Factory:  DataSourceDataCatalogEncryptionSettings,
			TypeName: "aws_glue_data_catalog_encryption_settings",
		},
		{
			Factory:  DataSourceScript,
			TypeName: "aws_glue_script",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceCatalogDatabase,
			TypeName: "aws_glue_catalog_database",
			Name:     "Database",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceCatalogTable,
			TypeName: "aws_glue_catalog_table",
		},
		{
			Factory:  ResourceClassifier,
			TypeName: "aws_glue_classifier",
		},
		{
			Factory:  ResourceConnection,
			TypeName: "aws_glue_connection",
			Name:     "Connection",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceCrawler,
			TypeName: "aws_glue_crawler",
			Name:     "Crawler",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceDataCatalogEncryptionSettings,
			TypeName: "aws_glue_data_catalog_encryption_settings",
		},
		{
			Factory:  ResourceDataQualityRuleset,
			TypeName: "aws_glue_data_quality_ruleset",
			Name:     "Data Quality Ruleset",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceDevEndpoint,
			TypeName: "aws_glue_dev_endpoint",
			Name:     "Dev Endpoint",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceJob,
			TypeName: "aws_glue_job",
			Name:     "Job",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceMLTransform,
			TypeName: "aws_glue_ml_transform",
			Name:     "ML Transform",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourcePartition,
			TypeName: "aws_glue_partition",
		},
		{
			Factory:  ResourcePartitionIndex,
			TypeName: "aws_glue_partition_index",
		},
		{
			Factory:  ResourceRegistry,
			TypeName: "aws_glue_registry",
			Name:     "Registry",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceResourcePolicy,
			TypeName: "aws_glue_resource_policy",
		},
		{
			Factory:  ResourceSchema,
			TypeName: "aws_glue_schema",
			Name:     "Schema",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceSecurityConfiguration,
			TypeName: "aws_glue_security_configuration",
		},
		{
			Factory:  ResourceTrigger,
			TypeName: "aws_glue_trigger",
			Name:     "Trigger",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceUserDefinedFunction,
			TypeName: "aws_glue_user_defined_function",
		},
		{
			Factory:  ResourceWorkflow,
			TypeName: "aws_glue_workflow",
			Name:     "Workflow",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Glue
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*glue_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return glue_sdkv2.NewFromConfig(cfg,
		glue_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}