// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a managed prefix list resource.
//
// > **NOTE on `maxEntries`:** When you reference a Prefix List in a resource,
// the maximum number of entries for the prefix lists counts as the same number of rules
// or entries for the resource. For example, if you create a prefix list with a maximum
// of 20 entries and you reference that prefix list in a security group rule, this counts
// as 20 rules for the security group.
//
// ## Example Usage
//
// Basic usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ec2.NewManagedPrefixList(ctx, "example", &ec2.ManagedPrefixListArgs{
// 			AddressFamily: pulumi.String("IPv4"),
// 			MaxEntries:    pulumi.Int(5),
// 			Entries: ec2.ManagedPrefixListEntryArray{
// 				&ec2.ManagedPrefixListEntryArgs{
// 					Cidr:        pulumi.Any(aws_vpc.Example.Cidr_block),
// 					Description: pulumi.String("Primary"),
// 				},
// 				&ec2.ManagedPrefixListEntryArgs{
// 					Cidr:        pulumi.Any(aws_vpc_ipv4_cidr_block_association.Example.Cidr_block),
// 					Description: pulumi.String("Secondary"),
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"Env": pulumi.String("live"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Prefix Lists can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:ec2/managedPrefixList:ManagedPrefixList default pl-0570a1d2d725c16be
// ```
type ManagedPrefixList struct {
	pulumi.CustomResourceState

	// The address family (`IPv4` or `IPv6`) of
	// this prefix list.
	AddressFamily pulumi.StringOutput `pulumi:"addressFamily"`
	// The ARN of the prefix list.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Can be specified multiple times for each prefix list entry.
	// Each entry block supports fields documented below. Different entries may have
	// overlapping CIDR blocks, but a particular CIDR should not be duplicated.
	Entries ManagedPrefixListEntryArrayOutput `pulumi:"entries"`
	// The maximum number of entries that
	// this prefix list can contain.
	MaxEntries pulumi.IntOutput `pulumi:"maxEntries"`
	// The name of this resource. The name must not start with `com.amazonaws`.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the AWS account that owns this prefix list.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// A map of tags to assign to this resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The latest version of this prefix list.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewManagedPrefixList registers a new resource with the given unique name, arguments, and options.
func NewManagedPrefixList(ctx *pulumi.Context,
	name string, args *ManagedPrefixListArgs, opts ...pulumi.ResourceOption) (*ManagedPrefixList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AddressFamily == nil {
		return nil, errors.New("invalid value for required argument 'AddressFamily'")
	}
	if args.MaxEntries == nil {
		return nil, errors.New("invalid value for required argument 'MaxEntries'")
	}
	var resource ManagedPrefixList
	err := ctx.RegisterResource("aws:ec2/managedPrefixList:ManagedPrefixList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedPrefixList gets an existing ManagedPrefixList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedPrefixList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedPrefixListState, opts ...pulumi.ResourceOption) (*ManagedPrefixList, error) {
	var resource ManagedPrefixList
	err := ctx.ReadResource("aws:ec2/managedPrefixList:ManagedPrefixList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedPrefixList resources.
type managedPrefixListState struct {
	// The address family (`IPv4` or `IPv6`) of
	// this prefix list.
	AddressFamily *string `pulumi:"addressFamily"`
	// The ARN of the prefix list.
	Arn *string `pulumi:"arn"`
	// Can be specified multiple times for each prefix list entry.
	// Each entry block supports fields documented below. Different entries may have
	// overlapping CIDR blocks, but a particular CIDR should not be duplicated.
	Entries []ManagedPrefixListEntry `pulumi:"entries"`
	// The maximum number of entries that
	// this prefix list can contain.
	MaxEntries *int `pulumi:"maxEntries"`
	// The name of this resource. The name must not start with `com.amazonaws`.
	Name *string `pulumi:"name"`
	// The ID of the AWS account that owns this prefix list.
	OwnerId *string `pulumi:"ownerId"`
	// A map of tags to assign to this resource.
	Tags map[string]string `pulumi:"tags"`
	// The latest version of this prefix list.
	Version *int `pulumi:"version"`
}

type ManagedPrefixListState struct {
	// The address family (`IPv4` or `IPv6`) of
	// this prefix list.
	AddressFamily pulumi.StringPtrInput
	// The ARN of the prefix list.
	Arn pulumi.StringPtrInput
	// Can be specified multiple times for each prefix list entry.
	// Each entry block supports fields documented below. Different entries may have
	// overlapping CIDR blocks, but a particular CIDR should not be duplicated.
	Entries ManagedPrefixListEntryArrayInput
	// The maximum number of entries that
	// this prefix list can contain.
	MaxEntries pulumi.IntPtrInput
	// The name of this resource. The name must not start with `com.amazonaws`.
	Name pulumi.StringPtrInput
	// The ID of the AWS account that owns this prefix list.
	OwnerId pulumi.StringPtrInput
	// A map of tags to assign to this resource.
	Tags pulumi.StringMapInput
	// The latest version of this prefix list.
	Version pulumi.IntPtrInput
}

func (ManagedPrefixListState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedPrefixListState)(nil)).Elem()
}

type managedPrefixListArgs struct {
	// The address family (`IPv4` or `IPv6`) of
	// this prefix list.
	AddressFamily string `pulumi:"addressFamily"`
	// Can be specified multiple times for each prefix list entry.
	// Each entry block supports fields documented below. Different entries may have
	// overlapping CIDR blocks, but a particular CIDR should not be duplicated.
	Entries []ManagedPrefixListEntry `pulumi:"entries"`
	// The maximum number of entries that
	// this prefix list can contain.
	MaxEntries int `pulumi:"maxEntries"`
	// The name of this resource. The name must not start with `com.amazonaws`.
	Name *string `pulumi:"name"`
	// A map of tags to assign to this resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ManagedPrefixList resource.
type ManagedPrefixListArgs struct {
	// The address family (`IPv4` or `IPv6`) of
	// this prefix list.
	AddressFamily pulumi.StringInput
	// Can be specified multiple times for each prefix list entry.
	// Each entry block supports fields documented below. Different entries may have
	// overlapping CIDR blocks, but a particular CIDR should not be duplicated.
	Entries ManagedPrefixListEntryArrayInput
	// The maximum number of entries that
	// this prefix list can contain.
	MaxEntries pulumi.IntInput
	// The name of this resource. The name must not start with `com.amazonaws`.
	Name pulumi.StringPtrInput
	// A map of tags to assign to this resource.
	Tags pulumi.StringMapInput
}

func (ManagedPrefixListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedPrefixListArgs)(nil)).Elem()
}

type ManagedPrefixListInput interface {
	pulumi.Input

	ToManagedPrefixListOutput() ManagedPrefixListOutput
	ToManagedPrefixListOutputWithContext(ctx context.Context) ManagedPrefixListOutput
}

func (ManagedPrefixList) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedPrefixList)(nil)).Elem()
}

func (i ManagedPrefixList) ToManagedPrefixListOutput() ManagedPrefixListOutput {
	return i.ToManagedPrefixListOutputWithContext(context.Background())
}

func (i ManagedPrefixList) ToManagedPrefixListOutputWithContext(ctx context.Context) ManagedPrefixListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedPrefixListOutput)
}

type ManagedPrefixListOutput struct {
	*pulumi.OutputState
}

func (ManagedPrefixListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedPrefixListOutput)(nil)).Elem()
}

func (o ManagedPrefixListOutput) ToManagedPrefixListOutput() ManagedPrefixListOutput {
	return o
}

func (o ManagedPrefixListOutput) ToManagedPrefixListOutputWithContext(ctx context.Context) ManagedPrefixListOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ManagedPrefixListOutput{})
}