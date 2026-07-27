{
  "Marker": null,
  "ReplicationGroups": [
    {
      "ARN": "arn:aws:elasticache:us-east-1:000000000000:replicationgroup:test-replication-group",
      "AtRestEncryptionEnabled": false,
      "AuthTokenEnabled": false,
      "AuthTokenLastModifiedDate": null,
      "AutoMinorVersionUpgrade": false,
      "AutomaticFailover": "disabled",
      "CacheNodeType": "cache.t3.micro",
      "ClusterEnabled": false,
      "ClusterMode": "",
      "ConfigurationEndpoint": {
        "Address": "test-replication-group.ea40a98e.clustercfg.us-east-1.cache.amazonaws.com",
        "Port": 6379
      },
      "DataTiering": "",
      "Description": "Test replication group",
      "Durability": "",
      "EffectiveDurability": "",
      "Engine": null,
      "GlobalReplicationGroupInfo": null,
      "IpDiscovery": "",
      "KmsKeyId": null,
      "LogDeliveryConfigurations": null,
      "MemberClusters": [],
      "MemberClustersOutpostArns": null,
      "MultiAZ": "disabled",
      "NetworkType": "",
      "NodeGroups": [
        {
          "NodeGroupId": "0001",
          "NodeGroupMembers": [
            {
              "CacheClusterId": "test-replication-group-0001-0001",
              "CacheNodeId": "0001",
              "CurrentRole": "primary",
              "PreferredAvailabilityZone": "us-east-1a",
              "PreferredOutpostArn": null,
              "ReadEndpoint": {
                "Address": "test-replication-group-0001-0001.bc6153cc.us-east-1.cache.amazonaws.com",
                "Port": 6379
              }
            }
          ],
          "PrimaryEndpoint": {
            "Address": "test-replication-group-0001.680c4de6.us-east-1.cache.amazonaws.com",
            "Port": 6379
          },
          "ReaderEndpoint": {
            "Address": "test-replication-group-0001-ro.f45ce032.us-east-1.cache.amazonaws.com",
            "Port": 6379
          },
          "Slots": null,
          "Status": "available"
        }
      ],
      "PendingModifiedValues": null,
      "ReplicationGroupCreateTime": "2026-07-27T03:22:59.488Z",
      "ReplicationGroupId": "test-replication-group",
      "SnapshotRetentionLimit": 0,
      "SnapshotWindow": null,
      "SnapshottingClusterId": null,
      "Status": "available",
      "StorageEncryptionType": "",
      "TransitEncryptionEnabled": false,
      "TransitEncryptionMode": "",
      "UserGroupIds": null
    }
  ],
  "ResultMetadata": {}
}