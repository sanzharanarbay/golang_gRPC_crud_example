package requests

import (
	"github.com/sanzharanarbay/golang_gRPC_crud_eample/internal/models"
	"net/url"
)

type ProductRequest struct {
	model *models.Product
}

func NewProductRequest(model *models.Product) *ProductRequest {
	return &ProductRequest{
		model: model,
	}
}

func(p *ProductRequest) ValidateProduct() url.Values{
	errs := url.Values{}

	if p.model.Name == "" {
		errs.Add("name", "The name field is required!")
	}

	// check the title field is between 3 to 120 chars
	if len(p.model.Name) < 3 || len(p.model.Name) > 255 {
		errs.Add("name", "The name field must be between 3-255 chars!")
	}

	if p.model.Description == "" {
		errs.Add("description", "The description field is required!")
	}

	// check the title field is between 3 to 120 chars
	if len(p.model.Description) < 3 || len(p.model.Description) > 255 {
		errs.Add("description", "The description field must be between 3-255 chars!")
	}

	if p.model.CategoryId == 0 {
		errs.Add("category_id", "The category_id field is required!")
	}

	if p.model.Price == 0 {
		errs.Add("price", "The price field is required!")
	}

	return errs
}

