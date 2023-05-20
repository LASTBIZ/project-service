package services

import (
	project1 "lastbiz/project-service/internal/project"
	"lastbiz/project-service/pkg/project"
)

type Service struct {
	storage *project1.Storage
	project.UnimplementedProjectServiceServer
}

func NewProjectService(storage *project1.Storage) project.ProjectServiceServer {
	return &Service{
		storage: storage,
	}
}
