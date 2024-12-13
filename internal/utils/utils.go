package utils

import "encoding/json"

func Swap2Json(data interface{}) string {
	if data == nil {
		data = make([]int, 0)
		marshal, _ := json.Marshal(data)
		return string(marshal)
	}
	marshal, _ := json.Marshal(data)
	return string(marshal)
}

func NewPointer[T any](value T) *T {
	return &value
}

func GetTotalPage(total int64, pageSize int) int {
	pageSize2 := int64(pageSize)
	return int((total + pageSize2 - 1) / pageSize2)
}
