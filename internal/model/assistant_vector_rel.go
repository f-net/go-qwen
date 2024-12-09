package model

// AssistantVectorStore 助手-向量存储关联模型
type AssistantVectorStore struct {
	AssistantId   int64 `json:"assistantId"`
	VectorStoreId int64 `json:"vectorStoreId"`
}
