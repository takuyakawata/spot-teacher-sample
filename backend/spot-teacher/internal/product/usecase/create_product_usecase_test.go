package usecase_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/takuyakawta/spot-teacher-sample/backend/spot-teacher/internal/product/domain"
	"github.com/takuyakawta/spot-teacher-sample/backend/spot-teacher/internal/product/usecase"
	"github.com/takuyakawta/spot-teacher-sample/backend/spot-teacher/internal/product/usecase/mock"
)

func TestCreateProductUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックリポジトリ生成
	mockRepo := mock.NewMockProductRepository(ctrl)

	// テストデータ
	name := domain.ProductName("Pen")
	price := domain.ProductPrice(1.23)
	desc := "A nice pen"
	input := usecase.CreateProductInput{
		Name:        name,
		Description: &desc,
		Price:       price,
	}

	// 成功ケース：Create の期待設定
	mockRepo.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, p *domain.Product) (*domain.Product, error) {
			// 実際の実装同様に ID を入れて返す
			p.ID = domain.ProductID(1)
			return p, nil
		})

	uc := usecase.NewCreateProductUseCase(mockRepo)
	got, err := uc.Execute(input)

	assert.NoError(t, err)
	assert.Equal(t, domain.ProductID(1), got.ID)
	assert.Equal(t, "Pen", got.Name.Value())
	assert.Equal(t, &desc, got.Description)
	assert.Equal(t, price, got.Price)
}
