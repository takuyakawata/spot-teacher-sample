package usecase

import (
	"context"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
)

// ListSchools retrieves all schools
type ListSchoolsUseCase struct {
	repo domain.SchoolRepository
}

func NewListSchoolsUseCase(repo domain.SchoolRepository) ListSchoolsUseCase {
	return ListSchoolsUseCase{repo: repo}
}

func (u *ListSchoolsUseCase) ListSchools(ctx context.Context) ([]*domain.School, error) {
	return u.repo.FindAll(ctx)
}
