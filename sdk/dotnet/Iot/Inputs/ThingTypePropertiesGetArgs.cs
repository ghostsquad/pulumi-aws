// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iot.Inputs
{

    public sealed class ThingTypePropertiesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the thing type.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("searchableAttributes")]
        private InputList<string>? _searchableAttributes;

        /// <summary>
        /// A list of searchable thing attribute names.
        /// </summary>
        public InputList<string> SearchableAttributes
        {
            get => _searchableAttributes ?? (_searchableAttributes = new InputList<string>());
            set => _searchableAttributes = value;
        }

        public ThingTypePropertiesGetArgs()
        {
        }
    }
}