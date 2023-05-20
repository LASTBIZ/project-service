package services

import (
	"context"
	"gorm.io/gorm"
	project1 "lastbiz/project-service/internal/project"
	"lastbiz/project-service/pkg/project"
	"net/http"
	"unsafe"
)

func (s Service) CreateProject(ctx context.Context, request *project.CreateProjectRequest) (*project.CreateProjectResponse, error) {
	if request.GetProject() == nil {
		return &project.CreateProjectResponse{
			Status: http.StatusConflict,
			Error:  "project not found",
		}, nil
	}

	pr := request.GetProject()

	if pr.GetImages() == nil {
		return &project.CreateProjectResponse{
			Status: http.StatusConflict,
			Error:  "images required",
		}, nil
	}

	if pr.GetBudgets() == nil {
		return &project.CreateProjectResponse{
			Status: http.StatusConflict,
			Error:  "budgets required",
		}, nil
	}

	if pr.GetRoadmaps() == nil {
		return &project.CreateProjectResponse{
			Status: http.StatusConflict,
			Error:  "roadmaps required",
		}, nil
	}

	roadmaps := make([]project1.Roadmap, 0)

	for _, roadMap := range pr.GetRoadmaps() {
		r := &project1.GRPCRoadmap{
			RoadMap: roadMap,
		}
		roadmaps = append(roadmaps, r.ToRoadmap())
	}

	createProject := project1.Project{
		Images: project1.Image{
			MainImageUrl:  pr.Images.GetMainImageUrl(),
			RoadMapImgUrl: pr.Images.GetRoadmapImgUrl(),
		},
		Dates: project1.Date{
			StartDate: pr.GetDates().GetStartDate().AsTime(),
			EndDate:   pr.GetDates().GetEndDate().AsTime(),
		},
		Budgets: project1.Budget{
			Need:    pr.GetBudgets().GetNeed(),
			Current: 0,
		},
		Location:    pr.GetLocation(),
		Name:        pr.GetName(),
		Description: pr.GetDescription(),
		Industry:    pr.GetIndustry(),
		Roadmaps:    roadmaps,
	}

	err := createProject.Validate()
	if err != nil {
		return &project.CreateProjectResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	err = s.storage.CreateProject(createProject)

	if err != nil {
		return &project.CreateProjectResponse{
			Status: http.StatusInternalServerError,
			Error:  "error create project",
		}, nil
	}

	return &project.CreateProjectResponse{
		Status: http.StatusOK,
	}, nil
}

func (s Service) DeleteProject(ctx context.Context, request *project.DeleteProjectRequest) (*project.DeleteProjectResponse, error) {
	if request.GetProjectId() == 0 {
		return &project.DeleteProjectResponse{
			Error:  "project_id not found",
			Status: http.StatusConflict,
		}, nil
	}

	err := s.storage.DeleteProject(uint32(request.GetProjectId()))
	if err != nil {
		return &project.DeleteProjectResponse{
			Error:  "error delete project",
			Status: http.StatusInternalServerError,
		}, nil
	}

	return &project.DeleteProjectResponse{
		Status: http.StatusOK,
	}, nil
}

func (s Service) UpdateProject(ctx context.Context, request *project.UpdateProjectRequest) (*project.UpdateProjectResponse, error) {
	if request.GetProject() == nil {
		return &project.UpdateProjectResponse{
			Status: http.StatusConflict,
			Error:  "project not found",
		}, nil
	}

	pr := request.GetProject()

	if pr.GetId() == 0 {
		return &project.UpdateProjectResponse{
			Status: http.StatusConflict,
			Error:  "id required",
		}, nil
	}

	if pr.GetImages() == nil {
		return &project.UpdateProjectResponse{
			Status: http.StatusConflict,
			Error:  "images required",
		}, nil
	}

	if pr.GetBudgets() == nil {
		return &project.UpdateProjectResponse{
			Status: http.StatusConflict,
			Error:  "budgets required",
		}, nil
	}

	roadmaps := make([]project1.Roadmap, 0)

	for _, roadMap := range pr.GetRoadmaps() {
		r := *(*project1.GRPCRoadmap)(unsafe.Pointer(roadMap))
		roadmaps = append(roadmaps, r.ToRoadmap())
	}

	createProject := project1.Project{
		ID: pr.GetId(),
		Images: project1.Image{
			MainImageUrl:  pr.Images.GetMainImageUrl(),
			RoadMapImgUrl: pr.Images.GetRoadmapImgUrl(),
		},
		Dates: project1.Date{
			StartDate: pr.GetDates().GetStartDate().AsTime(),
			EndDate:   pr.GetDates().GetEndDate().AsTime(),
		},
		Budgets: project1.Budget{
			Need:    pr.GetBudgets().GetNeed(),
			Current: 0,
		},
		Location:    pr.GetLocation(),
		Name:        pr.GetName(),
		Description: pr.GetDescription(),
		Industry:    pr.GetIndustry(),
		Roadmaps:    roadmaps,
	}

	err := createProject.Validate()
	if err != nil {
		return &project.UpdateProjectResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	err = s.storage.UpdateProject(uint32(createProject.ID), createProject)

	if err != nil {
		return &project.UpdateProjectResponse{
			Status: http.StatusInternalServerError,
			Error:  "error update project",
		}, nil
	}

	return &project.UpdateProjectResponse{
		Status: http.StatusOK,
	}, nil
}

func (s Service) GetProjectByID(ctx context.Context, request *project.GetProjectByIdRequest) (*project.GetProjectByIdResponse, error) {
	if request.GetProjectId() == 0 {
		return &project.GetProjectByIdResponse{
			Status: http.StatusConflict,
			Error:  "id required",
		}, nil
	}

	projectID := request.GetProjectId()
	pr, err := s.storage.GetProjectByID(uint32(projectID))
	if err != nil {
		return &project.GetProjectByIdResponse{
			Status: http.StatusInternalServerError,
			Error:  "project not found",
		}, nil
	}
	return &project.GetProjectByIdResponse{
		Status:  http.StatusOK,
		Project: pr.ToRPC(),
	}, nil
}

func (s Service) GetProjects(ctx context.Context, request *project.GetProjectsRequest) (*project.GetProjectsResponse, error) {
	page := request.GetPage()
	if page <= 0 {
		page = 1
	}

	pageSize := request.GetPageSize()
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	projects, err := s.storage.GetProjects(s.paninate(int(offset), int(pageSize)))
	if err != nil {
		return &project.GetProjectsResponse{
			Status: http.StatusInternalServerError,
			Error:  "projects not found",
		}, nil
	}

	prs := make([]*project.Project, 0)

	for _, p := range *projects {
		prs = append(prs, p.ToRPC())
	}

	return &project.GetProjectsResponse{
		Status:   http.StatusOK,
		Projects: prs,
	}, nil
}

func (s Service) paninate(offset, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pageSize)
	}
}
