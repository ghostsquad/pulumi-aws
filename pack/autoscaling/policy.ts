// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class Policy extends fabric.Resource {
    public readonly adjustmentType: fabric.Property<string>;
    public /*out*/ readonly arn: fabric.Property<string>;
    public readonly autoscalingGroupName: fabric.Property<string>;
    public readonly cooldown?: fabric.Property<number>;
    public readonly estimatedInstanceWarmup?: fabric.Property<number>;
    public readonly metricAggregationType: fabric.Property<string>;
    public readonly minAdjustmentMagnitude?: fabric.Property<number>;
    public readonly minAdjustmentStep?: fabric.Property<number>;
    public readonly name: fabric.Property<string>;
    public readonly policyType?: fabric.Property<string>;
    public readonly scalingAdjustment?: fabric.Property<number>;
    public readonly stepAdjustment?: fabric.Property<{ metricIntervalLowerBound?: string, metricIntervalUpperBound?: string, scalingAdjustment: number }[]>;

    constructor(urnName: string, args: PolicyArgs) {
        if (args.adjustmentType === undefined) {
            throw new Error("Missing required property 'adjustmentType'");
        }
        if (args.autoscalingGroupName === undefined) {
            throw new Error("Missing required property 'autoscalingGroupName'");
        }
        super("aws:autoscaling/policy:Policy", urnName, {
            "adjustmentType": args.adjustmentType,
            "autoscalingGroupName": args.autoscalingGroupName,
            "cooldown": args.cooldown,
            "estimatedInstanceWarmup": args.estimatedInstanceWarmup,
            "metricAggregationType": args.metricAggregationType,
            "minAdjustmentMagnitude": args.minAdjustmentMagnitude,
            "minAdjustmentStep": args.minAdjustmentStep,
            "name": args.name,
            "policyType": args.policyType,
            "scalingAdjustment": args.scalingAdjustment,
            "stepAdjustment": args.stepAdjustment,
            "arn": undefined,
        });
    }
}

export interface PolicyArgs {
    readonly adjustmentType: fabric.PropertyValue<string>;
    readonly autoscalingGroupName: fabric.PropertyValue<string>;
    readonly cooldown?: fabric.PropertyValue<number>;
    readonly estimatedInstanceWarmup?: fabric.PropertyValue<number>;
    readonly metricAggregationType?: fabric.PropertyValue<string>;
    readonly minAdjustmentMagnitude?: fabric.PropertyValue<number>;
    readonly minAdjustmentStep?: fabric.PropertyValue<number>;
    readonly name?: fabric.PropertyValue<string>;
    readonly policyType?: fabric.PropertyValue<string>;
    readonly scalingAdjustment?: fabric.PropertyValue<number>;
    readonly stepAdjustment?: fabric.PropertyValue<{ metricIntervalLowerBound?: fabric.PropertyValue<string>, metricIntervalUpperBound?: fabric.PropertyValue<string>, scalingAdjustment: fabric.PropertyValue<number> }>[];
}
