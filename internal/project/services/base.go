package services

import (
	project1 "lastbiz/project-service/internal/project"
	"lastbiz/project-service/pkg/project/pb"
)

type Service struct {
	storage *project1.Storage
	pb.UnimplementedProjectServiceServer
}

func NewProjectService(storage *project1.Storage) pb.ProjectServiceServer {
	return &Service{
		storage: storage,
	}
}
