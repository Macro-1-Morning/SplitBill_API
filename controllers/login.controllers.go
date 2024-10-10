package controllers

import(
	"github.com/labstack/echo/v4"
	"splitbill_api/helpers"
	// "vp_week11_echo/models"
	"net/http"
	// "github.com/dgrijalva/jwt-go"
	// "time"
)
func GenerateHashPassword(c echo.Context) error{
	password := c.Param("password")
	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)

}