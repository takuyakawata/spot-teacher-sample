package domain

import (
	"context" // コンテキストを使用してリクエストのキャンセルやタイムアウトを管理
)

type CompanyRepository interface {
	Create(ctx context.Context, company *Company) (*Company, error)
	Update(ctx context.Context, company *Company) (*Company, error)
	Delete(ctx context.Context, id CompanyID) error
	FindByID(ctx context.Context, id CompanyID) (*Company, error)
	FindAll(ctx context.Context) ([]*Company, error)
}
