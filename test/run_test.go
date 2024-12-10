package test

import (
	"context"
	"encoding/json"
	"github.com/sashabaranov/go-openai"
	"qwen/test/utils"
	"testing"
)

// TestCreateRunMessage 创建任务
func TestCreateRunMessage(t *testing.T) {
	client := utils.GetAssistantClient()
	response, err := client.CreateRun(context.Background(), threadId, openai.RunRequest{
		AssistantID: "asst_348ef069-e95e-4810-aecf-4c2410276023",
		Model:       QwenPlus,
	})
	if err != nil {
		t.Fatal(err)
	}
	marshal, _ := json.Marshal(response)
	t.Log(string(marshal))
}

// TestGetRunMessageList
func TestGetRunMessageList(t *testing.T) {
	limit := 10
	client := utils.GetAssistantClient()
	var response, err = client.ListRuns(context.Background(), threadId, openai.Pagination{
		Limit: &limit,
	})
	if err != nil {
		t.Fatal(err)
	}
	marshal, _ := json.Marshal(response)
	t.Log(string(marshal))
	//	run_7d3e15e3-7ce2-4d84-9b5f-4bd983a83ba8
}

// 列出运行任务的关联步骤
func TestGetRunMessageStep(t *testing.T) {

	limit := 10
	client := utils.GetAssistantClient()
	response, err := client.ListRunSteps(context.Background(), threadId, "run_7d3e15e3-7ce2-4d84-9b5f-4bd983a83ba8",
		openai.Pagination{
			Limit: &limit,
		})
	if err != nil {
		t.Fatal(err)
	}
	marshal, _ := json.Marshal(response)
	t.Log(string(marshal))

}

func TestRetrieveRun(t *testing.T) {

	client := utils.GetAssistantClient()
	response, err := client.RetrieveRun(context.Background(), threadId, "run_7d3e15e3-7ce2-4d84-9b5f-4bd983a83ba8")
	if err != nil {
		t.Fatal(err)
	}
	marshal, _ := json.Marshal(response)
	t.Log(string(marshal))

}

// TestRetrieveRunListMessage 获取提问的回答
func TestRetrieveRunListMessage(t *testing.T) {
	numMessages := 1
	client := utils.GetAssistantClient()
	response, err := client.ListMessage(context.Background(), threadId, &numMessages,
		nil, nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	marshal, _ := json.Marshal(response)
	t.Log(string(marshal))
}
