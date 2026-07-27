{
  "NextShardIterator": "YXJuOmF3czpkeW5hbW9kYjp1cy1lYXN0LTE6MDAwMDAwMDAwMDAwOnRhYmxlL3Rlc3Qtc3RyZWFtcy1yZWNvcmRzL3N0cmVhbS8yMDI2LTA3LTI3VDAzOjIyOjU3LjUwNDpzaGFyZElkLTAwMDAwMDAwMDAwMDoxOjE3ODUxMjI1Nzc2MTM3MDcyMDQ=",
  "Records": [
    {
      "AwsRegion": "us-east-1",
      "Dynamodb": {
        "ApproximateCreationDateTime": "2026-07-27T03:22:57Z",
        "Keys": {
          "pk": {
            "Value": "stream-item-1"
          }
        },
        "NewImage": {
          "data": {
            "Value": "hello"
          },
          "pk": {
            "Value": "stream-item-1"
          }
        },
        "OldImage": null,
        "SequenceNumber": "000000000000000000004",
        "SizeBytes": 100,
        "StreamViewType": "NEW_AND_OLD_IMAGES"
      },
      "EventID": "2c3acacb-34c1-4826-b761-7da4524d090d",
      "EventName": "INSERT",
      "EventSource": "aws:dynamodb",
      "EventVersion": "1.1",
      "UserIdentity": null
    }
  ],
  "ResultMetadata": {}
}