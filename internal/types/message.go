package types

type Message struct {
	AssistantId int64  `json:"assistantId"`
	ThreadId    int64  `json:"threadId"`
	Message     string `json:"message"`
}
