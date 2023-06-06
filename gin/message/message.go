package message

import "net/http"

type ResponseBody struct {
	Code    int32
	Message string
	Data    any
}

func New(code int32, message string, data any) ResponseBody {
	return ResponseBody{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func OK(data any) ResponseBody {
	msg := ResponseBody{
		Code:    http.StatusOK,
		Message: "ok",
		Data:    data,
	}

	return msg
}

func Error(code int32, message string) ResponseBody {
	msg := ResponseBody{
		Code:    code,
		Message: message,
		Data:    nil,
	}

	return msg
}
