package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID            int     // could be substituted by SKU
	Category      string  `json:"category"`
	Name          string  `json:"name"`
	StockQuantity int     `json:"stock_quantity"`
	Discount      float64 `json:"discount"`
	NettPrice     float64 `json:"nett_price"`
}
