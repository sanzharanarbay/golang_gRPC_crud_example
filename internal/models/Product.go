package models

import "time"

type Product struct {
	ID          int64    `json:"id"`
	Name        string    `json:"name" validate:"required,min=3,max=255`
	Description        string   `json:"description" validate:"required,min=3,max=500`
	CategoryId int64    `json:"category_id" validate:"required,numeric`
	Price float64    `json:"price" validate:"required,numeric`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
