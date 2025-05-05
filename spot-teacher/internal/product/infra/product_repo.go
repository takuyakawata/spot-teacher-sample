package infra

import (
	"context"
	"errors"
	"fmt"
	"github.com/takuyakawta/spot-teacher-sample/db/ent"

	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/product/domain"
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

func (r *productRepository) Update(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	primitiveID := int(product.ID.Value())
	updateCmd := r.client.Product.UpdateOneID(primitiveID)
	updateCmd.SetName(product.Name.Value())
	updateCmd.SetPrice(product.Price.Value())
	if product.Description != nil {
		updateCmd.SetDescription(*product.Description)
	}
	updatedEntProduct, err := updateCmd.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to update product: %w", err)
	}
	return mapEntProductToDomain(updatedEntProduct)
}

func (r *productRepository) Delete(ctx context.Context, id domain.ProductID) error {
	err := r.client.Product.DeleteOneID(int(id)).Exec(ctx) // 削除クエリを実行

	if err != nil {
		if ent.IsNotFound(err) {
			return fmt.Errorf("infra.ent: product to delete with id %v not found: %w", id, err)
		}
		return fmt.Errorf("infra.ent: failed to delete product with id %v: %w", id, err)
	}

	return nil
}

func (r *productRepository) FindByID(ctx context.Context, id domain.ProductID) (*domain.Product, error) {
	entProduct, err := r.client.Product.Get(ctx, int(id.Value()))
	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to find product by id %d: %w", id, err)
	}
	return mapEntProductToDomain(entProduct)
}

func (r *productRepository) FindAll(ctx context.Context) ([]*domain.Product, error) {
	products, err := r.client.Product.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to find all products: %w", err)
	}
	domainProducts := make([]*domain.Product, 0, len(products))
	for _, entP := range products {
		domainP, mapErr := mapEntProductToDomain(entP)
		if mapErr != nil {
			return nil, fmt.Errorf("failed to map product (ent ID: %v) in FindAll: %w", entP.ID, mapErr)
		}
		// 変換成功したドメインオブジェクトをスライスに追加
		domainProducts = append(domainProducts, domainP)
	}
	return domainProducts, nil
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
		return nil, nil
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
