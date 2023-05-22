package models

type Category struct {
	ID               uint32      `gorm:"primaryKey; not null; unique_index" json:"id"`
	Name             string      `gorm:"type:varchar(50);not null" json:"name"`
	ParentCategoryID int32       `json:"parent_id"`
	ParentCategory   *Category   `json:"-"`
	SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID;references:ID" json:"sub_category"`
	Level            int32       `gorm:"column:level;default:1;type:int" json:"level"`
	IsTab            bool        `gorm:"default:false" json:"is_tab"`
}

func (Category) TableName() string {
	return "category"
}

type CategoryInfo struct {
	ID                 uint32
	Name               string
	ParentCategory     int32
	ParentCategoryName string
	Level              int32
	IsTab              bool
}

type CategoryList struct {
	Category    *CategoryInfo
	SubCategory []*CategoryInfo
}
