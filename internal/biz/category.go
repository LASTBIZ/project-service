package biz

import "context"

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
