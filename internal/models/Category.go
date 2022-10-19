package models

import "time"

type Category struct {
	ID          int64    `json:"id"`
	Name        string    `json:"name" validate:"required,min=3,max=255`
	Keyword     string   `json:"keyword" validate:"required,min=3,max=100`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
