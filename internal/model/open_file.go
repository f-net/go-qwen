package model

// OpenaiFile openai文件属性模型
type OpenaiFile struct {
	Id            int64  `json:"id"`
	FileAppId     string `json:"fileAppId"`
	FileName      string `json:"fileName"`
	Extension     string `json:"extension"`
	FileBytes     int64  `json:"fileBytes"`
	VectorStoreId int64  `json:"vectorStoreId"`
	Status        string `json:"status"`
	Purpose       string `json:"purpose"`
	CreatedAt     int64  `json:"createdAt" default:"now"`
	UpdatedAt     int64  `json:"updatedAt" default:"now"`
}
