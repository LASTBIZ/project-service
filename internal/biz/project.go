package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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

type ProjectRepo interface {
	CreateProject(ctx context.Context, project *Project) (*Project, error)
	DeleteProject(ctx context.Context, id uint64) (bool, error)
	UpdateProject(ctx context.Context, project *Project) (bool, error)
	GetProjectById(ctx context.Context, id uint64) (*Project, error)
	GetProjectByCategoryID(ctx context.Context, categoryID uint64) ([]*Project, error)
}

type ProjectUseCase struct {
	repo ProjectRepo
	log  *log.Helper
}

func NewProjectUseCase(repo ProjectRepo, logger log.Logger) *ProjectUseCase {
	return &ProjectUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (up *ProjectUseCase) Create(ctx context.Context, p *Project) (*Project, error) {
	return up.Create(ctx, p)
}

func (up *ProjectUseCase) Update(ctx context.Context, p *Project) (bool, error) {
	return up.Update(ctx, p)
}

func (up *ProjectUseCase) Delete(ctx context.Context, id uint64) (bool, error) {
	return up.Delete(ctx, id)
}

func (up *ProjectUseCase) GetProjectById(ctx context.Context, id uint64) (*Project, error) {
	return up.GetProjectById(ctx, id)
}

func (up *ProjectUseCase) GetProjectByCategoryId(ctx context.Context, categoryID uint64) ([]*Project, error) {
	return up.GetProjectByCategoryId(ctx, categoryID)
}
