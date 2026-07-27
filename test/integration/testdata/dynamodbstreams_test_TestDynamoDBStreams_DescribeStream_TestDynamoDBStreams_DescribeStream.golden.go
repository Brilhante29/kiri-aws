{
  "StreamDescription": {
    "CreationRequestDateTime": "2026-07-27T03:22:57Z",
    "KeySchema": [
      {
        "AttributeName": "pk",
        "KeyType": "HASH"
      }
    ],
    "LastEvaluatedShardId": null,
    "Shards": [
      {
        "ParentShardId": null,
        "SequenceNumberRange": {
          "EndingSequenceNumber": null,
          "StartingSequenceNumber": "000000000000000000002"
        },
        "ShardId": "shardId-000000000000"
      }
    ],
    "StreamArn": "arn:aws:dynamodb:us-east-1:000000000000:table/test-streams-describe/stream/2026-07-27T03:22:57.489",
    "StreamLabel": "2026-07-27T03:22:57.489",
    "StreamStatus": "ENABLED",
    "StreamViewType": "NEW_AND_OLD_IMAGES",
    "TableName": "test-streams-describe"
  },
  "ResultMetadata": {}
}