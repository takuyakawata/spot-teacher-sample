package usecase

import (
	"fmt"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/product/domain"
)

type FindProductUseCase struct {
	productRepo domain.ProductRepository
}

func NewFindProductUseCase(productRepo domain.ProductRepository) *FindProductUseCase {
	return &FindProductUseCase{
		productRepo: productRepo,
	}
}

func (uc *FindProductUseCase) Execute(id domain.ProductID) (*domain.Product, error) {
	product, err := uc.productRepo.FindByID(nil, id)
	if err != nil {
		// エラーハンドリングを適切に行う
		return nil, fmt.Errorf("failed to find product:%w", err)
	}
	return product, nil
}
