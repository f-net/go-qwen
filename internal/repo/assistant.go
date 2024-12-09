package repo

import (
	"context"
	"gorm.io/gorm"
	"qwen/internal/model"
)

type (
	assistantRepo struct {
		db *gorm.DB
	}
	IAssistantRepo interface {
		Create(ctx context.Context, assistant *model.Assistant) error
	}
)

func NewAssistantRepo(db *gorm.DB) IAssistantRepo {
	return &assistantRepo{
		db: db,
	}
}

func (repo *assistantRepo) Create(ctx context.Context, assistant *model.Assistant) error {
	return repo.db.Create(assistant).Error
}
