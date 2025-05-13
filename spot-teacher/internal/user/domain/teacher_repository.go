package domain

import (
	"context"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
)

type TeacherRepository interface {
	Create(ctx context.Context, teacher *Teacher) error
	FindByID(ctx context.Context, id TeacherID) (*Teacher, error)
	FindByEmail(ctx context.Context, email sharedDomain.EmailAddress) (*Teacher, error)
}

type CompanyMemberRepository interface {
}

type adminUserRepository interface {
}
