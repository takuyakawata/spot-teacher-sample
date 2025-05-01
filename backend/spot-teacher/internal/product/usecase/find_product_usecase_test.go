package usecase_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takuyakawta/spot-teacher-sample/backend/spot-teacher/internal/product/domain"
	"github.com/takuyakawta/spot-teacher-sample/backend/spot-teacher/internal/product/usecase"
	"github.com/takuyakawta/spot-teacher-sample/backend/spot-teacher/internal/product/usecase/mock"
	"testing"
)

func TestFindProductUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックリポジトリ生成
	mockRepo := mock.NewMockProductRepository(ctrl)

	// テストデータ
	id := domain.ProductID(1)
	name := domain.ProductName("Pen")
	price := domain.ProductPrice(1.23)
	desc := "A nice pen"

	product := &domain.Product{
		ID:          id,
		Name:        name,
		Description: &desc,
		Price:       price,
	}

	// 成功ケース：FindByID の期待設定
	mockRepo.EXPECT().
		FindByID(gomock.Any(), id).
		Return(product, nil).Times(1)

	uc := usecase.NewFindProductUseCase(mockRepo)
	got, err := uc.Execute(id)

	assert.NoError(t, err)
	assert.Equal(t, product, got)
	assert.Equal(t, id, got.ID)
	assert.Equal(t, name, got.Name)
	assert.Equal(t, &desc, got.Description)
	assert.Equal(t, price, got.Price)
}
