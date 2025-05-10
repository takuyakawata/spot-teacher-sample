package usecase_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"

	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/user/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/user/usecase"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/user/usecase/mock"
)

func TestCreateTeacherUseCase_Execute(t *testing.T) {
	// Test cases
	tests := []struct {
		name          string
		input         usecase.CreateTeacherUseCaseInput
		mockSetup     func(mockRepo *mock.MockTeacherRepository)
		expectedError string
	}{
		{
			name: "Success",
			input: usecase.CreateTeacherUseCaseInput{
				Email:           sharedDomain.EmailAddress("test@example.com"),
				FirstName:       domain.TeacherName("John"),
				LastName:        domain.TeacherName("Doe"),
				Password:        sharedDomain.Password("password"),
				ConfirmPassword: sharedDomain.Password("password"),
			},
			mockSetup: func(mockRepo *mock.MockTeacherRepository) {
				// Email doesn't exist
				mockRepo.EXPECT().
					FindByEmail(gomock.Any(), sharedDomain.EmailAddress("test@example.com")).
					Return(nil, nil)

				// Create succeeds
				mockRepo.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			expectedError: "",
		},
		{
			name: "Email already exists",
			input: usecase.CreateTeacherUseCaseInput{
				Email:           sharedDomain.EmailAddress("existing@example.com"),
				FirstName:       domain.TeacherName("John"),
				LastName:        domain.TeacherName("Doe"),
				Password:        sharedDomain.Password("password"),
				ConfirmPassword: sharedDomain.Password("password"),
			},
			mockSetup: func(mockRepo *mock.MockTeacherRepository) {
				// Email exists
				existingTeacher := &domain.Teacher{
					Email: sharedDomain.EmailAddress("existing@example.com"),
				}
				mockRepo.EXPECT().
					FindByEmail(gomock.Any(), sharedDomain.EmailAddress("existing@example.com")).
					Return(existingTeacher, nil)
			},
			expectedError: "email already exists",
		},
		{
			name: "Passwords do not match",
			input: usecase.CreateTeacherUseCaseInput{
				Email:           sharedDomain.EmailAddress("test@example.com"),
				FirstName:       domain.TeacherName("John"),
				LastName:        domain.TeacherName("Doe"),
				Password:        sharedDomain.Password("password1"),
				ConfirmPassword: sharedDomain.Password("password2"),
			},
			mockSetup: func(mockRepo *mock.MockTeacherRepository) {
				// Email doesn't exist
				mockRepo.EXPECT().
					FindByEmail(gomock.Any(), sharedDomain.EmailAddress("test@example.com")).
					Return(nil, nil)
			},
			expectedError: "passwords do not match",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Create mock repository
			mockRepo := mock.NewMockTeacherRepository(ctrl)

			// Setup mock expectations
			if tt.mockSetup != nil {
				tt.mockSetup(mockRepo)
			}

			// Create use case
			uc := usecase.NewCreateTeacherUseCase(mockRepo)

			// Execute use case
			err := uc.Execute(tt.input)

			// Check results
			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
