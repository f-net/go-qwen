package config

import (
	"github.com/sashabaranov/go-openai"
	"qwen/internal/types"
)

var chatClient *openai.Client
var assistantClient *openai.Client

func GetChatClient() *openai.Client {
	return chatClient
}

func GetAssistantClient() *openai.Client {
	return assistantClient
}

func InitOpenaiClient() {
	chatConfig := openai.DefaultConfig(GetApiKey())
	chatConfig.BaseURL = types.QwenChatUrl
	chatClient = openai.NewClientWithConfig(chatConfig)

	assistantConfig := openai.DefaultConfig(GetApiKey())
	assistantConfig.BaseURL = types.QwenAssistantUrl
	assistantClient = openai.NewClientWithConfig(assistantConfig)
}

func GetApiKey() string {
	return types.ApiKey + config.BaiLian.Apikey
}
