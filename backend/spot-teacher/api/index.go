package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()
	// ここで /products, /healthz… など自由に定義
	e.GET("/healthz", func(c echo.Context) error { return c.NoContent(http.StatusOK) })
	// ... ほかハンドラ登録 ...
	e.ServeHTTP(w, r)
}
