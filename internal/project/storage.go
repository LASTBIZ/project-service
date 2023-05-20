package project

import "gorm.io/gorm"

type Storage struct {
	db *gorm.DB
}

func NewProjectStorage(db *gorm.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s Storage) CreateProject(project Project) error {
	return s.db.Create(&project).Error
}

func (s Storage) UpdateProject(projectID uint32, project Project) error {
	return s.db.Model(&Project{}).Where("id = ?", projectID).
		Updates(Project{
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
	return s.db.Where("id = ?", projectID).Delete(&Project{}).Error
}

func (s Storage) GetProjectByID(projectID uint32) (*Project, error) {
	var getProject Project
	err := s.db.Where("id = ?", projectID).First(&getProject).Error
	return &getProject, err
}

// GetProjects
// paginate example db.Offset(offset).Limit(pageSize)
func (s Storage) GetProjects(paginate func(db *gorm.DB) *gorm.DB) (*[]Project, error) {
	var getProjects []Project
	err := s.db.Scopes(paginate).First(&getProjects).Error
	return &getProjects, err
}

func (s Storage) AddInvestor(investor Investor) error {
	return s.db.Create(&investor).Error
}

func (s Storage) RemoveInvestor(investorID uint32) error {
	return s.db.Where("id = ?", investorID).Delete(&Investor{}).Error
}

func (s Storage) AddMoney(investorID, money uint32) error {
	return s.db.Model(&Investor{}).Where("id = ?", investorID).
		Update("money", gorm.Expr("money + ?", money)).Error
}

func (s Storage) RemoveMoney(investorID, money uint32) error {
	return s.db.Model(&Investor{}).Where("id = ?", investorID).
		Update("money", gorm.Expr("money - ?", money)).Error
}

func (s Storage) SetMoney(investorID, money uint32) error {
	return s.db.Model(&Investor{}).Where("id = ?", investorID).
		Update("money", money).Error
}

func (s Storage) InvestProject(money, projectID uint32) error {
	return s.db.Model(&Project{}).Where("id = ?", projectID).
		Update("current", gorm.Expr("current + ?", money)).Error
}

func (s Storage) AddRoadmap(roadmap Roadmap) error {
	return s.db.Create(&roadmap).Error
}

func (s Storage) DeleteRoadmap(roadmapID uint32) error {
	return s.db.Where("id = ?", roadmapID).
		Delete(&Roadmap{}).
		Error
}
