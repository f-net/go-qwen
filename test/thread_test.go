package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"qwen/test/utils"
	"testing"
)

// 建立线程会话
func TestCreateThread(t *testing.T) {
	ms := []openai.ThreadMessage{
		{
			Role:    "user",
			Content: "你好",
		},
	}

	client := utils.GetAssistantClient()

	request := openai.ThreadRequest{
		Messages: ms,
	}

	response, err := client.CreateThread(context.Background(), request)
	if err != nil {
		t.Fatal()
	}
	fmt.Println(response)
	//thread_7e006a91-c100-4dbe-9622-e6a056fff55d
}
func TestGetThread(t *testing.T) {
	client := utils.GetAssistantClient()

	response, err := client.RetrieveThread(context.Background(), "thread_7e006a91-c100-4dbe-9622-e6a056fff55d")
	if err != nil {
		t.Fatal()
	}
	marshal, _ := json.Marshal(response)
	t.Log(string(marshal))
	//thread_cc15ce96-056c-4ede-9ee5-8885c76c2714
}

func TestDeleteThread(t *testing.T) {
	client := utils.GetAssistantClient()

	response, err := client.DeleteThread(context.Background(), "thread_cc15ce96-056c-4ede-9ee5-8885c76c2714")
	if err != nil {
		t.Fatal()
	}
	fmt.Println(response)
	//thread_cc15ce96-056c-4ede-9ee5-8885c76c2714
}

func TestModifyThread(t *testing.T) {

	client := utils.GetAssistantClient()

	request := openai.ModifyThreadRequest{
		Metadata: map[string]interface{}{
			"modified": "true",
			"user":     "system",
		},
	}

	response, err := client.ModifyThread(context.Background(),
		"thread_7e006a91-c100-4dbe-9622-e6a056fff55d", request)
	if err != nil {
		t.Fatal()
	}
	fmt.Println(response)
}
