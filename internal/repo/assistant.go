package repo

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"qwen/internal/model"
	"qwen/internal/types"
)

type (
	assistantRepo struct {
		db *gorm.DB
	}
	IAssistantRepo interface {
		Create(ctx context.Context, assistant *model.Assistant) error
		Save(ctx context.Context, assistant *model.Assistant) error
		First(ctx context.Context, id int64) (*model.Assistant, error)
		Search(ctx context.Context, req *types.ListAssistantReq) (*[]model.Assistant, int64, error)
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
func (repo *assistantRepo) Save(ctx context.Context, assistant *model.Assistant) error {
	return repo.db.Save(assistant).Error
}
func (repo *assistantRepo) First(ctx context.Context, id int64) (*model.Assistant, error) {
	var assistant model.Assistant
	err := repo.db.First(&assistant).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &assistant, nil
}
func (repo *assistantRepo) Search(ctx context.Context, req *types.ListAssistantReq) (*[]model.Assistant, int64, error) {
	var assistants []model.Assistant
	var total int64

	query := repo.db.Model(&model.Assistant{})
	if req.Name != "" {
		query = query.Where("name like ?", "%"+req.Name+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Limit(req.Size).Offset(req.Offset).Find(&assistants).Error
	if err != nil {
		return nil, 0, err
	}

	return &assistants, total, nil
}
