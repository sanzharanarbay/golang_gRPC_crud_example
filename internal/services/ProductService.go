package services

import (
	"github.com/sanzharanarbay/golang_gRPC_crud_eample/internal/models"
	"github.com/sanzharanarbay/golang_gRPC_crud_eample/internal/repositories"
)

type ProductService struct {
	productRepository *repositories.ProductRepository
}

func NewProductService(productRepository *repositories.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (p *ProductService) IsProductAvailable(id int) (bool, error) {
	product, err := p.productRepository.GetProductById(id)
	if err == nil && product != nil {
		return true, nil
	}
	return false, err
}

func (p *ProductService) GetSingleProduct(ID int) (*models.Product, error) {
	product, err := p.productRepository.GetProductById(ID)
	return product, err
}

func (p *ProductService) GetAllProducts() (*[]models.Product, error) {
	productList, err := p.productRepository.GetAllProducts()
	return productList, err
}

func (p *ProductService) InsertProduct(product *models.Product) (bool, error) {
	state, err := p.productRepository.SaveProduct(product)
	return state, err
}

func (p *ProductService) UpdateProduct(product *models.Product, ID int) (bool, error) {
	found, err := p.IsProductAvailable(ID)
	if found == false {
		return false, err
	}
	state, err := p.productRepository.UpdateProduct(product, ID)
	return state, err
}

func (p *ProductService) DeleteProduct(id int) (bool, error) {
	var err error
	found, err := p.IsProductAvailable(id)
	if found == false {
		return false, err
	}
	state, err := p.productRepository.DeleteProduct(id)
	return state, err
}