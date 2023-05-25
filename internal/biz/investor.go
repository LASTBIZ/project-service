package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Investor struct {
	ID        uint64
	Money     uint64
	FullName  string
	UserID    uint32
	ProjectID *uint64
}

type InvestorRepo interface {
	CreateInvestor(ctx context.Context, investor *Investor) (*Investor, error)
	DeleteInvestor(ctx context.Context, id uint64) (bool, error)
	GetInvestorById(ctx context.Context, id uint64) (*Investor, error)
	ListInvestorByProjectId(ctx context.Context, projectId uint64, pageNum, pageSize int) ([]*Investor, error)
	AddMoneyInvestor(ctx context.Context, money int, id uint64) (bool, error)
	RemoveMoneyInvestor(ctx context.Context, money int, id uint64) (bool, error)
	SetMoneyInvestor(ctx context.Context, money int, id uint64) (bool, error)
	InvestProject(ctx context.Context, money int, id, projectId uint64) (bool, error)
}

type InvestorUseCase struct {
	repo InvestorRepo
	log  *log.Helper
}

func NewInvestorUseCase(repo InvestorRepo, logger log.Logger) *InvestorUseCase {
	return &InvestorUseCase{repo: repo, log: log.NewHelper(logger)}
}
