package dto

type ProductDetailResponse struct {
	Id          int     `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Description string  `json:"description,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
	CreatedAt   string  `json:"created_at,omitempty"`
	UpdatedAt   string  `json:"updated_at,omitempty"`
}
