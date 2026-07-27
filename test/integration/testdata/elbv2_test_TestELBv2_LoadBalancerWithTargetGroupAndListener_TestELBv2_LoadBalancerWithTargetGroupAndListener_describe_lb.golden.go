{
  "LoadBalancers": [
    {
      "AvailabilityZones": [
        {
          "LoadBalancerAddresses": null,
          "OutpostId": null,
          "SourceNatIpv6Prefixes": null,
          "SubnetId": "subnet-12345678",
          "ZoneName": "us-east-1a"
        },
        {
          "LoadBalancerAddresses": null,
          "OutpostId": null,
          "SourceNatIpv6Prefixes": null,
          "SubnetId": "subnet-87654321",
          "ZoneName": "us-east-1b"
        }
      ],
      "CanonicalHostedZoneId": "Z35SXDOTRQ7X7K",
      "CreatedTime": "2026-07-27T03:22:59.775Z",
      "CustomerOwnedIpv4Pool": null,
      "DNSName": "test-full-lb-f2a7877a.us-east-1.elb.amazonaws.com",
      "EnablePrefixForIpv6SourceNat": "",
      "EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic": null,
      "IpAddressType": "ipv4",
      "IpamPools": null,
      "LoadBalancerArn": "arn:aws:elasticloadbalancing:us-east-1:000000000000:loadbalancer/app/test-full-lb/f2a7877a-a365-46d",
      "LoadBalancerName": "test-full-lb",
      "Scheme": "internet-facing",
      "SecurityGroups": [],
      "State": {
        "Code": "active",
        "Reason": null
      },
      "Type": "application",
      "VpcId": "vpc-f066667e"
    }
  ],
  "NextMarker": null,
  "ResultMetadata": {}
}