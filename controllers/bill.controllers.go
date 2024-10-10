package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	// "vp_week11_echo/helpers"
	"io"
	"mime/multipart"
	"os"
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
	result, err := models.EditBill(id, receipt_image, restaurant_name, subtotal, total_discount, service_charge, tax, other, grand_total, isSettled, date_created, payment_id, user_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func CreateBill(c echo.Context) error {
	// receipt_image := c.FormValue("receipt_image")
	receipt_image, err := c.FormFile("receipt_image")
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

	if err != nil {
		return c.JSON(http.StatusBadRequest, &models.Response{
			Message: "Invalid data! The data type must be images!",
		})
	}

	// Define the file path to save the uploaded image.
	pathImage := "/Users/marshalikorawung/SplitBill_API/images/" + receipt_image.Filename

	// Save the uploaded file to the specified path.
	if err := saveUploadedFile(receipt_image, pathImage); err != nil {
		return c.JSON(http.StatusInternalServerError, &models.Response{
			Message: "An internal server error occurred when saving the image. Please try again in a few moments!",
		})
	}

	// Construct the URL for the saved picture.
	baseURL := "http://127.0.0.1:8080"
	pictureURL := baseURL + "/images/" + receipt_image.Filename

	result, err := models.CreateBill(pictureURL, restaurant_name, subtotal, total_discount, service_charge, tax, other, grand_total, isSettled, date_created, payment_id, user_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

// saveUploadedFile function to handle file uploads.
func saveUploadedFile(file *multipart.FileHeader, path string) error {
	// Open the uploaded file.
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Create a destination file for the uploaded content.
	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy the uploaded content to the destination file.
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}
