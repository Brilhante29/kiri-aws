{
  "Policy": {
    "DateCreated": "2026-07-27T03:22:56.416544548Z",
    "DateModified": "2026-07-27T03:22:56.416544548Z",
    "DefaultPolicy": null,
    "Description": "Test policy for EBS snapshots",
    "ExecutionRoleArn": "arn:aws:iam::123456789012:role/dlm-role",
    "PolicyArn": "arn:aws:dlm:us-east-1:123456789012:policy/policy-cd506b92-3449-445",
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
            "Times": [
              "03:00"
            ]
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
    "PolicyId": "policy-cd506b92-3449-445",
    "State": "ENABLED",
    "StatusMessage": null,
    "Tags": null
  },
  "ResultMetadata": {}
}