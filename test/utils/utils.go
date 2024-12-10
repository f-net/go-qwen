package utils

import (
	"github.com/sashabaranov/go-openai"
	"qwen/internal/config"
	"qwen/test"
)

func GetChatClient() *openai.Client {
	config := openai.DefaultConfig(GetApiKey())
	config.BaseURL = test.QwenChatUrl
	return openai.NewClientWithConfig(config)
}
func GetAssistantClient() *openai.Client {
	config := openai.DefaultConfig(GetApiKey())
	config.BaseURL = test.QwenAssistantUrl
	return openai.NewClientWithConfig(config)
}

func GetApiKey() string {
	config.InitConfig()
	return test.ApiKey + config.GetConfig().BaiLian.Apikey
}

func NewPointer[T any](value T) *T {
	return &value
}
