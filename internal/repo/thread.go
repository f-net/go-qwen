package repo

import (
	"context"
	"gorm.io/gorm"
	"qwen/internal/model"
	"qwen/internal/types"
)

type (
	threadRepo struct {
		db *gorm.DB
	}
	IThreadRepo interface {
		First(ctx context.Context, id int64) (*model.AssistantThread, error)
		Create(ctx context.Context, thread *model.AssistantThread) error
		List(ctx context.Context, req *types.ListAssistantThreadReq) ([]*model.AssistantThread, int64, error)
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

func (repo *threadRepo) List(ctx context.Context, req *types.ListAssistantThreadReq) ([]*model.AssistantThread, int64, error) {
	var (
		err   error
		total int64
		list  []*model.AssistantThread
	)
	query := repo.db.Model(&model.AssistantThread{})
	if req.AssistantId != 0 {
		query = query.Where("assistant_id", req.AssistantId)
	}
	if req.Name != "" {
		query = query.Where("name like ?", "%"+req.Name+"%")
	}
	err = query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Limit(req.Size).Offset(req.Offset).Find(&list).Error
	return list, total, err
}
