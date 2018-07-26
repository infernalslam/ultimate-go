package main

import (
	"log"
	"net/http"

	"github.com/infernalslam/simple-api/controllers/account"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/db")
	if err != nil {
		log.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()
	e := echo.New()
	e.HideBanner = true
	e.GET("/", createAccount)
	e.Logger.Fatal(e.Start(":1323"))
}

func createAccount(c echo.Context) error {
	response := account.CreateAccount()
	return c.String(http.StatusOK, response)
}
