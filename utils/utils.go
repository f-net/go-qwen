package utils

import (
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"os"
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
	godotenv.Load()
	return test.ApiKey + os.Getenv("apikey")
}

func NewPointer[T any](value T) *T {
	return &value
}
