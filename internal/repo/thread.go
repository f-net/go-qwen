package repo

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"qwen/internal/model"
	"qwen/internal/types"
)

type (
	threadRepo struct {
		db *gorm.DB
	}
	IThreadRepo interface {
		First(ctx context.Context, id int64) (*model.Thread, error)
		Create(ctx context.Context, thread *model.Thread) error
		List(ctx context.Context, req *types.ListAssistantThreadReq) ([]*model.Thread, int64, error)
	}
)

func NewThreadRepo(db *gorm.DB) IThreadRepo {
	return &threadRepo{
		db: db,
	}
}

func (repo *threadRepo) First(ctx context.Context, id int64) (*model.Thread, error) {
	var assistant model.Thread
	err := repo.db.Where("id", id).First(&assistant).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &assistant, nil
}
func (repo *threadRepo) Create(ctx context.Context, thread *model.Thread) error {

	return repo.db.Create(thread).Error
}

func (repo *threadRepo) List(ctx context.Context, req *types.ListAssistantThreadReq) ([]*model.Thread, int64, error) {
	var (
		err   error
		total int64
		list  []*model.Thread
	)
	query := repo.db.Model(&model.Thread{}).Debug().Order("id desc")
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
	err = query.Find(&list).Error
	return list, total, err
}
