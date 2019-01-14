// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2transitgateway

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an EC2 Transit Gateway Route Table propagation.
type RouteTablePropagation struct {
	s *pulumi.ResourceState
}

// NewRouteTablePropagation registers a new resource with the given unique name, arguments, and options.
func NewRouteTablePropagation(ctx *pulumi.Context,
	name string, args *RouteTablePropagationArgs, opts ...pulumi.ResourceOpt) (*RouteTablePropagation, error) {
	if args == nil || args.TransitGatewayAttachmentId == nil {
		return nil, errors.New("missing required argument 'TransitGatewayAttachmentId'")
	}
	if args == nil || args.TransitGatewayRouteTableId == nil {
		return nil, errors.New("missing required argument 'TransitGatewayRouteTableId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["transitGatewayAttachmentId"] = nil
		inputs["transitGatewayRouteTableId"] = nil
	} else {
		inputs["transitGatewayAttachmentId"] = args.TransitGatewayAttachmentId
		inputs["transitGatewayRouteTableId"] = args.TransitGatewayRouteTableId
	}
	inputs["resourceId"] = nil
	inputs["resourceType"] = nil
	s, err := ctx.RegisterResource("aws:ec2transitgateway/routeTablePropagation:RouteTablePropagation", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RouteTablePropagation{s: s}, nil
}

// GetRouteTablePropagation gets an existing RouteTablePropagation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteTablePropagation(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RouteTablePropagationState, opts ...pulumi.ResourceOpt) (*RouteTablePropagation, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["resourceId"] = state.ResourceId
		inputs["resourceType"] = state.ResourceType
		inputs["transitGatewayAttachmentId"] = state.TransitGatewayAttachmentId
		inputs["transitGatewayRouteTableId"] = state.TransitGatewayRouteTableId
	}
	s, err := ctx.ReadResource("aws:ec2transitgateway/routeTablePropagation:RouteTablePropagation", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RouteTablePropagation{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *RouteTablePropagation) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *RouteTablePropagation) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Identifier of the resource
func (r *RouteTablePropagation) ResourceId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceId"])
}

// Type of the resource
func (r *RouteTablePropagation) ResourceType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceType"])
}

// Identifier of EC2 Transit Gateway Attachment.
func (r *RouteTablePropagation) TransitGatewayAttachmentId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["transitGatewayAttachmentId"])
}

// Identifier of EC2 Transit Gateway Route Table.
func (r *RouteTablePropagation) TransitGatewayRouteTableId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["transitGatewayRouteTableId"])
}

// Input properties used for looking up and filtering RouteTablePropagation resources.
type RouteTablePropagationState struct {
	// Identifier of the resource
	ResourceId interface{}
	// Type of the resource
	ResourceType interface{}
	// Identifier of EC2 Transit Gateway Attachment.
	TransitGatewayAttachmentId interface{}
	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId interface{}
}

// The set of arguments for constructing a RouteTablePropagation resource.
type RouteTablePropagationArgs struct {
	// Identifier of EC2 Transit Gateway Attachment.
	TransitGatewayAttachmentId interface{}
	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId interface{}
}