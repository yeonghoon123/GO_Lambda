package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"fmt"
)

// language 데이터 구조체
type Language_data struct {
	Ko string `json:"ko"`
	En string `json:"en"`
	Ja string `json:"ja"`
}

// 클라이언트에서 받는 데이터 구조체
type GetItem struct {
	Id              string        `json:"id"`
	Stt_text        string        `json:"sttText"`
	Language_code   string        `json:"languageCode"`
	Tts_base64      Language_data `json:"ttsBase64"`
	Translator_text Language_data `json:"translatorText"`
	Language_name   Language_data `json:"languageName"`
}

// dynamodb 에 저장될 데이터 구조체
type CreateItem struct {
	Id              string
	Stt_text        string
	Language_code   string
	Tts_base64      Language_data
	Translator_text Language_data
	Language_name   Language_data
}

// 새로운 db 데이터 생성
func CreateTableItem(data GetItem) error {
	// 세션 생성
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// DynamoDB client 생성
	svc := dynamodb.New(sess) // dynamodb 연결

	// 저장할 데이터
	item := CreateItem{
		Id:              data.Id,
		Stt_text:        data.Stt_text,
		Language_code:   data.Language_code,
		Tts_base64:      data.Tts_base64,
		Translator_text: data.Translator_text,
		Language_name:   data.Language_name,
	}

	av, err := dynamodbattribute.MarshalMap(item) // dynamo db 데이터 변환
	if err != nil {
		return err
	}

	fmt.Println(av)

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

// dynamoDB에서 scan해올 데이터 구조체
type ScanItem struct {
	Id              string        `json:"Id"`
	Stt_text        string        `json:"Stt_text"`
	Language_code   string        `json:"Language_code"`
	Tts_base64      Language_data `json:"Tts_base64"`
	Translator_text Language_data `json:"Translator_text"`
	Language_name   Language_data `json:"Language_name"`
}

// dynamoDB 데이터 목록 가져오기
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

// 저장된 데이터 id 구조체
type DeleteItem struct {
	Id string `json:"Id"`
}

// 저장된 데이터 삭제
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
