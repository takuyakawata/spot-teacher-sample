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

// MockSchoolUsecaseForList is a mock implementation of the SchoolUsecase interface for list tests
type MockSchoolUsecaseForList struct {
	mock.Mock
}

// Ensure MockSchoolUsecaseForList implements usecase.SchoolUsecase
var _ usecase.SchoolUsecase = (*MockSchoolUsecaseForList)(nil)

func (m *MockSchoolUsecaseForList) CreateSchool(ctx context.Context, school *domain.School) (*domain.School, error) {
	args := m.Called(ctx, school)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.School), args.Error(1)
}

func (m *MockSchoolUsecaseForList) GetSchool(ctx context.Context, id domain.SchoolID) (*domain.School, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.School), args.Error(1)
}

func (m *MockSchoolUsecaseForList) UpdateSchool(ctx context.Context, school *domain.School) (*domain.School, error) {
	args := m.Called(ctx, school)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.School), args.Error(1)
}

func (m *MockSchoolUsecaseForList) DeleteSchool(ctx context.Context, id domain.SchoolID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockSchoolUsecaseForList) ListSchools(ctx context.Context) ([]*domain.School, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.School), args.Error(1)
}

// MockListSchoolsPresenter is a mock implementation of the ListSchoolsPresenter interface
type MockListSchoolsPresenter struct {
	mock.Mock
}

func (m *MockListSchoolsPresenter) Present(c echo.Context, schools []*domain.School) error {
	args := m.Called(c, schools)
	return args.Error(0)
}

func TestHandleListSchools_Success(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/schools", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create mocks
	mockUsecase := new(MockSchoolUsecaseForList)
	mockPresenter := new(MockListSchoolsPresenter)

	// Create handler
	handler := NewListSchoolsHandler(mockUsecase, mockPresenter)

	// Setup expectations
	// Create sample schools
	schoolName1, _ := domain.NewSchoolName("Test School 1")
	phoneNumber1, _ := sharedDomain.NewPhoneNumber("03-1234-5678")
	postCode1, _ := sharedDomain.NewPostCode("123-4567")
	emailAddr1, _ := sharedDomain.NewEmailAddress("test1@example.com")
	url1, _ := sharedDomain.NewURL("https://example1.com")

	address1 := sharedDomain.Address{
		Prefecture: sharedDomain.Prefecture(13),
		City:       "Tokyo",
		Street:     nil,
		PostCode:   postCode1,
	}

	school1, _ := domain.NewSchool(
		domain.SchoolID(1),
		domain.SchoolType("elementary"),
		schoolName1,
		&emailAddr1,
		phoneNumber1,
		address1,
		*url1,
	)

	schoolName2, _ := domain.NewSchoolName("Test School 2")
	phoneNumber2, _ := sharedDomain.NewPhoneNumber("03-8765-4321")
	postCode2, _ := sharedDomain.NewPostCode("765-4321")
	emailAddr2, _ := sharedDomain.NewEmailAddress("test2@example.com")
	url2, _ := sharedDomain.NewURL("https://example2.com")

	address2 := sharedDomain.Address{
		Prefecture: sharedDomain.Prefecture(14),
		City:       "Yokohama",
		Street:     nil,
		PostCode:   postCode2,
	}

	school2, _ := domain.NewSchool(
		domain.SchoolID(2),
		domain.SchoolType("high"),
		schoolName2,
		&emailAddr2,
		phoneNumber2,
		address2,
		*url2,
	)

	schools := []*domain.School{school1, school2}

	mockUsecase.On("ListSchools", mock.Anything).Return(schools, nil)
	mockPresenter.On("Present", c, schools).Return(nil)

	// Call the handler
	err := handler.HandleListSchools(c)

	// Assert
	assert.NoError(t, err)
	mockUsecase.AssertExpectations(t)
	mockPresenter.AssertExpectations(t)
}

func TestHandleListSchools_EmptyList(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/schools", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create mocks
	mockUsecase := new(MockSchoolUsecaseForList)
	mockPresenter := new(MockListSchoolsPresenter)

	// Create handler
	handler := NewListSchoolsHandler(mockUsecase, mockPresenter)

	// Setup expectations
	schools := []*domain.School{}

	mockUsecase.On("ListSchools", mock.Anything).Return(schools, nil)
	mockPresenter.On("Present", c, schools).Return(nil)

	// Call the handler
	err := handler.HandleListSchools(c)

	// Assert
	assert.NoError(t, err)
	mockUsecase.AssertExpectations(t)
	mockPresenter.AssertExpectations(t)
}

func TestHandleListSchools_Error(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/schools", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create mocks
	mockUsecase := new(MockSchoolUsecaseForList)
	mockPresenter := new(MockListSchoolsPresenter)

	// Create handler
	handler := NewListSchoolsHandler(mockUsecase, mockPresenter)

	// Setup expectations
	mockUsecase.On("ListSchools", mock.Anything).Return(nil, errors.New("database error"))

	// Call the handler
	err := handler.HandleListSchools(c)

	// Assert
	assert.NoError(t, err) // The error is handled within the handler
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUsecase.AssertExpectations(t)
}
