{
  "NextShardIterator": "YXJuOmF3czpkeW5hbW9kYjp1cy1lYXN0LTE6MDAwMDAwMDAwMDAwOnRhYmxlL3Rlc3Qtc3RyZWFtcy1tdWx0aS1vcHMvc3RyZWFtLzIwMjYtMDctMjdUMDM6MjI6NTcuNjQ1OnNoYXJkSWQtMDAwMDAwMDAwMDAwOjM6MTc4NTEyMjU3Nzc1ODgwODIzOQ==",
  "Records": [
    {
      "AwsRegion": "us-east-1",
      "Dynamodb": {
        "ApproximateCreationDateTime": "2026-07-27T03:22:57Z",
        "Keys": {
          "pk": {
            "Value": "multi-1"
          }
        },
        "NewImage": {
          "data": {
            "Value": "original"
          },
          "pk": {
            "Value": "multi-1"
          }
        },
        "OldImage": null,
        "SequenceNumber": "000000000000000000006",
        "SizeBytes": 100,
        "StreamViewType": "NEW_AND_OLD_IMAGES"
      },
      "EventID": "a81120d4-ff57-40e0-92d5-27845cf2e987",
      "EventName": "INSERT",
      "EventSource": "aws:dynamodb",
      "EventVersion": "1.1",
      "UserIdentity": null
    },
    {
      "AwsRegion": "us-east-1",
      "Dynamodb": {
        "ApproximateCreationDateTime": "2026-07-27T03:22:57Z",
        "Keys": {
          "pk": {
            "Value": "multi-1"
          }
        },
        "NewImage": {
          "data": {
            "Value": "updated"
          },
          "pk": {
            "Value": "multi-1"
          }
        },
        "OldImage": {
          "data": {
            "Value": "original"
          },
          "pk": {
            "Value": "multi-1"
          }
        },
        "SequenceNumber": "000000000000000000007",
        "SizeBytes": 100,
        "StreamViewType": "NEW_AND_OLD_IMAGES"
      },
      "EventID": "e3cda079-5a74-443c-991e-be545ef5664b",
      "EventName": "MODIFY",
      "EventSource": "aws:dynamodb",
      "EventVersion": "1.1",
      "UserIdentity": null
    },
    {
      "AwsRegion": "us-east-1",
      "Dynamodb": {
        "ApproximateCreationDateTime": "2026-07-27T03:22:57Z",
        "Keys": {
          "pk": {
            "Value": "multi-1"
          }
        },
        "NewImage": null,
        "OldImage": {
          "data": {
            "Value": "updated"
          },
          "pk": {
            "Value": "multi-1"
          }
        },
        "SequenceNumber": "000000000000000000008",
        "SizeBytes": 100,
        "StreamViewType": "NEW_AND_OLD_IMAGES"
      },
      "EventID": "a76e1997-857d-46b0-ad4b-ebfce0759787",
      "EventName": "REMOVE",
      "EventSource": "aws:dynamodb",
      "EventVersion": "1.1",
      "UserIdentity": null
    }
  ],
  "ResultMetadata": {}
}