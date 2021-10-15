package handler

import (
	"github.com/gamberooni/go-supermarket/store"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	CustomerStore    store.CustomerStore
	ProductStore     store.ProductStore
	TransactionStore store.TransactionStore
	Validator        *validator.Validate
}

func NewHandler(cs store.CustomerStore, ps store.ProductStore, ts store.TransactionStore, v *validator.Validate) *Handler {
	return &Handler{
		CustomerStore:    cs,
		ProductStore:     ps,
		TransactionStore: ts,
		Validator:        v,
	}
}
