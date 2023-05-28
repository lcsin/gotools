package message

import "net/http"

type messageBody struct {
	Code    int32
	Message string
	Data    any
}

func OK(data any) messageBody {
	msg := messageBody{
		Code:    http.StatusOK,
		Message: "ok",
		Data:    data,
	}

	return msg
}

func Error(code int32, message string) messageBody {
	msg := messageBody{
		Code:    code,
		Message: message,
		Data:    nil,
	}

	return msg
}
