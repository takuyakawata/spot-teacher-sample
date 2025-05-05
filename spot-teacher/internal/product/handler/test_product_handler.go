package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// CreateProductRequest は商品作成APIへのリクエストボディを表す DTO
type TestProductRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Price       int64   `json:"price"`
}

// CreateProductResponse は商品作成APIの成功レスポンスを表す DTO
type TestProductResponse struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Price       int64   `json:"price"`
}

type TestProductHandler struct {
}

func NewTestProductHandler() *TestProductHandler {
	return &TestProductHandler{}
}

func (h *TestProductHandler) HandleTestProduct(c echo.Context) error {
	//helloと返す
	return c.String(http.StatusOK, "hello")
}
