// to generate dummy data for the db

package main

import (
	"time"

	"github.com/gamberooni/go-supermarket/model"
	"github.com/gamberooni/go-supermarket/util"
)

func main() {
	// initialize db
	db := util.InitDB()

	hashedPassword, _ := util.HashPassword("p4ssw0rd")

	customers := []model.Customer{
		{Name: "Tom", Email: "tom@jerry.com", Birthday: "1996-03-14", PhoneNumber: "012-3456789", Password: string(hashedPassword)},
		{Name: "Jerry", Email: "jerry@tom.com", Birthday: "1999-01-30", PhoneNumber: "012-9876543"},
		{Name: "Donald", Email: "donald@duck.com", Birthday: "1976-06-12"},
		{Name: "Mickey", PhoneNumber: "019-8765432", Birthday: "1987-07-17"},
	}

	products := []model.Product{
		{Category: "vegetables", Name: "Onion", StockQuantity: 30, Discount: 0, NettPrice: 0.95},
		{Category: "vegetables", Name: "Bak Choy", StockQuantity: 13, Discount: 0.05, NettPrice: 1.95},
		{Category: "vegetables", Name: "Asparagus", StockQuantity: 7, Discount: 0, NettPrice: 3.00},
		{Category: "drinks", Name: "Farmfresh Milk", StockQuantity: 5, Discount: 0.1, NettPrice: 4.50},
		{Category: "drinks", Name: "Heineken Beer", StockQuantity: 15, Discount: 0, NettPrice: 6.00},
		{Category: "snacks", Name: "Lays Chips", StockQuantity: 10, Discount: 0, NettPrice: 2.00},
		{Category: "dried-food", Name: "Raisins", StockQuantity: 11, Discount: 0.25, NettPrice: 8.00},
		{Category: "dried-food", Name: "Anchovies", StockQuantity: 7, Discount: 0.05, NettPrice: 5.00},
	}

	transactions := []model.Transaction{
		{
			Timestamp:     time.Date(2021, 10, 12, 20, 34, 58, 651387237, time.UTC),
			PaymentMethod: "VISA",
			CustomerID:    1,
			Purchases: []model.Purchase{
				{
					Quantity:  1,
					ProductID: 1,
				},
				{
					Quantity:  2,
					ProductID: 3,
				},
			},
		},
		{
			Timestamp:     time.Date(2021, 10, 13, 12, 55, 21, 551926242, time.UTC),
			PaymentMethod: "MasterCard",
			CustomerID:    1,
			Purchases: []model.Purchase{
				{
					Quantity:  2,
					ProductID: 3,
				},
			},
		},
		{
			Timestamp:     time.Date(2021, 10, 13, 15, 12, 26, 431236751, time.UTC),
			PaymentMethod: "Cash",
			CustomerID:    2,
			Purchases: []model.Purchase{
				{
					Quantity:  1,
					ProductID: 6,
				},
			},
		},
		{
			Timestamp:     time.Date(2021, 10, 15, 18, 21, 43, 836712921, time.UTC),
			PaymentMethod: "VISA",
			CustomerID:    3,
			Purchases: []model.Purchase{
				{
					Quantity:  2,
					ProductID: 5,
				},
				{
					Quantity:  3,
					ProductID: 3,
				},
			},
		},
		{
			Timestamp:     time.Date(2021, 10, 15, 18, 21, 43, 836712921, time.UTC),
			PaymentMethod: "Cash",
			CustomerID:    4,
			Purchases: []model.Purchase{
				{
					Quantity:  10,
					ProductID: 2,
				},
				{
					Quantity:  10,
					ProductID: 7,
				},
			},
		},
	}

	for _, c := range customers {
		db.Create(&c)
	}

	for _, p := range products {
		db.Create(&p)
	}

	for _, t := range transactions {
		db.Create(&t)
	}
}
