package usecase

import (
	"context"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
)

type CreateSchoolUseCase struct {
	repo domain.SchoolRepository
}

func NewCreateSchoolUseCase(repo domain.SchoolRepository) *CreateSchoolUseCase {
	return &CreateSchoolUseCase{repo: repo}
}

// CreateSchool creates a new school
func (u *CreateSchoolUseCase) CreateSchool(ctx context.Context, school *domain.School) (*domain.School, error) {
	existingSchool, err := u.repo.FindByName(ctx, school.Name)
	if err != nil {
		return nil, err
	}

	if existingSchool != nil {
		return nil, &domain.SchoolAlreadyExistsError{Name: existingSchool.Name}
	}

	return u.repo.Create(ctx, school)
}
