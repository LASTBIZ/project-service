package services

import (
	"context"
	"gorm.io/gorm"
	"lastbiz/project-service/internal/project/models"
	"lastbiz/project-service/pkg/project/pb"
	"net/http"
	"unsafe"
)

func (s Service) CreateProject(ctx context.Context, request *pb.CreateProjectRequest) (*pb.CreateProjectResponse, error) {
	if request.GetProject() == nil {
		return &pb.CreateProjectResponse{
			Status: http.StatusConflict,
			Error:  "project not found",
		}, nil
	}

	pr := request.GetProject()

	if pr.GetImages() == nil {
		return &pb.CreateProjectResponse{
			Status: http.StatusConflict,
			Error:  "images required",
		}, nil
	}

	if pr.GetBudgets() == nil {
		return &pb.CreateProjectResponse{
			Status: http.StatusConflict,
			Error:  "budgets required",
		}, nil
	}

	if pr.GetRoadmaps() == nil {
		return &pb.CreateProjectResponse{
			Status: http.StatusConflict,
			Error:  "roadmaps required",
		}, nil
	}

	roadmaps := make([]models.Roadmap, 0)

	for _, roadMap := range pr.GetRoadmaps() {
		r := &models.GRPCRoadmap{
			RoadMap: roadMap,
		}
		roadmaps = append(roadmaps, r.ToRoadmap())
	}

	createProject := models.Project{
		Images: models.Image{
			MainImageUrl:  pr.Images.GetMainImageUrl(),
			RoadMapImgUrl: pr.Images.GetRoadmapImgUrl(),
		},
		Dates: models.Date{
			StartDate: pr.GetDates().GetStartDate().AsTime(),
			EndDate:   pr.GetDates().GetEndDate().AsTime(),
		},
		Budgets: models.Budget{
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
		return &pb.CreateProjectResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	err = s.storage.CreateProject(createProject)

	if err != nil {
		return &pb.CreateProjectResponse{
			Status: http.StatusInternalServerError,
			Error:  "error create project",
		}, nil
	}

	return &pb.CreateProjectResponse{
		Status: http.StatusOK,
	}, nil
}

func (s Service) DeleteProject(ctx context.Context, request *pb.DeleteProjectRequest) (*pb.DeleteProjectResponse, error) {
	if request.GetProjectId() == 0 {
		return &pb.DeleteProjectResponse{
			Error:  "project_id not found",
			Status: http.StatusConflict,
		}, nil
	}

	err := s.storage.DeleteProject(uint32(request.GetProjectId()))
	if err != nil {
		return &pb.DeleteProjectResponse{
			Error:  "error delete project",
			Status: http.StatusInternalServerError,
		}, nil
	}

	return &pb.DeleteProjectResponse{
		Status: http.StatusOK,
	}, nil
}

func (s Service) UpdateProject(ctx context.Context, request *pb.UpdateProjectRequest) (*pb.UpdateProjectResponse, error) {
	if request.GetProject() == nil {
		return &pb.UpdateProjectResponse{
			Status: http.StatusConflict,
			Error:  "project not found",
		}, nil
	}

	pr := request.GetProject()

	if pr.GetId() == 0 {
		return &pb.UpdateProjectResponse{
			Status: http.StatusConflict,
			Error:  "id required",
		}, nil
	}

	if pr.GetImages() == nil {
		return &pb.UpdateProjectResponse{
			Status: http.StatusConflict,
			Error:  "images required",
		}, nil
	}

	if pr.GetBudgets() == nil {
		return &pb.UpdateProjectResponse{
			Status: http.StatusConflict,
			Error:  "budgets required",
		}, nil
	}

	roadmaps := make([]models.Roadmap, 0)

	for _, roadMap := range pr.GetRoadmaps() {
		r := *(*models.GRPCRoadmap)(unsafe.Pointer(roadMap))
		roadmaps = append(roadmaps, r.ToRoadmap())
	}

	createProject := models.Project{
		ID: pr.GetId(),
		Images: models.Image{
			MainImageUrl:  pr.Images.GetMainImageUrl(),
			RoadMapImgUrl: pr.Images.GetRoadmapImgUrl(),
		},
		Dates: models.Date{
			StartDate: pr.GetDates().GetStartDate().AsTime(),
			EndDate:   pr.GetDates().GetEndDate().AsTime(),
		},
		Budgets: models.Budget{
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
		return &pb.UpdateProjectResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	err = s.storage.UpdateProject(uint32(createProject.ID), createProject)

	if err != nil {
		return &pb.UpdateProjectResponse{
			Status: http.StatusInternalServerError,
			Error:  "error update project",
		}, nil
	}

	return &pb.UpdateProjectResponse{
		Status: http.StatusOK,
	}, nil
}

func (s Service) GetProjectByID(ctx context.Context, request *pb.GetProjectByIdRequest) (*pb.GetProjectByIdResponse, error) {
	if request.GetProjectId() == 0 {
		return &pb.GetProjectByIdResponse{
			Status: http.StatusConflict,
			Error:  "id required",
		}, nil
	}

	projectID := request.GetProjectId()
	pr, err := s.storage.GetProjectByID(uint32(projectID))
	if err != nil {
		return &pb.GetProjectByIdResponse{
			Status: http.StatusInternalServerError,
			Error:  "project not found",
		}, nil
	}
	return &pb.GetProjectByIdResponse{
		Status:  http.StatusOK,
		Project: pr.ToRPC(),
	}, nil
}

func (s Service) GetProjects(ctx context.Context, request *pb.GetProjectsRequest) (*pb.GetProjectsResponse, error) {
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
		return &pb.GetProjectsResponse{
			Status: http.StatusInternalServerError,
			Error:  "projects not found",
		}, nil
	}

	prs := make([]*pb.Project, 0)

	for _, p := range *projects {
		prs = append(prs, p.ToRPC())
	}

	return &pb.GetProjectsResponse{
		Status:   http.StatusOK,
		Projects: prs,
	}, nil
}

func (s Service) paninate(offset, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pageSize)
	}
}
