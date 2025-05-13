package usecase

import (
	"context"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
)

type UpdateSchoolUseCase struct {
	repo domain.SchoolRepository
}

func NewUpdateSchoolUseCase(repo domain.SchoolRepository) *UpdateSchoolUseCase {
	return &UpdateSchoolUseCase{repo: repo}
}

// UpdateSchool updates an existing school
func (u *UpdateSchoolUseCase) UpdateSchool(ctx context.Context, school *domain.School) (*domain.School, error) {
	// Check if the school exists
	_, err := u.repo.FindByID(ctx, school.ID)
	if err != nil {
		return nil, err
	}

	return u.repo.Update(ctx, school)
}
