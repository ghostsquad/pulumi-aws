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
    public sealed class BucketV2ReplicationConfigurationRuleFilter
    {
        /// <summary>
        /// Object keyname prefix identifying one or more objects to which the rule applies
        /// </summary>
        public readonly string? Prefix;
        /// <summary>
        /// A map of tags to assign to the bucket. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;

        [OutputConstructor]
        private BucketV2ReplicationConfigurationRuleFilter(
            string? prefix,

            ImmutableDictionary<string, string>? tags)
        {
            Prefix = prefix;
            Tags = tags;
        }
    }
}
