package models

import (
	"time"

	"github.com/Huyvtph13755/go/config"
)

type Order struct {
	ID        int64     `json:"id"`
	Total     float64   `json:"total"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      *User     `json:"-"`
}

func (o *Order) CalcTotal() {
	var order_details []*OrderDetail
	config.Database.Find(&order_details, "order_id = ?", o.ID)
	var total float64
	for _, od := range order_details {
		var product *Product
		config.Database.First(&product, "id = ?", od.ProductID)
		total += product.Price + float64(od.Quantity)
	}
	o.Total = total
	config.Database.Save(&o)
}
