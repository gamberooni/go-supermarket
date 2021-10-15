package store

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gamberooni/go-supermarket/model"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

const Key = "secret"

type CustomerStore struct {
	db *gorm.DB
}

// return store instance to interact with db
func NewCustomerStore(db *gorm.DB) *CustomerStore {
	return &CustomerStore{
		db: db,
	}
}

func (cs *CustomerStore) GetAllCustomers() ([]model.Customer, error) {
	customers := []model.Customer{}
	result := cs.db.Find(&customers)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // return nil if error is record not found - not raised as error
		}

		return nil, result.Error
	}
	return customers, nil
}

func (cs *CustomerStore) GetCustomerById(id int) (*model.Customer, error) {
	customer := model.Customer{}
	err := cs.db.First(&customer, id).Error

	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (cs *CustomerStore) Signup(c *model.Customer) error {
	err := cs.db.Create(&c).Error

	if err != nil {
		return err
	}

	return nil
}

func createJwtToken(c *model.Customer) (string, error) {
	claims := model.JwtClaims{
		c.Name,
		jwt.StandardClaims{
			Id:        strconv.Itoa(c.ID),
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // token expires after 24 hours
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString([]byte(Key))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (cs *CustomerStore) Login(c *model.Customer) (*http.Cookie, error) {
	result := cs.db.Where("email = ? AND password = ?", c.Email, c.Password).Find(&c)
	if result.RowsAffected == 0 {
		return nil, errors.New("invalid credentials")
	}

	// create jwt token
	token, err := createJwtToken(c)
	if err != nil {
		return nil, err
	}

	// create cookie to store jwt token
	jwtCookie := &http.Cookie{}
	jwtCookie.Name = "JWTCookie"
	jwtCookie.Value = token
	jwtCookie.Expires = time.Now().Add(48 * time.Hour)
	// Http-only helps mitigate the risk of client side script accessing the protected cookie
	jwtCookie.HttpOnly = true

	c.JWTToken = token

	return jwtCookie, nil
}

func (cs *CustomerStore) DeleteCustomerById(id int) error {
	customer := model.Customer{}            // initialize an empty variable of Customer type
	err := cs.db.First(&customer, id).Error // get customer by id from db
	if err != nil {
		return err
	}
	cs.db.Delete(&customer) // delete the customer with the specified id
	return nil
}

func (cs *CustomerStore) UpdateCustomerById(id int, c *model.Customer) (*model.Customer, error) {
	customer := model.Customer{}
	err := cs.db.First(&customer, id).Error // get customer by id from db

	if err != nil {
		return nil, err
	}

	// update the customer with the specified id from db with the new values
	cs.db.Model(&customer).Updates(model.Customer{Name: c.Name, Email: c.Email, PhoneNumber: c.PhoneNumber})

	return &customer, nil
}
