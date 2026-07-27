{
  "CostCategory": {
    "CostCategoryArn": "arn:aws:ce::123456789012:costcategory/4720626c-6a1d-4724-a900-ef485789a43e",
    "EffectiveStart": "2026-07-27T03:22:54Z",
    "Name": "test-cost-category",
    "RuleVersion": "CostCategoryExpression.v1",
    "Rules": [
      {
        "InheritedValue": null,
        "Rule": {
          "And": null,
          "CostCategories": null,
          "Dimensions": {
            "Key": "LINKED_ACCOUNT",
            "MatchOptions": null,
            "Values": [
              "123456789012"
            ]
          },
          "Not": null,
          "Or": null,
          "Tags": null
        },
        "Type": "",
        "Value": "Development"
      }
    ],
    "DefaultValue": "Other",
    "EffectiveEnd": null,
    "ProcessingStatus": [
      {
        "Component": "COST_EXPLORER",
        "Status": "APPLIED"
      }
    ],
    "SplitChargeRules": null
  },
  "ResultMetadata": {}
}