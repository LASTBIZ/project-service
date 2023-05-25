package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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
	CategoryID    string `gorm:"not null"`
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
	//TODO implement me
	panic("implement me")
}

func (p projectRepo) DeleteProject(ctx context.Context, id uint64) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (p projectRepo) UpdateProject(ctx context.Context, project *biz.Project) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (p projectRepo) GetProjectById(ctx context.Context, id uint64) (*biz.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p projectRepo) GetProjectByCategoryID(ctx context.Context, categoryID uint64) ([]*biz.Project, error) {
	//TODO implement me
	panic("implement me")
}
