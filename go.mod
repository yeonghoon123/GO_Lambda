module yeonghoon123/main

go 1.21.3

require (
	github.com/aws/aws-lambda-go v1.41.0
	yeonghoon123/GO_Lambda/dynamodb v0.0.0-00010101000000-000000000000
)

require (
	github.com/aws/aws-sdk-go v1.47.8 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
)

replace yeonghoon123/GO_Lambda/dynamodb => ./dynamodb
