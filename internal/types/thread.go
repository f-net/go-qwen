package types

type ListAssistantThreadReq struct {
	ListQuery
	AssistantId int64  `json:"assistantId"`
	Name        string `json:"name"`
}

type ListAssistantThreadResp struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
