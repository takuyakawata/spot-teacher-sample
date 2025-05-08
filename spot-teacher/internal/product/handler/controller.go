package handler

import "github.com/labstack/echo/v4"

type ProductHandler struct {
	Create *CreateProductHandler
	Test   *TestProductHandler
}

func (h ProductHandler) RegisterRoutes(e *echo.Echo) {
	apiGroup := e.Group("/spot-teacher")
	g := apiGroup.Group("/product")
	g.POST("", h.Create.HandleCreateProduct)
	g.GET("/test", h.Test.HandleTestProduct)
}
