package store

import (
	"errors"

	"github.com/gamberooni/go-supermarket/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionStore struct {
	db *gorm.DB
}

// return store instance to interact with db
func NewTransactionStore(db *gorm.DB) *TransactionStore {
	return &TransactionStore{
		db: db,
	}
}

func (ts *TransactionStore) GetAllTransactions() ([]model.Transaction, error) {
	transactions := []model.Transaction{}
	err := ts.db.Preload("Customer").Find(&transactions).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // return nil if error is record not found - not raised as error
		}

		return nil, err
	}
	return transactions, nil
}

func (ts *TransactionStore) GetTransactionById(id int) (*model.Transaction, error) {
	transaction := model.Transaction{}
	err := ts.db.Preload("Purchases.Product").Preload(clause.Associations).First(&transaction, id).Error

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (ts *TransactionStore) AddTransaction(t *model.Transaction) error {
	// validate the purchase quantity
	for _, p := range t.Purchases {
		product := model.Product{}
		productError := ts.db.First(&product, p.ProductID).Error
		if productError != nil {
			return productError
		}
		if p.Quantity > product.StockQuantity {
			return errors.New("purchase quantity exceeds stock quantity")
		}
	}

	result := ts.db.Create(&t)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ts *TransactionStore) CalculateTransactionAmountById(id int) (float64, error) {
	t := model.Transaction{}
	err := ts.db.Preload("Purchases").First(&t, id).Error
	if err != nil {
		return 0.0, err
	}

	amount := 0.0
	for _, p := range t.Purchases {
		amount += p.PricePerUnit * float64(p.Quantity)
	}

	ts.db.Model(&t).Update("amount", amount)

	return t.Amount, nil
}
