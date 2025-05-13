package usecase

import (
	"context"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
)

type GetSchoolUseCase struct {
	repo domain.SchoolRepository
}

func NewGetSchoolUseCase(repo domain.SchoolRepository) *GetSchoolUseCase {
	return &GetSchoolUseCase{repo: repo}
}

// GetSchool retrieves a school by ID
func (u *GetSchoolUseCase) GetSchool(ctx context.Context, id domain.SchoolID) (*domain.School, error) {
	return u.repo.FindByID(ctx, id)
}
