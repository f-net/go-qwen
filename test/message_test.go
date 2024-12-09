package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"qwen/utils"
	"testing"
)

var threadId = "thread_7e006a91-c100-4dbe-9622-e6a056fff55d"

func TestCreateMessage(t *testing.T) {
	client := utils.GetAssistantClient()
	msg, err := client.CreateMessage(context.Background(), threadId, openai.MessageRequest{
		Role:    "user",
		Content: "写一篇200字的冬天堆雪人文章",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(msg.Content[0].Text)
}
func TestGetMessageList(t *testing.T) {
	client := utils.GetAssistantClient()
	msg, err := client.ListMessage(context.Background(), threadId, nil, nil, nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(msg)
	fmt.Println(*msg.Messages[0].Content[0].Text)
}

func TestGetMessage(t *testing.T) {
	client := utils.GetAssistantClient()
	msg, err := client.RetrieveMessage(context.Background(), threadId,
		"message_4a3a69e5-9d0c-46d2-93dc-d1809ab5019a")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(*msg.Content[0].Text)
	marshal, _ := json.Marshal(msg)
	t.Log(string(marshal))
}

// TestModifyMessage 修改元数据
func TestModifyMessage(t *testing.T) {
	client := utils.GetAssistantClient()
	msg, err := client.ModifyMessage(context.Background(), threadId,
		"message_4a3a69e5-9d0c-46d2-93dc-d1809ab5019a", map[string]string{
			"modified": "true",
			"user":     "user",
		})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(msg.Content[0].Text)
}
