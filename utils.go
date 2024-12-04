package qwen

import "github.com/sashabaranov/go-openai"

func GetChatClient() *openai.Client {
	config := openai.DefaultConfig(ApiKey)
	config.BaseURL = QwenChatUrl
	return openai.NewClientWithConfig(config)
}
func GetAssistanceClient() *openai.Client {
	config := openai.DefaultConfig(ApiKey)
	config.BaseURL = QwenChatUrl
	return openai.NewClientWithConfig(config)
}
