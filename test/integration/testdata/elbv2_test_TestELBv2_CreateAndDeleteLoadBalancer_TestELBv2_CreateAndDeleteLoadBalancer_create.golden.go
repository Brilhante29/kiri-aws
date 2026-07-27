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
      "CreatedTime": "2026-07-27T03:22:59.665Z",
      "CustomerOwnedIpv4Pool": null,
      "DNSName": "test-load-balancer-df56e1cf.us-east-1.elb.amazonaws.com",
      "EnablePrefixForIpv6SourceNat": "",
      "EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic": null,
      "IpAddressType": "ipv4",
      "IpamPools": null,
      "LoadBalancerArn": "arn:aws:elasticloadbalancing:us-east-1:000000000000:loadbalancer/app/test-load-balancer/df56e1cf-9d84-4a4",
      "LoadBalancerName": "test-load-balancer",
      "Scheme": "internet-facing",
      "SecurityGroups": [],
      "State": {
        "Code": "active",
        "Reason": null
      },
      "Type": "application",
      "VpcId": "vpc-39de632a"
    }
  ],
  "ResultMetadata": {}
}