// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class OpenIdConnectProvider extends fabric.Resource {
    public /*out*/ readonly arn: fabric.Property<string>;
    public readonly clientIdList: fabric.Property<string[]>;
    public readonly thumbprintList: fabric.Property<string[]>;
    public readonly url: fabric.Property<string>;

    constructor(urnName: string, args: OpenIdConnectProviderArgs) {
        if (args.clientIdList === undefined) {
            throw new Error("Missing required property 'clientIdList'");
        }
        if (args.thumbprintList === undefined) {
            throw new Error("Missing required property 'thumbprintList'");
        }
        if (args.url === undefined) {
            throw new Error("Missing required property 'url'");
        }
        super("aws:iam/openIdConnectProvider:OpenIdConnectProvider", urnName, {
            "clientIdList": args.clientIdList,
            "thumbprintList": args.thumbprintList,
            "url": args.url,
            "arn": undefined,
        });
    }
}

export interface OpenIdConnectProviderArgs {
    readonly clientIdList: fabric.PropertyValue<fabric.PropertyValue<string>>[];
    readonly thumbprintList: fabric.PropertyValue<fabric.PropertyValue<string>>[];
    readonly url: fabric.PropertyValue<string>;
}
