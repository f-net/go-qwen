package types

type CreateMessageReq struct {
	AssistantId int64  `json:"assistantId"`
	ThreadId    int64  `json:"threadId"`
	Question    string `json:"question"`
}
type CreateMessageResp struct {
	ThreadId int64  `json:"threadId"`
	Answer   string `json:"answer"`
}

type GetMessageListReq struct {
	ListQuery
	AssistantId int64  `form:"assistantId"`
	ThreadId    int64  `form:"threadId"`
	Name        string `form:"name"`
}

type GetMessageListResp struct {
	Id       int64  `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
