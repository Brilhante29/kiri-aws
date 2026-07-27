{
  "Certificate": {
    "AcmeAccountId": null,
    "AcmeEndpointArn": null,
    "CertificateArn": "arn:aws:acm:us-east-1:000000000000:certificate/63235cf2-5e93-4f55-96df-0e80a639ca23",
    "CertificateAuthorityArn": null,
    "CertificateKeyPairOrigin": "",
    "CreatedAt": "2026-07-27T03:22:52.826Z",
    "DomainName": "example.com",
    "DomainValidationOptions": [
      {
        "DomainName": "example.com",
        "HttpRedirect": null,
        "ResourceRecord": {
          "Name": "_acme-challenge.example.com",
          "Type": "CNAME",
          "Value": "_63235cf2.acm-validations.aws."
        },
        "ValidationDomain": "example.com",
        "ValidationEmails": null,
        "ValidationMethod": "DNS",
        "ValidationStatus": "PENDING_VALIDATION"
      },
      {
        "DomainName": "www.example.com",
        "HttpRedirect": null,
        "ResourceRecord": {
          "Name": "_acme-challenge.www.example.com",
          "Type": "CNAME",
          "Value": "_63235cf2.acm-validations.aws."
        },
        "ValidationDomain": "www.example.com",
        "ValidationEmails": null,
        "ValidationMethod": "DNS",
        "ValidationStatus": "PENDING_VALIDATION"
      },
      {
        "DomainName": "api.example.com",
        "HttpRedirect": null,
        "ResourceRecord": {
          "Name": "_acme-challenge.api.example.com",
          "Type": "CNAME",
          "Value": "_63235cf2.acm-validations.aws."
        },
        "ValidationDomain": "api.example.com",
        "ValidationEmails": null,
        "ValidationMethod": "DNS",
        "ValidationStatus": "PENDING_VALIDATION"
      }
    ],
    "ExtendedKeyUsages": null,
    "FailureReason": "",
    "ImportedAt": null,
    "InUseBy": null,
    "IssuedAt": null,
    "Issuer": null,
    "KeyAlgorithm": "RSA_2048",
    "KeyUsages": null,
    "ManagedBy": "",
    "NotAfter": null,
    "NotBefore": null,
    "Options": null,
    "RenewalEligibility": "INELIGIBLE",
    "RenewalSummary": null,
    "RevocationReason": "",
    "RevokedAt": null,
    "Serial": "8974b32e20d3eb0316a053c06d5ac733",
    "SignatureAlgorithm": null,
    "Status": "PENDING_VALIDATION",
    "Subject": "CN=example.com",
    "SubjectAlternativeNames": [
      "www.example.com",
      "api.example.com"
    ],
    "Type": "AMAZON_ISSUED"
  },
  "ResultMetadata": {}
}