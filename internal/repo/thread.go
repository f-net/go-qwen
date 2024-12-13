package repo

import (
	"context"
	"gorm.io/gorm"
	"qwen/internal/model"
)

type (
	threadRepo struct {
		db *gorm.DB
	}
	IThreadRepo interface {
		First(ctx context.Context, id int64) (*model.AssistantThread, error)
		Create(ctx context.Context, thread *model.AssistantThread) error
	}
)

func NewThreadRepo(db *gorm.DB) IThreadRepo {
	return &threadRepo{
		db: db,
	}
}

func (repo *threadRepo) First(ctx context.Context, id int64) (*model.AssistantThread, error) {
	var assistant model.AssistantThread
	err := repo.db.Where("id", id).First(&assistant).Error
	if err != nil {
		return nil, err
	}
	return &assistant, nil
}
func (repo *threadRepo) Create(ctx context.Context, thread *model.AssistantThread) error {

	return repo.db.Create(thread).Error
}
