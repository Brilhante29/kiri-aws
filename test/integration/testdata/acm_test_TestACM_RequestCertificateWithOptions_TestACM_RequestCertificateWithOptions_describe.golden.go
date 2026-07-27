{
  "Certificate": {
    "AcmeAccountId": null,
    "AcmeEndpointArn": null,
    "CertificateArn": "arn:aws:acm:us-east-1:000000000000:certificate/df4e5e07-58dc-464c-822c-32a6e29ceada",
    "CertificateAuthorityArn": null,
    "CertificateKeyPairOrigin": "",
    "CreatedAt": "2026-07-27T03:22:52.915Z",
    "DomainName": "options-test.example.com",
    "DomainValidationOptions": [
      {
        "DomainName": "options-test.example.com",
        "HttpRedirect": null,
        "ResourceRecord": {
          "Name": "_acme-challenge.options-test.example.com",
          "Type": "CNAME",
          "Value": "_df4e5e07.acm-validations.aws."
        },
        "ValidationDomain": "options-test.example.com",
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
    "KeyAlgorithm": "EC_prime256v1",
    "KeyUsages": null,
    "ManagedBy": "",
    "NotAfter": null,
    "NotBefore": null,
    "Options": {
      "CertificateTransparencyLoggingPreference": "ENABLED",
      "Export": ""
    },
    "RenewalEligibility": "INELIGIBLE",
    "RenewalSummary": null,
    "RevocationReason": "",
    "RevokedAt": null,
    "Serial": "572348a938779f1bd576a026fa19ec83",
    "SignatureAlgorithm": null,
    "Status": "PENDING_VALIDATION",
    "Subject": "CN=options-test.example.com",
    "SubjectAlternativeNames": null,
    "Type": "AMAZON_ISSUED"
  },
  "ResultMetadata": {}
}