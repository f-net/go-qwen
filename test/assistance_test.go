package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"qwen/utils"
	"testing"
)

// TestCreateAssistant 创建assistant
func TestCreateAssistant(t *testing.T) {
	var (
		model        = "qwen-plus"
		name         = "小机灵鬼"
		instructions = "你是一个智能助手，你帮助用户回答问题"
		description  = "小机灵鬼"

		tools []openai.AssistantTool
	)
	tools = append(tools, openai.AssistantTool{
		Type:     "code_interpreter",
		Function: nil,
	})
	client := utils.GetAssistantClient()

	var assistantRequest = openai.AssistantRequest{
		Model:        model,
		Name:         &name,
		Instructions: &instructions,
		Description:  &description,
	}
	assistantRequest.Tools = tools
	assistant, err := client.CreateAssistant(context.Background(), assistantRequest)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(assistant)
}

// 更新assistant
func TestUpdateAssistant(t *testing.T) {
	assistantId := "asst_348ef069-e95e-4810-aecf-4c2410276023"
	var (
		model        = "qwen-plus"
		name         = "小机灵鬼edit"
		instructions = "你是一个智能助手，你帮助用户回答问题edit"
		description  = "小机灵鬼edit"

		tools []openai.AssistantTool
	)
	tools = append(tools, openai.AssistantTool{
		Type:     "code_interpreter",
		Function: nil,
	})
	client := utils.GetAssistantClient()
	var assistantRequest = openai.AssistantRequest{
		Model:        model,
		Name:         &name,
		Instructions: &instructions,
		Description:  &description,
	}
	assistant, err := client.ModifyAssistant(context.Background(), assistantId, assistantRequest)
	if err != nil {
		t.Fatal()
	}
	fmt.Println(assistant)
}

// 获取assistantList
func TestGetAssistantList(t *testing.T) {
	limit := 10
	order := "desc"
	client := utils.GetAssistantClient()
	assistant, err := client.ListAssistants(context.Background(), &limit, &order, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range assistant.Assistants {
		fmt.Println(v)
	}
}

// 获取assistant
func TestGetAssistant(t *testing.T) {
	assistantId := "asst_348ef069-e95e-4810-aecf-4c2410276023"
	client := utils.GetAssistantClient()
	assistant, err := client.RetrieveAssistant(context.Background(), assistantId)
	if err != nil {
		t.Fatal(err)
	}
	marshal, _ := json.Marshal(assistant)
	t.Log(string(marshal))
}
