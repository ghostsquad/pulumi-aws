// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudFront.Outputs
{

    [OutputType]
    public sealed class CachePolicyParametersInCacheKeyAndForwardedToOrigin
    {
        /// <summary>
        /// Object that determines whether any cookies in viewer requests (and if so, which cookies) are included in the cache key and automatically included in requests that CloudFront sends to the origin. See Cookies Config for more information.
        /// </summary>
        public readonly Outputs.CachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig CookiesConfig;
        /// <summary>
        /// A flag that can affect whether the Accept-Encoding HTTP header is included in the cache key and included in requests that CloudFront sends to the origin.
        /// </summary>
        public readonly bool? EnableAcceptEncodingBrotli;
        /// <summary>
        /// A flag that can affect whether the Accept-Encoding HTTP header is included in the cache key and included in requests that CloudFront sends to the origin.
        /// </summary>
        public readonly bool? EnableAcceptEncodingGzip;
        /// <summary>
        /// Object that determines whether any HTTP headers (and if so, which headers) are included in the cache key and automatically included in requests that CloudFront sends to the origin. See Headers Config for more information.
        /// </summary>
        public readonly Outputs.CachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig HeadersConfig;
        /// <summary>
        /// Object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in the cache key and automatically included in requests that CloudFront sends to the origin. See Query String Config for more information.
        /// </summary>
        public readonly Outputs.CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig QueryStringsConfig;

        [OutputConstructor]
        private CachePolicyParametersInCacheKeyAndForwardedToOrigin(
            Outputs.CachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig cookiesConfig,

            bool? enableAcceptEncodingBrotli,

            bool? enableAcceptEncodingGzip,

            Outputs.CachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig headersConfig,

            Outputs.CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig queryStringsConfig)
        {
            CookiesConfig = cookiesConfig;
            EnableAcceptEncodingBrotli = enableAcceptEncodingBrotli;
            EnableAcceptEncodingGzip = enableAcceptEncodingGzip;
            HeadersConfig = headersConfig;
            QueryStringsConfig = queryStringsConfig;
        }
    }
}
