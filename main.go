package main

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/infernalslam/simple-api/config"
	"github.com/infernalslam/simple-api/controllers/account"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

func main() {
	db := config.InitDatabase()
	fmt.Println(reflect.TypeOf(db))
	e := echo.New()
	e.HideBanner = true
	e.GET("/", createAccount)
	e.Logger.Fatal(e.Start(":1323"))
}

func createAccount(c echo.Context) error {
	response := account.CreateAccount()
	return c.String(http.StatusOK, response)
}
