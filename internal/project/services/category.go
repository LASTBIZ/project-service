package services

import (
	"context"
	"encoding/json"
	"google.golang.org/protobuf/types/known/emptypb"
	"lastbiz/project-service/internal/project/models"
	"lastbiz/project-service/pkg/project/pb"
)

func (s Service) GetAllCategoryList(ctx context.Context, _ *emptypb.Empty) (*pb.CategoryListResponse, error) {
	cate, err := s.storage.Category()
	if err != nil {
		return nil, err
	}
	jsonData, _ := json.Marshal(cate)
	res := &pb.CategoryListResponse{
		JsonData: string(jsonData),
	}
	return res, nil
}

func (s Service) GetSubCategory(ctx context.Context, req *pb.CategoryListRequest) (*pb.SubCategoryListResponse, error) {
	list, err := s.storage.SubCategoryList(req.Id)
	if err != nil {
		return nil, err
	}

	categoryListRes := pb.SubCategoryListResponse{}
	categoryListRes.Info = &pb.CategoryInfoResponse{
		Id:             int32(list.Category.ID),
		Name:           list.Category.Name,
		ParentCategory: list.Category.ParentCategory,
		Level:          list.Category.Level,
		IsTab:          list.Category.IsTab,
	}

	var subCategoryResponse []*pb.CategoryInfoResponse
	for _, subC := range list.SubCategory {
		subCategoryResponse = append(subCategoryResponse, &pb.CategoryInfoResponse{
			Id:             int32(subC.ID),
			Name:           subC.Name,
			ParentCategory: subC.ParentCategory,
			Level:          subC.Level,
			IsTab:          subC.IsTab,
		})
	}

	categoryListRes.SubCategory = subCategoryResponse
	return &categoryListRes, nil
}

func (s Service) CreateCategory(ctx context.Context, req *pb.CategoryInfoRequest) (*pb.CategoryInfoResponse, error) {
	result, err := s.storage.AddCategory(&models.CategoryInfo{
		Name:           req.Name,
		ParentCategory: req.ParentCategory,
		Level:          req.Level,
		IsTab:          req.IsTab,
	})

	if err != nil {
		return nil, err
	}

	return &pb.CategoryInfoResponse{
		Id:             int32(result.ID),
		Name:           result.Name,
		ParentCategory: result.ParentCategory,
		Level:          result.Level,
		IsTab:          result.IsTab,
	}, nil
}

func (s Service) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*emptypb.Empty, error) {
	err := s.storage.DeleteCategory(req.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s Service) UpdateCategory(ctx context.Context, req *pb.CategoryInfoRequest) (*emptypb.Empty, error) {
	err := s.storage.UpdateCategory(&models.CategoryInfo{
		ID:             uint32(req.Id),
		Name:           req.Name,
		ParentCategory: req.ParentCategory,
		Level:          req.Level,
		IsTab:          req.IsTab,
	})
	return &emptypb.Empty{}, err
}
