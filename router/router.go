package router

import (
	"net/http"

	"github.com/gamberooni/go-supermarket/handler"
	"github.com/gamberooni/go-supermarket/router/middleware"
	"github.com/gamberooni/go-supermarket/store"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *echo.Echo {
	// create new echo instance
	e := echo.New()

	// stores to interact with db via gorm
	cs := store.NewCustomerStore(db)
	ps := store.NewProductStore(db)
	ts := store.NewTransactionStore(db)

	// validator for struct fields - use a single instance of Validate, it caches struct info
	var Validator *validator.Validate
	Validator = validator.New()

	// handler
	h := handler.NewHandler(*cs, *ps, *ts, Validator)

	e.POST("/signup", h.Signup)
	e.POST("/login", h.Login)
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "OK",
		})
	})

	// create groups
	apiGroup := e.Group("/api")

	// initialize middlewares
	middleware.SetMainMiddlewares(e)
	middleware.SetApiMiddlewares(apiGroup)

	// customers endpoints
	apiGroup.GET("/customers", h.GetAllCustomers)
	apiGroup.GET("/customers/:id", h.GetCustomerById)
	apiGroup.PUT("/customers/:id", h.UpdateCustomerById)
	apiGroup.DELETE("/customers/:id", h.DeleteCustomerById)

	// products endpoints
	apiGroup.GET("/products", h.GetAllProducts)
	apiGroup.POST("/products", h.AddProduct)
	apiGroup.GET("/products/:id", h.GetProductById)
	apiGroup.GET("/products/category/:category", h.GetProductByCategory)
	apiGroup.PUT("/products/:id", h.UpdateProductById)
	apiGroup.DELETE("/products/:id", h.DeleteProductById)

	// transactions endpoints
	apiGroup.GET("/transactions", h.GetAllTransactions)
	apiGroup.POST("/transactions", h.AddTransaction)
	apiGroup.GET("/transactions/:id", h.GetTransactionById)
	apiGroup.PUT("/transactions/:id/amount", h.CalculateTransactionAmountById)

	return e
}
