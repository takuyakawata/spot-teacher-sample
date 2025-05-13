package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/usecase"
	sh "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/hander"
	"net/http"
	"strconv"
)

// DeleteSchoolHandler handles delete school request
type DeleteSchoolHandler struct {
	useCase usecase.SchoolUsecase
}

// NewDeleteSchoolHandler creates a new DeleteSchoolHandler
func NewDeleteSchoolHandler(
	uc usecase.SchoolUsecase,
) *DeleteSchoolHandler {
	return &DeleteSchoolHandler{
		useCase: uc,
	}
}

// HandleDeleteSchool handles the delete school request
func (h *DeleteSchoolHandler) HandleDeleteSchool(c echo.Context) error {
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
	err = h.useCase.DeleteSchool(c.Request().Context(), schoolID)
	if err != nil {
		// Check if the error is because the school has associated teachers
		if err.Error() == "cannot delete school with associated teachers" {
			return sh.ErrorJSON(c, http.StatusBadRequest, "Cannot delete school with associated teachers", err.Error())
		}
		return sh.ErrorJSON(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	// Return success response
	return c.NoContent(http.StatusNoContent)
}
