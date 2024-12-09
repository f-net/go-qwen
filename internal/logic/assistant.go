package logic

import (
	"context"
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
		Create(ctx context.Context, req *types.Assistant) error
	}
)

func NewAssistantLogic(db *gorm.DB) IAssistantLogic {
	return &assistantLogic{
		assistantRepo: repo.NewAssistantRepo(db),
		aiClient:      config.GetAssistantClient(),
	}
}

func (l *assistantLogic) Create(ctx context.Context, req *types.Assistant) error {
	assistant := &model.Assistant{
		Name:          req.Name,
		Instructions:  req.Instructions,
		Description:   req.Description,
		Model:         req.Model,
		Tools:         utils.Swap2Json(req.Tools),
		ToolResources: utils.Swap2Json(req.ToolResources),
		Remark:        req.Remark,
	}

	var assistantRequest = openai.AssistantRequest{
		Model:        req.Model,
		Name:         &req.Name,
		Instructions: &req.Instructions,
		Description:  &req.Description,
	}

	assistantRequest.Tools = req.Tools
	remoteAssistant, err := l.aiClient.CreateAssistant(context.Background(), assistantRequest)
	if err != nil {
		return err
	}

	assistant.AssistantAppId = remoteAssistant.ID
	err = l.assistantRepo.Create(ctx, assistant)
	return err
}
