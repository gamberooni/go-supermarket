package model

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID            int
	Timestamp     time.Time
	PaymentMethod string
	Amount        float64
	// CompanyID is implicitly used to create a FK relationship between Transaction and Customer tables
	// Customer struct needs to be included
	CustomerID int
	Customer   Customer
	Purchases  []Purchase
}

type Purchase struct {
	gorm.Model
	Quantity      int `json:"quantity"`
	PricePerUnit  float64
	TransactionID int
	ProductID     int `json:"product_id"`
	Product       Product
}

func (p *Purchase) BeforeSave(tx *gorm.DB) error {
	// query product price from products table based on product id before saving a purchase record
	product := Product{}
	productError := tx.First(&product, p.ProductID).Error
	if productError != nil {
		return productError
	}
	p.PricePerUnit = product.NettPrice * (1 - product.Discount)

	return nil
}

func (t *Transaction) BeforeSave(tx *gorm.DB) error {
	// amount := 0.0
	for _, p := range t.Purchases {
		// calculate the amount for a transaction by looping through each purchase in that transaction
		// amount += p.PricePerUnit * float64(p.Quantity)
		// tx.Model(&t).Preload("Purchases").Update("amount", amount)

		// get product by id
		product := Product{}
		tx.First(&product, p.ProductID)
		// update stock quantity
		product.StockQuantity -= p.Quantity
		tx.Model(&product).Update("stock_quantity", product.StockQuantity)
	}

	return nil
}
