package utils

import "encoding/json"

func Swap2Json(data interface{}) json.RawMessage {
	marshal, _ := json.Marshal(data)
	return marshal
}

func NewPointer[T any](value T) *T {
	return &value
}
