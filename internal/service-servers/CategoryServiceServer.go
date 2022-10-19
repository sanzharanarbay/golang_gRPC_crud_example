package service_servers

import (
	"context"
	"fmt"
	"github.com/sanzharanarbay/golang_gRPC_crud_eample/configs/database"
	"github.com/sanzharanarbay/golang_gRPC_crud_eample/internal/models"
	"github.com/sanzharanarbay/golang_gRPC_crud_eample/internal/repositories"
	serializers "github.com/sanzharanarbay/golang_gRPC_crud_eample/internal/serializers"
	"github.com/sanzharanarbay/golang_gRPC_crud_eample/internal/services"
)

type CategoryServiceServer  struct {

}

var db = database.ConnectDB()
var catRepo = repositories.NewCategoryRepository(db)
var catService = services.NewCategoryService(catRepo)


func (c *CategoryServiceServer) ReadCategory(ctx context.Context, req *serializers.ReadCategoryRequest) (*serializers.ReadCategoryResponse, error){
	category , _ := catService.GetSingleCategory(int(req.GetId()))
	response := &serializers.ReadCategoryResponse{
		Category: &serializers.Category{
			Id:       category.ID,
			Name: category.Name,
			Keyword: category.Keyword,
			CreatedAt: category.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: category.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	}
	return response, nil
}

func (c *CategoryServiceServer) ListCategories(ctx context.Context, req *serializers.ListCategoryRequest) (*serializers.ListCategoryResponse, error){
	categoryList , _ := catService.GetAllCategories()

	var categories []*serializers.Category

	for _, element := range *categoryList {
		response := &serializers.Category{
				Id:       element.ID,
				Name: element.Name,
				Keyword: element.Keyword,
				CreatedAt: element.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: element.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		categories = append(categories, response)
	}

	response := &serializers.ListCategoryResponse{
		Categories: categories,
	}

	return response, nil
}

func (c *CategoryServiceServer) CreateCategory(ctx context.Context, req *serializers.CreateCategoryReq) (*serializers.CreateCategoryResponse, error){
	categoryReq := models.Category{
		Name: req.GetName(),
		Keyword: req.GetKeyword(),
	}

	created, err := catService.InsertCategory(&categoryReq)

	if err != nil {
		panic(err)
	}

	if created == true {
		fmt.Println("Category Saved Successfully")
	}

	response := &serializers.CreateCategoryResponse{
	Status: 1,
	Message: "Category Saved Successfully",
	}
	return response, nil
}

func (c *CategoryServiceServer) UpdateCategory(ctx context.Context, req *serializers.UpdateCategoryReq) (*serializers.UpdateCategoryResponse, error){
	categoryReq := models.Category{
		Name: req.GetName(),
		Keyword: req.GetKeyword(),
	}

	updated, err := catService.UpdateCategory(&categoryReq, int(req.GetId()))

	if err != nil {
		panic(err)
	}

	if updated == true {
		fmt.Println("Category Updated Successfully")
	}

	response := &serializers.UpdateCategoryResponse{
		Status: 1,
		Message: "Category Updated Successfully",
	}
	return response, nil
}

func (c *CategoryServiceServer) DeleteCategory(ctx context.Context, req *serializers.DeleteCategoryRequest) (*serializers.DeleteCategoryResponse, error){

	deleted, err := catService.DeleteCategory(int(req.GetId()))

	if err != nil {
		panic(err)
	}

	if deleted == true {
		fmt.Println("Category Deleted Successfully")
	}

	response := &serializers.DeleteCategoryResponse{
		Status: 1,
		Message: "Category Deleted Successfully",
	}
	return response, nil
}



