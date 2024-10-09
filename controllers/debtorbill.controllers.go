package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	// "vp_week11_echo/helpers"
	"vp_week11_echo/models"
)

func ReadDebtorBill(c echo.Context) error {
	id := c.FormValue("id")

	result, err := models.ReadDebtorBill(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func DeleteDebtorBill(c echo.Context) error {
	id := c.FormValue("id")

	result, err := models.DeleteDebtorBill(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func EditDebtorBill(c echo.Context) error {
	id := c.FormValue("debtorbill_id")
	total := c.FormValue("total")
	tax := c.FormValue("tax")
	service_charge := c.FormValue("service_charge")
	discount := c.FormValue("discount")
	other := c.FormValue("other")
	isConfirmed := c.FormValue("isConfirmed")
	bill_id := c.FormValue("bill_id")
	result, err := models.EditDebtorBill(id, total, tax, service_charge, discount, other, isConfirmed, bill_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func CreateDebtorBill(c echo.Context) error {
	total := c.FormValue("total")
	tax := c.FormValue("tax")
	service_charge := c.FormValue("service_charge")
	discount := c.FormValue("discount")
	other := c.FormValue("other")
	isConfirmed := c.FormValue("isConfirmed")
	bill_id := c.FormValue("bill_id")
	result, err := models.CreateDebtorBill(total, tax, service_charge, discount, other, isConfirmed, bill_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}
