package routes
// go mod init vp_week11_echo 
// GO111MODULE=on go get github.com/labstack/echo/v4

import (
	// "net/http"
	"github.com/labstack/echo/v4"
	"vp_week11_echo/controllers"
	// "vp_week11_echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	
	//Bill
	e.GET("/getAllBill", controllers.ReadAllBill)
	e.GET("/getBill", controllers.ReadBillByBillId)
	e.POST("/createBill", controllers.CreateBill)
	e.PATCH("/editBill", controllers.EditBill)
	e.DELETE("/deleteBill", controllers.DeleteBill)

	//Debtor Bill
	e.GET("/getDebtorBill", controllers.ReadDebtorBill)
	e.POST("/createDebtorBill", controllers.CreateDebtorBill)
	e.PATCH("/editDebtorBill", controllers.EditDebtorBill)
	e.DELETE("/deleteDebtorBill", controllers.DeleteDebtorBill)

	return e

}

