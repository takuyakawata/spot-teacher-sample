package handler

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/usecase"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockSchoolUsecaseForGet is a mock implementation of the SchoolUsecase interface for get tests
type MockSchoolUsecaseForGet struct {
	mock.Mock
}

// Ensure MockSchoolUsecaseForGet implements usecase.SchoolUsecase
var _ usecase.SchoolUsecase = (*MockSchoolUsecaseForGet)(nil)

func (m *MockSchoolUsecaseForGet) CreateSchool(ctx context.Context, school *domain.School) (*domain.School, error) {
	args := m.Called(ctx, school)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.School), args.Error(1)
}

func (m *MockSchoolUsecaseForGet) GetSchool(ctx context.Context, id domain.SchoolID) (*domain.School, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.School), args.Error(1)
}

func (m *MockSchoolUsecaseForGet) UpdateSchool(ctx context.Context, school *domain.School) (*domain.School, error) {
	args := m.Called(ctx, school)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.School), args.Error(1)
}

func (m *MockSchoolUsecaseForGet) DeleteSchool(ctx context.Context, id domain.SchoolID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockSchoolUsecaseForGet) ListSchools(ctx context.Context) ([]*domain.School, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.School), args.Error(1)
}

// MockGetSchoolPresenter is a mock implementation of the GetSchoolPresenter interface
type MockGetSchoolPresenter struct {
	mock.Mock
}

func (m *MockGetSchoolPresenter) Present(c echo.Context, school *domain.School) error {
	args := m.Called(c, school)
	return args.Error(0)
}

func TestHandleGetSchool_Success(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/schools/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Create mocks
	mockUsecase := new(MockSchoolUsecaseForGet)
	mockPresenter := new(MockGetSchoolPresenter)

	// Create handler
	handler := NewGetSchoolHandler(mockUsecase, mockPresenter)

	// Setup expectations
	schoolID, _ := domain.NewSchoolID(1)
	schoolName, _ := domain.NewSchoolName("Test School")
	phoneNumber, _ := sharedDomain.NewPhoneNumber("03-1234-5678")
	postCode, _ := sharedDomain.NewPostCode("123-4567")
	emailAddr, _ := sharedDomain.NewEmailAddress("test@example.com")
	url, _ := sharedDomain.NewURL("https://example.com")

	address := sharedDomain.Address{
		Prefecture: sharedDomain.Prefecture(13),
		City:       "Tokyo",
		Street:     nil,
		PostCode:   postCode,
	}

	school, _ := domain.NewSchool(
		schoolID,
		domain.SchoolType("elementary"),
		schoolName,
		&emailAddr,
		phoneNumber,
		address,
		*url,
	)

	mockUsecase.On("GetSchool", mock.Anything, schoolID).Return(school, nil)
	mockPresenter.On("Present", c, school).Return(nil)

	// Call the handler
	err := handler.HandleGetSchool(c)

	// Assert
	assert.NoError(t, err)
	mockUsecase.AssertExpectations(t)
	mockPresenter.AssertExpectations(t)
}

func TestHandleGetSchool_InvalidID(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/schools/:id")
	c.SetParamNames("id")
	c.SetParamValues("invalid")

	// Create mocks
	mockUsecase := new(MockSchoolUsecaseForGet)
	mockPresenter := new(MockGetSchoolPresenter)

	// Create handler
	handler := NewGetSchoolHandler(mockUsecase, mockPresenter)

	// Call the handler
	err := handler.HandleGetSchool(c)

	// Assert
	assert.NoError(t, err) // The error is handled within the handler
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestHandleGetSchool_NotFound(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/schools/:id")
	c.SetParamNames("id")
	c.SetParamValues("999")

	// Create mocks
	mockUsecase := new(MockSchoolUsecaseForGet)
	mockPresenter := new(MockGetSchoolPresenter)

	// Create handler
	handler := NewGetSchoolHandler(mockUsecase, mockPresenter)

	// Setup expectations
	schoolID, _ := domain.NewSchoolID(999)
	mockUsecase.On("GetSchool", mock.Anything, schoolID).Return(nil, errors.New("school not found"))

	// Call the handler
	err := handler.HandleGetSchool(c)

	// Assert
	assert.NoError(t, err) // The error is handled within the handler
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUsecase.AssertExpectations(t)
}
