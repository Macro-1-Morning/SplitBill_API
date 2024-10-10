package main
// go mod init vp_week11_echo 
// GO111MODULE=on go get github.com/labstack/echo/v4
// go get github.com/tkanos/gonfig
// go get -u github.com/go-sql-driver/mysql
//go get golang.org/x/crypto/bcrypt
//go get github.com/dgrijalva/jwt-go
//go get github.com/labstack/echo/v4/middleware
//go get github.com/go-playground/validator
import (
	"splitbill_api/routes"
	"splitbill_api/db"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start("127.0.0.1:8080"))
}
