package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model

	ID               string `gorm:"primaryKey; not null; unique_index"`
	ParentCategoryID int32
	ParentCategory   *Category
	SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID;references:ID"`
	Level            int32       `gorm:"column:level;default:1;type:int"`
	IsTab            bool        `gorm:"default:false"`
}

func (Category) TableName() string {
	return "category"
}
