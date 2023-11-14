package main

import (
	"context"
	"encoding/json"
	"yeonghoon123/GO_Lambda/dynamodb"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type ResponseData struct {
	Status  bool
	Message string
	Data    []dynamodb.ScanItem
}

// 람다 핸들러
func HandleRequest(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	var httpMethod = request.RequestContext.HTTP.Method                           // http method 데이터
	var requestBody = request.Body                                                // client에서 받은 body 데이터
	var responseData = ResponseData{Status: false, Message: "create item failed"} // response 초기 데이터

	// method별 실행될 스위치
	switch httpMethod {
	case "GET":
		scanItem, err := dynamodb.GetSaveDataList()
		if err != nil {
			responseData = ResponseData{
				Status:  false,
				Message: "Scanning Data failed",
			}
		} else {
			responseData = ResponseData{
				Status:  true,
				Message: "Scanning Data failed success",
				Data:    scanItem,
			}
		}

	case "POST":
		// JSON 데이터를 Go의 구조체로 변환하는 예시
		var getItemInput dynamodb.GetItem
		if err := json.Unmarshal([]byte(requestBody), &getItemInput); err != nil {
			// 에러 처리
			responseData = ResponseData{
				Status:  false,
				Message: "Error unmarshal response data to JSON",
			}
		}

		err := dynamodb.CreateTableItem(getItemInput)

		if err != nil {
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
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Body:       "Error marshaling response data to JSON",
		}, err
	}

	var response = events.APIGatewayV2HTTPResponse{}

	if responseData.Status {
		// 응답 생성
		response = events.APIGatewayV2HTTPResponse{
			StatusCode: 200,
			Body:       string(responseBody),
		}
	} else {
		response = events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Body:       string(responseBody),
		}
	}

	return response, nil
}

func main() {
	lambda.Start(HandleRequest)
}
