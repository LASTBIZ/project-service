package data

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"project-service/internal/biz"
)

type Category struct {
	ID               uint32      `gorm:"primaryKey; not null; unique_index" json:"id"`
	Name             string      `gorm:"type:varchar(50);not null" json:"name"`
	ParentCategoryID uint32      `json:"parent_id"`
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

func (c categoryRepo) GetAllCategories(ctx context.Context) (string, error) {
	/*
		[
			{
				"id":xxx,
				"name":"",
				"level":1,
				"parent":13xxx,
				"sub_category":[
					"id":xxx,
					"name":"",
					"level":1,
					"is_tab":false,
					"sub_category":[]
				]
			}
		]
	*/
	var categories []Category
	c.data.db.Where(&Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categories)
	b, _ := json.Marshal(&categories)
	return string(b), nil
}

func (c categoryRepo) CreateCategory(ctx context.Context, category *biz.Category) (*biz.Category, error) {
	cMap := Category{}
	cMap.Name = category.Name
	cMap.Level = category.Level
	if category.Level != 1 {
		var categories Category
		if res := c.data.db.First(&categories, category.ParentCategoryID); res.RowsAffected == 0 {
			return nil, errors.NotFound("CATEGORY_NOT_FOUND", "商品分类不存在")
		}
		cMap.ParentCategoryID = category.ParentCategoryID
	}

	result := c.data.db.Model(&Category{}).Create(cMap)
	if result.Error != nil {
		return nil, errors.InternalServer("CATEGORY_CREATE_ERROR", result.Error.Error())
	}

	return &biz.Category{
		Name:             cMap.Name,
		ParentCategoryID: cMap.ParentCategoryID,
		Level:            cMap.Level,
	}, nil
}

func (c categoryRepo) DeleteCategory(ctx context.Context, id uint32) error {
	if res := c.data.db.Delete(&Category{}, id); res.RowsAffected == 0 {
		return errors.InternalServer("DELETE_CATEGORY_ERROR", res.Error.Error())
	}
	return nil
}

func (c categoryRepo) UpdateCategory(ctx context.Context, cat *biz.Category) error {
	var category Category
	if result := c.data.db.First(&category, cat.ID); result.RowsAffected == 0 {
		return errors.NotFound("CATEGORY_NOT_FOUND", "category not found")
	}

	if cat.Name != "" {
		category.Name = cat.Name
	}
	if cat.ParentCategoryID != 0 {
		category.ParentCategoryID = cat.ParentCategoryID
	}
	if cat.Level != 0 {
		category.Level = cat.Level
	}
	result := c.data.db.Save(&category)
	if result.Error != nil {
		return errors.InternalServer("CATEGORY_UPDATE_ERROR", "error update category")
	}
	return nil
}

//func ModelToResponseCategory(cat Category) *biz.Category {
//	catInfoRsp := &biz.Category{
//		ID:               cat.ID,
//		Name:             cat.Name,
//		Level:            cat.Level,
//		ParentCategoryID: uint32(cat.ParentCategoryID),
//	}
//	if cat.ParentCategory != nil {
//		catInfoRsp.ParentCategory = ModelToResponseCategory(*cat.ParentCategory)
//	}
//	subCategory := make([]*biz.Category, 0)
//	for _, cv := range cat.SubCategory {
//		if cv != nil {
//			subCategory = append(subCategory, ModelToResponseCategory(*cv))
//		}
//	}
//	catInfoRsp.SubCategory = subCategory
//	return catInfoRsp
//}
