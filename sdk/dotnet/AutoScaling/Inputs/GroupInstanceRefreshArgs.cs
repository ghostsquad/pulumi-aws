// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AutoScaling.Inputs
{

    public sealed class GroupInstanceRefreshArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Override default parameters for Instance Refresh.
        /// </summary>
        [Input("preferences")]
        public Input<Inputs.GroupInstanceRefreshPreferencesArgs>? Preferences { get; set; }

        /// <summary>
        /// The strategy to use for instance refresh. The only allowed value is `Rolling`. See [StartInstanceRefresh Action](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_StartInstanceRefresh.html#API_StartInstanceRefresh_RequestParameters) for more information.
        /// </summary>
        [Input("strategy", required: true)]
        public Input<string> Strategy { get; set; } = null!;

        [Input("triggers")]
        private InputList<string>? _triggers;

        /// <summary>
        /// Set of additional property names that will trigger an Instance Refresh. A refresh will always be triggered by a change in any of `launch_configuration`, `launch_template`, or `mixed_instances_policy`.
        /// </summary>
        public InputList<string> Triggers
        {
            get => _triggers ?? (_triggers = new InputList<string>());
            set => _triggers = value;
        }

        public GroupInstanceRefreshArgs()
        {
        }
    }
}