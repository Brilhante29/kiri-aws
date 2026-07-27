{
  "Route": {
    "MeshName": "route-mesh",
    "Metadata": {
      "Arn": "arn:aws:appmesh:us-east-1:123456789012:mesh/route-mesh/virtualRouter/route-vr/route/test-route",
      "CreatedAt": "2026-07-27T03:22:53.795Z",
      "LastUpdatedAt": "2026-07-27T03:22:53.795Z",
      "MeshOwner": "123456789012",
      "ResourceOwner": "123456789012",
      "Uid": "ba6af896-6afd-4a71-919c-92c39a6c6f6c",
      "Version": 1
    },
    "RouteName": "test-route",
    "Spec": {
      "GrpcRoute": null,
      "Http2Route": null,
      "HttpRoute": {
        "Action": {
          "WeightedTargets": [
            {
              "VirtualNode": "route-vn",
              "Weight": 100,
              "Port": null
            }
          ]
        },
        "Match": {
          "Headers": null,
          "Method": "",
          "Path": null,
          "Port": null,
          "Prefix": "/",
          "QueryParameters": null,
          "Scheme": ""
        },
        "RetryPolicy": null,
        "Timeout": null
      },
      "Priority": null,
      "TcpRoute": null
    },
    "Status": {
      "Status": "ACTIVE"
    },
    "VirtualRouterName": "route-vr"
  },
  "ResultMetadata": {}
}