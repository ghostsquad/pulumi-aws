// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Information about an RDS Certificate.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/rds"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := true
// 		_, err := rds.GetCertificate(ctx, &rds.GetCertificateArgs{
// 			LatestValidTill: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetCertificate(ctx *pulumi.Context, args *GetCertificateArgs, opts ...pulumi.InvokeOption) (*GetCertificateResult, error) {
	var rv GetCertificateResult
	err := ctx.Invoke("aws:rds/getCertificate:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCertificate.
type GetCertificateArgs struct {
	// Certificate identifier. For example, `rds-ca-2019`.
	Id *string `pulumi:"id"`
	// When enabled, returns the certificate with the latest `ValidTill`.
	LatestValidTill *bool `pulumi:"latestValidTill"`
}

// A collection of values returned by getCertificate.
type GetCertificateResult struct {
	// Amazon Resource Name (ARN) of the certificate.
	Arn string `pulumi:"arn"`
	// Type of certificate. For example, `CA`.
	CertificateType string `pulumi:"certificateType"`
	// Boolean whether there is an override for the default certificate identifier.
	CustomerOverride bool `pulumi:"customerOverride"`
	// If there is an override for the default certificate identifier, when the override expires.
	CustomerOverrideValidTill string `pulumi:"customerOverrideValidTill"`
	Id                        string `pulumi:"id"`
	LatestValidTill           *bool  `pulumi:"latestValidTill"`
	// Thumbprint of the certificate.
	Thumbprint string `pulumi:"thumbprint"`
	// [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) of certificate starting validity date.
	ValidFrom string `pulumi:"validFrom"`
	// [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) of certificate ending validity date.
	ValidTill string `pulumi:"validTill"`
}