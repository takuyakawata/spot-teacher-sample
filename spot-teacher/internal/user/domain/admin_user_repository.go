package domain

import (
	"context"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
)

type AdminUserRepository interface {
	Create(ctx context.Context, adminUser *AdminUser) error
	FindByID(ctx context.Context, id AdminUserID) (*AdminUser, error)
	FindByEmail(ctx context.Context, email sharedDomain.EmailAddress) (*AdminUser, error)
}
