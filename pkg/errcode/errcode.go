package errcode

import (
	"encoding/json"
	"log"
	"strings"
)

type Error struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Detail  []string `json:"detail"`
}

var codes = map[int]string{}

func NewError(code int, message string, detail ...string) *Error {
	if _, ok := codes[code]; ok {
		log.Fatalf("errcode %d already exists", code)
	}
	return &Error{
		Code:    code,
		Message: message,
		Detail:  detail,
	}
}

func (e *Error) Error() string {
	buffer := strings.Builder{}
	marshal, _ := json.Marshal(e)
	buffer.Write(marshal)
	return buffer.String()
}
