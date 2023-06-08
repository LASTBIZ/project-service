package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Investor struct {
	ID        uint64
	Money     uint64
	FullName  string
	UserID    uint64
	ProjectID *uint64
}

type ProjectInvestor struct {
	ID         uint64
	ProjectID  uint64
	InvestorID uint64
	Investor   Investor
	Money      uint64
}

type InvestorRepo interface {
	CreateInvestor(ctx context.Context, investor *Investor) (*Investor, error)
	DeleteInvestor(ctx context.Context, id uint64) (bool, error)
	GetInvestorById(ctx context.Context, id uint64) (*Investor, error)
	ListInvestorByProjectId(ctx context.Context, projectId *uint64, pageNum, pageSize int) ([]*ProjectInvestor, int, error)
	AddMoneyInvestor(ctx context.Context, money int, id uint64) (bool, error)
	RemoveMoneyInvestor(ctx context.Context, money int, id uint64) (bool, error)
	SetMoneyInvestor(ctx context.Context, money int, id uint64) (bool, error)
}

type InvestorUseCase struct {
	repo InvestorRepo
	log  *log.Helper
}

func NewInvestorUseCase(repo InvestorRepo, logger log.Logger) *InvestorUseCase {
	return &InvestorUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (iu *InvestorUseCase) CreateInvestor(ctx context.Context, investor *Investor) (*Investor, error) {
	return iu.repo.CreateInvestor(ctx, investor)
}

func (iu *InvestorUseCase) DeleteInvestor(ctx context.Context, id uint64) (bool, error) {
	return iu.repo.DeleteInvestor(ctx, id)
}

func (iu *InvestorUseCase) GetInvestorById(ctx context.Context, id uint64) (*Investor, error) {
	return iu.repo.GetInvestorById(ctx, id)
}

func (iu *InvestorUseCase) ListInvestorByProjectId(ctx context.Context, projectId *uint64, pageNum, pageSize int) ([]*ProjectInvestor, int, error) {
	return iu.repo.ListInvestorByProjectId(ctx, projectId, pageNum, pageSize)
}

func (iu *InvestorUseCase) AddMoneyInvestor(ctx context.Context, money int, id uint64) (bool, error) {
	return iu.repo.AddMoneyInvestor(ctx, money, id)
}

func (iu *InvestorUseCase) RemoveMoneyInvestor(ctx context.Context, money int, id uint64) (bool, error) {
	return iu.repo.RemoveMoneyInvestor(ctx, money, id)
}

func (iu *InvestorUseCase) SetMoneyInvestor(ctx context.Context, money int, id uint64) (bool, error) {
	return iu.repo.SetMoneyInvestor(ctx, money, id)
}
