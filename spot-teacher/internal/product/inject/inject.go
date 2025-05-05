package inject

import (
	"github.com/google/wire"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/product/handler"
)

var productSet = wire.NewSet(
	handler.NewCreateProductHandler,
	wire.Struct(new(handler.ProductHandler), "*"),
)

func InitializeProductHandler() *handler.ProductHandler {
	wire.Build(productSet)
	return &handler.ProductHandler{}
}
