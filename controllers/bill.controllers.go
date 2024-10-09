package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	// "vp_week11_echo/helpers"
	"vp_week11_echo/models"
)

func ReadBillByBillId(c echo.Context) error {
	id := c.FormValue("id")

	result, err := models.ReadBillByBillId(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func ReadAllBill(c echo.Context) error {
	id := c.FormValue("id")

	result, err := models.ReadAllBill(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func DeleteBill(c echo.Context) error {
	id := c.FormValue("id")

	result, err := models.DeleteBill(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func EditBill(c echo.Context) error {
	id := c.FormValue("bill_id")
	receipt_image := c.FormValue("receipt_image")
	restaurant_name := c.FormValue("restaurant_name")
	subtotal := c.FormValue("subtotal")
	total_discount := c.FormValue("total_discount")
	service_charge := c.FormValue("service_charge")
	tax := c.FormValue("tax")
	other := c.FormValue("other")
	grand_total := c.FormValue("grand_total")
	isSettled := c.FormValue("isSettled")
	date_created := c.FormValue("date_created")
	payment_id := c.FormValue("payment_id")
	user_id := c.FormValue("user_id")
	result, err := models.EditBill(id, receipt_image, restaurant_name, subtotal, total_discount, service_charge, tax, other, grand_total,isSettled,date_created,payment_id, user_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func CreateBill(c echo.Context) error {
	receipt_image := c.FormValue("receipt_image")
	restaurant_name := c.FormValue("restaurant_name")
	subtotal := c.FormValue("subtotal")
	total_discount := c.FormValue("total_discount")
	service_charge := c.FormValue("service_charge")
	tax := c.FormValue("tax")
	other := c.FormValue("other")
	grand_total := c.FormValue("grand_total")
	isSettled := c.FormValue("isSettled")
	date_created := c.FormValue("date_created")
	payment_id := c.FormValue("payment_id")
	user_id := c.FormValue("user_id")
	result, err := models.CreateBill(receipt_image, restaurant_name, subtotal, total_discount, service_charge, tax, other, grand_total,isSettled,date_created,payment_id, user_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}
