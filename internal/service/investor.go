package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes/empty"
	"project-service/internal/biz"

	pb "project-service/api/investor"
)

type InvestorService struct {
	ui  *biz.InvestorUseCase
	log *log.Helper
	pb.UnimplementedInvestorServer
}

func NewInvestorService(ui *biz.InvestorUseCase, logger log.Logger) *InvestorService {
	return &InvestorService{ui: ui, log: log.NewHelper(logger)}
}

func (s *InvestorService) CreateInvestor(ctx context.Context, req *pb.CreateInvestorRequest) (*empty.Empty, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	_, err = s.ui.CreateInvestor(ctx, &biz.Investor{
		FullName: req.FullName,
		UserID:   req.UserId,
	})

	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
func (s *InvestorService) DeleteInvestor(ctx context.Context, req *pb.DeleteInvestorRequest) (*empty.Empty, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	_, err = s.ui.DeleteInvestor(ctx, req.Id)

	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
func (s *InvestorService) AddMoneyInvestor(ctx context.Context, req *pb.AddMoneyInvestorRequest) (*empty.Empty, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	_, err = s.ui.AddMoneyInvestor(ctx, int(req.Money), req.UserId)

	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
func (s *InvestorService) RemoveMoneyInvestor(ctx context.Context, req *pb.RemoveMoneyInvestorRequest) (*empty.Empty, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	_, err = s.ui.RemoveMoneyInvestor(ctx, int(req.Money), req.UserId)

	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
func (s *InvestorService) SetMoneyInvestor(ctx context.Context, req *pb.SetMoneyInvestorRequest) (*empty.Empty, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	_, err = s.ui.SetMoneyInvestor(ctx, int(req.Money), req.UserId)

	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
func (s *InvestorService) ListInvestor(ctx context.Context, req *pb.ListInvestorRequest) (*pb.ListInvestorReply, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	li, total, err := s.ui.ListInvestorByProjectId(ctx, &req.ProjectId, int(req.PageNum), int(req.PageSize))
	if err != nil {
		return nil, err
	}
	investorsResponse := make([]*pb.ListInvestorReply_Investor, 0)
	if total == 0 {
		return &pb.ListInvestorReply{
			Investors: investorsResponse,
			Total:     0,
		}, nil
	}

	for _, inv := range li {
		investorsResponse = append(investorsResponse, &pb.ListInvestorReply_Investor{
			FullName: inv.Investor.FullName,
			Sum:      int64(inv.Money),
		})
	}

	return &pb.ListInvestorReply{}, nil
}
