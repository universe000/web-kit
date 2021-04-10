package view

import (
	"encoding/json"
)

const (
	RETURN_SUCCESS = 1
	RETURN_FAILURE = 0
	ERRCODE        = "500"
)

type JsonResponse struct {
	ReturnCode int             `json:"returnCode"`
	ErrorCode  string          `json:"errorCode"`
	ReturnMsg  string          `json:"returnMsg"`
	Content    json.RawMessage `json:"content"`
}

func newJsonResponse(returnCode int, returnMsg string, content interface{}) JsonResponse {
	contentBytes, err := json.Marshal(content)
	if err != nil {
		return Failed("content cast fail:" + err.Error())
	}
	return JsonResponse{ReturnCode: returnCode, ReturnMsg: returnMsg, Content: contentBytes}
}

func Success(content interface{}) JsonResponse {
	return newJsonResponse(RETURN_SUCCESS, "success", content)
}

func Failed(returnMsg string) JsonResponse {
	return JsonResponse{ReturnCode: RETURN_FAILURE, ErrorCode: ERRCODE, ReturnMsg: returnMsg}
}

func FailedWithCode(errCode string, returnMsg string) JsonResponse {
	return JsonResponse{ReturnCode: RETURN_FAILURE, ErrorCode: errCode, ReturnMsg: returnMsg}
}
