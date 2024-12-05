package qwen

import (
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"os"
)

func GetChatClient() *openai.Client {
	config := openai.DefaultConfig(GetApiKey())
	config.BaseURL = QwenChatUrl
	return openai.NewClientWithConfig(config)
}
func GetAssistantClient() *openai.Client {
	config := openai.DefaultConfig(GetApiKey())
	config.BaseURL = QwenAssistantUrl
	return openai.NewClientWithConfig(config)
}

func GetApiKey() string {
	godotenv.Load()
	return ApiKey + os.Getenv("apikey")
}
