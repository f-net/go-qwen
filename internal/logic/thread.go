package logic

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
	"qwen/internal/config"
	"qwen/internal/repo"
	"qwen/internal/types"
)

type (
	assistantThreadLogic struct {
		assistantRepo repo.IAssistantRepo
		threadRepo    repo.IThreadRepo
		aiClient      *openai.Client
	}
	IAssistantThreadLogic interface {
		List(ctx context.Context, req *types.ListAssistantThreadReq) (*types.ListCommonResp, error)
	}
)

func NewAssistantThreadLogic(db *gorm.DB) IAssistantThreadLogic {
	return &assistantThreadLogic{
		assistantRepo: repo.NewAssistantRepo(db),
		threadRepo:    repo.NewThreadRepo(db),
		aiClient:      config.GetAssistantClient(),
	}
}

func (l *assistantThreadLogic) List(ctx context.Context, req *types.ListAssistantThreadReq) (*types.ListCommonResp, error) {
	list, total, err := l.threadRepo.List(ctx, req)
	if err != nil {
		return nil, err
	}
	var resp = make([]*types.ListAssistantThreadResp, 0)
	for _, thread := range list {
		resp = append(resp, &types.ListAssistantThreadResp{
			Id:   thread.Id,
			Name: thread.Name,
		})
	}
	return &types.ListCommonResp{
		List:      list,
		Page:      req.Page,
		Size:      req.Size,
		Total:     total,
		TotalPage: 1,
	}, nil
}
