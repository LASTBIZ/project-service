package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes/empty"
	"project-service/internal/biz"

	pb "project-service/api/category"
)

type CategoryService struct {
	uc  *biz.CategoryUseCase
	log *log.Helper
	pb.UnimplementedCategoryServer
}

func NewCategoryService(uc *biz.CategoryUseCase, logger log.Logger) *CategoryService {
	return &CategoryService{uc: uc, log: log.NewHelper(logger)}
}

func (s *CategoryService) GetAllCategoryList(ctx context.Context, req *empty.Empty) (*pb.CategoryListReply, error) {
	jsonData, err := s.uc.GetAllCategories(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.CategoryListReply{
		JsonData: jsonData,
	}, nil
}

func (s *CategoryService) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*empty.Empty, error) {
	parentId := int32(0)
	if req.ParentId != nil {
		parentId = *req.ParentId
	}
	_, err := s.uc.CreateCategory(ctx, &biz.Category{
		Name:             req.Name,
		Level:            req.Level,
		ParentCategoryID: uint32(parentId),
	})

	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *CategoryService) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*empty.Empty, error) {
	err := s.uc.DeleteCategory(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *CategoryService) UpdateCategory(ctx context.Context, req *pb.CategoryInfoRequest) (*empty.Empty, error) {
	parentId := int32(0)
	if req.ParentId != nil {
		parentId = *req.ParentId
	}
	err := s.uc.UpdateCategory(ctx, &biz.Category{
		ID:               req.Id,
		Name:             req.Name,
		Level:            req.Level,
		ParentCategoryID: uint32(parentId),
	})
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
