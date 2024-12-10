package model

import "encoding/json"

type Assistant struct {
	Id            int64           `json:"id"`
	Name          string          `json:"name"`
	Instructions  string          `json:"instructions"`
	RemoteId      string          `json:"remoteId"`
	Model         string          `json:"model"`
	Tools         json.RawMessage `json:"tools"`
	ToolResources json.RawMessage `json:"toolResources"` // 工具资源
	Remark        string          `json:"remark"`        // 备注
	CreatedAt     int64           `json:"createdAt" default:"now"`
	UpdatedAt     int64           `json:"updatedAt" default:"now"`
}
