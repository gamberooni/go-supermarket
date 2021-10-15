package handler

import (
	"net/http"
	"strconv"

	"github.com/gamberooni/go-supermarket/model"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllTransactions(c echo.Context) error {
	transactions, err := h.TransactionStore.GetAllTransactions()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, transactions)
}

func (h *Handler) GetTransactionById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction, getError := h.TransactionStore.GetTransactionById(id)
	if getError != nil {
		return c.JSON(http.StatusInternalServerError, getError)
	}

	return c.JSON(http.StatusOK, transaction)
}

func (h *Handler) AddTransaction(c echo.Context) error {
	transaction := model.Transaction{}
	err := c.Bind(&transaction)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	createError := h.TransactionStore.AddTransaction(&transaction)
	if createError != nil {
		return c.JSON(http.StatusInternalServerError, createError)
	}

	return c.JSON(http.StatusOK, transaction)
}

func (h *Handler) CalculateTransactionAmountById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	amount, err := h.TransactionStore.CalculateTransactionAmountById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, map[string]float64{"amount": amount})
}
