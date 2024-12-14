package repo

import (
	"context"
	"gorm.io/gorm"
	"qwen/internal/model"
	"qwen/internal/types"
)

type (
	messageRepo struct {
		db *gorm.DB
	}
	IMessageRepo interface {
		Create(ctx context.Context, Message *model.Message) error
		List(ctx context.Context, req *types.GetMessageListReq) ([]*model.Message, int64, error)
	}
)

func NewMessageRepo(db *gorm.DB) IMessageRepo {
	return &messageRepo{
		db: db,
	}
}

func (repo *messageRepo) Create(ctx context.Context, Message *model.Message) error {

	return repo.db.Create(Message).Error
}

func (repo *messageRepo) List(ctx context.Context, req *types.GetMessageListReq) ([]*model.Message, int64, error) {
	var (
		err   error
		total int64
		list  []*model.Message
	)
	query := repo.db.Model(&model.Message{}).Debug()
	if req.AssistantId != 0 {
		query = query.Where("assistant_id", req.AssistantId)
	}
	if req.ThreadId != 0 {
		query = query.Where("thread_id", req.ThreadId)
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
