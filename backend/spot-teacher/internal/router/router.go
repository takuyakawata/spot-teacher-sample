package router

import "github.com/labstack/echo/v4"

// RouteRegistrar は自分のルートを登録できるもの
type RouteRegistrar interface {
	RegisterRoutes(e *echo.Echo)
}

// RegisterAll は渡されたすべての RouteRegistrar を登録する
func RegisterAll(e *echo.Echo, regs ...RouteRegistrar) {
	for _, r := range regs {
		r.RegisterRoutes(e)
	}
}
