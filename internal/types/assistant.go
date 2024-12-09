package types

import (
	"github.com/sashabaranov/go-openai"
)

type Assistant struct {
	Name           string                 `json:"name"`
	Instructions   string                 `json:"instructions"`
	Description    string                 `json:"description"` // 描述信息
	AssistantAppId string                 `json:"assistantAppId"`
	Model          string                 `json:"model"`
	Tools          []openai.AssistantTool `json:"tools"`
	ToolResources  openai.ToolResources   `json:"toolResources"` // 工具资源
	Remark         string                 `json:"remark"`        // 备注
	CreatedAt      int64                  `json:"createdAt" default:"now"`
	UpdatedAt      int64                  `json:"updatedAt" default:"now"`
}
