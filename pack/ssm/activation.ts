// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class Activation extends fabric.Resource {
    public readonly description?: fabric.Property<string>;
    public readonly expirationDate?: fabric.Property<string>;
    public /*out*/ readonly expired: fabric.Property<string>;
    public readonly iamRole: fabric.Property<string>;
    public readonly name: fabric.Property<string>;
    public /*out*/ readonly registrationCount: fabric.Property<number>;
    public readonly registrationLimit?: fabric.Property<number>;

    constructor(urnName: string, args: ActivationArgs) {
        if (args.iamRole === undefined) {
            throw new Error("Missing required property 'iamRole'");
        }
        super("aws:ssm/activation:Activation", urnName, {
            "description": args.description,
            "expirationDate": args.expirationDate,
            "iamRole": args.iamRole,
            "name": args.name,
            "registrationLimit": args.registrationLimit,
            "expired": undefined,
            "registrationCount": undefined,
        });
    }
}

export interface ActivationArgs {
    readonly description?: fabric.PropertyValue<string>;
    readonly expirationDate?: fabric.PropertyValue<string>;
    readonly iamRole: fabric.PropertyValue<string>;
    readonly name?: fabric.PropertyValue<string>;
    readonly registrationLimit?: fabric.PropertyValue<number>;
}
