// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class Cluster extends fabric.Resource {
    public readonly applications?: fabric.Property<string[]>;
    public readonly autoscalingRole?: fabric.Property<string>;
    public readonly bootstrapAction?: fabric.Property<{ args?: string[], name: string, path: string }[]>;
    public /*out*/ readonly clusterState: fabric.Property<string>;
    public readonly configurations?: fabric.Property<string>;
    public readonly coreInstanceCount?: fabric.Property<number>;
    public readonly coreInstanceType: fabric.Property<string>;
    public readonly ec2Attributes?: fabric.Property<{ additionalMasterSecurityGroups?: string, additionalSlaveSecurityGroups?: string, emrManagedMasterSecurityGroup?: string, emrManagedSlaveSecurityGroup?: string, instanceProfile: string, keyName?: string, serviceAccessSecurityGroup?: string, subnetId?: string }[]>;
    public readonly keepJobFlowAliveWhenNoSteps: fabric.Property<boolean>;
    public readonly logUri?: fabric.Property<string>;
    public readonly masterInstanceType: fabric.Property<string>;
    public /*out*/ readonly masterPublicDns: fabric.Property<string>;
    public readonly name: fabric.Property<string>;
    public readonly releaseLabel: fabric.Property<string>;
    public readonly securityConfiguration?: fabric.Property<string>;
    public readonly serviceRole: fabric.Property<string>;
    public readonly tags?: fabric.Property<{[key: string]: any}>;
    public readonly terminationProtection: fabric.Property<boolean>;
    public readonly visibleToAllUsers?: fabric.Property<boolean>;

    constructor(urnName: string, args: ClusterArgs) {
        if (args.masterInstanceType === undefined) {
            throw new Error("Missing required property 'masterInstanceType'");
        }
        if (args.releaseLabel === undefined) {
            throw new Error("Missing required property 'releaseLabel'");
        }
        if (args.serviceRole === undefined) {
            throw new Error("Missing required property 'serviceRole'");
        }
        super("aws:emr/cluster:Cluster", urnName, {
            "applications": args.applications,
            "autoscalingRole": args.autoscalingRole,
            "bootstrapAction": args.bootstrapAction,
            "configurations": args.configurations,
            "coreInstanceCount": args.coreInstanceCount,
            "coreInstanceType": args.coreInstanceType,
            "ec2Attributes": args.ec2Attributes,
            "keepJobFlowAliveWhenNoSteps": args.keepJobFlowAliveWhenNoSteps,
            "logUri": args.logUri,
            "masterInstanceType": args.masterInstanceType,
            "name": args.name,
            "releaseLabel": args.releaseLabel,
            "securityConfiguration": args.securityConfiguration,
            "serviceRole": args.serviceRole,
            "tags": args.tags,
            "terminationProtection": args.terminationProtection,
            "visibleToAllUsers": args.visibleToAllUsers,
            "clusterState": undefined,
            "masterPublicDns": undefined,
        });
    }
}

export interface ClusterArgs {
    readonly applications?: fabric.PropertyValue<fabric.PropertyValue<string>>[];
    readonly autoscalingRole?: fabric.PropertyValue<string>;
    readonly bootstrapAction?: fabric.PropertyValue<{ args?: fabric.PropertyValue<fabric.PropertyValue<string>>[], name: fabric.PropertyValue<string>, path: fabric.PropertyValue<string> }>[];
    readonly configurations?: fabric.PropertyValue<string>;
    readonly coreInstanceCount?: fabric.PropertyValue<number>;
    readonly coreInstanceType?: fabric.PropertyValue<string>;
    readonly ec2Attributes?: fabric.PropertyValue<{ additionalMasterSecurityGroups?: fabric.PropertyValue<string>, additionalSlaveSecurityGroups?: fabric.PropertyValue<string>, emrManagedMasterSecurityGroup?: fabric.PropertyValue<string>, emrManagedSlaveSecurityGroup?: fabric.PropertyValue<string>, instanceProfile: fabric.PropertyValue<string>, keyName?: fabric.PropertyValue<string>, serviceAccessSecurityGroup?: fabric.PropertyValue<string>, subnetId?: fabric.PropertyValue<string> }>[];
    readonly keepJobFlowAliveWhenNoSteps?: fabric.PropertyValue<boolean>;
    readonly logUri?: fabric.PropertyValue<string>;
    readonly masterInstanceType: fabric.PropertyValue<string>;
    readonly name?: fabric.PropertyValue<string>;
    readonly releaseLabel: fabric.PropertyValue<string>;
    readonly securityConfiguration?: fabric.PropertyValue<string>;
    readonly serviceRole: fabric.PropertyValue<string>;
    readonly tags?: fabric.PropertyValue<{[key: string]: any}>;
    readonly terminationProtection?: fabric.PropertyValue<boolean>;
    readonly visibleToAllUsers?: fabric.PropertyValue<boolean>;
}
