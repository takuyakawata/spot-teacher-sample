package infra

import (
	"context"
	"errors"
	"fmt"
	"github.com/takuyakawta/spot-teacher-sample/backend/db/ent"
	"github.com/takuyakawta/spot-teacher-sample/backend/spot-teacher/internal/product/domain"
)

type productRepository struct {
	client *ent.Client
}

func NewProductRepository(client *ent.Client) domain.ProductRepository {
	return &productRepository{client: client}
}

func (r *productRepository) Create(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	createCmd := r.client.Product.Create()
	createCmd.SetName(product.Name.Value())
	createCmd.SetPrice(product.Price.Value())
	if product.Description != nil {
		createCmd.SetDescription(*product.Description)
	}
	createdEntProduct, err := createCmd.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to create product: %w", err)
	}
	return mapEntProductToDomain(createdEntProduct)
}

func (r *productRepository) FindByID(ctx context.Context, id domain.ProductID) (*domain.Product, error) {
	entProduct, err := r.client.Product.Get(ctx, int(id.Value()))
	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to find product by id %d: %w", id, err)
	}
	return mapEntProductToDomain(entProduct)
}

func mapEntProductToDomain(entP *ent.Product) (*domain.Product, error) {
	if entP == nil {
		// nil ポインタが渡された場合のエラー処理
		return nil, errors.New("infra.ent: cannot map nil ent.Product")
	}

	// ID の変換とバリデーション
	domainID, err := domain.NewProductID(int64(entP.ID))
	if err != nil {
		// DBから取得したIDがドメインルールを満たさない場合 (通常は起こりにくい)
		return nil, fmt.Errorf("infra.ent: invalid id %d from db: %w", entP.ID, err)
	}

	// Name の変換とバリデーション
	domainName, err := domain.NewProductName(entP.Name)
	if err != nil {
		return nil, fmt.Errorf("infra.ent: invalid name '%s' from db (id: %d): %w", entP.Name, entP.ID, err)
	}

	// Price の変換
	domainPrice, err := domain.NewProductPrice(entP.Price)
	if err != nil {
		return nil, fmt.Errorf("infra.ent: invalid price %f from db (id: %d): %w", entP.Price, entP.ID, err)
	}

	var descriptionPtr *string
	if entP.Description != "" {
		descValue := entP.Description
		descriptionPtr = &descValue
	}

	return &domain.Product{
		ID:          domainID,
		Name:        domainName,
		Description: descriptionPtr,
		Price:       domainPrice,
	}, nil
}
