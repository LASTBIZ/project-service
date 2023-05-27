package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Category struct {
	ID               uint32
	Name             string
	ParentCategory   *Category
	ParentCategoryID uint32
	SubCategory      []*Category
	Projects         []Project
	Level            int32
}

type CategoryRepo interface {
	GetAllCategories(ctx context.Context) (string, error)
	CreateCategory(ctx context.Context, category *Category) (*Category, error)
	DeleteCategory(ctx context.Context, id uint32) error
	UpdateCategory(ctx context.Context, category *Category) error
}

type CategoryUseCase struct {
	repo CategoryRepo
	log  *log.Helper
}

func NewCategoryUseCase(repo CategoryRepo, logger log.Logger) *CategoryUseCase {
	return &CategoryUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (cu *CategoryUseCase) GetAllCategories(ctx context.Context) (string, error) {

	return cu.repo.GetAllCategories(ctx)
}

func (cu *CategoryUseCase) CreateCategory(ctx context.Context, category *Category) (*Category, error) {
	return cu.repo.CreateCategory(ctx, category)
}

func (cu *CategoryUseCase) DeleteCategory(ctx context.Context, id uint32) error {
	return cu.repo.DeleteCategory(ctx, id)
}

func (cu *CategoryUseCase) UpdateCategory(ctx context.Context, category *Category) error {
	return cu.repo.UpdateCategory(ctx, category)
}
