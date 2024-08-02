// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package waf

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	waf_sdkv2 "github.com/aws/aws-sdk-go-v2/service/waf"
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
			Factory:  dataSourceIPSet,
			TypeName: "aws_waf_ipset",
			Name:     "IPSet",
		},
		{
			Factory:  dataSourceRateBasedRule,
			TypeName: "aws_waf_rate_based_rule",
			Name:     "Rate Based Rule",
		},
		{
			Factory:  dataSourceRule,
			TypeName: "aws_waf_rule",
			Name:     "Rule",
		},
		{
			Factory:  dataSourceSubscribedRuleGroup,
			TypeName: "aws_waf_subscribed_rule_group",
			Name:     "Subscribed Rule Group",
		},
		{
			Factory:  dataSourceWebACL,
			TypeName: "aws_waf_web_acl",
			Name:     "Web ACL",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceByteMatchSet,
			TypeName: "aws_waf_byte_match_set",
			Name:     "ByteMatchSet",
		},
		{
			Factory:  resourceGeoMatchSet,
			TypeName: "aws_waf_geo_match_set",
			Name:     "GeoMatchSet",
		},
		{
			Factory:  resourceIPSet,
			TypeName: "aws_waf_ipset",
			Name:     "IPSet",
		},
		{
			Factory:  resourceRateBasedRule,
			TypeName: "aws_waf_rate_based_rule",
			Name:     "Rate Based Rule",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceRegexMatchSet,
			TypeName: "aws_waf_regex_match_set",
			Name:     "Regex Match Set",
		},
		{
			Factory:  resourceRegexPatternSet,
			TypeName: "aws_waf_regex_pattern_set",
			Name:     "Regex Pattern Set",
		},
		{
			Factory:  resourceRule,
			TypeName: "aws_waf_rule",
			Name:     "Rule",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceRuleGroup,
			TypeName: "aws_waf_rule_group",
			Name:     "Rule Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceSizeConstraintSet,
			TypeName: "aws_waf_size_constraint_set",
			Name:     "Size Constraint Set",
		},
		{
			Factory:  resourceSQLInjectionMatchSet,
			TypeName: "aws_waf_sql_injection_match_set",
			Name:     "SqlInjectionMatchSet",
		},
		{
			Factory:  resourceWebACL,
			TypeName: "aws_waf_web_acl",
			Name:     "Web ACL",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceXSSMatchSet,
			TypeName: "aws_waf_xss_match_set",
			Name:     "XSS Match Set",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.WAF
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*waf_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return waf_sdkv2.NewFromConfig(cfg,
		waf_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}