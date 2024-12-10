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
		aiClient      *openai.Client
	}
	IAssistantLogic interface {
		Create(ctx context.Context, req *types.CreateAssistantReq) error
		Save(ctx context.Context, req *types.UpdateAssistantReq) error
		First(ctx context.Context, id int64) (*types.GetAssistantResp, error)
		List(ctx context.Context, req *types.ListAssistantReq) (*types.ListCommonResp, error)
	}
)

func NewAssistantLogic(db *gorm.DB) IAssistantLogic {
	return &assistantLogic{
		assistantRepo: repo.NewAssistantRepo(db),
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

	assistant.AssistantAppId = remoteAssistant.ID
	err = l.assistantRepo.Create(ctx, assistant)
	return err
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
	remoteAssistant, err := l.aiClient.ModifyAssistant(ctx, assistant.AssistantAppId, assistantRequest)
	if err != nil {
		return err
	}

	assistant.AssistantAppId = remoteAssistant.ID
	err = l.assistantRepo.Save(ctx, assistant)
	return err
}

func (l *assistantLogic) First(ctx context.Context, id int64) (*types.GetAssistantResp, error) {
	assistant, err := l.assistantRepo.First(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := &types.GetAssistantResp{
		AssistantAppId: assistant.AssistantAppId,
		Id:             assistant.Id,
		Instructions:   assistant.Instructions,
		Model:          assistant.Model,
		Name:           assistant.Name,
		Remark:         assistant.Remark,
	}
	_ = json.Unmarshal(assistant.ToolResources, &resp.ToolResources)
	_ = json.Unmarshal(assistant.Tools, &resp.Tools)

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
