# Changelog

## [0.26.0](https://github.com/Brilhante29/kiri-aws/compare/v0.25.3...v0.26.0) (2026-07-31)


### Features

* add project governance, CI workflow, docker-compose and honest upstream attribution ([02861df](https://github.com/Brilhante29/kiri-aws/commit/02861df357c8d778d8f3edc5b4eb58a4b886d33b))
* initial release of kiro-aws (Local AWS Emulator & Billing Time Machine) ([1b1ec26](https://github.com/Brilhante29/kiri-aws/commit/1b1ec2651f265bc1cdfe80d7fa172145b0fb1ab8))


### Bug fixes

* **brand:** replace leftover Kiro references with kiri-aws across docs, README, and tests ([027e4e5](https://github.com/Brilhante29/kiri-aws/commit/027e4e55e1721893b1ea102b2a7a4f43c514b4e5))
* **gamelift,lint:** clean up CBOR refactor so golangci-lint passes ([4434e36](https://github.com/Brilhante29/kiri-aws/commit/4434e36771fd0b54696e52f126945e948df98826))
* **gamelift:** resolve CBOR timestamp encoding and update golden schemas for SDK v2 ([66c8b9a](https://github.com/Brilhante29/kiri-aws/commit/66c8b9ae8778a68aec596bc79095ef64b302ac5b))
* **lint,test:** add missing package comments and refresh integration goldens ([11b8698](https://github.com/Brilhante29/kiri-aws/commit/11b8698a7926200627cfd23b351db55df0593c45))
* **security:** bound untrusted allocation sizes, restrict S3 redirect, pin base images ([5a0c3ee](https://github.com/Brilhante29/kiri-aws/commit/5a0c3ee1dbe9218aeac923cb4ed133896814dc65))
* **security:** stop sizing slice allocations from attacker-controlled counts ([9993fa1](https://github.com/Brilhante29/kiri-aws/commit/9993fa190519dbb2904409770b16f3c9b161b27c))


### Documentation

* add OpenSSF Scorecard and Go Report Card badges ([a7a43d8](https://github.com/Brilhante29/kiri-aws/commit/a7a43d8b1c122361a117fe0a3d0c4b8182b6b72c))

## [v0.25.3](https://github.com/Brilhante29/kiri-aws/compare/v0.25.2...v0.25.3) - 2026-06-18
- fix testcase by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/821
- fix(s3): externalize object bodies from snapshot to bound persistence memory by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/823
- release v0.25.3 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/824

## [v0.25.2](https://github.com/Brilhante29/kiri-aws/compare/v0.25.1...v0.25.2) - 2026-06-15
- fix(dynamodb): fail closed on unparseable FilterExpression and implement IN by @wakame1367 in https://github.com/Brilhante29/kiri-aws/pull/808
- fix(dynamodb): fail closed on KeyConditionExpression and implement attribute_type by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/810
- refactor(dynamodb): simplify storage/expression functions to drop complexity nolints by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/811
- refactor: simplify service handlers to drop complexity nolints by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/812
- refactor: consolidate duplicated readJSONRequest into service.ReadJSONRequest by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/813
- refactor: consolidate writeJSONResponse into service.WriteJSONResponse by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/814
- refactor: share identical ServiceError via service.CodedError alias by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/815
- fix(eventbridge): expand InvocationEndpoint path parameters on API destination delivery by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/817
- refactor: consolidate JSON-protocol writeError into service.WriteJSONError by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/818
- fix(json): emit lowercase "message" in error responses for 4 services by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/819
- release v0.25.2 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/820

## [v0.25.1](https://github.com/Brilhante29/kiri-aws/compare/v0.25.0...v0.25.1) - 2026-06-11
- feat(readme): auto-generate Supported Services catalog from service metadata by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/801
- fix(ci): improve govulncheck output parsing with JSON and reachability filter by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/805
- fix(lambda): queue async Event invocations with retry instead of fire-and-forget by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/804
- release v0.25.1 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/806

## [v0.25.0](https://github.com/Brilhante29/kiri-aws/compare/v0.24.1...v0.25.0) - 2026-06-04
- feat(apigatewayv2): implement API Gateway v2 (HTTP API) CRUD by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/772
- feat(apigateway): execute-api invoke for deployed stages (v1 + v2) by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/773
- feat(lambda): kiri-native Lambda Runtime API (run lambda.Start without RIE) by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/776
- feat(s3): support browser-based POST Object (presigned POST) by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/777
- fix(storage): debounce snapshot persistence to stop OOM under write load by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/799
- release v0.25.0 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/800

## [v0.24.1](https://github.com/Brilhante29/kiri-aws/compare/v0.24.0...v0.24.1) - 2026-05-30
- fix(cloudfront): support SHA256 signed cookie/URL verification by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/768
- release v0.24.1 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/770

## [v0.24.0](https://github.com/Brilhante29/kiri-aws/compare/v0.23.1...v0.24.0) - 2026-05-29
- feat(kms): support Sign, Verify, GetPublicKey and asymmetric keys by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/765
- release v0.24.0 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/766

## [v0.23.1](https://github.com/Brilhante29/kiri-aws/compare/v0.23.0...v0.23.1) - 2026-05-28
- fix(dynamodb): accept space-separated values for --key-schema and --attribute-definitions by @hirasawayuki in https://github.com/Brilhante29/kiri-aws/pull/761
- fix(dynamodb): support legacy KeyConditions in Query by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/763
- release v0.23.1 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/764

## [v0.23.0](https://github.com/Brilhante29/kiri-aws/compare/v0.22.0...v0.23.0) - 2026-05-28
- feat(cloudfront): Signed Cookie / Signed URL verification on Edge proxy by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/758
- release v0.23.0 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/760

## [v0.22.0](https://github.com/Brilhante29/kiri-aws/compare/v0.21.0...v0.22.0) - 2026-05-26
- Fix DynamoDB UpdateExpression invalid UTF-8 panic by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/674
- fix(s3): reject unordered multipart completion by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/680
- fix(range): reject suffix ranges on empty bodies by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/681
- Add fuzz coverage for S3 and DynamoDB edge cases by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/669
- fix(dynamodb): partition parallel scan segments by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/700
- fix(dynamodb): apply projection expressions to reads by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/699
- fix(s3): honor copy source version id by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/693
- fix(kinesis): validate record key inputs by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/685
- Proposal: env-var endpoint passthrough for RDS (and a model for ElastiCache / AMP / MSK / …) by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/579
- fix(s3): fix fuzz test build failure for parseCopySource by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/730
- delete tflog by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/733
- fix(s3): reduce CopyObject function length to satisfy funlen linter by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/731
- chore: add lefthook for pre-commit and pre-push hooks by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/735
- add toolchain by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/738
- delete integration test task by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/739
- feat(ssm): implement tag operations with storage persistence by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/732
- feat(eventbridge): implement tag operations with storage persistence by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/734
- feat(glue): implement tag operations with storage persistence by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/737
- feat(route53): implement tag operations with storage persistence by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/740
- feat(kms): implement policy and tag operations with storage persistence by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/741
- feat(lambda): implement tag and permission operations with storage persistence by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/742
- feat(cloudwatch): implement tag operations with storage persistence by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/743
- feat(sfn): implement tag and validation operations with storage persistence by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/744
- feat(secretsmanager): implement resource policy operations with storage persistence by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/745
- feat(dynamodb): implement tag and backup operations with storage persistence by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/746
- feat(cognito): implement MFA config operations with storage persistence by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/747
- ci(fuzz): add daily fuzzing CI workflow by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/748
- test(dynamodbstreams): add integration test by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/749
- refactor: eliminate hardcoded magic strings across services by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/751
- fix(ci): improve fuzz workflow and fix dynamodbstreams golden test by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/752
- feat: add write-through persistence to all services by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/753
- fix(s3): honor CopyObject tagging directive by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/691
- fix(dynamodb): validate key schema and write actions by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/679
- feat(cloudfront): PublicKey + KeyGroup CRUD (signed URL building blocks) by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/586
- fix(s3): honor CopyObject metadata directive by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/689
- fix(server): normalize AWS service name from User-Agent by @OliveiraNt in https://github.com/Brilhante29/kiri-aws/pull/755
- feat(sesv2): EmailTemplate CRUD + SendBulkEmail by @hc100 in https://github.com/Brilhante29/kiri-aws/pull/720
- ci(govulncheck): add daily vulnerability scan and pin toolchain by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/756
- release v0.22.0 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/757

## [v0.21.0](https://github.com/Brilhante29/kiri-aws/compare/v0.20.0...v0.21.0) - 2026-05-18
- feat(ssm): no-op stubs for ListTagsForResource / AddTagsToResource / RemoveTagsFromResource by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/626
- feat(sqs): support Policy attribute in SetQueueAttributes / GetQueueAttributes by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/627
- feat(sns): implement GetSubscriptionAttributes by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/632
- fix(kms): store key policy in CreateKey and PutKeyPolicy by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/633
- fix(s3): implement GetBucketCors handler by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/634
- fix(ssm): implement ParameterFilters in DescribeParameters by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/636
- feat(sns): implement SetSubscriptionAttributes by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/639
- feat(dynamodb): no-op stub for UpdateTable by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/640
- feat(elbv2): empty-payload stubs for read-only metadata APIs by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/567
- feat(cloudcontrol): AWS Cloud Control API + S3 Bucket / EC2 VPC+Subnet / IAM Role by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/569
- feat(cloudfront): RFC 9111-conformant edge cache layer (Cache-Control / Vary / Range / revalidation / swr) by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/583
- feat(dynamodb): support StreamSpecification in CreateTable and TableDescription by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/642
- fix(lambda): add GetFunctionConfiguration route by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/657
- feat(sqs): implement ChangeMessageVisibility and ChangeMessageVisibilityBatch by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/658
- feat(cloudwatch): implement SetAlarmState by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/660
- fix(sqs): allow batch receive of multiple FIFO messages from same group by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/662
- fix(sns): pass MessageGroupId and MessageDeduplicationId to SQS on FIFO publish by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/663
- feat(kinesis): implement DescribeStreamSummary by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/659
- fix(s3): support x-amz-tagging header and SSE-KMS in PutObject by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/661
- fix(sqs): requeue messages after visibility timeout expires by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/701
- feat(eventbridge): implement InputTransformer for event targets by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/702
- feat(s3): deliver ObjectCreated notifications to SQS queues by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/703
- feat(sfn): implement state machine execution engine by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/705
- feat(ses): add SES v1 Query protocol service by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/704
- fix(ses): use /_aws/ses mailbox endpoint matching harness convention by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/706
- fix(sqs): accept numeric maxReceiveCount in RedrivePolicy and notify DLQ on redrive by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/717
- feat(cloudwatch): notify SNS action targets on alarm state change by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/719
- fix(sns): evaluate FilterPolicy before delivering messages to subscriptions by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/718
- fix(lambda): echo payload on Invoke and fix EventSourceMapping UUID extraction by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/716
- fix(s3): implement ListObjectsV2 pagination with MaxKeys and ContinuationToken by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/715
- feat(lambda): implement AddPermission, RemovePermission, and ListTags by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/721
- fix(sns): parse MessageAttributes from Query protocol and support RawMessageDelivery by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/722
- feat(dynamodbstreams): add DynamoDB Streams service with cross-service event emission by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/723
- fix(awsquery): convert non-string attribute values in flattenAttributes by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/724
- fix(eventbridge): respect AWS_DEFAULT_REGION and resolve API destinations tolerantly by @tak848 in https://github.com/Brilhante29/kiri-aws/pull/726
- release v0.21.0 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/727

## [v0.20.0](https://github.com/Brilhante29/kiri-aws/compare/v0.19.0...v0.20.0) - 2026-05-12
- feat(s3): PutBucketLogging / GetBucketLogging by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/577
- feat(apigateway): bare-prefix routes + DeleteMethod + path-extract fix by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/602
- feat(server): KIRI_PPROF and KIRI_LOG_LEVEL env knobs by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/609
- feat(s3): conditional requests + presigned response-* overrides + ListObjectsV1 + Object ACL by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/587
- feat(eventbridge): Lambda targets + content filters + tag stubs by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/591
- fix(sqs): avoid nil map panic in ReceiveMessage after persisted reload by @hirasawayuki in https://github.com/Brilhante29/kiri-aws/pull/624
- fix(dynamodb): cli accepts object form for --local/global-secondary-indexes by @hirasawayuki in https://github.com/Brilhante29/kiri-aws/pull/623
- fix(s3): decode URL-encoded x-amz-copy-source before parsing by @hc100 in https://github.com/Brilhante29/kiri-aws/pull/617
- feat(s3): BucketWebsite + BucketLifecycle + Object Restore by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/588
- release v0.20.0 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/625

## [v0.19.0](https://github.com/Brilhante29/kiri-aws/compare/v0.18.2...v0.19.0) - 2026-05-11
- feat(ec2): implement CreateTags / DeleteTags / DescribeTags by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/547
- feat(iam): inline role policies, ListAttachedRolePolicies, OIDC provider by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/548
- feat(s3): PublicAccessBlock and BucketEncryption by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/549
- fix(ec2): correct InvalidRouteTableID.NotFound error code casing by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/550
- feat(elbv2): listener rule lifecycle (Create/Describe/Modify/Delete + SetRulePriorities) by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/556
- fix(elbv2): parse Targets.member.N from form in RegisterTargets / DeregisterTargets by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/559
- feat(route53): ListHostedZonesByName by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/560
- feat(logs): PutRetentionPolicy / DeleteRetentionPolicy by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/561
- feat(ecr): PutLifecyclePolicy / GetLifecyclePolicy / DeleteLifecyclePolicy by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/562
- fix(ec2): sort DescribeTags response for deterministic ordering by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/564
- feat(ec2): tags-on-create via TagSpecifications on Create* actions by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/553
- feat(iam): instance profile lifecycle (Create/Get/Delete/List + AddRole/RemoveRole) by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/555
- feat(elbv2): LoadBalancer / TargetGroup attribute APIs (Modify/Describe) by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/557
- feat(s3): default-response stubs for unconfigured bucket sub-resources by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/552
- feat(elbv2): DescribeListeners / ModifyListener / DescribeTargetHealth by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/558
- feat(ecr): no-op stubs for ListTagsForResource / TagResource / UntagResource by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/565
- feat(logs): no-op stubs for ListTagsForResource / TagResource / UntagResource by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/566
- feat(ec2): VPC/Subnet attribute mutation (Modify/Describe) by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/554
- fix(elbv2): align listener-rule wire format with AWS (ARN segment + condition typed configs) by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/568
- fix(server): honor KIRI_HOST/KIRI_PORT and --host/--port by @izumin5210 in https://github.com/Brilhante29/kiri-aws/pull/612
- fix(server): path-prefix boundary + AWS S3 virtual-hosted-style support by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/574
- perf(dynamodb): cache key once for Scan sort + pagination (~1.7x) by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/608
- feat(s3): PutBucketPolicy / GetBucketPolicy / DeleteBucketPolicy by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/576
- feat(s3): GetObject Range support + UploadPartCopy by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/584
- feat(ec2): real RevokeSecurityGroupIngress / RevokeSecurityGroupEgress by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/575
- fix(ec2): make terraform destroy work for IGW + subnet by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/596
- feat(iam): UpdateRole / UpdateAssumeRolePolicy / TagRole + unified Query dispatcher by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/578
- feat(elbv2): weighted forward (TargetGroupTuples) for ALB canary / blue-green by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/570
- feat(dynamodb): no-op stubs for ListTagsOfResource / Tag / Untag / DescribeContinuousBackups by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/590
- feat(secretsmanager): GetResourcePolicy / Put / Delete stubs by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/592
- feat(kms): GetKeyPolicy / ListResourceTags / GetKeyRotationStatus stubs by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/593
- feat(sfn): ValidateStateMachineDefinition + Versions/Aliases/Tags stubs by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/594
- feat(route53): ListTagsForResource / ChangeTagsForResource stubs by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/597
- feat(glue): GetTags / TagResource / UntagResource stubs by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/599
- feat(sns): GetTopicAttributes / SetTopicAttributes / Tag stubs by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/595
- feat(cloudwatch): register Query dispatcher + XML responses + tag stubs by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/600
- fix(lambda): terraform path-extract + refresh stubs + InvokeEndpoint-less stub-mode by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/601
- feat(cognito): populate UserPool refresh fields + GetUserPoolMfaConfig stub by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/603
- fix(server): wire SNS topic SQS subscriptions to actually deliver messages by @mizchi in https://github.com/Brilhante29/kiri-aws/pull/604
- fix(sqs): enforce message group lock for FIFO queues by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/613
- release v0.19.0 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/614

## [v0.18.2](https://github.com/Brilhante29/kiri-aws/compare/v0.18.1...v0.18.2) - 2026-05-08
- fix(s3): clean URL path to prevent 301 redirects and sort GetObjectTagging response by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/545
- release v0.18.2 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/544

## [v0.18.1](https://github.com/Brilhante29/kiri-aws/compare/v0.18.0...v0.18.1) - 2026-05-08
- fix(dynamodb): return HTTP 400 for ConditionalCheckFailedException by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/541
- release v0.18.1 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/542

## [v0.18.0](https://github.com/Brilhante29/kiri-aws/compare/v0.17.2...v0.18.0) - 2026-05-07
- sort changelog by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/536
- ci: replace manual actions/cache with setup-go built-in cache by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/538
- feat(s3): implement PutObjectTagging and GetObjectTagging by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/539
- release v0.18.0 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/540

## [v0.17.2](https://github.com/Brilhante29/kiri-aws/compare/v0.17.1...v0.17.2) - 2026-05-07
- feat(secretsmanager): support partial ARN lookup per AWS spec by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/523
- re-run v0.17.2 release by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/531
- ci: add Go module cache to lint and integration test workflows by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/534


## [v0.17.1](https://github.com/Brilhante29/kiri-aws/compare/v0.17.0...v0.17.1) - 2026-05-07
- fix(dynamodb): apply KeyConditionExpression range key filter and add BETWEEN by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/521

## [v0.17.0](https://github.com/Brilhante29/kiri-aws/compare/v0.16.1...v0.17.0) - 2026-05-07
- feat(dynamodb): support arithmetic expressions in SET clause by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/510
- fix(secretsmanager): support partial ARN lookup and use AWS_DEFAULT_REGION by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/514
- feat(sqs): implement long polling for ReceiveMessage by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/515
- feat(sfn): add SendTaskSuccess, SendTaskFailure, SendTaskHeartbeat stubs by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/516
- fix(secretsmanager): remove partial ARN match for AWS compatibility by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/518
- fix(sfn,lambda): use AWS_DEFAULT_REGION and fix Lambda routing by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/520
- release v0.17.0 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/519

## [v0.16.1](https://github.com/Brilhante29/kiri-aws/compare/v0.16.0...v0.16.1) - 2026-05-01
- fix(dynamodb): sort Query results by sort key and fix MarshalJSON for Item type by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/507

## [v0.16.0](https://github.com/Brilhante29/kiri-aws/compare/v0.15.0...v0.16.0) - 2026-05-01
- feat(cli): add ACM subcommands by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/497
- feat(cli): add Amplify subcommands by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/499
- feat(cli): add API Gateway subcommands by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/500
- feat(cli): add App Mesh subcommands by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/501
- feat(cli): add AppSync subcommands by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/502
- feat(cli): add Athena subcommands by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/503
- feat(cli): add Backup subcommands by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/504
- feat(dynamodb): support legacy AttributeUpdates in UpdateItem by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/505
- release v0.16.0 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/506

## [v0.15.0](https://github.com/Brilhante29/kiri-aws/compare/v0.14.3...v0.15.0) - 2026-04-30
- feat(cli): add AWS-compatible CLI subcommands for resource management by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/494
- release v0.15.0 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/496

## [v0.14.3](https://github.com/Brilhante29/kiri-aws/compare/v0.14.2...v0.14.3) - 2026-04-30
- fix(dynamodb): use pointer types for M and L fields in AttributeValue by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/491
- release v0.14.3 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/492

## [v0.14.2](https://github.com/Brilhante29/kiri-aws/compare/v0.14.1...v0.14.2) - 2026-04-30
- fix(dynamodb): preserve empty lists/maps/sets in AttributeValue serialization by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/488
- release v0.14.2 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/489

## [v0.14.1](https://github.com/Brilhante29/kiri-aws/compare/v0.14.0...v0.14.1) - 2026-04-30
- fix(dynamodb): handle space between function name and parenthesis in ConditionExpression by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/484
- fix(server): fallback Query protocol routing when User-Agent lacks service identifier by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/483
- fix(eventbridge): accept epoch seconds in PutEvents Time field by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/485
- release v0.14.1 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/487

## [v0.14.0](https://github.com/Brilhante29/kiri-aws/compare/v0.13.0...v0.14.0) - 2026-04-28
- feat(s3): implement CopyObject by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/480
- release v0.14.0 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/482

## [v0.13.0](https://github.com/Brilhante29/kiri-aws/compare/v0.12.0...v0.13.0) - 2026-04-28
- fix(cognito): generate valid UserPoolId format by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/471
- feat(sesv2): support reading sent email HTML body by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/472
- feat(eventbridge): deliver events to SQS targets and apply InputPath by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/474
- docs: rename awsim to kiri in CONTRIBUTING.md by @v420v in https://github.com/Brilhante29/kiri-aws/pull/475
- feat(s3): emit EventBridge events on PutObject and store notification/CORS config by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/476
- feat(s3): add CORS response headers and OPTIONS preflight support by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/477
- release v0.13.0 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/479

## [v0.12.0](https://github.com/Brilhante29/kiri-aws/compare/v0.11.0...v0.12.0) - 2026-04-23
- feat: support EmailContent.Raw in SES v2 SendEmail by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/450
- feat: allow raw email submission without explicit Destination by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/453
- feat: support init directory execution on startup by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/461
- feat(dynamodb): support ADD and DELETE clauses in UpdateItem by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/462
- feat(dynamodb): support Local Secondary Index (LSI) by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/463
- feat(dynamodb): support TTL auto-deletion of expired items by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/464
- feat(eventbridge): support API Destinations with actual HTTP delivery by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/465
- feat(sqs): support RedrivePolicy (Dead Letter Queue) by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/466
- feat(s3): stub support for PutBucketNotificationConfiguration and PutBucketCors by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/467
- release v0.12.0 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/468

## [v0.11.0](https://github.com/Brilhante29/kiri-aws/compare/v0.10.0...v0.11.0) - 2026-04-22
- feat: add DynamoDB ConditionExpression support by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/442
- feat: add EventBridge event pattern matching and delivery tracking by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/445
- feat: add DynamoDB TransactWriteItems/TransactGetItems support by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/448
- feat: add DynamoDB BatchWriteItem/BatchGetItem support by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/444
- feat: add DynamoDB Global Secondary Index (GSI) support by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/446
- release v0.11.0 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/449

## [v0.10.0](https://github.com/Brilhante29/kiri-aws/compare/v0.9.0...v0.10.0) - 2026-04-21
- feat(route53): add GetChange API by @apenney in https://github.com/Brilhante29/kiri-aws/pull/430
- feat(sesv2): add kiri-specific endpoint to retrieve sent emails by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/432
- feat: add Pinpoint SMS Voice v2 service with SendTextMessage by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/434
- release v0.10.0 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/435

## [v0.9.0](https://github.com/Brilhante29/kiri-aws/compare/v0.8.0...v0.9.0) - 2026-04-16
- feat(sqs): add ListQueueTags, TagQueue, and UntagQueue actions by @kahirokunn in https://github.com/Brilhante29/kiri-aws/pull/400
- feat(neptune): add Amazon Neptune service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/417
- fix(server): disambiguate Query protocol actions via User-Agent header by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/422
- fix(neptune): regenerate golden files after Query protocol routing fix by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/423
- feat(documentdb): add Amazon DocumentDB service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/424
- fix(ssm): return masked value for SecureString when WithDecryption is false by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/425
- ci: pin GitHub Actions with pinact by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/426
- feat(redshift): add Amazon Redshift service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/421
- feat(chart): add Helm chart for kiri by @kahirokunn in https://github.com/Brilhante29/kiri-aws/pull/399
- release v0.9.0 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/427
- Update chart by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/428

## [v0.8.0](https://github.com/Brilhante29/kiri-aws/compare/v0.7.0...v0.8.0) - 2026-04-07
- feat(sqs): add DeleteMessageBatch action by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/395
- fix(sqs): resolve queue by URL path to handle hostname mismatch by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/394
- fix(docker): run as non-root user by @kahirokunn in https://github.com/Brilhante29/kiri-aws/pull/398
- fix(server): use mime.ParseMediaType for Content-Type matching by @kahirokunn in https://github.com/Brilhante29/kiri-aws/pull/397
- feat(location): add Amazon Location Service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/401
- feat(macie2): add Amazon Macie2 service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/402
- release v0.8.0 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/403

## [v0.7.0](https://github.com/Brilhante29/kiri-aws/compare/v0.6.0...v0.7.0) - 2026-03-31
- feat(s3): add DeleteObjects (multi-object delete) support by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/388
- feat(dynamodb): add UpdateTimeToLive and DescribeTimeToLive actions by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/389
- feat(sqs): add SendMessageBatch action by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/387
- fix(ssm): return parameters without leading slash in GetParametersByPath by @tchssk in https://github.com/Brilhante29/kiri-aws/pull/386
- release v0.7.0 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/391

## [v0.6.0](https://github.com/Brilhante29/kiri-aws/compare/v0.5.4...v0.6.0) - 2026-03-27
- feat(storage): add optional persistent storage via KIRI_DATA_DIR by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/378
- feat(secretsmanager): add GetRandomPassword API by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/380
- release v0.6.0 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/382
- feat(server): add API action name to request logs by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/381

## [v0.5.4](https://github.com/Brilhante29/kiri-aws/compare/v0.5.3...v0.5.4) - 2026-03-24
- Fix/goreleaser cask token by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/373

## [v0.5.3](https://github.com/Brilhante29/kiri-aws/compare/v0.5.2...v0.5.3) - 2026-03-24
- Fix/goreleaser cask token by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/371

## [v0.5.2](https://github.com/Brilhante29/kiri-aws/compare/v0.5.1...v0.5.2) - 2026-03-24
- fix(ci): set repository token for homebrew cask in goreleaser by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/369

## [v0.5.1](https://github.com/Brilhante29/kiri-aws/compare/v0.5.0...v0.5.1) - 2026-03-24
- fix(ci): add HOMEBREW_TAP_GITHUB_TOKEN for goreleaser by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/367

## [v0.5.0](https://github.com/Brilhante29/kiri-aws/compare/v0.4.0...v0.5.0) - 2026-03-24
- test(integration): migrate all integration tests to golden test pattern by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/359
- feat: expose public Go API for in-process testing by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/361
- rename: awsim to kiri by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/362
- fix: regenerate golden files after rename to kiri by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/363
- feat: add Homebrew tap support via goreleaser by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/364
- release v0.5.0 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/365

## [v0.4.0](https://github.com/Brilhante29/kiri-aws/compare/v0.3.0...v0.4.0) - 2026-03-19
- feat(ebs): implement EBS Direct API remaining operations by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/356

## [v0.3.0](https://github.com/Brilhante29/kiri-aws/compare/v0.2.0...v0.3.0) - 2026-03-16
- feat(kafka): implement AWS MSK service by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/305
- feat(rds): add ModifyDBCluster for DocumentDB support by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/307
- feat(entityresolution): implement AWS Entity Resolution service by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/308
- feat(memorydb): implement AWS MemoryDB service by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/309
- feat(glacier): implement AWS Glacier service by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/310
- feat(elasticbeanstalk): implement AWS Elastic Beanstalk service by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/311
- feat(ebs): implement AWS EBS direct API service by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/312
- feat(sts): implement AWS Security Token Service by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/313
- feat(amplify): implement AWS Amplify service by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/314
- feat(backup): implement AWS Backup service by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/315
- feat(dataexchange): implement AWS Data Exchange service by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/316
- docs: add logo and update service list by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/317
- feat(codeguru-reviewer): implement AWS CodeGuru Reviewer service by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/318
- feat(codeguru-profiler): implement AWS CodeGuru Profiler service by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/319
- release v0.3.0 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/320

## [v0.2.0](https://github.com/Brilhante29/kiri-aws/compare/v0.1.3...v0.2.0) - 2026-03-06
- feat(mq): implement Amazon MQ service by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/290
- feat(ds): implement AWS Directory Service by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/292
- refactor: remove unused Prefix() method from Service interface by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/294
- feat(s3control): implement AWS S3 Control service by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/295
- feat(route53resolver): implement AWS Route 53 Resolver service by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/296
- feat(securitylake): implement AWS Security Lake service by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/297
- feat(finspace): implement AWS FinSpace service by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/298
- feat(comprehend): implement AWS Comprehend service by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/299
- feat(resiliencehub): implement AWS Resilience Hub service by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/300
- feat(ce): implement AWS Cost Explorer service by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/301
- feat(rekognition): implement AWS Rekognition service by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/302
- docs: update README with current service list and examples by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/303
- Release v0.2.0 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/304

## [v0.1.3](https://github.com/Brilhante29/kiri-aws/compare/v0.1.2...v0.1.3) - 2026-02-26
- fix(kinesis): use NextToken parameter in ListStreams by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/273
- fix(eks): fix nextToken handling and add maxResults validation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/274
- fix(globalaccelerator): implement pagination in ListAccelerators by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/275
- fix(route53): add pagination support to ListHostedZones by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/276
- fix(appsync): read nextToken and maxResults parameters in ListGraphqlApis by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/277
- fix(eventbridge): add ManagedBy field to ListEventBuses response by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/278
- fix(cloudfront): add missing fields to ListDistributions response by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/279
- fix(organizations): add State field to ListAccounts response by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/280
- fix(ecs): add maxResults validation to ListClusters by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/281
- fix(ecr): implement sorting and pagination in ListImages by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/282
- fix(secretsmanager): add Type field and rotation metadata to ListSecrets response by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/283
- fix(cognito): populate LambdaConfig in DescribeUserPool response by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/284
- fix(acm): add missing fields to ListCertificates response by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/285
- fix(cloudwatch): add OwningAccounts field to ListMetrics response by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/286
- fix(s3tables): add missing fields and fix pagination by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/287
- refactor(test): migrate from testify to golden and separate test dependencies by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/289

## [v0.1.2](https://github.com/Brilhante29/kiri-aws/compare/v0.1.1...v0.1.2) - 2026-02-26

## [v0.1.1](https://github.com/Brilhante29/kiri-aws/compare/v0.1.0...v0.1.1) - 2026-02-26
- fix(s3): align ListBuckets response format with AWS API by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/255

## [v0.1.0](https://github.com/Brilhante29/kiri-aws/compare/v0.0.2...v0.1.0) - 2026-02-25
- feat(rds): add RDS service basic implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/233
- feat(elasticache): add ElastiCache service basic implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/235
- feat(route53): add Route 53 service basic implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/237
- feat(ec2): add VPC service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/238
- feat(elbv2): add ELB v2 service basic implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/239
- feat(cloudformation): add CloudFormation service basic implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/240
- feat(cloudtrail): add CloudTrail service basic implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/241
- feat(configservice): add AWS Config service basic implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/242
- feat(gamelift): add GameLift service basic implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/243
- feat(servicequotas): add Service Quotas service basic implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/244
- feat(organizations): add Organizations service basic implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/245
- feat(forecast): add Forecast service basic implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/246
- feat(pipes): add EventBridge Pipes service basic implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/247
- feat(emrserverless): add EMR Serverless service basic implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/248
- feat(appmesh): add App Mesh service basic implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/249
- chore: bump version to v0.1.0 by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/250
- feat(scheduler): add EventBridge Scheduler service basic implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/251
- feat(dlm): add DLM service basic implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/252
- feat(sagemaker): add SageMaker service basic implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/253

## [v0.0.2](https://github.com/Brilhante29/kiri-aws/compare/v0.0.1...v0.0.2) - 2026-02-17
- feat(eks): add basic EKS service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/207
- feat(ecs): add basic ECS service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/206
- feat(iam): add IAM service basic implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/209
- feat(s3tables): add S3 Tables service basic implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/210
- feat(athena): add Athena service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/211
- feat(codeconnections): add CodeConnections service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/212
- feat(kms): add KMS service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/213
- feat(globalaccelerator): add Global Accelerator service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/214
- feat(cognito): add Cognito Identity Provider service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/215
- feat(eventbridge): add EventBridge service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/216
- feat(kinesis): add basic Kinesis service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/217
- feat(sfn): add basic Step Functions service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/218
- feat(ecr): add basic ECR service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/219
- feat(apigateway): add API Gateway service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/220
- feat(sesv2): add basic SES v2 service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/221
- feat(acm): add basic ACM service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/222
- feat(glue): add basic Glue service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/223
- feat(xray): add X-Ray service basic implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/224
- fix(lint): resolve lint issues in ACM and X-Ray services by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/228
- fix(acm,glue): use JSONProtocolService interface by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/229
- fix: ACM and Glue JSON protocol service conflict by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/230
- feat(batch): add Batch service basic implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/226
- feat(cloudfront): add CloudFront service basic implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/231
- feat(cloudwatch): add CloudWatch metrics service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/232
- feat(firehose): add Data Firehose service basic implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/227
- feat(appsync): add AppSync service basic implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/225

## [v0.0.1](https://github.com/Brilhante29/kiri-aws/commits/v0.0.1) - 2026-02-06
- feat: add project foundation with base interfaces by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/186
- refactor(service): add auto-registration via init() by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/187
- feat(s3): add basic S3 service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/188
- feat(sqs): add basic SQS service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/189
- feat(dynamodb): add basic DynamoDB service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/191
- fix(ci): use integration tests for coverage reporting by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/192
- feat(secretsmanager): add basic Secrets Manager service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/193
- feat(sns): add basic SNS service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/194
- feat(lambda): add basic Lambda service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/195
- feat(lambda): add HTTP proxy support for Lambda invocations by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/196
- feat(sqs): add FIFO queue support by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/198
- feat(s3): add presigned URL support by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/200
- feat(s3): add versioning support by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/201
- feat(lambda): add EventSourceMapping API support by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/202
- feat(ssm): add SSM Parameter Store support by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/203
- feat(s3): add multipart upload support by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/199
- feat(cloudwatchlogs): add CloudWatch Logs service implementation by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/204
- feat(ec2): implement EC2 service with basic instance operations by @sivchari in https://github.com/Brilhante29/kiri-aws/pull/205
