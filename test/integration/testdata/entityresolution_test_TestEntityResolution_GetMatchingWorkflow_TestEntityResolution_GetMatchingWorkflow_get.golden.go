{
  "CreatedAt": "2026-07-27T03:23:00Z",
  "InputSourceConfig": [
    {
      "InputSourceARN": null,
      "SchemaName": "test-schema",
      "ApplyNormalization": null
    }
  ],
  "OutputSourceConfig": [
    {
      "Output": [
        {
          "Name": "id",
          "Hashed": null
        }
      ],
      "ApplyNormalization": null,
      "CustomerProfilesIntegrationConfig": null,
      "KMSArn": null,
      "OutputS3Path": "s3://bucket/output/"
    }
  ],
  "ResolutionTechniques": {
    "ResolutionType": "RULE_MATCHING",
    "EnableRealTimeMatching": null,
    "ProviderProperties": null,
    "RuleBasedProperties": null,
    "RuleConditionProperties": null
  },
  "RoleArn": "arn:aws:iam::000000000000:role/test-role",
  "UpdatedAt": "2026-07-27T03:23:00Z",
  "WorkflowArn": "arn:aws:entityresolution:us-east-1:000000000000:matchingworkflow/test-get-matching-workflow",
  "WorkflowName": "test-get-matching-workflow",
  "Description": null,
  "IncrementalRunConfig": null,
  "Tags": null,
  "ResultMetadata": {}
}