// Code generated by "internal/generate/listpages/main.go -AWSSDKVersion=2 -ListOps=DescribeSpotFleetInstances,DescribeSpotFleetRequestHistory,DescribeVpcEndpointServices"; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func describeSpotFleetInstancesPages(ctx context.Context, conn *ec2.Client, input *ec2.DescribeSpotFleetInstancesInput, fn func(*ec2.DescribeSpotFleetInstancesOutput, bool) bool) error {
	for {
		output, err := conn.DescribeSpotFleetInstances(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
func describeSpotFleetRequestHistoryPages(ctx context.Context, conn *ec2.Client, input *ec2.DescribeSpotFleetRequestHistoryInput, fn func(*ec2.DescribeSpotFleetRequestHistoryOutput, bool) bool) error {
	for {
		output, err := conn.DescribeSpotFleetRequestHistory(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
func describeVPCEndpointServicesPages(ctx context.Context, conn *ec2.Client, input *ec2.DescribeVpcEndpointServicesInput, fn func(*ec2.DescribeVpcEndpointServicesOutput, bool) bool) error {
	for {
		output, err := conn.DescribeVpcEndpointServices(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}