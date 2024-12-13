package types

import (
	"github.com/sashabaranov/go-openai"
)

type CreateAssistantReq struct {
	Name          string                 `json:"name"`
	Instructions  string                 `json:"instructions"`
	Model         string                 `json:"model"`
	Tools         []openai.AssistantTool `json:"tools"`
	ToolResources openai.ToolResources   `json:"toolResources"` // 工具资源
	Remark        string                 `json:"remark"`        // 备注
}
type UpdateAssistantReq struct {
	Id int64 `json:"id"`
	CreateAssistantReq
}

type GetAssistantResp struct {
	Id            int64                  `json:"id"`
	Name          string                 `json:"name"`
	Instructions  string                 `json:"instructions"`
	RemoteId      string                 `json:"assistantAppId"`
	Model         string                 `json:"model"`
	Tools         []openai.AssistantTool `json:"tools"`
	ToolResources openai.ToolResources   `json:"toolResources"` // 工具资源
	Remark        string                 `json:"remark"`        // 备注
}

type ListAssistantReq struct {
	ListQuery
	Name string `json:"name"`
}

type ListAssistantResp struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Instructions string `json:"instructions"`
	Model        string `json:"model"`
	Remark       string `json:"remark"` // 备注
	CreatedAt    int64  `json:"createdAt" default:"now"`
	UpdatedAt    int64  `json:"updatedAt" default:"now"`
}

type GetMessageResp struct {
	Message string `json:"message"`
}
