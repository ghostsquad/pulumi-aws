// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ssm.Outputs
{

    [OutputType]
    public sealed class MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig
    {
        /// <summary>
        /// An Amazon Resource Name (ARN) for a Simple Notification Service (SNS) topic. Run Command pushes notifications about command status changes to this topic.
        /// </summary>
        public readonly string? NotificationArn;
        /// <summary>
        /// The different events for which you can receive notifications. Valid values: `All`, `InProgress`, `Success`, `TimedOut`, `Cancelled`, and `Failed`
        /// </summary>
        public readonly ImmutableArray<string> NotificationEvents;
        /// <summary>
        /// When specified with `Command`, receive notification when the status of a command changes. When specified with `Invocation`, for commands sent to multiple instances, receive notification on a per-instance basis when the status of a command changes. Valid values: `Command` and `Invocation`
        /// </summary>
        public readonly string? NotificationType;

        [OutputConstructor]
        private MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig(
            string? notificationArn,

            ImmutableArray<string> notificationEvents,

            string? notificationType)
        {
            NotificationArn = notificationArn;
            NotificationEvents = notificationEvents;
            NotificationType = notificationType;
        }
    }
}