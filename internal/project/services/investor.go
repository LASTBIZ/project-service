package services

import (
	"context"
	project1 "lastbiz/project-service/internal/project"
	"lastbiz/project-service/pkg/project"
	"net/http"
)

func (s Service) AddInvestor(ctx context.Context, req *project.AddInvestorRequest) (*project.InvestorResponse, error) {
	if req.GetInvestor() == nil {
		return &project.InvestorResponse{
			Status: http.StatusConflict,
			Error:  "investor not found",
		}, nil
	}

	investor := req.GetInvestor()

	createInvestor := project1.Investor{
		Money:    0,
		FullName: investor.GetFullName(),
		UserID:   investor.GetUserId(),
	}

	err := s.storage.AddInvestor(createInvestor)

	if err != nil {
		return &project.InvestorResponse{
			Status: http.StatusInternalServerError,
			Error:  "error create investor",
		}, nil
	}

	return &project.InvestorResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s Service) RemoveInvestor(ctx context.Context, req *project.RemoveInvestorRequest) (*project.InvestorResponse, error) {
	if req.GetInvestorId() == 0 {
		return &project.InvestorResponse{
			Status: http.StatusConflict,
			Error:  "investor_id not found",
		}, nil
	}

	err := s.storage.RemoveInvestor(req.GetInvestorId())

	if err != nil {
		return &project.InvestorResponse{
			Status: http.StatusInternalServerError,
			Error:  "error remove investor",
		}, nil
	}

	return &project.InvestorResponse{
		Status: http.StatusOK,
	}, nil
}

func (s Service) AddMoney(ctx context.Context, req *project.MoneyRequest) (*project.InvestorResponse, error) {
	if req.GetMoney() <= 0 {
		return &project.InvestorResponse{
			Status: http.StatusConflict,
			Error:  "money not found",
		}, nil
	}

	if req.GetInvestorId() == 0 {
		return &project.InvestorResponse{
			Status: http.StatusConflict,
			Error:  "investor_id not found",
		}, nil
	}

	err := s.storage.AddMoney(req.GetInvestorId(), req.GetMoney())

	if err != nil {
		return &project.InvestorResponse{
			Status: http.StatusInternalServerError,
			Error:  "error add money",
		}, nil
	}

	return &project.InvestorResponse{
		Status: http.StatusOK,
	}, nil
}

func (s Service) RemoveMoney(ctx context.Context, req *project.MoneyRequest) (*project.InvestorResponse, error) {
	if req.GetMoney() <= 0 {
		return &project.InvestorResponse{
			Status: http.StatusConflict,
			Error:  "money not found",
		}, nil
	}

	if req.GetInvestorId() == 0 {
		return &project.InvestorResponse{
			Status: http.StatusConflict,
			Error:  "investor_id not found",
		}, nil
	}

	err := s.storage.RemoveMoney(req.GetInvestorId(), req.GetMoney())

	if err != nil {
		return &project.InvestorResponse{
			Status: http.StatusInternalServerError,
			Error:  "error remove money",
		}, nil
	}

	return &project.InvestorResponse{
		Status: http.StatusOK,
	}, nil
}

func (s Service) SetMoney(ctx context.Context, req *project.MoneyRequest) (*project.InvestorResponse, error) {
	if req.GetMoney() <= 0 {
		return &project.InvestorResponse{
			Status: http.StatusConflict,
			Error:  "money not found",
		}, nil
	}

	if req.GetInvestorId() == 0 {
		return &project.InvestorResponse{
			Status: http.StatusConflict,
			Error:  "investor_id not found",
		}, nil
	}

	err := s.storage.SetMoney(req.GetInvestorId(), req.GetMoney())

	if err != nil {
		return &project.InvestorResponse{
			Status: http.StatusInternalServerError,
			Error:  "error set money",
		}, nil
	}

	return &project.InvestorResponse{
		Status: http.StatusOK,
	}, nil
}

func (s Service) InvestProject(ctx context.Context, req *project.InvestProjectRequest) (*project.InvestorResponse, error) {
	if req.GetMoney() <= 0 {
		return &project.InvestorResponse{
			Status: http.StatusConflict,
			Error:  "money not found",
		}, nil
	}

	if req.GetInvestorId() == 0 {
		return &project.InvestorResponse{
			Status: http.StatusConflict,
			Error:  "investor_id not found",
		}, nil
	}

	if req.GetProjectId() == 0 {
		return &project.InvestorResponse{
			Status: http.StatusConflict,
			Error:  "project_id not found",
		}, nil
	}

	err := s.storage.RemoveMoney(req.GetInvestorId(), req.GetMoney())
	if err != nil {
		return &project.InvestorResponse{
			Status: http.StatusInternalServerError,
			Error:  "error invest project",
		}, nil
	}

	err = s.storage.InvestProject(req.GetInvestorId(), req.GetMoney())

	if err != nil {
		return &project.InvestorResponse{
			Status: http.StatusInternalServerError,
			Error:  "error invest project",
		}, nil
	}

	return &project.InvestorResponse{
		Status: http.StatusOK,
	}, nil
}
