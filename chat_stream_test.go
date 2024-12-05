package qwen

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
	"testing"
)

func TestChatStreamQwen(t *testing.T) {

	client := GetChatClient()
	request := openai.ChatCompletionRequest{
		Model: "qwen-plus",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "你好，帮我生成60字的文章，关于春游去故宫的文章",
			},
		},
		Stream: true,
	}
	response, err := client.CreateChatCompletionStream(
		context.Background(), request)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Close()

	for {
		recv, err := response.Recv()
		if err == io.EOF {
			fmt.Println(recv)
			break // 结束循环
		}
		if err != nil {
			log.Fatal(err)
		}

		// 检查 recv 是否包含有效的 Content
		if recv.Choices[0].Delta.Content != "" {
			fmt.Println(recv.Choices[0].Delta.Content)
		}
	}
}
