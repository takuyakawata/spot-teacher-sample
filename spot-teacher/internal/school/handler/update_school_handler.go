package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/usecase"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
	sh "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/hander"
	"net/http"
	"strconv"
)

// UpdateSchoolRequest is a DTO for update school request
type UpdateSchoolRequest struct {
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

// UpdateSchoolPresenter defines the interface for presenting update school response
type UpdateSchoolPresenter interface {
	Present(c echo.Context, school *domain.School) error
}

// DefaultUpdateSchoolPresenter is the default implementation of UpdateSchoolPresenter
type DefaultUpdateSchoolPresenter struct{}

// Present presents the update school response
func (p *DefaultUpdateSchoolPresenter) Present(c echo.Context, school *domain.School) error {
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

	return c.JSON(http.StatusOK, resp)
}

// UpdateSchoolHandler handles update school request
type UpdateSchoolHandler struct {
	useCase   usecase.SchoolUsecase
	presenter UpdateSchoolPresenter
}

// NewUpdateSchoolHandler creates a new UpdateSchoolHandler
func NewUpdateSchoolHandler(
	uc usecase.SchoolUsecase,
	p UpdateSchoolPresenter,
) *UpdateSchoolHandler {
	return &UpdateSchoolHandler{
		useCase:   uc,
		presenter: p,
	}
}

// HandleUpdateSchool handles the update school request
func (h *UpdateSchoolHandler) HandleUpdateSchool(c echo.Context) error {
	// Parse school ID from path parameter
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return sh.ErrorJSON(c, http.StatusBadRequest, "Invalid school ID", err.Error())
	}

	// Create domain ID
	schoolID, err := domain.NewSchoolID(id)
	if err != nil {
		return sh.ErrorJSON(c, http.StatusBadRequest, "Invalid school ID", err.Error())
	}

	// Bind request
	var req UpdateSchoolRequest
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
		schoolID,
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
	updatedSchool, err := h.useCase.UpdateSchool(c.Request().Context(), school)
	if err != nil {
		return sh.ErrorJSON(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	// Present the response
	return h.presenter.Present(c, updatedSchool)
}
