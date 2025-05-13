package handler

import (
	"bytes"
	"context"
	"encoding/json"
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

// MockSchoolUsecase is a mock implementation of the SchoolUsecase interface
type MockSchoolUsecase struct {
	mock.Mock
}

// Ensure MockSchoolUsecase implements usecase.SchoolUsecase
var _ usecase.SchoolUsecase = (*MockSchoolUsecase)(nil)

func (m *MockSchoolUsecase) CreateSchool(ctx context.Context, school *domain.School) (*domain.School, error) {
	args := m.Called(ctx, school)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.School), args.Error(1)
}

func (m *MockSchoolUsecase) GetSchool(ctx context.Context, id domain.SchoolID) (*domain.School, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.School), args.Error(1)
}

func (m *MockSchoolUsecase) UpdateSchool(ctx context.Context, school *domain.School) (*domain.School, error) {
	args := m.Called(ctx, school)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.School), args.Error(1)
}

func (m *MockSchoolUsecase) DeleteSchool(ctx context.Context, id domain.SchoolID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockSchoolUsecase) ListSchools(ctx context.Context) ([]*domain.School, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.School), args.Error(1)
}

// MockCreateSchoolPresenter is a mock implementation of the CreateSchoolPresenter interface
type MockCreateSchoolPresenter struct {
	mock.Mock
}

func (m *MockCreateSchoolPresenter) Present(c echo.Context, school *domain.School) error {
	args := m.Called(c, school)
	return args.Error(0)
}

func TestHandleCreateSchool_Success(t *testing.T) {
	// Setup
	e := echo.New()
	reqBody := CreateSchoolRequest{
		SchoolType:  "elementary",
		Name:        "Test School",
		Email:       "test@example.com",
		PhoneNumber: "03-1234-5678",
		Address: struct {
			Prefecture int     `json:"prefecture"`
			City       string  `json:"city"`
			Street     *string `json:"street,omitempty"`
			PostCode   string  `json:"postCode"`
		}{
			Prefecture: 13,
			City:       "Tokyo",
			Street:     nil,
			PostCode:   "123-4567",
		},
		URL: "https://example.com",
	}

	reqJSON, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/schools", bytes.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create mocks
	mockUsecase := new(MockSchoolUsecase)
	mockPresenter := new(MockCreateSchoolPresenter)

	// Create handler
	handler := NewCreateSchoolHandler(mockUsecase, mockPresenter)

	// Setup expectations
	schoolName, _ := domain.NewSchoolName(reqBody.Name)
	phoneNumber, _ := sharedDomain.NewPhoneNumber(reqBody.PhoneNumber)
	postCode, _ := sharedDomain.NewPostCode(reqBody.Address.PostCode)
	emailAddr, _ := sharedDomain.NewEmailAddress(reqBody.Email)
	url, _ := sharedDomain.NewURL(reqBody.URL)

	address := sharedDomain.Address{
		Prefecture: sharedDomain.Prefecture(reqBody.Address.Prefecture),
		City:       reqBody.Address.City,
		Street:     reqBody.Address.Street,
		PostCode:   postCode,
	}

	expectedSchool, _ := domain.NewSchool(
		domain.SchoolID(0),
		domain.SchoolType(reqBody.SchoolType),
		schoolName,
		&emailAddr,
		phoneNumber,
		address,
		*url,
	)

	createdSchool, _ := domain.NewSchool(
		domain.SchoolID(1),
		domain.SchoolType(reqBody.SchoolType),
		schoolName,
		&emailAddr,
		phoneNumber,
		address,
		*url,
	)

	mockUsecase.On("CreateSchool", mock.Anything, mock.MatchedBy(func(s *domain.School) bool {
		// We can't directly compare the schools because the one passed to CreateSchool
		// will have ID 0, but we can compare the other fields
		return s.SchoolType == expectedSchool.SchoolType &&
			s.Name.Value() == expectedSchool.Name.Value() &&
			s.Email.Value() == expectedSchool.Email.Value() &&
			s.PhoneNumber.Value() == expectedSchool.PhoneNumber.Value()
	})).Return(createdSchool, nil)

	mockPresenter.On("Present", c, createdSchool).Return(nil)

	// Call the handler
	err := handler.HandleCreateSchool(c)

	// Assert
	assert.NoError(t, err)
	mockUsecase.AssertExpectations(t)
	mockPresenter.AssertExpectations(t)
}

func TestHandleCreateSchool_InvalidRequest(t *testing.T) {
	// Setup
	e := echo.New()
	reqBody := `{"invalid": json`
	req := httptest.NewRequest(http.MethodPost, "/schools", bytes.NewReader([]byte(reqBody)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create mocks
	mockUsecase := new(MockSchoolUsecase)
	mockPresenter := new(MockCreateSchoolPresenter)

	// Create handler
	handler := NewCreateSchoolHandler(mockUsecase, mockPresenter)

	// Call the handler
	err := handler.HandleCreateSchool(c)

	// Assert
	assert.NoError(t, err) // The error is handled within the handler
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestHandleCreateSchool_UsecaseError(t *testing.T) {
	// Setup
	e := echo.New()
	reqBody := CreateSchoolRequest{
		SchoolType:  "elementary",
		Name:        "Test School",
		Email:       "test@example.com",
		PhoneNumber: "03-1234-5678",
		Address: struct {
			Prefecture int     `json:"prefecture"`
			City       string  `json:"city"`
			Street     *string `json:"street,omitempty"`
			PostCode   string  `json:"postCode"`
		}{
			Prefecture: 13,
			City:       "Tokyo",
			Street:     nil,
			PostCode:   "123-4567",
		},
		URL: "https://example.com",
	}

	reqJSON, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/schools", bytes.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create mocks
	mockUsecase := new(MockSchoolUsecase)
	mockPresenter := new(MockCreateSchoolPresenter)

	// Create handler
	handler := NewCreateSchoolHandler(mockUsecase, mockPresenter)

	// Setup expectations
	mockUsecase.On("CreateSchool", mock.Anything, mock.Anything).Return(nil, errors.New("usecase error"))

	// Call the handler
	err := handler.HandleCreateSchool(c)

	// Assert
	assert.NoError(t, err) // The error is handled within the handler
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUsecase.AssertExpectations(t)
}
