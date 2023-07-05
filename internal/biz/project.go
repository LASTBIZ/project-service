package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Project struct {
	ID            uint64
	RoadMapImgURL string
	MainImageUrl  string
	StartDate     time.Time
	EndDate       time.Time
	Location      string
	Name          string
	Description   string
	CategoryID    uint32
	Category      Category
	CurrentBudget uint64
	NeedBudget    uint64
	Investors     []Investor
	Roadmaps      []RoadMap
}

type ProjectRepo interface {
	CreateProject(ctx context.Context, project *Project) (*Project, error)
	DeleteProject(ctx context.Context, id uint64) (bool, error)
	UpdateProject(ctx context.Context, project *Project) (bool, error)
	GetProjectById(ctx context.Context, id uint64) (*Project, error)
	GetProjectByCategoryID(ctx context.Context, categoryID uint32, keyWords string, pageNum, pageSize int) ([]*Project, int, error)
	InvestProject(ctx context.Context, id uint64, investorID uint64, money int) error
	Video(ctx context.Context, id uint64, video string) (bool, error)
	ScreenShot(ctx context.Context, id uint64, screenShot string) (bool, error)
	InLive(ctx context.Context, id uint64, inLive bool) (bool, error)
}

type ProjectUseCase struct {
	repo ProjectRepo
	log  *log.Helper
}

func NewProjectUseCase(repo ProjectRepo, logger log.Logger) *ProjectUseCase {
	return &ProjectUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (up *ProjectUseCase) Create(ctx context.Context, p *Project) (*Project, error) {
	return up.repo.CreateProject(ctx, p)
}

func (up *ProjectUseCase) Update(ctx context.Context, p *Project) (bool, error) {
	return up.repo.UpdateProject(ctx, p)
}

func (up *ProjectUseCase) Delete(ctx context.Context, id uint64) (bool, error) {
	return up.repo.DeleteProject(ctx, id)
}

func (up *ProjectUseCase) GetProjectById(ctx context.Context, id uint64) (*Project, error) {
	return up.repo.GetProjectById(ctx, id)
}

func (up *ProjectUseCase) GetProjectByCategoryId(ctx context.Context, keywords string, categoryID uint32, pageNum, pageSize int) ([]*Project, int, error) {
	return up.repo.GetProjectByCategoryID(ctx, categoryID, keywords, pageNum, pageSize)
}

func (up *ProjectUseCase) InvestProject(ctx context.Context, id uint64, investorId uint64, money int) error {
	return up.repo.InvestProject(ctx, id, investorId, money)
}

func (up *ProjectUseCase) Video(ctx context.Context, id uint64, video string) (bool, error) {
	return up.repo.Video(ctx, id, video)
}

func (up *ProjectUseCase) ScreenShot(ctx context.Context, id uint64, screenShot string) (bool, error) {
	return up.repo.ScreenShot(ctx, id, screenShot)
}

func (up *ProjectUseCase) InLive(ctx context.Context, id uint64, inLive bool) (bool, error) {
	return up.repo.InLive(ctx, id, inLive)
}
