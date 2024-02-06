package entity

import "time"

type Product struct {
	Id          int       `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Price       float64   `json:"price,omitempty"`
	Description string    `json:"description,omitempty"`
	Quantity    int       `json:"quantity,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
