package project

import (
	"context"
	"lastbiz/project-service/pkg/project"
)

type Service struct {
	storage Storage
	project.UnimplementedProjectServiceServer
}

func NewProjectService() project.ProjectServiceServer {
	return &Service{}
}

func (s Service) CreateProject(ctx context.Context, request *project.CreateProjectRequest) (*project.CreateProjectResponse, error) {
}

func (s Service) DeleteProject(ctx context.Context, request *project.DeleteProjectRequest) (*project.DeleteProjectResponse, error) {

}

func (s Service) UpdateProject(ctx context.Context, request *project.UpdateProjectRequest) (*project.UpdateProjectResponse, error) {

}

func (s Service) GetProjectByID(ctx context.Context, request *project.GetProjectByIdRequest) (*project.GetProjectByIdResponse, error) {

}

func (s Service) GetProjects(ctx context.Context, request *project.GetProjectsRequest) (*project.GetProjectsResponse, error) {

}
