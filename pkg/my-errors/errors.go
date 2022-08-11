package myerrors

import "net/http"

type MyError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewAssertError() *MyError {
	return &MyError{
		Type: "Bad Request",
		Code: http.StatusBadRequest,
	}
}
