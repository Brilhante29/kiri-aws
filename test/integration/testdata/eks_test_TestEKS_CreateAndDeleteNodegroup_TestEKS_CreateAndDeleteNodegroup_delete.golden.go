{
  "Nodegroup": {
    "AmiType": "AL2_x86_64",
    "CapacityType": "ON_DEMAND",
    "ClusterName": "test-nodegroup-cluster",
    "CreatedAt": "2026-07-27T03:22:59.333Z",
    "DiskSize": null,
    "Health": {
      "Issues": null
    },
    "InstanceTypes": [
      "t3.medium"
    ],
    "Labels": null,
    "LaunchTemplate": null,
    "ModifiedAt": "2026-07-27T03:22:59.333Z",
    "NodeRepairConfig": null,
    "NodeRole": "arn:aws:iam::123456789012:role/eks-nodegroup-role",
    "NodegroupArn": "arn:aws:eks:us-east-1:123456789012:nodegroup/test-nodegroup-cluster/test-nodegroup/bd95153a",
    "NodegroupName": "test-nodegroup",
    "ReleaseVersion": "1.29-20231116",
    "RemoteAccess": null,
    "Resources": {
      "AutoScalingGroups": [
        {
          "Name": "eks-test-nodegroup-2271f275"
        }
      ],
      "RemoteAccessSecurityGroup": null
    },
    "ScalingConfig": {
      "DesiredSize": 2,
      "MaxSize": 3,
      "MinSize": 1
    },
    "Status": "DELETING",
    "Subnets": [
      "subnet-12345678",
      "subnet-87654321"
    ],
    "Tags": null,
    "Taints": null,
    "UpdateConfig": null,
    "Version": "1.29",
    "WarmPoolConfig": null
  },
  "ResultMetadata": {}
}