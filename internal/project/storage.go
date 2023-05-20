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
