{
  "Nodegroup": {
    "AmiType": "AL2_x86_64",
    "CapacityType": "ON_DEMAND",
    "ClusterName": "test-describe-nodegroup-cluster",
    "CreatedAt": "2026-07-27T03:22:59.37Z",
    "DiskSize": null,
    "Health": {
      "Issues": null
    },
    "InstanceTypes": [
      "t3.medium"
    ],
    "Labels": null,
    "LaunchTemplate": null,
    "ModifiedAt": "2026-07-27T03:22:59.37Z",
    "NodeRepairConfig": null,
    "NodeRole": "arn:aws:iam::123456789012:role/eks-nodegroup-role",
    "NodegroupArn": "arn:aws:eks:us-east-1:123456789012:nodegroup/test-describe-nodegroup-cluster/test-describe-nodegroup/346ff1e4",
    "NodegroupName": "test-describe-nodegroup",
    "ReleaseVersion": "1.29-20231116",
    "RemoteAccess": null,
    "Resources": {
      "AutoScalingGroups": [
        {
          "Name": "eks-test-describe-nodegroup-52539ac4"
        }
      ],
      "RemoteAccessSecurityGroup": null
    },
    "ScalingConfig": {
      "DesiredSize": 1,
      "MaxSize": 2,
      "MinSize": 1
    },
    "Status": "ACTIVE",
    "Subnets": [
      "subnet-12345678"
    ],
    "Tags": null,
    "Taints": null,
    "UpdateConfig": null,
    "Version": "1.29",
    "WarmPoolConfig": null
  },
  "ResultMetadata": {}
}