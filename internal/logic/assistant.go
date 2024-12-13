package logic

import (
	"context"
	"encoding/json"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
	"qwen/internal/config"
	"qwen/internal/model"
	"qwen/internal/repo"
	"qwen/internal/types"
	"qwen/internal/utils"
)

type (
	assistantLogic struct {
		assistantRepo repo.IAssistantRepo
		threadRepo    repo.IThreadRepo
		aiClient      *openai.Client
	}
	IAssistantLogic interface {
		Create(ctx context.Context, req *types.CreateAssistantReq) error
		Delete(ctx context.Context, id int64) error
		Save(ctx context.Context, req *types.UpdateAssistantReq) error
		First(ctx context.Context, id int64) (*types.GetAssistantResp, error)
		List(ctx context.Context, req *types.ListAssistantReq) (*types.ListCommonResp, error)
	}
)

func NewAssistantLogic(db *gorm.DB) IAssistantLogic {
	return &assistantLogic{
		assistantRepo: repo.NewAssistantRepo(db),
		threadRepo:    repo.NewThreadRepo(db),
		aiClient:      config.GetAssistantClient(),
	}
}

func (l *assistantLogic) Create(ctx context.Context, req *types.CreateAssistantReq) error {
	assistant := &model.Assistant{
		Name:          req.Name,
		Instructions:  req.Instructions,
		Model:         req.Model,
		Tools:         utils.Swap2Json(req.Tools),
		ToolResources: utils.Swap2Json(req.ToolResources),
		Remark:        req.Remark,
	}

	var assistantRequest = openai.AssistantRequest{
		Model:        req.Model,
		Name:         &req.Name,
		Instructions: &req.Instructions,
		Description:  &req.Remark,
	}

	assistantRequest.Tools = req.Tools
	remoteAssistant, err := l.aiClient.CreateAssistant(ctx, assistantRequest)
	if err != nil {
		return err
	}

	assistant.RemoteId = remoteAssistant.ID
	err = l.assistantRepo.Create(ctx, assistant)
	return err
}

func (l *assistantLogic) Delete(ctx context.Context, id int64) error {

	assistant, err := l.assistantRepo.First(ctx, id)
	if err != nil {
		return err
	}

	if assistant.Id == 0 {
		return nil
	}

	err = l.assistantRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	_, err = l.aiClient.DeleteAssistant(ctx, assistant.RemoteId)
	if err != nil {
		return err
	}

	return nil
}

func (l *assistantLogic) Save(ctx context.Context, req *types.UpdateAssistantReq) error {
	assistant, err := l.assistantRepo.First(ctx, req.Id)
	if err != nil {
		return err
	}

	assistant.Id = req.Id
	assistant.Name = req.Name
	assistant.Instructions = req.Instructions
	assistant.Model = req.Model
	assistant.Tools = utils.Swap2Json(req.Tools)
	assistant.ToolResources = utils.Swap2Json(req.ToolResources)
	assistant.Remark = req.Remark

	var assistantRequest = openai.AssistantRequest{
		Model:        req.Model,
		Name:         &req.Name,
		Instructions: &req.Instructions,
		Description:  &req.Remark,
	}

	assistantRequest.Tools = req.Tools
	remoteAssistant, err := l.aiClient.ModifyAssistant(ctx, assistant.RemoteId, assistantRequest)
	if err != nil {
		return err
	}

	assistant.RemoteId = remoteAssistant.ID
	err = l.assistantRepo.Save(ctx, assistant)
	return err
}

func (l *assistantLogic) First(ctx context.Context, id int64) (*types.GetAssistantResp, error) {
	assistant, err := l.assistantRepo.First(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := &types.GetAssistantResp{
		RemoteId:     assistant.RemoteId,
		Id:           assistant.Id,
		Instructions: assistant.Instructions,
		Model:        assistant.Model,
		Name:         assistant.Name,
		Remark:       assistant.Remark,
	}
	_ = json.Unmarshal([]byte(assistant.ToolResources), &resp.ToolResources)
	_ = json.Unmarshal([]byte(assistant.Tools), &resp.Tools)

	return resp, nil
}

func (l *assistantLogic) List(ctx context.Context, req *types.ListAssistantReq) (*types.ListCommonResp, error) {
	var resp = make([]types.ListAssistantResp, 0)
	assistants, total, err := l.assistantRepo.Search(ctx, req)
	if err != nil {
		return nil, err
	}
	for _, assistant := range *assistants {
		resp = append(resp, types.ListAssistantResp{
			CreatedAt:    assistant.CreatedAt,
			Id:           assistant.Id,
			Instructions: assistant.Instructions,
			Model:        assistant.Model,
			Name:         assistant.Name,
			Remark:       assistant.Remark,
		})
	}

	return &types.ListCommonResp{
		List:      resp,
		Total:     total,
		Page:      req.Page,
		Size:      req.Size,
		TotalPage: utils.GetTotalPage(total, req.Size),
	}, nil
}

func (l *assistantLogic) SendMessage(ctx context.Context, req *types.Message) (*types.GetAssistantResp, error) {
	assistant, err := l.assistantRepo.First(ctx, req.AssistantId)
	if err != nil {
		return nil, err
	}

	thread, err := l.threadRepo.First(ctx, req.ThreadId)
	if err != nil {
		return nil, nil
	}
	if thread.Id == 0 {
		threadResp, err := l.aiClient.CreateThread(
			ctx,
			openai.ThreadRequest{
				Messages: []openai.ThreadMessage{
					{
						Role:    openai.ThreadMessageRoleAssistant,
						Content: req.Message,
					},
				},
			},
		)
		if err != nil {
			return nil, err
		}

		tmpThread := &model.AssistantThread{
			Name:        req.Message,
			AssistantId: req.AssistantId,
			RemoteId:    threadResp.ID,
		}

		err = l.threadRepo.Create(ctx, tmpThread)
		if err != nil {
			return nil, err
		}
		thread = tmpThread
	}

	l.aiClient.CreateThreadAndRun()

	return resp, nil
}
