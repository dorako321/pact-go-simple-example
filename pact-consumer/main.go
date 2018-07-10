package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func index(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

func getAnimal(c echo.Context) error {
	// APIにアクセス
	animal_name := "サーバル" // TODO: api叩く処理

	// レスポンスを加工して返却
	return c.String(http.StatusOK, "そうだね！"+animal_name+"ちゃんだね！")
}

func main() {
	// 8000番でサーバの起動
	e := echo.New()
	e.GET("/animal/1", getAnimal)
	e.Logger.Fatal(e.Start(":8000"))
}