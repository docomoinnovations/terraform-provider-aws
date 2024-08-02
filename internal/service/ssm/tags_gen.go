// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	awstypes "github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/logging"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/types/option"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// listTags lists ssm service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func listTags(ctx context.Context, conn *ssm.Client, identifier, resourceType string, optFns ...func(*ssm.Options)) (tftags.KeyValueTags, error) {
	input := &ssm.ListTagsForResourceInput{
		ResourceId:   aws.String(identifier),
		ResourceType: awstypes.ResourceTypeForTagging(resourceType),
	}

	output, err := conn.ListTagsForResource(ctx, input, optFns...)

	if err != nil {
		return tftags.New(ctx, nil), err
	}

	return KeyValueTags(ctx, output.TagList), nil
}

// ListTags lists ssm service tags and set them in Context.
// It is called from outside this package.
func (p *servicePackage) ListTags(ctx context.Context, meta any, identifier, resourceType string) error {
	tags, err := listTags(ctx, meta.(*conns.AWSClient).SSMClient(ctx), identifier, resourceType)

	if err != nil {
		return err
	}

	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = option.Some(tags)
	}

	return nil
}

// []*SERVICE.Tag handling

// Tags returns ssm service tags.
func Tags(tags tftags.KeyValueTags) []awstypes.Tag {
	result := make([]awstypes.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := awstypes.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from ssm service tags.
func KeyValueTags(ctx context.Context, tags []awstypes.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.ToString(tag.Key)] = tag.Value
	}

	return tftags.New(ctx, m)
}

// getTagsIn returns ssm service tags from Context.
// nil is returned if there are no input tags.
func getTagsIn(ctx context.Context) []awstypes.Tag {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn.UnwrapOrDefault()); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// setTagsOut sets ssm service tags in Context.
func setTagsOut(ctx context.Context, tags []awstypes.Tag) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = option.Some(KeyValueTags(ctx, tags))
	}
}

// createTags creates ssm service tags for new resources.
func createTags(ctx context.Context, conn *ssm.Client, identifier, resourceType string, tags []awstypes.Tag, optFns ...func(*ssm.Options)) error {
	if len(tags) == 0 {
		return nil
	}

	return updateTags(ctx, conn, identifier, resourceType, nil, KeyValueTags(ctx, tags), optFns...)
}

// updateTags updates ssm service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func updateTags(ctx context.Context, conn *ssm.Client, identifier, resourceType string, oldTagsMap, newTagsMap any, optFns ...func(*ssm.Options)) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	ctx = tflog.SetField(ctx, logging.KeyResourceId, identifier)

	removedTags := oldTags.Removed(newTags)
	removedTags = removedTags.IgnoreSystem(names.SSM)
	if len(removedTags) > 0 {
		input := &ssm.RemoveTagsFromResourceInput{
			ResourceId:   aws.String(identifier),
			ResourceType: awstypes.ResourceTypeForTagging(resourceType),
			TagKeys:      removedTags.Keys(),
		}

		_, err := conn.RemoveTagsFromResource(ctx, input, optFns...)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	updatedTags := oldTags.Updated(newTags)
	updatedTags = updatedTags.IgnoreSystem(names.SSM)
	if len(updatedTags) > 0 {
		input := &ssm.AddTagsToResourceInput{
			ResourceId:   aws.String(identifier),
			ResourceType: awstypes.ResourceTypeForTagging(resourceType),
			Tags:         Tags(updatedTags),
		}

		_, err := conn.AddTagsToResource(ctx, input, optFns...)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}

// UpdateTags updates ssm service tags.
// It is called from outside this package.
func (p *servicePackage) UpdateTags(ctx context.Context, meta any, identifier, resourceType string, oldTags, newTags any) error {
	return updateTags(ctx, meta.(*conns.AWSClient).SSMClient(ctx), identifier, resourceType, oldTags, newTags)
}