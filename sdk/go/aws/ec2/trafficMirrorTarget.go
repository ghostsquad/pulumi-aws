// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ec2

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an Traffic mirror target.
// Read [limits and considerations](https://docs.aws.amazon.com/vpc/latest/mirroring/traffic-mirroring-considerations.html) for traffic mirroring
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_traffic_mirror_target.html.markdown.
type TrafficMirrorTarget struct {
	pulumi.CustomResourceState

	// A description of the traffic mirror session.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The network interface ID that is associated with the target.
	NetworkInterfaceId pulumi.StringPtrOutput `pulumi:"networkInterfaceId"`
	// The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.
	NetworkLoadBalancerArn pulumi.StringPtrOutput `pulumi:"networkLoadBalancerArn"`
}

// NewTrafficMirrorTarget registers a new resource with the given unique name, arguments, and options.
func NewTrafficMirrorTarget(ctx *pulumi.Context,
	name string, args *TrafficMirrorTargetArgs, opts ...pulumi.ResourceOption) (*TrafficMirrorTarget, error) {
	if args == nil {
		args = &TrafficMirrorTargetArgs{}
	}
	var resource TrafficMirrorTarget
	err := ctx.RegisterResource("aws:ec2/trafficMirrorTarget:TrafficMirrorTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrafficMirrorTarget gets an existing TrafficMirrorTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrafficMirrorTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficMirrorTargetState, opts ...pulumi.ResourceOption) (*TrafficMirrorTarget, error) {
	var resource TrafficMirrorTarget
	err := ctx.ReadResource("aws:ec2/trafficMirrorTarget:TrafficMirrorTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrafficMirrorTarget resources.
type trafficMirrorTargetState struct {
	// A description of the traffic mirror session.
	Description *string `pulumi:"description"`
	// The network interface ID that is associated with the target.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.
	NetworkLoadBalancerArn *string `pulumi:"networkLoadBalancerArn"`
}

type TrafficMirrorTargetState struct {
	// A description of the traffic mirror session.
	Description pulumi.StringPtrInput
	// The network interface ID that is associated with the target.
	NetworkInterfaceId pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.
	NetworkLoadBalancerArn pulumi.StringPtrInput
}

func (TrafficMirrorTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficMirrorTargetState)(nil)).Elem()
}

type trafficMirrorTargetArgs struct {
	// A description of the traffic mirror session.
	Description *string `pulumi:"description"`
	// The network interface ID that is associated with the target.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.
	NetworkLoadBalancerArn *string `pulumi:"networkLoadBalancerArn"`
}

// The set of arguments for constructing a TrafficMirrorTarget resource.
type TrafficMirrorTargetArgs struct {
	// A description of the traffic mirror session.
	Description pulumi.StringPtrInput
	// The network interface ID that is associated with the target.
	NetworkInterfaceId pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.
	NetworkLoadBalancerArn pulumi.StringPtrInput
}

func (TrafficMirrorTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficMirrorTargetArgs)(nil)).Elem()
}