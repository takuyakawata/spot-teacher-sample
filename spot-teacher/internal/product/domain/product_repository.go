package domain

import (
	"context" // コンテキストを使用してリクエストのキャンセルやタイムアウトを管理
)

type ProductRepository interface {
	Create(ctx context.Context, product *Product) (*Product, error)
	Update(ctx context.Context, product *Product) (*Product, error)
	Delete(ctx context.Context, id ProductID) error
	FindByID(ctx context.Context, id ProductID) (*Product, error)
	FindAll(ctx context.Context) ([]*Product, error)
}
