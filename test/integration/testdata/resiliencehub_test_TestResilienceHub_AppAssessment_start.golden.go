{
  "Assessment": {
    "AssessmentArn": "arn:aws:resiliencehub:us-east-1:123456789012:app-assessment/ea0e2528-39b6-4243-8e44-7d59546f5882",
    "AssessmentStatus": "Success",
    "Invoker": "User",
    "AppArn": "arn:aws:resiliencehub:us-east-1:123456789012:app/21f12511-2e70-4baa-bf66-c22e130cf0c3",
    "AppVersion": "release",
    "AssessmentName": "test-assessment",
    "Compliance": {
      "Hardware": {
        "ComplianceStatus": "PolicyMet",
        "AchievableRpoInSecs": 3600,
        "AchievableRtoInSecs": 3600,
        "CurrentRpoInSecs": 3600,
        "CurrentRtoInSecs": 3600,
        "Message": null,
        "RpoDescription": null,
        "RpoReferenceId": null,
        "RtoDescription": null,
        "RtoReferenceId": null
      },
      "Software": {
        "ComplianceStatus": "PolicyBreached",
        "AchievableRpoInSecs": 3600,
        "AchievableRtoInSecs": 3600,
        "CurrentRpoInSecs": 86400,
        "CurrentRtoInSecs": 86400,
        "Message": null,
        "RpoDescription": null,
        "RpoReferenceId": null,
        "RtoDescription": null,
        "RtoReferenceId": null
      }
    },
    "ComplianceStatus": "PolicyBreached",
    "Cost": null,
    "DriftStatus": "",
    "EndTime": "2026-07-27T03:23:14Z",
    "Message": null,
    "Policy": null,
    "ResiliencyScore": {
      "DisruptionScore": {
        "Hardware": 100,
        "Software": 50
      },
      "Score": 75,
      "ComponentScore": null
    },
    "ResourceErrorsDetails": null,
    "StartTime": "2026-07-27T03:23:14Z",
    "Summary": null,
    "Tags": null,
    "VersionName": null
  },
  "ResultMetadata": {}
}