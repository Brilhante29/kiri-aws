{
  "Listeners": [
    {
      "AlpnPolicy": null,
      "Certificates": null,
      "DefaultActions": [
        {
          "Type": "forward",
          "AuthenticateCognitoConfig": null,
          "AuthenticateOidcConfig": null,
          "FixedResponseConfig": null,
          "ForwardConfig": null,
          "JwtValidationConfig": null,
          "Order": null,
          "RedirectConfig": null,
          "TargetGroupArn": "arn:aws:elasticloadbalancing:us-east-1:000000000000:targetgroup/test-listener-tg/08eb7df6-e39f-411"
        }
      ],
      "ListenerArn": "arn:aws:elasticloadbalancing:us-east-1:000000000000:listener/app/test-listener-lb/cc4ecdcc-f61c-443/d701ed2a-838f-41b",
      "LoadBalancerArn": "arn:aws:elasticloadbalancing:us-east-1:000000000000:loadbalancer/app/test-listener-lb/cc4ecdcc-f61c-443",
      "MutualAuthentication": null,
      "Port": 80,
      "Protocol": "HTTP",
      "SslPolicy": null
    }
  ],
  "ResultMetadata": {}
}