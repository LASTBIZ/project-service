package services

import (
	"context"
	project1 "lastbiz/project-service/internal/project"
	"lastbiz/project-service/pkg/project"
	"net/http"
)

func (s Service) AddRoadmap(ctx context.Context, req *project.CreateRoadmapRequest) (*project.RoadmapResponse, error) {
	if req.GetProjectId() == 0 {
		return &project.RoadmapResponse{
			Status: http.StatusConflict,
			Error:  "project_id is required",
		}, nil
	}

	if req.GetRoadmap() == nil {
		return &project.RoadmapResponse{
			Status: http.StatusConflict,
			Error:  "roadmap is null",
		}, nil
	}

	r := req.GetRoadmap()

	createRoadmap := project1.Roadmap{
		Dates: project1.Date{
			StartDate: r.GetDates().GetStartDate().AsTime(),
			EndDate:   r.GetDates().GetEndDate().AsTime(),
		},
		Name:        r.GetName(),
		Description: r.GetDescription(),
		Job:         r.GetJob(),
		Target:      r.GetTarget(),
		ProjectID:   uint64(req.GetProjectId()),
	}

	err := createRoadmap.Validate()
	if err != nil {
		return &project.RoadmapResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	err = s.storage.AddRoadmap(createRoadmap)
	if err != nil {
		return &project.RoadmapResponse{
			Status: http.StatusInternalServerError,
			Error:  "error create roadmap",
		}, nil
	}

	return &project.RoadmapResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s Service) RemoveRoadmap(ctx context.Context, req *project.DeleteRoadmapRequest) (*project.RoadmapResponse, error) {
	if req.GetRoadmapId() == 0 {
		return &project.RoadmapResponse{
			Status: http.StatusInternalServerError,
			Error:  "roadmap_id is required",
		}, nil
	}

	err := s.storage.DeleteRoadmap(req.GetRoadmapId())
	if err != nil {
		return &project.RoadmapResponse{
			Status: http.StatusInternalServerError,
			Error:  "error delete roadmap",
		}, nil
	}

	return &project.RoadmapResponse{
		Status: http.StatusOK,
	}, nil
}
