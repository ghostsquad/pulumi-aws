// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.S3.Outputs
{

    [OutputType]
    public sealed class BucketV2ServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault
    {
        /// <summary>
        /// (optional) The AWS KMS master key ID used for the SSE-KMS encryption.
        /// </summary>
        public readonly string? KmsMasterKeyId;
        /// <summary>
        /// (required) The server-side encryption algorithm used.
        /// </summary>
        public readonly string? SseAlgorithm;

        [OutputConstructor]
        private BucketV2ServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault(
            string? kmsMasterKeyId,

            string? sseAlgorithm)
        {
            KmsMasterKeyId = kmsMasterKeyId;
            SseAlgorithm = sseAlgorithm;
        }
    }
}
