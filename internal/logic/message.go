package logic

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
	"qwen/internal/config"
	"qwen/internal/model"
	"qwen/internal/repo"
	"qwen/internal/types"
	"qwen/internal/utils"
	"time"
)

type (
	messageLogic struct {
		assistantRepo repo.IAssistantRepo
		threadRepo    repo.IThreadRepo
		messageRepo   repo.IMessageRepo
		aiClient      *openai.Client
	}
	IMessageLogic interface {
		List(ctx context.Context, req *types.GetMessageListReq) (*types.ListCommonResp, error)
		SendMessage(ctx context.Context, req *types.CreateMessageReq) (string, int64, error)
	}
)

func NewMessageLogic(db *gorm.DB) IMessageLogic {
	return &messageLogic{
		assistantRepo: repo.NewAssistantRepo(db),
		threadRepo:    repo.NewThreadRepo(db),
		aiClient:      config.GetAssistantClient(),
		messageRepo:   repo.NewMessageRepo(db),
	}
}

func (l *messageLogic) List(ctx context.Context, req *types.GetMessageListReq) (*types.ListCommonResp, error) {
	list, total, err := l.messageRepo.List(ctx, req)
	if err != nil {
		return nil, err
	}
	var resp = make([]*types.GetMessageListResp, 0)
	for _, thread := range list {
		resp = append(resp, &types.GetMessageListResp{
			Id:       thread.Id,
			Question: thread.Question,
			Answer:   thread.Answer,
		})
	}
	return &types.ListCommonResp{
		List:      list,
		Page:      req.Page,
		Size:      req.Size,
		Total:     total,
		TotalPage: utils.GetTotalPage(total, req.Size),
	}, nil
}

func (l *messageLogic) SendMessage(ctx context.Context, req *types.CreateMessageReq) (string, int64, error) {

	assistant, err := l.assistantRepo.First(ctx, req.AssistantId)
	if err != nil {
		return "", 0, err
	}
	if assistant.Id == 0 {
		return "", 0, nil
	}

	thread, err := l.threadRepo.First(ctx, req.ThreadId)
	if err != nil {
		return "", 0, err
	}
	if thread == nil || thread.Id == 0 {
		threadRes, err := l.aiClient.CreateThread(ctx, openai.ThreadRequest{
			Messages: []openai.ThreadMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: req.Question,
				},
			},
		})
		if err != nil {
			return "", 0, err
		}
		fmt.Println(threadRes.ID)
		thread = &model.Thread{
			AssistantId: req.AssistantId,
			Name:        req.Question,
			RemoteId:    threadRes.ID,
		}

		err = l.threadRepo.Create(ctx, thread)
		if err != nil {
			return "", 0, err
		}
	}

	messageResp, err := l.aiClient.CreateMessage(ctx, thread.RemoteId, openai.MessageRequest{
		Role:    "user",
		Content: req.Question,
	})
	if err != nil {
		return "", 0, err
	}
	runRes, err := l.aiClient.CreateRun(ctx, thread.RemoteId, openai.RunRequest{
		AssistantID: assistant.RemoteId,
	})
	if err != nil {
		return "", 0, err
	}
	for runRes.Status == openai.RunStatusQueued || runRes.Status == openai.RunStatusInProgress {
		//fmt.Println("排队和正在运行")
		runRes, err = l.aiClient.RetrieveRun(ctx, runRes.ThreadID, runRes.ID)
		if err != nil {
			return "", 0, err
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

	messages, err := l.aiClient.ListMessage(ctx, runRes.ThreadID, utils.NewPointer(1), nil, nil, nil, utils.NewPointer(runRes.ID))
	if err != nil {
		return "", 0, err
	}

	answer := messages.Messages[0].Content[0].Text

	err = l.messageRepo.Create(ctx, &model.Message{
		AssistantId: req.AssistantId,
		Question:    req.Question,
		Answer:      answer.Value,
		RemoteId:    messageResp.ID,
		ThreadId:    thread.Id,
	})

	return answer.Value, thread.Id, nil
}
