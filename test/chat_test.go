package test

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"qwen/test/utils"
	"testing"
)

func TestChatQwen(t *testing.T) {
	client := utils.GetChatClient()
	request := openai.ChatCompletionRequest{
		Model: "qwen-plus",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "你好，帮我生成200字的文章，关于春游去故宫的文章",
			},
		},
	}
	response, err := client.CreateChatCompletion(
		context.Background(), request)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(response.Choices[0].Message.Content)
}
