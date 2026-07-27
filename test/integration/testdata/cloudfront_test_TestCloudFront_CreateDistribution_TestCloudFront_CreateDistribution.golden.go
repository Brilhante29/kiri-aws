{
  "Distribution": {
    "ARN": "arn:aws:cloudfront::000000000000:distribution/Ef64a9b8f-db62",
    "DistributionConfig": {
      "CallerReference": "test-create-distribution",
      "Comment": "Test distribution",
      "DefaultCacheBehavior": {
        "TargetOriginId": "myS3Origin",
        "ViewerProtocolPolicy": "allow-all",
        "AllowedMethods": null,
        "CachePolicyId": "658327ea-f89d-4fab-a63d-7e88639e58f6",
        "Compress": null,
        "DefaultTTL": null,
        "FieldLevelEncryptionId": null,
        "ForwardedValues": null,
        "FunctionAssociations": null,
        "GrpcConfig": null,
        "LambdaFunctionAssociations": null,
        "MaxTTL": null,
        "MinTTL": null,
        "OriginRequestPolicyId": null,
        "RealtimeLogConfigArn": null,
        "ResponseHeadersPolicyId": null,
        "SmoothStreaming": null,
        "TrustedKeyGroups": null,
        "TrustedSigners": null
      },
      "Enabled": true,
      "Origins": {
        "Items": [
          {
            "DomainName": "mybucket.s3.amazonaws.com",
            "Id": "myS3Origin",
            "ConnectionAttempts": null,
            "ConnectionTimeout": null,
            "CustomHeaders": null,
            "CustomOriginConfig": null,
            "OriginAccessControlId": null,
            "OriginPath": null,
            "OriginShield": null,
            "ResponseCompletionTimeout": null,
            "S3OriginConfig": {
              "OriginAccessIdentity": "",
              "OriginReadTimeout": null
            },
            "VpcOriginConfig": null
          }
        ],
        "Quantity": 1
      },
      "Aliases": null,
      "AnycastIpListId": null,
      "CacheBehaviors": null,
      "CacheTagConfig": null,
      "ConnectionFunctionAssociation": null,
      "ConnectionMode": "",
      "ContinuousDeploymentPolicyId": null,
      "CustomErrorResponses": null,
      "DefaultRootObject": null,
      "HttpVersion": "http2",
      "IsIPV6Enabled": null,
      "Logging": null,
      "OriginGroups": null,
      "PriceClass": "PriceClass_All",
      "Restrictions": null,
      "Staging": null,
      "TenantConfig": null,
      "ViewerCertificate": {
        "ACMCertificateArn": null,
        "Certificate": null,
        "CertificateSource": "",
        "CloudFrontDefaultCertificate": true,
        "IAMCertificateId": null,
        "MinimumProtocolVersion": "TLSv1",
        "SSLSupportMethod": ""
      },
      "ViewerMtlsConfig": null,
      "WebACLId": null
    },
    "DomainName": "Ef64a9b8f-db62.cloudfront.net",
    "Id": "Ef64a9b8f-db62",
    "InProgressInvalidationBatches": null,
    "LastModifiedTime": "2026-07-27T03:23:26Z",
    "Status": "InProgress",
    "ActiveTrustedKeyGroups": {
      "Enabled": false,
      "Quantity": 0,
      "Items": null
    },
    "ActiveTrustedSigners": {
      "Enabled": false,
      "Quantity": 0,
      "Items": null
    },
    "AliasICPRecordals": null
  },
  "ETag": "E95966b09-85a8-4bb9-b5b5-d4a7a31f",
  "Location": "/2020-05-31/distribution/Ef64a9b8f-db62",
  "ResultMetadata": {}
}