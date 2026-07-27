{
  "Policy": {
    "DateCreated": "2026-07-27T03:22:56.435528286Z",
    "DateModified": "2026-07-27T03:22:56.436700519Z",
    "DefaultPolicy": null,
    "Description": "Updated test policy",
    "ExecutionRoleArn": "arn:aws:iam::123456789012:role/dlm-role",
    "PolicyArn": "arn:aws:dlm:us-east-1:123456789012:policy/policy-0d3019ae-8551-460",
    "PolicyDetails": {
      "Actions": null,
      "CopyTags": null,
      "CreateInterval": null,
      "CrossRegionCopyTargets": null,
      "EventSource": null,
      "Exclusions": null,
      "ExtendDeletion": null,
      "Parameters": null,
      "PolicyLanguage": "",
      "PolicyType": "",
      "ResourceLocations": null,
      "ResourceType": "",
      "ResourceTypes": [
        "VOLUME"
      ],
      "RetainInterval": null,
      "Schedules": [
        {
          "ArchiveRule": null,
          "CopyTags": null,
          "CreateRule": {
            "CronExpression": null,
            "Interval": 24,
            "IntervalUnit": "HOURS",
            "Location": "",
            "Scripts": null,
            "Times": null
          },
          "CrossRegionCopyRules": null,
          "DeprecateRule": null,
          "FastRestoreRule": null,
          "Name": "Daily snapshots",
          "RetainRule": {
            "Count": 7,
            "Interval": null,
            "IntervalUnit": ""
          },
          "ShareRules": null,
          "TagsToAdd": null,
          "VariableTags": null
        }
      ],
      "TargetTags": [
        {
          "Key": "Backup",
          "Value": "true"
        }
      ]
    },
    "PolicyId": "policy-0d3019ae-8551-460",
    "State": "DISABLED",
    "StatusMessage": null,
    "Tags": null
  },
  "ResultMetadata": {}
}