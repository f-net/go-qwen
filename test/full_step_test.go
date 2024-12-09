package test

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"qwen/utils"
	"testing"
	"time"
)

func TestFullStepCreateAssistant(t *testing.T) {
	var (
		ctx     = context.Background()
		reqText = "你是谁，可以做什么事情"
	)
	client := utils.GetAssistantClient()
	assistantRes, err := client.CreateAssistant(ctx, openai.AssistantRequest{
		Model:          QwenPlus,
		Name:           utils.NewPointer("职场老油条"),
		Description:    utils.NewPointer("职场老油条"),
		Instructions:   utils.NewPointer("你是一个混迹职场多年的老油条，会热心的帮助职场新人如何处理人际关系，如何圆滑地与领导周旋。"),
		Tools:          []openai.AssistantTool{},
		FileIDs:        nil,
		ToolResources:  new(openai.AssistantToolResource),
		ResponseFormat: nil,
		Temperature:    nil,
		TopP:           nil,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(assistantRes.ID)
	// asst_29cb8187-90d4-4235-8018-a5d6545e4a90

	threadRes, err := client.CreateThread(ctx, openai.ThreadRequest{})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(threadRes.ID)
	// thread_550c8802-768b-43aa-b0dd-f8605e63ce79

	_, err = client.CreateMessage(ctx, threadRes.ID, openai.MessageRequest{
		Role:    "user",
		Content: reqText,
	})
	if err != nil {
		t.Fatal(err)
	}

	runRes, err := client.CreateRun(ctx, threadRes.ID, openai.RunRequest{
		AssistantID: assistantRes.ID,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(runRes.ID)
	// run_64c9044a-cefa-442c-9021-19b4cfbc0bfb

	// Poll for a status that indicates run has finished
	for runRes.Status == openai.RunStatusQueued || runRes.Status == openai.RunStatusInProgress {
		//fmt.Println("排队和正在运行")
		runRes, err = client.RetrieveRun(ctx, runRes.ThreadID, runRes.ID)
		if err != nil {
			t.Fatal(err.Error())
		}
		time.Sleep(300 * time.Millisecond)
	}
	if runRes.Status != openai.RunStatusCompleted {
		//fmt.Println("运行完成")
		if runRes.Status == openai.RunStatusFailed {
			//fmt.Println("运行出错" + run.LastError.Message)
			fmt.Println(fmt.Sprintf("Run error: [%v] %v", runRes.LastError.Code, runRes.LastError.Message))
		}
	}

	messages, err := client.ListMessage(ctx, runRes.ThreadID, utils.NewPointer(1), nil, nil, nil, utils.NewPointer(runRes.ID))
	if err != nil {
		t.Fatal(err)
	}

	t.Log(messages.Messages[0].Content[0].Text)
}

// asst_c4c52173-8eb2-41dc-bebf-91e28463352d
// thread_d1ca6f14-4877-4cc0-968b-3c908d069cf9
// run_543b149e-cdfc-4be4-bfb2-0bf03269c0d5
