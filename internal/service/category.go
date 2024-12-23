package service

import (
	"context"

	"github.com/iPingOi/grpccycle/internal/database"
	"github.com/iPingOi/grpccycle/internal/pb"
)

type CategoryService struct{
	pb.UnsafeCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(CategoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: CategoryDB,}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id: category.ID,
		Name: category.Name,
		Description: category.Description,
	}
	
	return categoryResponse, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var categoryResponses []*pb.Category
	for _, category := range categories {
		categoryResponses = append(categoryResponses, &pb.Category{
			Id: category.ID,
			Name: category.Name,
			Description: category.Description,
		})
	}

	return &pb.CategoryList{Categories: categoryResponses}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.CategoryGetRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Find(in.Id)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id: category.ID,
		Name: category.Name,
		Description: category.Description,
	}

	return categoryResponse, nil
}