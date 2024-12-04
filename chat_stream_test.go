package qwen

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
	"os"
	"testing"
)

func TestChatStreamQwen(t *testing.T) {
	path := "chat.txt"

	// 创建或打开文件以写入模式
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
	}
	defer file.Close()

	client := GetChatClient()
	request := openai.ChatCompletionRequest{
		Model: "qwen-plus",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "你好，帮我生成200字的文章，关于春游去故宫的文章",
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
			break // 结束循环
		}
		if err != nil {
			log.Fatal(err)
		}

		// 检查 recv 是否包含有效的 Content
		if recv.Choices[0].Delta.Content != "" {
			_, err := fmt.Fprintf(file, "%s", recv.Choices[0].Delta.Content)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	fmt.Println("完成写入 chat.txt")
}
