package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/usecase"
	sh "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/hander"
	"net/http"
)

// SchoolResponse is a DTO for school response
type SchoolResponse struct {
	ID          int64   `json:"id"`
	SchoolType  string  `json:"schoolType"`
	Name        string  `json:"name"`
	Email       *string `json:"email,omitempty"`
	PhoneNumber string  `json:"phoneNumber"`
	Address     struct {
		Prefecture int     `json:"prefecture"`
		City       string  `json:"city"`
		Street     *string `json:"street,omitempty"`
		PostCode   string  `json:"postCode"`
	} `json:"address"`
	URL *string `json:"url,omitempty"`
}

// ListSchoolsResponse is a DTO for list schools response
type ListSchoolsResponse struct {
	Schools []SchoolResponse `json:"schools"`
}

// ListSchoolsPresenter defines the interface for presenting list schools response
type ListSchoolsPresenter interface {
	Present(c echo.Context, schools []*domain.School) error
}

// DefaultListSchoolsPresenter is the default implementation of ListSchoolsPresenter
type DefaultListSchoolsPresenter struct{}

// Present presents the list schools response
func (p *DefaultListSchoolsPresenter) Present(c echo.Context, schools []*domain.School) error {
	resp := ListSchoolsResponse{
		Schools: make([]SchoolResponse, len(schools)),
	}

	for i, school := range schools {
		schoolResp := SchoolResponse{
			ID:          school.ID.Value(),
			SchoolType:  string(school.SchoolType),
			Name:        school.Name.Value(),
			PhoneNumber: school.PhoneNumber.Value(),
		}

		if school.Email != nil {
			email := school.Email.Value()
			schoolResp.Email = &email
		}

		schoolResp.Address.Prefecture = int(school.Address.Prefecture)
		schoolResp.Address.City = school.Address.City
		schoolResp.Address.PostCode = school.Address.PostCode.Value()
		if school.Address.Street != nil {
			schoolResp.Address.Street = school.Address.Street
		}

		if school.URL.String() != "" {
			url := school.URL.String()
			schoolResp.URL = &url
		}

		resp.Schools[i] = schoolResp
	}

	return c.JSON(http.StatusOK, resp)
}

// ListSchoolsHandler handles list schools request
type ListSchoolsHandler struct {
	useCase   usecase.SchoolUsecase
	presenter ListSchoolsPresenter
}

// NewListSchoolsHandler creates a new ListSchoolsHandler
func NewListSchoolsHandler(
	uc usecase.SchoolUsecase,
	p ListSchoolsPresenter,
) *ListSchoolsHandler {
	return &ListSchoolsHandler{
		useCase:   uc,
		presenter: p,
	}
}

// HandleListSchools handles the list schools request
func (h *ListSchoolsHandler) HandleListSchools(c echo.Context) error {
	// Call the usecase
	schools, err := h.useCase.ListSchools(c.Request().Context())
	if err != nil {
		return sh.ErrorJSON(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	// Present the response
	return h.presenter.Present(c, schools)
}
