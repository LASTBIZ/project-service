package models

import (
	"errors"
	"lastbiz/project-service/pkg/project/pb"
	"strings"
)

type RoadMapSlice []Roadmap

func (p RoadMapSlice) Len() int {
	return len(p)
}

func (p RoadMapSlice) Less(i, j int) bool {
	return p[i].Dates.StartDate.Before(p[j].Dates.StartDate)
}

func (p RoadMapSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type Roadmap struct {
	ID          uint64 `gorm:"primaryKey"`
	Dates       Date   `gorm:"embedded"`
	Name        string
	Description string `gorm:"type:text"`
	Job         string `gorm:"type:text"`
	Target      string `gorm:"type:text"`
	ProjectID   uint64
}

func (r Roadmap) Validate() error {
	if r.Dates.StartDate.IsZero() {
		return errors.New("start date is required")
	}
	if r.Dates.EndDate.IsZero() {
		return errors.New("end date is required")
	}
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("name is required")
	}
	if strings.TrimSpace(r.Description) == "" {
		return errors.New("description is required")
	}
	if strings.TrimSpace(r.Job) == "" {
		return errors.New("job is required")
	}
	if strings.TrimSpace(r.Target) == "" {
		return errors.New("target is required")
	}
	return nil
}

func (r Roadmap) ToRPC() *pb.RoadMap {
	return &pb.RoadMap{
		Id:          r.ID,
		Dates:       r.Dates.ToRPC(),
		Name:        r.Name,
		Description: r.Description,
		Job:         r.Job,
		Target:      r.Target,
	}
}

type GRPCRoadmap struct {
	*pb.RoadMap
}

func (r *GRPCRoadmap) ToRoadmap() Roadmap {
	return Roadmap{
		ID: r.GetId(),
		Dates: Date{
			StartDate: r.Dates.GetStartDate().AsTime(),
			EndDate:   r.Dates.GetEndDate().AsTime(),
		},
		Name:        r.Name,
		Description: r.Description,
		Job:         r.Job,
		Target:      r.Target,
	}
}
