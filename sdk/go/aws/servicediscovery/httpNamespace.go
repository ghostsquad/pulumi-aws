// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicediscovery

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type HttpNamespace struct {
	s *pulumi.ResourceState
}

// NewHttpNamespace registers a new resource with the given unique name, arguments, and options.
func NewHttpNamespace(ctx *pulumi.Context,
	name string, args *HttpNamespaceArgs, opts ...pulumi.ResourceOpt) (*HttpNamespace, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["name"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["name"] = args.Name
	}
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:servicediscovery/httpNamespace:HttpNamespace", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &HttpNamespace{s: s}, nil
}

// GetHttpNamespace gets an existing HttpNamespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHttpNamespace(ctx *pulumi.Context,
	name string, id pulumi.ID, state *HttpNamespaceState, opts ...pulumi.ResourceOpt) (*HttpNamespace, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["description"] = state.Description
		inputs["name"] = state.Name
	}
	s, err := ctx.ReadResource("aws:servicediscovery/httpNamespace:HttpNamespace", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &HttpNamespace{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *HttpNamespace) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *HttpNamespace) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The ARN that Amazon Route 53 assigns to the namespace when you create it.
func (r *HttpNamespace) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// The description that you specify for the namespace when you create it.
func (r *HttpNamespace) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// The name of the http namespace.
func (r *HttpNamespace) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Input properties used for looking up and filtering HttpNamespace resources.
type HttpNamespaceState struct {
	// The ARN that Amazon Route 53 assigns to the namespace when you create it.
	Arn interface{}
	// The description that you specify for the namespace when you create it.
	Description interface{}
	// The name of the http namespace.
	Name interface{}
}

// The set of arguments for constructing a HttpNamespace resource.
type HttpNamespaceArgs struct {
	// The description that you specify for the namespace when you create it.
	Description interface{}
	// The name of the http namespace.
	Name interface{}
}