package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"

	pb "project-service/api/roadmap"
)

type RoadmapService struct {
	pb.UnimplementedRoadmapServer
}

func NewRoadmapService() *RoadmapService {
	return &RoadmapService{}
}

func (s *RoadmapService) AddRoadMap(ctx context.Context, req *pb.CreateRoadmapRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
func (s *RoadmapService) UpdateRoadMap(ctx context.Context, req *pb.UpdateRoadmapRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
func (s *RoadmapService) RemoveRoadMap(ctx context.Context, req *pb.DeleteRoadmapRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
func (s *RoadmapService) ListRoadmap(ctx context.Context, req *pb.ListRoadmapRequest) (*pb.ListRoadmapReply, error) {
	return &pb.ListRoadmapReply{}, nil
}
