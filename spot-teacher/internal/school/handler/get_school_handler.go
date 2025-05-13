package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/usecase"
	sh "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/hander"
	"net/http"
	"strconv"
)

// GetSchoolPresenter defines the interface for presenting get school response
type GetSchoolPresenter interface {
	Present(c echo.Context, school *domain.School) error
}

// DefaultGetSchoolPresenter is the default implementation of GetSchoolPresenter
type DefaultGetSchoolPresenter struct{}

// Present presents the get school response
func (p *DefaultGetSchoolPresenter) Present(c echo.Context, school *domain.School) error {
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

// GetSchoolHandler handles get school request
type GetSchoolHandler struct {
	useCase   usecase.SchoolUsecase
	presenter GetSchoolPresenter
}

// NewGetSchoolHandler creates a new GetSchoolHandler
func NewGetSchoolHandler(
	uc usecase.SchoolUsecase,
	p GetSchoolPresenter,
) *GetSchoolHandler {
	return &GetSchoolHandler{
		useCase:   uc,
		presenter: p,
	}
}

// HandleGetSchool handles the get school request
func (h *GetSchoolHandler) HandleGetSchool(c echo.Context) error {
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

	// Call the usecase
	school, err := h.useCase.GetSchool(c.Request().Context(), schoolID)
	if err != nil {
		return sh.ErrorJSON(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	// Present the response
	return h.presenter.Present(c, school)
}
