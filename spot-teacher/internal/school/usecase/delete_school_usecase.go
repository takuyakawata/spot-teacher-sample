package usecase

import (
	"context"
	"errors"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
	domain2 "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/user/domain"
)

// DeleteSchool deletes a school by ID
// Returns an error if the school has associated teachers

type DeleteSchoolUseCase struct {
	schoolRepo  domain.SchoolRepository
	teacherRepo domain2.TeacherRepository
}

func NewDeleteSchoolUseCase(
	repo domain.SchoolRepository,
	teacherRepo domain2.TeacherRepository,
) *DeleteSchoolUseCase {
	return &DeleteSchoolUseCase{
		schoolRepo:  repo,
		teacherRepo: teacherRepo,
	}
}

func (u *DeleteSchoolUseCase) DeleteSchool(ctx context.Context, id domain.SchoolID) error {
	// Check if the school exists
	_, err := u.schoolRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// Check if the school has associated teachers
	teachers, err := u.teacherRepo.FindBySchoolID(ctx, id)
	if err != nil {
		return err
	}

	if len(teachers) > 0 {
		return errors.New("cannot delete school with associated teachers")
	}

	return u.schoolRepo.Delete(ctx, id)
}
