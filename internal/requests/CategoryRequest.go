package requests

import (
	"github.com/sanzharanarbay/golang_gRPC_crud_eample/internal/models"
	"net/url"
)

type CategoryRequest struct {
	model *models.Category
}

func NewCategoryRequest(model *models.Category) *CategoryRequest {
	return &CategoryRequest{
		model: model,
	}
}

func(c *CategoryRequest) ValidateCategory() url.Values{
	errs := url.Values{}

	if c.model.Name == "" {
		errs.Add("name", "The name field is required!")
	}

	// check the title field is between 3 to 120 chars
	if len(c.model.Name) < 3 || len(c.model.Name) > 255 {
		errs.Add("name", "The name field must be between 3-255 chars!")
	}

	if c.model.Keyword == "" {
		errs.Add("keyword", "The keyword field is required!")
	}

	// check the title field is between 3 to 120 chars
	if len(c.model.Keyword) < 3 || len(c.model.Keyword) > 255 {
		errs.Add("content", "The keyword field must be between 3-255 chars!")
	}

	return errs
}
