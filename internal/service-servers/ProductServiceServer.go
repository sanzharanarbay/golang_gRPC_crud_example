package service_servers

import (
	"context"
	"fmt"
	"github.com/sanzharanarbay/golang_gRPC_crud_eample/internal/models"
	"github.com/sanzharanarbay/golang_gRPC_crud_eample/internal/repositories"
	"github.com/sanzharanarbay/golang_gRPC_crud_eample/internal/requests"
	serializers "github.com/sanzharanarbay/golang_gRPC_crud_eample/internal/serializers"
	"github.com/sanzharanarbay/golang_gRPC_crud_eample/internal/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductServiceServer  struct {

}

var productRepo = repositories.NewProductRepository(db)
var productService = services.NewProductService(productRepo)


func (p *ProductServiceServer) ReadProduct(ctx context.Context, req *serializers.ReadProductRequest) (*serializers.ReadProductResponse, error){
	product , _ := productService.GetSingleProduct(int(req.GetId()))

	if product == nil{
		err := status.Error(codes.NotFound, "The product was not found")
		return nil, err
	}

	response := &serializers.ReadProductResponse{
		Product: &serializers.Product{
			Id:       product.ID,
			Name: product.Name,
			Description: product.Description,
			CategoryId: product.CategoryId,
			Price: product.Price,
			CreatedAt: product.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: product.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	}
	return response, nil
}

func (p *ProductServiceServer) ListProducts(ctx context.Context, req *serializers.ListProductRequest) (*serializers.ListProductResponse, error){
	productList , _ := productService.GetAllProducts()

	var products []*serializers.Product

	for _, element := range *productList {
		response := &serializers.Product{
			Id:       element.ID,
			Name: element.Name,
			Description: element.Description,
			CategoryId: element.CategoryId,
			Price: element.Price,
			CreatedAt: element.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: element.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		products = append(products, response)
	}

	response := &serializers.ListProductResponse{
		Products: products,
	}

	return response, nil
}

func (p *ProductServiceServer) CreateProduct(ctx context.Context, req *serializers.CreateProductReq) (*serializers.CreateProductResponse, error){
	productReq := models.Product{
		Name: req.GetName(),
		Description: req.GetDescription(),
		CategoryId: req.GetCategoryId(),
		Price: req.GetPrice(),
	}

	validErrs := requests.NewProductRequest(&productReq).ValidateProduct()

	if len(validErrs) > 0{
		return nil, status.Error(codes.InvalidArgument, validErrs.Encode())
	}

	category , _ := catService.GetSingleCategory(int(req.GetCategoryId()))

	if category == nil{
		err := status.Error(codes.NotFound, "The category by ID was not found")
		return nil, err
	}

	created, err := productService.InsertProduct(&productReq)

	if err != nil {
		panic(err)
	}

	if created == true {
		fmt.Println("Product Saved Successfully")
	}

	response := &serializers.CreateProductResponse{
		Status: 1,
		Message: "Product Saved Successfully",
	}
	return response, nil
}

func (p *ProductServiceServer) UpdateProduct(ctx context.Context, req *serializers.UpdateProductReq) (*serializers.UpdateProductResponse, error){
	productReq := models.Product{
		Name: req.GetName(),
		Description: req.GetDescription(),
		CategoryId: req.GetCategoryId(),
		Price: req.GetPrice(),
	}

	product , _ := productService.GetSingleProduct(int(req.GetId()))

	if product == nil{
		err := status.Error(codes.NotFound, "The product was not found")
		return nil, err
	}

	validErrs := requests.NewProductRequest(&productReq).ValidateProduct()

	if len(validErrs) > 0{
		return nil, status.Error(codes.InvalidArgument, validErrs.Encode())
	}

	category , _ := catService.GetSingleCategory(int(req.GetCategoryId()))

	if category == nil{
		err := status.Error(codes.NotFound, "The category by ID was not found")
		return nil, err
	}


	updated, err := productService.UpdateProduct(&productReq, int(req.GetId()))

	if err != nil {
		panic(err)
	}

	if updated == true {
		fmt.Println("Product Updated Successfully")
	}

	response := &serializers.UpdateProductResponse{
		Status: 1,
		Message: "Product Updated Successfully",
	}
	return response, nil
}

func (p *ProductServiceServer) DeleteProduct(ctx context.Context, req *serializers.DeleteProductRequest) (*serializers.DeleteProductResponse, error){

	product , _ := productService.GetSingleProduct(int(req.GetId()))

	if product == nil{
		err := status.Error(codes.NotFound, "The product was not found")
		return nil, err
	}

	deleted, err := productService.DeleteProduct(int(req.GetId()))

	if err != nil {
		panic(err)
	}

	if deleted == true {
		fmt.Println("Product Deleted Successfully")
	}

	response := &serializers.DeleteProductResponse{
		Status: 1,
		Message: "Product Deleted Successfully",
	}
	return response, nil
}