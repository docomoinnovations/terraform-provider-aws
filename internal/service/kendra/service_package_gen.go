// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package kendra

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	kendra_sdkv2 "github.com/aws/aws-sdk-go-v2/service/kendra"
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
			Factory:  DataSourceExperience,
			TypeName: "aws_kendra_experience",
		},
		{
			Factory:  DataSourceFaq,
			TypeName: "aws_kendra_faq",
		},
		{
			Factory:  DataSourceIndex,
			TypeName: "aws_kendra_index",
		},
		{
			Factory:  DataSourceQuerySuggestionsBlockList,
			TypeName: "aws_kendra_query_suggestions_block_list",
		},
		{
			Factory:  DataSourceThesaurus,
			TypeName: "aws_kendra_thesaurus",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceDataSource,
			TypeName: "aws_kendra_data_source",
			Name:     "Data Source",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceExperience,
			TypeName: "aws_kendra_experience",
		},
		{
			Factory:  ResourceFaq,
			TypeName: "aws_kendra_faq",
			Name:     "FAQ",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceIndex,
			TypeName: "aws_kendra_index",
			Name:     "Index",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceQuerySuggestionsBlockList,
			TypeName: "aws_kendra_query_suggestions_block_list",
			Name:     "Query Suggestions Block List",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceThesaurus,
			TypeName: "aws_kendra_thesaurus",
			Name:     "Thesaurus",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Kendra
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*kendra_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return kendra_sdkv2.NewFromConfig(cfg,
		kendra_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}