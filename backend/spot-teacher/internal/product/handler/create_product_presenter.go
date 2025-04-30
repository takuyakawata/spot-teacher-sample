package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/takuyakawta/spot-teacher-sample/backend/spot-teacher/internal/product/domain"
	"net/http"
)

type CreateProductPresenter interface {
	Present(c echo.Context, p domain.Product) error
}

type createProductPresenter struct{}

func NewCreateProductPresenter() CreateProductPresenter {
	return &createProductPresenter{}
}

func (p createProductPresenter) Present(c echo.Context, product domain.Product) error {
	// dto.CreateProductOutput をそのまま JSON シリアライズ
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"id":          product.ID,
		"name":        product.Name,
		"description": product.Description,
		"price":       product.Price,
	})
}
