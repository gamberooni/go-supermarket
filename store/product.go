package store

import (
	"github.com/gamberooni/go-supermarket/model"
	"gorm.io/gorm"
)

type ProductStore struct {
	db *gorm.DB
}

// return store instance to interact with db
func NewProductStore(db *gorm.DB) *ProductStore {
	return &ProductStore{
		db: db,
	}
}

func (ps *ProductStore) GetAllProducts() ([]model.Product, error) {
	products := []model.Product{}
	err := ps.db.Find(&products).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // return nil if error is record not found - not raised as error
		}

		return nil, err
	}
	return products, nil
}

func (ps *ProductStore) GetProductById(id int) (*model.Product, error) {
	product := model.Product{}
	err := ps.db.First(&product, id).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (ps *ProductStore) GetProductsByCategory(category string) ([]model.Product, error) {
	products := []model.Product{}
	err := ps.db.Where("category = ?", category).Find(&products).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // return nil if error is record not found - not raised as error
		}

		return nil, err
	}
	return products, nil
}

func (ps *ProductStore) AddProduct(p *model.Product) error {
	err := ps.db.Create(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (ps *ProductStore) DeleteProductById(id int) error {
	product := model.Product{}
	err := ps.db.First(&product, id).Error
	if err != nil {
		return err
	}
	ps.db.Delete(&product)
	return nil
}

func (ps *ProductStore) UpdateProductById(id int, p *model.Product) (*model.Product, error) {
	product := model.Product{}
	err := ps.db.First(&product, id).Error

	if err != nil {
		return nil, err
	}

	// update values
	ps.db.Model(&product).Updates(model.Product{Name: p.Name, Category: p.Category, StockQuantity: p.StockQuantity, Discount: p.Discount, NettPrice: p.NettPrice})

	return &product, nil
}
