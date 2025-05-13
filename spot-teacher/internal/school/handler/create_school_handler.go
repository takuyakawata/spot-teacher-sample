package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/usecase"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
	sh "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/hander"
	"net/http"
)

// CreateSchoolRequest is a DTO for create school request
type CreateSchoolRequest struct {
	SchoolType  string `json:"schoolType"`
	Name        string `json:"name"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phoneNumber"`
	Address     struct {
		Prefecture int     `json:"prefecture"`
		City       string  `json:"city"`
		Street     *string `json:"street,omitempty"`
		PostCode   string  `json:"postCode"`
	} `json:"address"`
	URL string `json:"url,omitempty"`
}

// CreateSchoolPresenter defines the interface for presenting create school response
type CreateSchoolPresenter interface {
	Present(c echo.Context, school *domain.School) error
}

// DefaultCreateSchoolPresenter is the default implementation of CreateSchoolPresenter
type DefaultCreateSchoolPresenter struct{}

// Present presents the create school response
func (p *DefaultCreateSchoolPresenter) Present(c echo.Context, school *domain.School) error {
	resp := SchoolResponse{
		ID:          school.ID.Value(),
		SchoolType:  string(school.SchoolType),
		Name:        school.Name.Value(),
		PhoneNumber: school.PhoneNumber.Value(),
	}

	if school.Email != nil {
		email := school.Email.Value()
		resp.Email = &email
	}

	resp.Address.Prefecture = int(school.Address.Prefecture)
	resp.Address.City = school.Address.City
	resp.Address.PostCode = school.Address.PostCode.Value()
	if school.Address.Street != nil {
		resp.Address.Street = school.Address.Street
	}

	if school.URL.String() != "" {
		url := school.URL.String()
		resp.URL = &url
	}

	return c.JSON(http.StatusCreated, resp)
}

// CreateSchoolHandler handles create school request
type CreateSchoolHandler struct {
	useCase   usecase.SchoolUsecase
	presenter CreateSchoolPresenter
}

// NewCreateSchoolHandler creates a new CreateSchoolHandler
func NewCreateSchoolHandler(
	uc usecase.SchoolUsecase,
	p CreateSchoolPresenter,
) *CreateSchoolHandler {
	return &CreateSchoolHandler{
		useCase:   uc,
		presenter: p,
	}
}

// HandleCreateSchool handles the create school request
func (h *CreateSchoolHandler) HandleCreateSchool(c echo.Context) error {
	// Bind request
	var req CreateSchoolRequest
	if err := c.Bind(&req); err != nil {
		return sh.ErrorJSON(c, http.StatusBadRequest, "Invalid request", err.Error())
	}

	// Create domain entities
	schoolName, err := domain.NewSchoolName(req.Name)
	if err != nil {
		return sh.ErrorJSON(c, http.StatusBadRequest, "Invalid school name", err.Error())
	}

	phoneNumber, err := sharedDomain.NewPhoneNumber(req.PhoneNumber)
	if err != nil {
		return sh.ErrorJSON(c, http.StatusBadRequest, "Invalid phone number", err.Error())
	}

	postCode, err := sharedDomain.NewPostCode(req.Address.PostCode)
	if err != nil {
		return sh.ErrorJSON(c, http.StatusBadRequest, "Invalid post code", err.Error())
	}

	address := sharedDomain.Address{
		Prefecture: sharedDomain.Prefecture(req.Address.Prefecture),
		City:       req.Address.City,
		Street:     req.Address.Street,
		PostCode:   postCode,
	}

	var url sharedDomain.URL
	if req.URL != "" {
		urlPtr, err := sharedDomain.NewURL(req.URL)
		if err != nil {
			return sh.ErrorJSON(c, http.StatusBadRequest, "Invalid URL", err.Error())
		}
		if urlPtr != nil {
			url = *urlPtr
		}
	}

	var email *sharedDomain.EmailAddress
	if req.Email != "" {
		emailAddr, err := sharedDomain.NewEmailAddress(req.Email)
		if err != nil {
			return sh.ErrorJSON(c, http.StatusBadRequest, "Invalid email", err.Error())
		}
		email = &emailAddr
	}

	// Create school entity
	school, err := domain.NewSchool(
		domain.SchoolID(0), // ID will be assigned by the repository
		domain.SchoolType(req.SchoolType),
		schoolName,
		email,
		phoneNumber,
		address,
		url,
	)
	if err != nil {
		return sh.ErrorJSON(c, http.StatusBadRequest, "Invalid school data", err.Error())
	}

	// Call the usecase
	createdSchool, err := h.useCase.CreateSchool(c.Request().Context(), school)
	if err != nil {
		return sh.ErrorJSON(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	// Present the response
	return h.presenter.Present(c, createdSchool)
}
