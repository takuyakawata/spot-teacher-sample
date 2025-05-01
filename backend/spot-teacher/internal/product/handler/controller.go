package handler

import "github.com/labstack/echo/v4"

type ProductHandler struct {
	Create *CreateProductHandler
}

func (h ProductHandler) RegisterRoutes(e *echo.Echo) {
	g := e.Group("/products")
	g.POST("", h.Create.HandleCreateProduct)
}
