package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/timestamppb"
	"project-service/internal/biz"

	pb "project-service/api/roadmap"
)

type RoadmapService struct {
	ur  *biz.RoadmapUseCase
	log *log.Helper
	pb.UnimplementedRoadmapServer
}

func NewRoadmapService(ur *biz.RoadmapUseCase, logger log.Logger) *RoadmapService {
	return &RoadmapService{
		ur:  ur,
		log: log.NewHelper(logger),
	}
}

func (s *RoadmapService) AddRoadMap(ctx context.Context, req *pb.CreateRoadmapRequest) (*empty.Empty, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	_, err = s.ur.CreateRoadmap(ctx, &biz.RoadMap{
		StartDate:   req.GetStartDate().AsTime(),
		EndDate:     req.GetEndDate().AsTime(),
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Job:         req.GetJob(),
		ProjectID:   req.GetProjectId(),
	})
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
func (s *RoadmapService) UpdateRoadMap(ctx context.Context, req *pb.UpdateRoadmapRequest) (*empty.Empty, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	_, err = s.ur.UpdateRoadmap(ctx, &biz.RoadMap{
		StartDate:   req.GetStartDate().AsTime(),
		EndDate:     req.GetEndDate().AsTime(),
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Job:         req.GetJob(),
		ID:          req.GetId(),
	})

	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
func (s *RoadmapService) RemoveRoadMap(ctx context.Context, req *pb.DeleteRoadmapRequest) (*empty.Empty, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	_, err = s.ur.DeleteRoadmap(ctx, req.GetId())
	return &empty.Empty{}, nil
}
func (s *RoadmapService) ListRoadmap(ctx context.Context, req *pb.ListRoadmapRequest) (*pb.ListRoadmapReply, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	roadmaps, total, err := s.ur.ListRoadmap(ctx, uint64(req.GetProjectId()), int(req.GetPageNum()), int(req.GetPageSize()))
	if err != nil {
		return nil, err
	}
	roadmapsResponse := make([]*pb.ListRoadmapReply_Roadmap, 0)
	if total == 0 {
		return &pb.ListRoadmapReply{
			Roadmaps: roadmapsResponse,
			Total:    0,
		}, nil
	}

	for _, rm := range roadmaps {
		roadmapsResponse = append(roadmapsResponse, &pb.ListRoadmapReply_Roadmap{
			Id:          rm.ID,
			StartDate:   timestamppb.New(rm.StartDate),
			EndDate:     timestamppb.New(rm.EndDate),
			Name:        rm.Name,
			Description: rm.Description,
			Job:         rm.Job,
			ProjectId:   rm.ProjectID,
		})
	}

	return &pb.ListRoadmapReply{
		Roadmaps: roadmapsResponse,
		Total:    uint32(total),
	}, nil
}
