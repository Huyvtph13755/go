package models

import "time"

type Product struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Price       float64        `json:"price"`
	Discount    int64          `json:"discount"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	OrderDetail []*OrderDetail `json:"order_details"`
}
