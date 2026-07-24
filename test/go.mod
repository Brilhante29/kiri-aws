module github.com/Brilhante29/kiri-aws/test

go 1.25.0

require (
	github.com/aws/aws-lambda-go v1.54.0
	github.com/aws/aws-sdk-go v1.55.8
	github.com/aws/aws-sdk-go-v2 v1.43.0
	github.com/aws/aws-sdk-go-v2/config v1.32.31
	github.com/aws/aws-sdk-go-v2/credentials v1.19.30
	github.com/aws/aws-sdk-go-v2/service/acm v1.43.0
	github.com/aws/aws-sdk-go-v2/service/amplify v1.41.0
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.42.0
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.37.0
	github.com/aws/aws-sdk-go-v2/service/appmesh v1.38.0
	github.com/aws/aws-sdk-go-v2/service/appsync v1.56.0
	github.com/aws/aws-sdk-go-v2/service/athena v1.60.0
	github.com/aws/aws-sdk-go-v2/service/backup v1.59.0
	github.com/aws/aws-sdk-go-v2/service/batch v1.68.0
	github.com/aws/aws-sdk-go-v2/service/cloudcontrol v1.32.0
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.75.0
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.67.0
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.58.0
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.65.0
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.80.0
	github.com/aws/aws-sdk-go-v2/service/codeconnections v1.13.0
	github.com/aws/aws-sdk-go-v2/service/codeguruprofiler v1.32.0
	github.com/aws/aws-sdk-go-v2/service/codegurureviewer v1.37.0
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.66.0
	github.com/aws/aws-sdk-go-v2/service/comprehend v1.43.0
	github.com/aws/aws-sdk-go-v2/service/configservice v1.68.0
	github.com/aws/aws-sdk-go-v2/service/costexplorer v1.67.0
	github.com/aws/aws-sdk-go-v2/service/dataexchange v1.44.0
	github.com/aws/aws-sdk-go-v2/service/directoryservice v1.41.0
	github.com/aws/aws-sdk-go-v2/service/dlm v1.39.0
	github.com/aws/aws-sdk-go-v2/service/docdb v1.51.0
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.61.0
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.36.0
	github.com/aws/aws-sdk-go-v2/service/ebs v1.36.0
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.317.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.60.0
	github.com/aws/aws-sdk-go-v2/service/ecs v1.89.0
	github.com/aws/aws-sdk-go-v2/service/eks v1.90.0
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.56.0
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.37.0
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.58.1
	github.com/aws/aws-sdk-go-v2/service/emrserverless v1.44.0
	github.com/aws/aws-sdk-go-v2/service/entityresolution v1.30.0
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.48.0
	github.com/aws/aws-sdk-go-v2/service/finspace v1.36.0
	github.com/aws/aws-sdk-go-v2/service/firehose v1.46.0
	github.com/aws/aws-sdk-go-v2/service/forecast v1.44.0
	github.com/aws/aws-sdk-go-v2/service/gamelift v1.59.0
	github.com/aws/aws-sdk-go-v2/service/glacier v1.35.0
	github.com/aws/aws-sdk-go-v2/service/globalaccelerator v1.38.0
	github.com/aws/aws-sdk-go-v2/service/glue v1.149.0
	github.com/aws/aws-sdk-go-v2/service/iam v1.56.0
	github.com/aws/aws-sdk-go-v2/service/kafka v1.56.0
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.46.0
	github.com/aws/aws-sdk-go-v2/service/kms v1.55.0
	github.com/aws/aws-sdk-go-v2/service/lambda v1.100.0
	github.com/aws/aws-sdk-go-v2/service/location v1.54.0
	github.com/aws/aws-sdk-go-v2/service/macie2 v1.54.0
	github.com/aws/aws-sdk-go-v2/service/memorydb v1.36.0
	github.com/aws/aws-sdk-go-v2/service/mq v1.39.0
	github.com/aws/aws-sdk-go-v2/service/neptune v1.48.0
	github.com/aws/aws-sdk-go-v2/service/organizations v1.53.0
	github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2 v1.32.0
	github.com/aws/aws-sdk-go-v2/service/pipes v1.26.0
	github.com/aws/aws-sdk-go-v2/service/rds v1.123.0
	github.com/aws/aws-sdk-go-v2/service/redshift v1.65.0
	github.com/aws/aws-sdk-go-v2/service/rekognition v1.54.0
	github.com/aws/aws-sdk-go-v2/service/resiliencehub v1.38.0
	github.com/aws/aws-sdk-go-v2/service/route53 v1.65.2
	github.com/aws/aws-sdk-go-v2/service/route53resolver v1.48.0
	github.com/aws/aws-sdk-go-v2/service/s3 v1.106.0
	github.com/aws/aws-sdk-go-v2/service/s3control v1.73.0
	github.com/aws/aws-sdk-go-v2/service/s3tables v1.18.0
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.261.0
	github.com/aws/aws-sdk-go-v2/service/scheduler v1.20.0
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.44.0
	github.com/aws/aws-sdk-go-v2/service/securitylake v1.28.0
	github.com/aws/aws-sdk-go-v2/service/servicequotas v1.37.0
	github.com/aws/aws-sdk-go-v2/service/ses v1.37.0
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.66.0
	github.com/aws/aws-sdk-go-v2/service/sfn v1.45.0
	github.com/aws/aws-sdk-go-v2/service/sns v1.42.0
	github.com/aws/aws-sdk-go-v2/service/sqs v1.46.0
	github.com/aws/aws-sdk-go-v2/service/ssm v1.73.0
	github.com/aws/aws-sdk-go-v2/service/sts v1.45.0
	github.com/aws/aws-sdk-go-v2/service/xray v1.39.0
	github.com/aws/smithy-go v1.27.4
	github.com/sivchari/golden v0.3.0
)

require (
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.7.14 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.18.31 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.4.31 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.7.31 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.4.32 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.13 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.9.24 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.12.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.13.31 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.19.32 // indirect
	github.com/aws/aws-sdk-go-v2/service/signin v1.5.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.33.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.38.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
)

replace github.com/Brilhante29/kiri-aws => ../
