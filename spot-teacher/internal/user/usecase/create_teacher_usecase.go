package usecase

import (
	"errors"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/user/domain"
)

type CreateTeacherUseCase struct {
	TeacherRepository domain.TeacherRepository
}

func NewCreateTeacherUseCase(teacherRepository domain.TeacherRepository) *CreateTeacherUseCase {
	return &CreateTeacherUseCase{TeacherRepository: teacherRepository}
}

type CreateTeacherUseCaseInput struct {
	Email           sharedDomain.EmailAddress
	FirstName       domain.TeacherName
	LastName        domain.TeacherName
	Password        sharedDomain.Password
	ConfirmPassword sharedDomain.Password
}

func (uc *CreateTeacherUseCase) Execute(input CreateTeacherUseCaseInput) error {
	// Check if the email already exists
	existingTeacher, _ := uc.TeacherRepository.FindByEmail(nil, input.Email)
	if existingTeacher != nil {
		return errors.New("email already exists")
	}

	// Validate passwords
	if input.Password != input.ConfirmPassword {
		return errors.New("passwords do not match")
	}

	// Create new teacher entity
	teacher := &domain.Teacher{
		FirstName:  input.FirstName,
		FamilyName: input.LastName,
		Email:      input.Email,
		Password:   input.Password,
	}

	// Save the teacher in the repository
	return uc.TeacherRepository.Create(nil, teacher)
}
