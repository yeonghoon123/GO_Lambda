# GO_Lambda

GO lang으로 Lambda 함수 CRD 핸들러 구현

## Getting Started

### Installing

git clone하여 프로젝트 설치

    git clone https://github.com/yeonghoon123/translator_app.git

<br>

## Running the tests

### Sample Tests

main.go 파일 build

    SET GOOS=linux& go build main.go

빌드된 main 파일을 function.zip파일로 압축

    zip function.zip main

Lambda 업로드 command

    aws lambda update-function-code --function-name my-funtion-name --zip-file fileb://function.zip

<br>

압축한 function.zip 파일을 업로드 하거나, aws 커맨드를 이용하여 Lambda 업로드

## Built With

```
GO
API Gateway
Lambda
DynamoDB
```

## Versioning

v0.1 - GO 설치 완료 <br>
v0.2 - Lambda 함수 핸들러 초기 테스트 검증<br>
v0.3 - client body 데이터 변환 완료<br>
v0.4 - PutItem 기능 구현<br>
v0.5 - Error 로직 구현 <br>
v0.6 - Scan 기능 구현 <br>
v0.7 - DELETE 기능 구현 <br>
