package handler

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gamberooni/go-supermarket/model"
	"github.com/gamberooni/go-supermarket/util"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllCustomers(c echo.Context) error {
	customers, err := h.CustomerStore.GetAllCustomers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, customers)
}

func (h *Handler) GetCustomerById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	cat, getError := h.CustomerStore.GetCustomerById(id)
	if getError != nil {
		return c.JSON(http.StatusInternalServerError, getError)
	}

	return c.JSON(http.StatusOK, cat)
}

func (h *Handler) Signup(c echo.Context) error {
	// bind
	customer := model.Customer{}
	err := c.Bind(&customer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	// manually validate birthday regex
	birthdayRe, _ := regexp.Compile("[0-9]{4}-[0-9]{2}-[0-9]{2}")
	if !birthdayRe.MatchString(customer.Birthday) {
		return echo.NewHTTPError(http.StatusInternalServerError, "birthday does not match regexp pattern 'yyyy-MM-dd'")
	}

	// manually validate phone number regex
	phoneNumberRe, _ := regexp.Compile("01[0-9]{1}-[0-9]{7,8}")
	if !phoneNumberRe.MatchString(customer.PhoneNumber) {
		return echo.NewHTTPError(http.StatusInternalServerError, "phone number does not match regexp pattern '01x-xxxxxxx(x)'")
	}

	// validate using validator package
	validateError := h.Validator.Struct(customer)
	if validateError != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := validateError.(*validator.InvalidValidationError); ok {
			fmt.Println(validateError)

			return validateError
		}

		for _, err := range validateError.(validator.ValidationErrors) {

			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}

		// from here you can create your own error messages in whatever language you wish

		return validateError

	}

	// hash the customer's signup password
	hashedPassword, hashError := util.HashPassword(customer.Password)
	if hashError != nil {
		return c.JSON(http.StatusInternalServerError, hashError)
	}
	customer.Password = hashedPassword

	createError := h.CustomerStore.Signup(&customer)
	if createError != nil {
		return c.JSON(http.StatusInternalServerError, createError)
	}

	return c.JSON(http.StatusOK, customer)
}

func (h *Handler) Login(c echo.Context) error {
	// bind
	customer := model.Customer{}
	err := c.Bind(&customer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// validate
	validateError := h.Validator.StructPartial(customer, "Email", "Password")
	if validateError != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := validateError.(*validator.InvalidValidationError); ok {
			fmt.Println(validateError)
			return validateError
		}

		for _, err := range validateError.(validator.ValidationErrors) {
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}

		// from here you can create your own error messages in whatever language you wish
		return validateError
	}

	jwtCookie, loginError := h.CustomerStore.Login(&customer)
	if loginError != nil {
		return c.JSON(http.StatusUnauthorized, loginError)
	}

	// set jwtcookie
	c.SetCookie(jwtCookie)

	// remove password in response
	customer.Password = ""
	return c.JSON(http.StatusOK, customer)
}

func (h *Handler) DeleteCustomerById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))            // convert id from string to int
	err := h.CustomerStore.DeleteCustomerById((id)) // invoke the underlying customerstore method to delete customer by id
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, err)

}

func (h *Handler) UpdateCustomerById(c echo.Context) error {
	customer := model.Customer{}
	err := c.Bind(&customer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	result, updateError := h.CustomerStore.UpdateCustomerById(id, &customer)
	if updateError != nil {
		return c.JSON(http.StatusInternalServerError, updateError)
	}

	return c.JSON(http.StatusOK, result)
}
