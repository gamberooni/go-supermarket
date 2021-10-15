package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gamberooni/go-supermarket/model"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllProducts(c echo.Context) error {
	products, err := h.ProductStore.GetAllProducts()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, products)
}

func (h *Handler) GetProductById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	cat, getError := h.ProductStore.GetProductById(id)
	if getError != nil {
		return c.JSON(http.StatusInternalServerError, getError)
	}

	return c.JSON(http.StatusOK, cat)
}

func (h *Handler) GetProductByCategory(c echo.Context) error {
	category := c.Param("category")
	log.Print(category)
	products, err := h.ProductStore.GetProductsByCategory(category)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, products)
}

func (h *Handler) AddProduct(c echo.Context) error {
	product := model.Product{}
	err := c.Bind(&product)
	if err != nil {
		log.Printf("Failed processing AddProduct request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	createError := h.ProductStore.AddProduct(&product)
	if createError != nil {
		return c.JSON(http.StatusInternalServerError, createError)
	}
	log.Printf("Added new cat: %#v", product)

	return c.JSON(http.StatusOK, product)
}

func (h *Handler) DeleteProductById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.ProductStore.DeleteProductById((id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	log.Printf("Deleted cat with ID: %v", id)
	return c.JSON(http.StatusOK, err)
}

func (h *Handler) UpdateProductById(c echo.Context) error {
	product := model.Product{}
	err := c.Bind(&product)
	if err != nil {
		log.Printf("Failed processing UpdateProductById request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	result, updateError := h.ProductStore.UpdateProductById(id, &product)
	if updateError != nil {
		return c.JSON(http.StatusInternalServerError, updateError)
	}
	log.Printf("Updated customer with ID: %v", id)

	return c.JSON(http.StatusOK, result)
}
