package types

import (
	"github.com/sashabaranov/go-openai"
)

type Assistant struct {
	Name          string                 `json:"name"`
	Instructions  string                 `json:"instructions"`
	Model         string                 `json:"model"`
	Tools         []openai.AssistantTool `json:"tools"`
	ToolResources openai.ToolResources   `json:"toolResources"` // 工具资源
	Remark        string                 `json:"remark"`        // 备注
}
