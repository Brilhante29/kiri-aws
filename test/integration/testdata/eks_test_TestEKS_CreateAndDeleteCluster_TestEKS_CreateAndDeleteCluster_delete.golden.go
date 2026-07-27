{
  "Cluster": {
    "AccessConfig": null,
    "Arn": "arn:aws:eks:us-east-1:123456789012:cluster/test-cluster",
    "CertificateAuthority": {
      "Data": "ZmFrZS1jZXJ0aWZpY2F0ZS1hdXRob3JpdHktZGF0YQ=="
    },
    "ClientRequestToken": null,
    "ComputeConfig": null,
    "ConnectorConfig": null,
    "ControlPlaneScalingConfig": null,
    "CreatedAt": "2026-07-27T03:22:59.279Z",
    "DeletionProtection": null,
    "EncryptionConfig": null,
    "Endpoint": "https://61020a05.gr7.us-east-1.eks.amazonaws.com",
    "Health": {
      "Issues": null
    },
    "Id": null,
    "Identity": {
      "Oidc": {
        "Issuer": "https://oidc.eks.us-east-1.amazonaws.com/id/7141ed8e-57ba-493a-bb7c-0209bc8a"
      }
    },
    "KubernetesNetworkConfig": {
      "ElasticLoadBalancing": null,
      "IpFamily": "ipv4",
      "ServiceIpv4Cidr": "10.100.0.0/16",
      "ServiceIpv6Cidr": null
    },
    "Logging": null,
    "Name": "test-cluster",
    "OutpostConfig": null,
    "PlatformVersion": "eks.1",
    "RemoteNetworkConfig": null,
    "ResourcesVpcConfig": {
      "ClusterSecurityGroupId": "sg-4bbe4c1f-a87b-48b",
      "ControlPlaneEgressMode": "",
      "EndpointPrivateAccess": false,
      "EndpointPublicAccess": true,
      "PublicAccessCidrs": [
        "0.0.0.0/0"
      ],
      "SecurityGroupIds": null,
      "SubnetIds": [
        "subnet-12345678",
        "subnet-87654321"
      ],
      "VpcId": "vpc-c14d4734-f7d9-49d"
    },
    "RoleArn": "arn:aws:iam::123456789012:role/eks-cluster-role",
    "Status": "DELETING",
    "StorageConfig": null,
    "Tags": null,
    "UpgradePolicy": null,
    "Version": "1.29",
    "ZonalShiftConfig": null
  },
  "ResultMetadata": {}
}