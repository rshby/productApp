package dto

type CreateProductRequest struct {
	Name        string  `json:"name,omitempty" validate:"required"`
	Price       float64 `json:"price,omitempty" validate:"required,gt=0"`
	Description string  `json:"description,omitempty"`
	Quantity    int     `json:"quantity,omitempty" validate:"required,gt=0"`
}
