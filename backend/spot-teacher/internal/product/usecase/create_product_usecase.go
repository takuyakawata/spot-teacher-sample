package usecase

import "github.com/takuyakawta/spot-teacher-sample/backend/spot-teacher/internal/product/domain"

type CreateProductUseCase struct {
	productRepository domain.ProductRepository
}

func NewCreateProductUseCase(productRepository domain.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{productRepository: productRepository}
}

type CreateProductInput struct {
	Name        domain.ProductName  // 商品名
	Description *string             // 商品の説明 (nil許可)
	Price       domain.ProductPrice // 商品の価格 (非負)
}

type CreateProductOutput struct {
	product *domain.Product
}

func (uc *CreateProductUseCase) Execute(input CreateProductInput) (*domain.Product, error) {
	// ドメインモデルの生成
	productID, _ := domain.NewProductID(0) // IDはDBで自動生成されるため、0を指定
	productName := input.Name
	productPrice := input.Price

	product, _ := domain.NewProduct(
		productID,
		productName,
		input.Description,
		productPrice,
	)
	return uc.productRepository.Create(nil, product)
}
