package project

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"lastbiz/project-service/internal/project/models"
)

type Storage struct {
	db *gorm.DB
}

func NewProjectStorage(db *gorm.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s Storage) CreateProject(project models.Project) error {
	return s.db.Create(&project).Error
}

func (s Storage) UpdateProject(projectID uint32, project models.Project) error {
	return s.db.Model(&models.Project{}).Where("id = ?", projectID).
		Updates(models.Project{
			Location:    project.Location,
			Name:        project.Name,
			Images:      project.Images,
			Dates:       project.Dates,
			Budgets:     project.Budgets,
			Description: project.Description,
			Industry:    project.Industry,
		}).Error
}

func (s Storage) DeleteProject(projectID uint32) error {
	return s.db.Where("id = ?", projectID).Delete(&models.Project{}).Error
}

func (s Storage) GetProjectByID(projectID uint32) (*models.Project, error) {
	var getProject models.Project
	err := s.db.Where("id = ?", projectID).First(&getProject).Error
	return &getProject, err
}

// GetProjects
// paginate example db.Offset(offset).Limit(pageSize)
func (s Storage) GetProjects(paginate func(db *gorm.DB) *gorm.DB) (*[]models.Project, error) {
	var getProjects []models.Project
	err := s.db.Scopes(paginate).First(&getProjects).Error
	return &getProjects, err
}

func (s Storage) AddInvestor(investor models.Investor) error {
	return s.db.Create(&investor).Error
}

func (s Storage) RemoveInvestor(investorID uint32) error {
	return s.db.Where("id = ?", investorID).Delete(&models.Investor{}).Error
}

func (s Storage) AddMoney(investorID, money uint32) error {
	return s.db.Model(&models.Investor{}).Where("id = ?", investorID).
		Update("money", gorm.Expr("money + ?", money)).Error
}

func (s Storage) RemoveMoney(investorID, money uint32) error {
	return s.db.Model(&models.Investor{}).Where("id = ?", investorID).
		Update("money", gorm.Expr("money - ?", money)).Error
}

func (s Storage) SetMoney(investorID, money uint32) error {
	return s.db.Model(&models.Investor{}).Where("id = ?", investorID).
		Update("money", money).Error
}

func (s Storage) InvestProject(money, projectID uint32) error {
	return s.db.Model(&models.Project{}).Where("id = ?", projectID).
		Update("current", gorm.Expr("current + ?", money)).Error
}

func (s Storage) AddRoadmap(roadmap models.Roadmap) error {
	return s.db.Create(&roadmap).Error
}

func (s Storage) DeleteRoadmap(roadmapID uint32) error {
	return s.db.Where("id = ?", roadmapID).
		Delete(&models.Roadmap{}).
		Error
}

// todo category
func (s Storage) DeleteCategory(id int32) error {
	if res := s.db.Delete(&models.Category{}, id); res.RowsAffected == 0 {
		return errors.New("CATEGORY_NOT_FOUND")
	}
	return nil
}

func (s Storage) UpdateCategory(req *models.CategoryInfo) error {
	var category models.Category
	if result := s.db.First(&category, req.ID); result.RowsAffected == 0 {
		return errors.New("CATEGORY_NOT_FOUND")
	}

	if req.Name != "" {
		category.Name = req.Name
	}
	if req.ParentCategory != 0 {
		category.ParentCategoryID = req.ParentCategory
	}
	if req.Level != 0 {
		category.Level = req.Level
	}
	if req.IsTab {
		category.IsTab = req.IsTab
	}
	result := s.db.Save(&category)
	if result.Error != nil {
		return errors.New("CATEGORY_UPDATE_ERROR")
	}
	return nil
}

func (s Storage) AddCategory(req *models.CategoryInfo) (*models.CategoryInfo, error) {
	cMap := map[string]interface{}{}
	cMap["name"] = req.Name
	cMap["level"] = req.Level
	cMap["is_tab"] = req.IsTab

	if req.Level != 1 {
		var categories models.Category
		if res := s.db.First(&categories, req.ParentCategory); res.RowsAffected == 0 {
			return nil, errors.New("CATEGORY_NOT_FOUND")
		}
		cMap["parent_category_id"] = req.ParentCategory
	}

	result := s.db.Model(&models.Category{}).Create(&cMap)
	if result.Error != nil {
		return nil, errors.New("CATEGORY_CREATE_ERROR")
	}
	var value int32
	value, ok := cMap["parent_category_id"].(int32)
	if !ok {
		value = 0
	}
	res := &models.CategoryInfo{
		Name:           cMap["name"].(string),
		ParentCategory: value,
		Level:          cMap["level"].(int32),
		IsTab:          cMap["is_tab"].(bool),
	}
	return res, nil

}

func (s Storage) Category() ([]*models.Category, error) {
	var cate []*models.Category
	result := s.db.Where(&models.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&cate)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("CATEGORY_NOT_FOUND")
	}
	if result.Error != nil {
		return nil, errors.New("CATEGORY_NOT_FOUND")
	}

	return cate, nil
}

func (s Storage) GetCategoryByID(id int32) (*models.CategoryInfo, error) {
	var categories models.Category
	if res := s.db.First(&categories, id); res.RowsAffected == 0 {
		return nil, errors.New("CATEGORY_NOT_FOUND")
	}

	info := &models.CategoryInfo{
		ID:             categories.ID,
		Name:           categories.Name,
		ParentCategory: categories.ParentCategoryID,
		Level:          categories.Level,
		IsTab:          categories.IsTab,
	}
	return info, nil
}

func (s Storage) SubCategory(req models.CategoryInfo) ([]*models.CategoryInfo, error) {
	var subCategory []models.Category
	var subCategoryInfo []*models.CategoryInfo
	preload := "SubCategory"
	if req.Level == 1 {
		preload = "SubCategory.SubCategory"
	}

	if err := s.db.Where(&models.Category{ParentCategoryID: int32(req.ID)}).Preload(preload).Find(&subCategory).Error; err != nil {
		return nil, errors.New("CATEGORY_NOT_FOUND")
	}
	for _, v := range subCategory {
		categoryInfo := &models.CategoryInfo{
			ID:                 v.ID,
			Name:               v.Name,
			ParentCategory:     v.ParentCategoryID,
			ParentCategoryName: "",
			Level:              v.Level,
			IsTab:              v.IsTab,
		}
		if v.ParentCategoryID != 0 {
			categoryInfo.ParentCategoryName = v.ParentCategory.Name
		}
		subCategoryInfo = append(subCategoryInfo, categoryInfo)
	}

	return subCategoryInfo, nil
}

type Result struct {
	ID   int32
	Name string
}

func (s Storage) GetCategoryAll(level, id int32) ([]Result, error) {
	//categoryIds := make([]int32, 0)
	var subQuery string
	if level == 1 {
		subQuery = fmt.Sprintf("SELECT id, name FROM categories WHERE parent_category_id IN (SELECT id, name FROM categories WHERE parent_category_id=%d)", id)
	} else if level == 2 {
		subQuery = fmt.Sprintf("SELECT id, name FROM categories WHERE parent_category_id=%d", id)
	} else if level == 3 {
		subQuery = fmt.Sprintf("SELECT id, name FROM categories WHERE id=%d", id)
	}

	var results []Result
	if err := s.db.Table("categories").Model(models.Category{}).Raw(subQuery).Scan(&results).Error; err != nil {
		return nil, errors.New("CATEGORY_ERROR")
	}
	//for _, re := range results {
	//	categoryIds = append(categoryIds, re.ID)
	//}
	return results, nil
}

func (s Storage) SubCategoryList(cid int32) (*models.CategoryList, error) {
	cateInfo, err := s.GetCategoryByID(cid)
	if err != nil {
		return nil, err
	}

	category, err := s.SubCategory(*cateInfo)
	if err != nil {
		return nil, err
	}

	return &models.CategoryList{
		Category:    cateInfo,
		SubCategory: category,
	}, nil
}
