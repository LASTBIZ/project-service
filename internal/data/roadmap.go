package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"project-service/internal/biz"
	"time"
)

type RoadMap struct {
	ID          uint64 `gorm:"primaryKey"`
	StartDate   time.Time
	EndDate     time.Time
	Description string `gorm:"type:text"`
	Name        string
	Job         string `gorm:"type:text"`
	ProjectID   uint64
}

type roadmapRepo struct {
	data *Data
	log  *log.Helper
}

func NewRoadmapRepo(data *Data, logger log.Logger) biz.RoadmapRepo {
	return &roadmapRepo{data: data, log: log.NewHelper(logger)}
}

func (r roadmapRepo) CreateRoadmap(ctx context.Context, roadmap *biz.RoadMap) (*biz.RoadMap, error) {
	var rm RoadMap
	//result := r.data.db.Where(&RoadMap{StartDate: roadmap.StartDate, EndDate: roadmap.EndDate}).Find(&rm)
	//if result.RowsAffected == 1 {
	//	return nil, errors.New(500, "ROADMAP_EXISTS", "exists is name "+roadmap.Name)
	//}

	rm.StartDate = roadmap.StartDate
	rm.EndDate = roadmap.EndDate
	rm.Description = roadmap.Description
	rm.Name = roadmap.Name
	rm.Job = roadmap.Job
	rm.ProjectID = roadmap.ProjectID

	if err := r.data.db.Create(&rm); err.Error != nil {
		return nil, errors.InternalServer("ERROR_CREATE_ROADMAP", "error create roadmap")
	}

	res := ModelToResponseRoadmap(rm)
	return res, nil
}

func (r roadmapRepo) UpdateRoadMap(ctx context.Context, roadmap *biz.RoadMap) (bool, error) {
	var roadmapInfo RoadMap
	result := r.data.db.Where(&RoadMap{ID: roadmap.ID}).First(&roadmapInfo)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, errors.NotFound("ROADMAP_NOT_FOUND", "roadmap not found")
	}

	if result.RowsAffected == 0 {
		return false, errors.NotFound("ROADMAP_NOT_FOUND", "rows null")
	}

	roadmapInfo.Name = roadmap.Name
	roadmapInfo.StartDate = roadmap.StartDate
	roadmapInfo.EndDate = roadmap.EndDate
	roadmapInfo.Description = roadmap.Description
	roadmapInfo.Job = roadmap.Job

	if err := r.data.db.Save(&roadmapInfo).Error; err != nil {
		return false, errors.InternalServer("ROADMAP_UPDATE_ERROR", "error update roadmap")
	}

	return true, nil
}

func (r roadmapRepo) DeleteRoadmap(ctx context.Context, id uint64) (bool, error) {
	var roadmapInfo RoadMap
	result := r.data.db.Where(&RoadMap{ID: id}).First(&roadmapInfo)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, errors.NotFound("ROADMAP_NOT_FOUND", "roadmap not found")
	}

	if result.RowsAffected == 0 {
		return false, errors.NotFound("ROADMAP_NOT_FOUND", "rows null")
	}

	if err := r.data.db.Delete(&roadmapInfo).Error; err != nil {
		return false, errors.InternalServer("ROADMAP_DELETE_ERROR", "error delete roadmap")
	}

	return true, nil
}

func (r roadmapRepo) ListRoadmap(ctx context.Context, projectId uint64, pageNum, pageSize int) ([]*biz.RoadMap, int, error) {
	var roadmapsInfo []RoadMap
	result := r.data.db.Where(&RoadMap{ProjectID: projectId}).Find(&roadmapsInfo)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, errors.NotFound("ROADMAP_NOT_FOUND", "roadmap not found")
		}

		return nil, 0, errors.NotFound("ROADMAP_NOT_FOUND", err.Error())
	}

	total := int(result.RowsAffected)
	r.data.db.Scopes(paginate(pageNum, pageSize)).Find(&roadmapsInfo)
	rv := make([]*biz.RoadMap, 0)
	for _, u := range roadmapsInfo {
		rv = append(rv, ModelToResponseRoadmap(u))
	}
	return rv, total, nil
}

func ModelToResponseRoadmap(rm RoadMap) *biz.RoadMap {
	rmInfoRsp := &biz.RoadMap{
		ID:          rm.ID,
		StartDate:   rm.StartDate,
		EndDate:     rm.EndDate,
		Description: rm.Description,
		Name:        rm.Name,
		Job:         rm.Job,
		ProjectID:   rm.ProjectID,
	}

	return rmInfoRsp
}
