package usecase

import (
	"errors"
	schoolDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
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
	FirstName       sharedDomain.UserName
	FamilyName      sharedDomain.UserName
	SchoolID        schoolDomain.SchoolID
	Email           sharedDomain.EmailAddress
	PhoneNumber     *sharedDomain.PhoneNumber
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
		FirstName:   input.FirstName,
		FamilyName:  input.FamilyName,
		SchoolID:    input.SchoolID,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Password:    input.Password,
	}

	// Save the teacher in the repository
	return uc.TeacherRepository.Create(nil, teacher)
}
