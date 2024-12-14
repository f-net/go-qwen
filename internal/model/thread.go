package model

type Thread struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	AssistantId int64  `json:"assistantId"`
	RemoteId    string `json:"remoteId"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
}
