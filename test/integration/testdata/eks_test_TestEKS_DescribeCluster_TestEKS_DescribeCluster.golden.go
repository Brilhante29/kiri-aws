{
  "Cluster": {
    "AccessConfig": null,
    "Arn": "arn:aws:eks:us-east-1:123456789012:cluster/test-describe-cluster",
    "CertificateAuthority": {
      "Data": "ZmFrZS1jZXJ0aWZpY2F0ZS1hdXRob3JpdHktZGF0YQ=="
    },
    "ClientRequestToken": null,
    "ComputeConfig": null,
    "ConnectorConfig": null,
    "ControlPlaneScalingConfig": null,
    "CreatedAt": "2026-07-27T03:22:59.311Z",
    "DeletionProtection": null,
    "EncryptionConfig": null,
    "Endpoint": "https://292305aa.gr7.us-east-1.eks.amazonaws.com",
    "Health": {
      "Issues": null
    },
    "Id": null,
    "Identity": {
      "Oidc": {
        "Issuer": "https://oidc.eks.us-east-1.amazonaws.com/id/1f0125a6-95e2-450c-a2bb-281366c4"
      }
    },
    "KubernetesNetworkConfig": {
      "ElasticLoadBalancing": null,
      "IpFamily": "ipv4",
      "ServiceIpv4Cidr": "10.100.0.0/16",
      "ServiceIpv6Cidr": null
    },
    "Logging": null,
    "Name": "test-describe-cluster",
    "OutpostConfig": null,
    "PlatformVersion": "eks.1",
    "RemoteNetworkConfig": null,
    "ResourcesVpcConfig": {
      "ClusterSecurityGroupId": "sg-294b911a-6351-49b",
      "ControlPlaneEgressMode": "",
      "EndpointPrivateAccess": false,
      "EndpointPublicAccess": true,
      "PublicAccessCidrs": [
        "0.0.0.0/0"
      ],
      "SecurityGroupIds": null,
      "SubnetIds": [
        "subnet-12345678"
      ],
      "VpcId": "vpc-d064f1c2-bedd-4bd"
    },
    "RoleArn": "arn:aws:iam::123456789012:role/eks-cluster-role",
    "Status": "ACTIVE",
    "StorageConfig": null,
    "Tags": null,
    "UpgradePolicy": null,
    "Version": "1.29",
    "ZonalShiftConfig": null
  },
  "ResultMetadata": {}
}