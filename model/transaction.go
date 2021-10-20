package model

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID            int
	Timestamp     time.Time `json:"timestamp"`
	PaymentMethod string    `json:"payment_method" validate:"required,containsany=Cash,VISA,MasterCard"`
	Amount        float64   // calculated by calling PUT request on /api/transactions/:id/amount
	// CompanyID is implicitly used to create a FK relationship between Transaction and Customer tables
	// Customer struct needs to be included
	CustomerID int `json:"customer_id" validate:"required"`
	Customer   Customer
	Purchases  []Purchase `json:"purchases"`
}

type Purchase struct {
	gorm.Model
	Quantity      int     `json:"quantity" validate:"required"`
	PricePerUnit  float64 // automatically calculated using gorm hook
	TransactionID int     // automatically referenced back to parent table
	ProductID     int     `json:"product_id" validate:"required"`
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
