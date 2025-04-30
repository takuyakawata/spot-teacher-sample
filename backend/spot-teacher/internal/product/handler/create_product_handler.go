package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/takuyakawta/spot-teacher-sample/backend/spot-teacher/internal/product/domain"
	"github.com/takuyakawta/spot-teacher-sample/backend/spot-teacher/internal/product/usecase"
	sh "github.com/takuyakawta/spot-teacher-sample/backend/spot-teacher/internal/shared/hander"
	"net/http"
)

// CreateProductRequest は商品作成APIへのリクエストボディを表す DTO
type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Price       float64 `json:"price"`
}

// CreateProductResponse は商品作成APIの成功レスポンスを表す DTO
type CreateProductResponse struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Price       float64 `json:"price"`
}

type CreateProductHandler struct {
	useCase   *usecase.CreateProductUseCase
	presenter CreateProductPresenter
}

func NewCreateProductHandler(
	uc *usecase.CreateProductUseCase,
	p CreateProductPresenter,
) *CreateProductHandler {
	return &CreateProductHandler{
		useCase:   uc,
		presenter: p,
	}
}

func (h *CreateProductHandler) HandleCreateProduct(c echo.Context) error {
	var req CreateProductRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// UseCaseの実行
	product, err := h.useCase.Execute(usecase.CreateProductInput{
		Name:        domain.ProductName(req.Name),
		Description: req.Description,
		Price:       domain.ProductPrice(req.Price),
	})
	if err != nil {
		return sh.ErrorJSON(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return h.presenter.Present(c, *product)

}
