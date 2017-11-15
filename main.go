package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()

	// ルーティング
	e.GET("/", MainPage())

	// サーバー起動
	e.Start(":8080")
}

func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}