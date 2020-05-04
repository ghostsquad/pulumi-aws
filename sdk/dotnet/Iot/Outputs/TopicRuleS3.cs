// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iot.Outputs
{

    [OutputType]
    public sealed class TopicRuleS3
    {
        /// <summary>
        /// The Amazon S3 bucket name.
        /// </summary>
        public readonly string BucketName;
        /// <summary>
        /// The object key.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// The IAM role ARN that allows access to the CloudWatch alarm.
        /// </summary>
        public readonly string RoleArn;

        [OutputConstructor]
        private TopicRuleS3(
            string bucketName,

            string key,

            string roleArn)
        {
            BucketName = bucketName;
            Key = key;
            RoleArn = roleArn;
        }
    }
}