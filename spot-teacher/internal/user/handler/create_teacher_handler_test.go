package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	schoolDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/user/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/user/usecase"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockTeacherRepository is a mock implementation of the TeacherRepository interface
type MockTeacherRepository struct {
	mock.Mock
}

func (m *MockTeacherRepository) Create(ctx interface{}, teacher *domain.Teacher) error {
	args := m.Called(ctx, teacher)
	return args.Error(0)
}

func (m *MockTeacherRepository) FindByID(ctx interface{}, id domain.TeacherID) (*domain.Teacher, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Teacher), args.Error(1)
}

func (m *MockTeacherRepository) FindByEmail(ctx interface{}, email sharedDomain.EmailAddress) (*domain.Teacher, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Teacher), args.Error(1)
}

func (m *MockTeacherRepository) FindBySchoolID(ctx interface{}, schoolID schoolDomain.SchoolID) ([]*domain.Teacher, error) {
	args := m.Called(ctx, schoolID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Teacher), args.Error(1)
}

func TestHandleCreateTeacher_Success(t *testing.T) {
	// Setup
	e := echo.New()
	reqBody := CreateTeacherRequest{
		FirstName:       "John",
		FamilyName:      "Doe",
		SchoolID:        1,
		Email:           "john.doe@example.com",
		PhoneNumber:     "03-1234-5678",
		Password:        "password123",
		ConfirmPassword: "password123",
	}

	reqJSON, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/teachers", bytes.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create mock repository
	mockRepo := new(MockTeacherRepository)

	// Setup repository expectations
	mockRepo.On("FindByEmail", nil, sharedDomain.EmailAddress(reqBody.Email)).Return(nil, nil)
	mockRepo.On("Create", nil, mock.AnythingOfType("*domain.Teacher")).Return(nil)

	// Create usecase with mock repository
	uc := usecase.NewCreateTeacherUseCase(mockRepo)

	// Create handler
	handler := NewCreateTeacherHandler(uc)

	// Call the handler
	err := handler.HandleCreateTeacher(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockRepo.AssertExpectations(t)
}

func TestHandleCreateTeacher_InvalidRequest(t *testing.T) {
	// Setup
	e := echo.New()
	reqBody := `{"invalid": json`
	req := httptest.NewRequest(http.MethodPost, "/teachers", bytes.NewReader([]byte(reqBody)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create mock repository
	mockRepo := new(MockTeacherRepository)

	// Create usecase with mock repository
	uc := usecase.NewCreateTeacherUseCase(mockRepo)

	// Create handler
	handler := NewCreateTeacherHandler(uc)

	// Call the handler
	err := handler.HandleCreateTeacher(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestHandleCreateTeacher_EmailExists(t *testing.T) {
	// Setup
	e := echo.New()
	reqBody := CreateTeacherRequest{
		FirstName:       "John",
		FamilyName:      "Doe",
		SchoolID:        1,
		Email:           "john.doe@example.com",
		PhoneNumber:     "03-1234-5678",
		Password:        "password123",
		ConfirmPassword: "password123",
	}

	reqJSON, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/teachers", bytes.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create mock repository
	mockRepo := new(MockTeacherRepository)

	// Setup repository expectations
	existingTeacher := &domain.Teacher{
		Email: sharedDomain.EmailAddress(reqBody.Email),
	}
	mockRepo.On("FindByEmail", nil, sharedDomain.EmailAddress(reqBody.Email)).Return(existingTeacher, nil)

	// Create usecase with mock repository
	uc := usecase.NewCreateTeacherUseCase(mockRepo)

	// Create handler
	handler := NewCreateTeacherHandler(uc)

	// Call the handler
	err := handler.HandleCreateTeacher(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	mockRepo.AssertExpectations(t)
}

func TestHandleCreateTeacher_PasswordMismatch(t *testing.T) {
	// Setup
	e := echo.New()
	reqBody := CreateTeacherRequest{
		FirstName:       "John",
		FamilyName:      "Doe",
		SchoolID:        1,
		Email:           "john.doe@example.com",
		PhoneNumber:     "03-1234-5678",
		Password:        "password123",
		ConfirmPassword: "different-password",
	}

	reqJSON, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/teachers", bytes.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create mock repository
	mockRepo := new(MockTeacherRepository)

	// Setup repository expectations
	mockRepo.On("FindByEmail", nil, sharedDomain.EmailAddress(reqBody.Email)).Return(nil, nil)

	// Create usecase with mock repository
	uc := usecase.NewCreateTeacherUseCase(mockRepo)

	// Create handler
	handler := NewCreateTeacherHandler(uc)

	// Call the handler
	err := handler.HandleCreateTeacher(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	mockRepo.AssertExpectations(t)
}
