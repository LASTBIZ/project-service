package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type RoadMap struct {
	ID          uint64
	StartDate   time.Time
	EndDate     time.Time
	Name        string
	Description string
	Job         string
	ProjectID   uint64
}

type RoadmapRepo interface {
	CreateRoadmap(ctx context.Context, roadmap *RoadMap) (*RoadMap, error)
	UpdateRoadMap(ctx context.Context, roadmap *RoadMap) (bool, error)
	DeleteRoadmap(ctx context.Context, id uint64) (bool, error)
	ListRoadmap(ctx context.Context, projectId uint64, pageNum, pageSize int) ([]*RoadMap, int, error)
}

type RoadmapUseCase struct {
	repo RoadmapRepo
	log  *log.Helper
}

func NewRoadmapUseCase(repo RoadmapRepo, logger log.Logger) *RoadmapUseCase {
	return &RoadmapUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (ru *RoadmapUseCase) CreateRoadmap(ctx context.Context, roadmap *RoadMap) (*RoadMap, error) {
	return ru.repo.CreateRoadmap(ctx, roadmap)
}

func (ru *RoadmapUseCase) UpdateRoadmap(ctx context.Context, roadMap *RoadMap) (bool, error) {
	return ru.repo.UpdateRoadMap(ctx, roadMap)
}

func (ru *RoadmapUseCase) DeleteRoadmap(ctx context.Context, id uint64) (bool, error) {
	return ru.repo.DeleteRoadmap(ctx, id)
}

func (ru *RoadmapUseCase) ListRoadmap(ctx context.Context, projectId uint64, pageNum, pageSize int) ([]*RoadMap, int, error) {
	return ru.repo.ListRoadmap(ctx, projectId, pageNum, pageSize)
}
