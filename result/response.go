package result

import "time"

type BasicResponse struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
}

func Success(obj interface{}) BasicResponse {
	return BasicResponse{
		Code:      0,
		Msg:       "success",
		Timestamp: time.Now().UnixNano(),
		Data:      obj,
	}
}

func Failure(code int, msg string) BasicResponse {
	return BasicResponse{
		Code:      code,
		Msg:       msg,
		Timestamp: time.Now().UnixNano(),
		Data:      nil,
	}
}
