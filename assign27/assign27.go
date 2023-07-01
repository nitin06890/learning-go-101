package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	host      = "localhost"
	port      = "9000"
	seperator = ":"
)

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Shelby7!")
	})
	e.Logger.Fatal(e.Start(host + seperator + port))
}
