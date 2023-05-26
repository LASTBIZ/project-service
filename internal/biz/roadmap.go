package biz

import (
	"context"
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
