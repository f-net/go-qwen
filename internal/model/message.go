package model

type Message struct {
	Id          int64  `json:"id"`          // '唯一标识符'
	Question    string `json:"question"`    // '提问'
	Answer      string `json:"answer"`      // '回答'
	RemoteId    string `json:"remoteId"`    // '远程id'
	AssistantId int64  `json:"assistantId"` // '助手id'
	ThreadId    int64  `json:"threadId"`    // '会话id'
	CreatedAt   int64  `json:"createdAt"`   // '创建时间'
	UpdatedAt   int64  `json:"updatedAt"`   // '更新时间'
}
