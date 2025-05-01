package domain

import (
	"context" // コンテキストを使用してリクエストのキャンセルやタイムアウトを管理
)

type ProductRepository interface {
	// Create は新しい Product を永続化します。
	Create(ctx context.Context, product *Product) (*Product, error)

	// FindByID は ID を元に Product を取得します。見つからない場合はエラーを返します。
	FindByID(ctx context.Context, id ProductID) (*Product, error)

	// Update は既存の Product を更新します。
	Update(ctx context.Context, product *Product) (*Product, error)

	// Delete は指定された ID の Product を削除します。
	Delete(ctx context.Context, id int64) error

	// FindAll は全ての Product を取得します (PoC用、実際はページネーションなどが必要)
	FindAll(ctx context.Context) ([]*Product, error)
}
