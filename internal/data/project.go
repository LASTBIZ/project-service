package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"project-service/internal/biz"
	"time"
)

type Project struct {
	ID            uint64 `gorm:"primaryKey"`
	RoadMapImgURL string
	MainImageUrl  string
	StartDate     time.Time
	EndDate       time.Time
	Location      string
	Name          string
	Description   string `gorm:"type:text"`
	CategoryID    uint64 `gorm:"not null"`
	Investors     []Investor
	Roadmaps      []RoadMap
}

type projectRepo struct {
	data *Data
	log  *log.Helper
}

func NewProjectRepo(data *Data, logger log.Logger) biz.ProjectRepo {
	return &projectRepo{data: data, log: log.NewHelper(logger)}
}

func (p projectRepo) CreateProject(ctx context.Context, project *biz.Project) (*biz.Project, error) {
	pr := Project{}
	pr.RoadMapImgURL = project.RoadMapImgURL
	pr.MainImageUrl = project.MainImageUrl
	pr.StartDate = project.StartDate
	pr.EndDate = project.EndDate
	pr.Location = project.Location
	pr.Name = project.Name
	pr.Description = project.Description
	pr.CategoryID = project.CategoryID
	//TODO add roadmaps ?
	res := p.data.db.Create(&pr)
	if res.Error != nil {
		return nil, errors.InternalServer("CREATE_PROJECT_ERROR", "error create project")
	}

	projectInfoRes := p.modelToResponse(pr)
	return projectInfoRes, nil
}

func (p projectRepo) DeleteProject(ctx context.Context, id uint64) (bool, error) {
	var projectInfo Project
	result := p.data.db.Where(&Project{ID: id}).First(&projectInfo)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, errors.NotFound("PROJECT_NOT_FOUND", "project not found")
	}

	if result.RowsAffected == 0 {
		return false, errors.NotFound("PROJECT_NOT_FOUND", "rows null")
	}

	err := p.data.db.Delete(&projectInfo)
	if err.Error != nil {
		return false, errors.InternalServer("PROJECT_DELETE_ERROR", "project delete error")
	}

	return true, nil
}

func (p projectRepo) UpdateProject(ctx context.Context, project *biz.Project) (bool, error) {
	var projectInfo Project
	result := p.data.db.Where(&Project{ID: project.ID}).First(&projectInfo)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, errors.NotFound("PROJECT_NOT_FOUND", "project not found")
	}

	if result.RowsAffected == 0 {
		return false, errors.NotFound("PROJECT_NOT_FOUND", "rows null")
	}

	projectInfo.RoadMapImgURL = project.RoadMapImgURL
	projectInfo.MainImageUrl = project.MainImageUrl
	projectInfo.StartDate = project.StartDate
	projectInfo.EndDate = project.EndDate
	projectInfo.Location = project.Location
	projectInfo.Name = project.Name
	projectInfo.Description = project.Description
	projectInfo.CategoryID = project.CategoryID

	if err := p.data.db.Save(&projectInfo); err != nil {
		return false, errors.InternalServer("PROJECT_UPDATE_ERROR", "error update project")
	}

	return true, nil
}

func (p projectRepo) GetProjectById(ctx context.Context, id uint64) (*biz.Project, error) {
	var projectInfo Project
	if err := p.data.db.Where(&Project{ID: id}).First(&projectInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("PROJECT_NOT_FOUND", "project not found")
		}

		return nil, errors.NotFound("PROJECT_NOT_FOUND", err.Error())
	}

	re := p.modelToResponse(projectInfo)
	return re, nil
}

func paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (p projectRepo) GetProjectByCategoryID(ctx context.Context, categoryID uint64, pageNum, pageSize int) ([]*biz.Project, int, error) {
	var projectsInfo []Project
	result := p.data.db.Where(&Project{CategoryID: categoryID}).Find(&projectsInfo)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, errors.NotFound("PROJECT_NOT_FOUND", "project not found")
		}

		return nil, 0, errors.NotFound("PROJECT_NOT_FOUND", err.Error())
	}

	total := int(result.RowsAffected)
	p.data.db.Scopes(paginate(pageNum, pageSize)).Find(&projectsInfo)
	rv := make([]*biz.Project, 0)
	for _, u := range projectsInfo {
		rv = append(rv, p.modelToResponse(u))
	}
	return rv, total, nil
}

func (p projectRepo) modelToResponse(project Project) *biz.Project {
	projectInfoRsp := &biz.Project{
		ID:            project.ID,
		RoadMapImgURL: project.RoadMapImgURL,
		MainImageUrl:  project.MainImageUrl,
		StartDate:     project.StartDate,
		EndDate:       project.EndDate,
		Location:      project.Location,
		Name:          project.Name,
		Description:   project.Description,
		CategoryID:    project.CategoryID,
		//TODO investors and roadmaps
	}

	return projectInfoRsp
}
