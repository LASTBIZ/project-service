package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/timestamppb"
	"project-service/internal/biz"

	pb "project-service/api/project"
)

type ProjectService struct {
	up  *biz.ProjectUseCase
	log *log.Helper
	pb.UnimplementedProjectServer
}

func NewProjectService(up *biz.ProjectUseCase, logger log.Logger) *ProjectService {
	return &ProjectService{up: up, log: log.NewHelper(logger)}
}

func (s *ProjectService) CreateProject(ctx context.Context, req *pb.CreateProjectRequest) (*empty.Empty, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	_, err = s.up.Create(ctx, &biz.Project{
		RoadMapImgURL: req.RoadmapImgUrl,
		MainImageUrl:  req.MainImageUrl,
		StartDate:     req.StartDate.AsTime(),
		EndDate:       req.EndDate.AsTime(),
		Location:      req.Location,
		Name:          req.Name,
		Description:   req.Description,
		CategoryID:    req.CategoryId,
		NeedBudget:    req.NeedBudget,
	})
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
func (s *ProjectService) UpdateProject(ctx context.Context, req *pb.UpdateProjectRequest) (*empty.Empty, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	_, err = s.up.Update(ctx, &biz.Project{
		RoadMapImgURL: req.RoadmapImgUrl,
		MainImageUrl:  req.MainImageUrl,
		StartDate:     req.StartDate.AsTime(),
		EndDate:       req.EndDate.AsTime(),
		Location:      req.Location,
		Name:          req.Name,
		Description:   req.Description,
		CategoryID:    req.CategoryId,
		NeedBudget:    req.NeedBudget,
		ID:            req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
func (s *ProjectService) DeleteProject(ctx context.Context, req *pb.DeleteProjectRequest) (*empty.Empty, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	_, err = s.up.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
func (s *ProjectService) GetProject(ctx context.Context, req *pb.GetProjectRequest) (*pb.GetProjectReply, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	pr, err := s.up.GetProjectById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetProjectReply{
		RoadmapImgUrl: pr.RoadMapImgURL,
		MainImageUrl:  pr.MainImageUrl,
		StartDate:     timestamppb.New(pr.StartDate),
		EndDate:       timestamppb.New(pr.EndDate),
		Location:      pr.Location,
		Name:          pr.Name,
		Description:   pr.Description,
		CategoryId:    pr.CategoryID,
		NeedBudget:    pr.NeedBudget,
		Id:            pr.ID,
		CurrentBudget: pr.CurrentBudget,
	}, nil
}
func (s *ProjectService) GetProjectByCategoryID(ctx context.Context, req *pb.ListProjectRequest) (*pb.ListProjectReply, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	prs, total, err := s.up.GetProjectByCategoryId(ctx, req.Keywords, req.CategoryId, int(req.PageNum), int(req.PageSize))
	projectsResponse := make([]*pb.ProjectInfoResponse, 0)
	if total == 0 {
		return &pb.ListProjectReply{
			Projects: projectsResponse,
			Total:    0,
		}, nil
	}

	for _, pr := range prs {
		projectsResponse = append(projectsResponse,
			&pb.ProjectInfoResponse{
				MainImageUrl: pr.MainImageUrl,
				StartDate:    timestamppb.New(pr.StartDate),
				EndDate:      timestamppb.New(pr.EndDate),
				Location:     pr.Location,
				Name:         pr.Name,
				CategoryId:   pr.CategoryID,
				Id:           pr.ID,
			})
	}

	return &pb.ListProjectReply{
		Projects: projectsResponse,
		Total:    uint32(total),
	}, nil
}

func (s *ProjectService) InvestProject(ctx context.Context, req *pb.InvestProjectRequest) (*empty.Empty, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	err = s.up.InvestProject(ctx, req.Id, req.InvestorId, int(req.Money))
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
