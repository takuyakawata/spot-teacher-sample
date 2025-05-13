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

// MockSchoolUsecaseForUpdate is a mock implementation of the SchoolUsecase interface for update tests
type MockSchoolUsecaseForUpdate struct {
	mock.Mock
}

// Ensure MockSchoolUsecaseForUpdate implements usecase.SchoolUsecase
var _ usecase.SchoolUsecase = (*MockSchoolUsecaseForUpdate)(nil)

func (m *MockSchoolUsecaseForUpdate) CreateSchool(ctx context.Context, school *domain.School) (*domain.School, error) {
	args := m.Called(ctx, school)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.School), args.Error(1)
}

func (m *MockSchoolUsecaseForUpdate) GetSchool(ctx context.Context, id domain.SchoolID) (*domain.School, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.School), args.Error(1)
}

func (m *MockSchoolUsecaseForUpdate) UpdateSchool(ctx context.Context, school *domain.School) (*domain.School, error) {
	args := m.Called(ctx, school)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.School), args.Error(1)
}

func (m *MockSchoolUsecaseForUpdate) DeleteSchool(ctx context.Context, id domain.SchoolID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockSchoolUsecaseForUpdate) ListSchools(ctx context.Context) ([]*domain.School, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.School), args.Error(1)
}

// MockUpdateSchoolPresenter is a mock implementation of the UpdateSchoolPresenter interface
type MockUpdateSchoolPresenter struct {
	mock.Mock
}

func (m *MockUpdateSchoolPresenter) Present(c echo.Context, school *domain.School) error {
	args := m.Called(c, school)
	return args.Error(0)
}

func TestHandleUpdateSchool_Success(t *testing.T) {
	// Setup
	e := echo.New()
	reqBody := UpdateSchoolRequest{
		SchoolType:  "elementary",
		Name:        "Updated School",
		Email:       "updated@example.com",
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
		URL: "https://updated-example.com",
	}

	reqJSON, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/schools/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Create mocks
	mockUsecase := new(MockSchoolUsecaseForUpdate)
	mockPresenter := new(MockUpdateSchoolPresenter)

	// Create handler
	handler := NewUpdateSchoolHandler(mockUsecase, mockPresenter)

	// Setup expectations
	schoolID, _ := domain.NewSchoolID(1)
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
		schoolID,
		domain.SchoolType(reqBody.SchoolType),
		schoolName,
		&emailAddr,
		phoneNumber,
		address,
		*url,
	)

	updatedSchool, _ := domain.NewSchool(
		schoolID,
		domain.SchoolType(reqBody.SchoolType),
		schoolName,
		&emailAddr,
		phoneNumber,
		address,
		*url,
	)

	mockUsecase.On("UpdateSchool", mock.Anything, mock.MatchedBy(func(s *domain.School) bool {
		return s.ID == expectedSchool.ID &&
			s.SchoolType == expectedSchool.SchoolType &&
			s.Name.Value() == expectedSchool.Name.Value() &&
			s.Email.Value() == expectedSchool.Email.Value() &&
			s.PhoneNumber.Value() == expectedSchool.PhoneNumber.Value()
	})).Return(updatedSchool, nil)

	mockPresenter.On("Present", c, updatedSchool).Return(nil)

	// Call the handler
	err := handler.HandleUpdateSchool(c)

	// Assert
	assert.NoError(t, err)
	mockUsecase.AssertExpectations(t)
	mockPresenter.AssertExpectations(t)
}

func TestHandleUpdateSchool_InvalidID(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/schools/:id")
	c.SetParamNames("id")
	c.SetParamValues("invalid")

	// Create mocks
	mockUsecase := new(MockSchoolUsecaseForUpdate)
	mockPresenter := new(MockUpdateSchoolPresenter)

	// Create handler
	handler := NewUpdateSchoolHandler(mockUsecase, mockPresenter)

	// Call the handler
	err := handler.HandleUpdateSchool(c)

	// Assert
	assert.NoError(t, err) // The error is handled within the handler
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestHandleUpdateSchool_InvalidRequest(t *testing.T) {
	// Setup
	e := echo.New()
	reqBody := `{"invalid": json`
	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewReader([]byte(reqBody)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/schools/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Create mocks
	mockUsecase := new(MockSchoolUsecaseForUpdate)
	mockPresenter := new(MockUpdateSchoolPresenter)

	// Create handler
	handler := NewUpdateSchoolHandler(mockUsecase, mockPresenter)

	// Call the handler
	err := handler.HandleUpdateSchool(c)

	// Assert
	assert.NoError(t, err) // The error is handled within the handler
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestHandleUpdateSchool_UsecaseError(t *testing.T) {
	// Setup
	e := echo.New()
	reqBody := UpdateSchoolRequest{
		SchoolType:  "elementary",
		Name:        "Updated School",
		Email:       "updated@example.com",
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
		URL: "https://updated-example.com",
	}

	reqJSON, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/schools/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Create mocks
	mockUsecase := new(MockSchoolUsecaseForUpdate)
	mockPresenter := new(MockUpdateSchoolPresenter)

	// Create handler
	handler := NewUpdateSchoolHandler(mockUsecase, mockPresenter)

	// Setup expectations
	mockUsecase.On("UpdateSchool", mock.Anything, mock.Anything).Return(nil, errors.New("usecase error"))

	// Call the handler
	err := handler.HandleUpdateSchool(c)

	// Assert
	assert.NoError(t, err) // The error is handled within the handler
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUsecase.AssertExpectations(t)
}
