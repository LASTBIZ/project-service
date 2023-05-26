package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"project-service/internal/biz"
)

type Category struct {
	ID               uint32      `gorm:"primaryKey; not null; unique_index" json:"id"`
	Name             string      `gorm:"type:varchar(50);not null" json:"name"`
	ParentCategoryID int32       `json:"parent_id"`
	ParentCategory   *Category   `json:"-"`
	SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID;references:ID" json:"sub_category"`
	Projects         []Project
	Level            int32 `gorm:"column:level;default:1;type:int" json:"level"`
}

type categoryRepo struct {
	data *Data
	log  *log.Helper
}

func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	return &categoryRepo{data: data, log: log.NewHelper(logger)}
}

type Result struct {
	ID   int32
	Name string
}

func (c categoryRepo) GetAllCategories(ctx context.Context, level, id int32) ([]*biz.Category, error) {
	var subQuery string
	if level == 1 {
		subQuery = fmt.Sprintf("SELECT id, name FROM categories WHERE parent_category_id IN (SELECT id, name FROM categories WHERE parent_category_id=%d)", id)
	} else if level == 2 {
		subQuery = fmt.Sprintf("SELECT id, name FROM categories WHERE parent_category_id=%d", id)
	} else if level == 3 {
		subQuery = fmt.Sprintf("SELECT id, name FROM categories WHERE id=%d", id)
	}

	var results []Result
	if err := c.data.db.Table("categories").Model(Category{}).Raw(subQuery).Scan(&results).Error; err != nil {
		return nil, errors.InternalServer("CATEGORY_ERROR", "category get errors")
	}
	for _, re := range results {
		categoryIds = append(categoryIds, re.ID)
	}
	return results, nil
}

func (c categoryRepo) GetSubCategory(ctx context.Context, parentId uint32) ([]*biz.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c categoryRepo) CreateCategory(ctx context.Context, category *biz.Category) (*biz.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c categoryRepo) DeleteCategory(ctx context.Context, id uint32) error {
	//TODO implement me
	panic("implement me")
}

func (c categoryRepo) UpdateCategory(ctx context.Context, category *biz.Category) error {
	//TODO implement me
	panic("implement me")
}
