package main

import (
"github.com/labstack/echo"
"net/http"
)

func getAnimal(c echo.Context) error {
	return c.String(http.StatusOK, "{\"id\": 1, \"name\": \"サーバル\"}")
}

func main() {
	e := echo.New()
	e.GET("/api/v1/animal/1", getAnimal)
	e.Logger.Fatal(e.Start(":3005"))
}
