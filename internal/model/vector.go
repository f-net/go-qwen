package model

import "encoding/json"

// VectorStore 向量存储模型
type VectorStore struct {
	Id               int64           `json:"id"`
	VectorStoreAppId string          `json:"vectorStoreAppId"`
	Object           string          `json:"object"`
	Name             string          `json:"name"`
	Description      string          `json:"description"`
	UsageBytes       int             `json:"usageBytes"`
	FileTotalBytes   int64           `json:"fileTotalBytes"`
	FileCounts       json.RawMessage `json:"fileCounts"`
	ExpiresAfter     json.RawMessage `json:"expiresAfter"`
	ExpiresAt        int             `json:"expiresAt"`
	AssistantAmount  int64           `json:"assistantAmount"` // 统计关联的助手数量用字段
	FileAmount       int64           `json:"FileAmount"`      // 统计关联的助手数量用字段
	CreatedAt        int64           `json:"createdAt" default:"now"`
	UpdatedAt        int64           `json:"updatedAt" default:"now"`
}
