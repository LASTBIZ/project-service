package biz

import "context"

type Category struct {
	ID             uint32
	Name           string
	ParentCategory *Category
	SubCategory    []*Category
	Projects       []Project
	Level          int32
}

type CategoryRepo interface {
	GetAllCategories(ctx context.Context, level, id int32) ([]*Category, error)
	GetSubCategory(ctx context.Context, parentId uint32) ([]*Category, error)
	CreateCategory(ctx context.Context, category *Category) (*Category, error)
	DeleteCategory(ctx context.Context, id uint32) error
	UpdateCategory(ctx context.Context, category *Category) error
}
