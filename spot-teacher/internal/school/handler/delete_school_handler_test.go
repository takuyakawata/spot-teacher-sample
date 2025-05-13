package handler

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/usecase"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockSchoolUsecaseForDelete is a mock implementation of the SchoolUsecase interface for delete tests
type MockSchoolUsecaseForDelete struct {
	mock.Mock
}

// Ensure MockSchoolUsecaseForDelete implements usecase.SchoolUsecase
var _ usecase.SchoolUsecase = (*MockSchoolUsecaseForDelete)(nil)

func (m *MockSchoolUsecaseForDelete) CreateSchool(ctx context.Context, school *domain.School) (*domain.School, error) {
	args := m.Called(ctx, school)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.School), args.Error(1)
}

func (m *MockSchoolUsecaseForDelete) GetSchool(ctx context.Context, id domain.SchoolID) (*domain.School, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.School), args.Error(1)
}

func (m *MockSchoolUsecaseForDelete) UpdateSchool(ctx context.Context, school *domain.School) (*domain.School, error) {
	args := m.Called(ctx, school)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.School), args.Error(1)
}

func (m *MockSchoolUsecaseForDelete) DeleteSchool(ctx context.Context, id domain.SchoolID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockSchoolUsecaseForDelete) ListSchools(ctx context.Context) ([]*domain.School, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.School), args.Error(1)
}

func TestHandleDeleteSchool_Success(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/schools/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Create mock
	mockUsecase := new(MockSchoolUsecaseForDelete)

	// Create handler
	handler := NewDeleteSchoolHandler(mockUsecase)

	// Setup expectations
	schoolID, _ := domain.NewSchoolID(1)
	mockUsecase.On("DeleteSchool", mock.Anything, schoolID).Return(nil)

	// Call the handler
	err := handler.HandleDeleteSchool(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockUsecase.AssertExpectations(t)
}

func TestHandleDeleteSchool_InvalidID(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/schools/:id")
	c.SetParamNames("id")
	c.SetParamValues("invalid")

	// Create mock
	mockUsecase := new(MockSchoolUsecaseForDelete)

	// Create handler
	handler := NewDeleteSchoolHandler(mockUsecase)

	// Call the handler
	err := handler.HandleDeleteSchool(c)

	// Assert
	assert.NoError(t, err) // The error is handled within the handler
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestHandleDeleteSchool_WithAssociatedTeachers(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/schools/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Create mock
	mockUsecase := new(MockSchoolUsecaseForDelete)

	// Create handler
	handler := NewDeleteSchoolHandler(mockUsecase)

	// Setup expectations
	schoolID, _ := domain.NewSchoolID(1)
	mockUsecase.On("DeleteSchool", mock.Anything, schoolID).Return(errors.New("cannot delete school with associated teachers"))

	// Call the handler
	err := handler.HandleDeleteSchool(c)

	// Assert
	assert.NoError(t, err) // The error is handled within the handler
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	mockUsecase.AssertExpectations(t)
}

func TestHandleDeleteSchool_InternalError(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/schools/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Create mock
	mockUsecase := new(MockSchoolUsecaseForDelete)

	// Create handler
	handler := NewDeleteSchoolHandler(mockUsecase)

	// Setup expectations
	schoolID, _ := domain.NewSchoolID(1)
	mockUsecase.On("DeleteSchool", mock.Anything, schoolID).Return(errors.New("internal error"))

	// Call the handler
	err := handler.HandleDeleteSchool(c)

	// Assert
	assert.NoError(t, err) // The error is handled within the handler
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUsecase.AssertExpectations(t)
}
