package services

import (
	"github.com/sanzharanarbay/golang_gRPC_crud_eample/internal/models"
	"github.com/sanzharanarbay/golang_gRPC_crud_eample/internal/repositories"
)

type CategoryService struct {
	categoryRepository *repositories.CategoryRepository
}

func NewCategoryService(categoryRepository *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}

func (c *CategoryService) IsCategoryAvailable(id int) (bool, error) {
	category, err := c.categoryRepository.GetCategoryById(id)
	if err == nil && category != nil {
		return true, nil
	}
	return false, err
}

func (c *CategoryService) GetSingleCategory(ID int) (*models.Category, error) {
	category, err := c.categoryRepository.GetCategoryById(ID)
	return category, err
}

func (c *CategoryService) GetAllCategories() (*[]models.Category, error) {
	catList, err := c.categoryRepository.GetAllCategories()
	return catList, err
}

func (c *CategoryService) InsertCategory(category *models.Category) (bool, error) {
	state, err := c.categoryRepository.SaveCategory(category)
	return state, err
}

func (c *CategoryService) UpdateCategory(category *models.Category, ID int) (bool, error) {
	found, err := c.IsCategoryAvailable(ID)
	if found == false {
		return false, err
	}
	state, err := c.categoryRepository.UpdateCategory(category, ID)
	return state, err
}

func (c *CategoryService) DeleteCategory(id int) (bool, error) {
	var err error
	found, err := c.IsCategoryAvailable(id)
	if found == false {
		return false, err
	}
	state, err := c.categoryRepository.DeleteCategory(id)
	return state, err
}


