{
  "VirtualNode": {
    "MeshName": "vn-mesh",
    "Metadata": {
      "Arn": "arn:aws:appmesh:us-east-1:123456789012:mesh/vn-mesh/virtualNode/test-vn",
      "CreatedAt": "2026-07-27T03:22:53.671Z",
      "LastUpdatedAt": "2026-07-27T03:22:53.671Z",
      "MeshOwner": "123456789012",
      "ResourceOwner": "123456789012",
      "Uid": "caba06f4-9671-484a-943a-37d6dcf3e90a",
      "Version": 1
    },
    "Spec": {
      "BackendDefaults": null,
      "Backends": null,
      "Listeners": [
        {
          "PortMapping": {
            "Port": 8080,
            "Protocol": "http"
          },
          "ConnectionPool": null,
          "HealthCheck": null,
          "OutlierDetection": null,
          "Timeout": null,
          "Tls": null
        }
      ],
      "Logging": null,
      "ServiceDiscovery": {
        "Value": {
          "Hostname": "test.local",
          "IpPreference": "",
          "ResponseType": ""
        }
      }
    },
    "Status": {
      "Status": "ACTIVE"
    },
    "VirtualNodeName": "test-vn"
  },
  "ResultMetadata": {}
}