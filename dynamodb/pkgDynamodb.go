/*
프로그램ID: PK-DY-100
작성자: 김영훈
작성일: 2023.11.15
설명: Dynamo Control 함수 모음
버전: 0.7
*/

package dynamodb // 패키지 명

// 사용하는 모듈
import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// DY-ST-100. 클라이언트에서 받는 데이터 구조체
type GetItem struct {
	Id            string `json:"id"`
	Stt_text      string `json:"sttText"`
	Language_code string `json:"languageCode"`
}

// DY-ST-110.dynamodb 에 저장될 데이터 구조체
type CreateItem struct {
	Id            string
	Stt_text      string
	Language_code string
}

// DY-FN-100. 새로운 db 데이터 생성
func CreateTableItem(data GetItem) error {
	// 세션 생성
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// DynamoDB client 생성
	svc := dynamodb.New(sess) // dynamodb 연결

	// 저장할 데이터
	item := CreateItem{
		Id:            data.Id,
		Stt_text:      data.Stt_text,
		Language_code: data.Language_code,
	}

	av, err := dynamodbattribute.MarshalMap(item) // dynamo db 데이터 변환
	if err != nil {
		return err
	}

	tableName := "translator_app" // 데이터 생설할 데이터 이름

	// dynamo db에 데이터 기입
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	// putItem api 실행
	_, err = svc.PutItem(input)
	if err != nil {
		return err

	}

	return nil
}

// DY-ST-120. dynamoDB에서 scan해올 데이터 구조체
type ScanItem struct {
	Id            string `json:"Id"`
	Stt_text      string `json:"Stt_text"`
	Language_code string `json:"Language_code"`
}

// DY-FN-110. dynamoDB 데이터 목록 가져오기
func GetSaveDataList() ([]ScanItem, error) {
	tableName := "translator_app" // 데이터 생설할 데이터 이름

	// 세션 생성
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// DynamoDB client 생성
	svc := dynamodb.New(sess)

	// 스캔할 테이블 데이터 기입
	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	// scan api 실행
	result, err := svc.Scan(params)
	if err != nil {
		return nil, err
	}

	var items []ScanItem // 테이블에 있는 데이터 목록

	// 데이터 가공
	for _, item := range result.Items {
		myItem := ScanItem{}
		err := dynamodbattribute.UnmarshalMap(item, &myItem)
		if err != nil {
			return nil, err
		}
		items = append(items, myItem)
	}

	return items, nil
}

// DY-ST-130. 저장된 데이터 id 구조체
type DeleteItem struct {
	Id string `json:"Id"`
}

// DY-FN-120. 저장된 데이터 삭제
func DeleteSaveData(itemId DeleteItem) error {
	tableName := "translator_app" // 데이터 생설할 데이터 이름

	// 세션 생성
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// DynamoDB client 생성
	svc := dynamodb.New(sess)

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(itemId.Id),
			},
		},
		TableName: aws.String(tableName),
	}

	_, err := svc.DeleteItem(input)
	if err != nil {
		return err
	}
	return nil
}
