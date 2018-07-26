package main

import (
	"net/http"

	"github.com/infernalslam/simple-api/controllers/account"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.GET("/", createAccount)
	e.Logger.Fatal(e.Start(":1323"))
}

func createAccount(c echo.Context) error {
	response := account.CreateAccount()
	return c.String(http.StatusOK, response)
}
