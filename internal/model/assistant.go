package model

type Assistant struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Instructions   string `json:"instructions"`
	AssistantAppId string `json:"assistantAppId"`
	Model          string `json:"model"`
	Tools          string `json:"tools"`
	ToolResources  string `json:"toolResources"` // 工具资源
	Remark         string `json:"remark"`        // 备注
	CreatedAt      int64  `json:"createdAt" default:"now"`
	UpdatedAt      int64  `json:"updatedAt" default:"now"`
}
