/*
프로그램ID: MN-100
작성자: 김영훈
작성일: 2023.11.15
설명: Lambda 핸들러 함수
버전: 0.7
*/
package main // 패키지명

// 사용하는 모듈
import (
	"context"
	"encoding/json"
	"fmt"
	"yeonghoon123/GO_Lambda/dynamodb"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// MN-ST-110 api gateway에서 받는 데이터 구조체
type ResponseData struct {
	Status  bool
	Message string
	Data    []dynamodb.ScanItem
}

// MN-FN-110 람다 핸들러(context 데이터, api gateway request 데이터) return : api gateway response 데이터, error
func HandleRequest(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	var httpMethod = request.RequestContext.HTTP.Method                           // http method 데이터
	var requestBody = request.Body                                                // client에서 받은 body 데이터
	var responseData = ResponseData{Status: false, Message: "create item failed"} // response 초기 데이터

	// method별 실행될 스위치
	switch httpMethod {
	case "GET":
		scanItem, err := dynamodb.GetSaveDataList()
		if err != nil {
			fmt.Println(err)

			responseData = ResponseData{
				Status:  false,
				Message: "Scanning Data failed",
			}
		} else {
			responseData = ResponseData{
				Status:  true,
				Message: "Scanning Data success",
				Data:    scanItem,
			}
		}

	case "POST":
		// JSON 데이터를 Go의 구조체로 변환하는 예시
		var getItemInput dynamodb.GetItem
		if err := json.Unmarshal([]byte(requestBody), &getItemInput); err != nil {
			fmt.Println(err)

			// 에러 처리
			responseData = ResponseData{
				Status:  false,
				Message: "Error unmarshal response data to JSON",
			}
		}

		err := dynamodb.CreateTableItem(getItemInput)

		if err != nil {
			fmt.Println(err)

			responseData = ResponseData{
				Status:  false,
				Message: "Create item failed",
			}
		} else {
			responseData = ResponseData{
				Status:  true,
				Message: "Create item success",
			}
		}

	case "DELETE":
		// JSON 데이터를 Go의 구조체로 변환하는 예시
		var DeleteItemInput dynamodb.DeleteItem
		if err := json.Unmarshal([]byte(requestBody), &DeleteItemInput); err != nil {
			// 에러 처리
			responseData = ResponseData{
				Status:  false,
				Message: "Error unmarshal response data to JSON",
			}
		}

		err := dynamodb.DeleteSaveData(DeleteItemInput)

		if err != nil {
			fmt.Println(err)

			responseData = ResponseData{
				Status:  false,
				Message: "Delete item failed",
			}
		} else {
			responseData = ResponseData{
				Status:  true,
				Message: "Delete item success",
			}
		}
	}

	// 응답 데이터를 JSON으로 마샬링
	responseBody, err := json.Marshal(responseData)
	if err != nil {
		fmt.Println(err)

		return events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Body:       "Error marshaling response data to JSON",
		}, err
	}

	// 응답 생성
	var response = events.APIGatewayV2HTTPResponse{}

	if responseData.Status {
		response = events.APIGatewayV2HTTPResponse{
			StatusCode: 200,
			Body:       string(responseBody),
		}
	} else {
		fmt.Println(err)

		response = events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Body:       string(responseBody),
		}
	}

	return response, nil
}

// MN-FN-120 람다 핸들러 실행
func main() {
	lambda.Start(HandleRequest)
}
