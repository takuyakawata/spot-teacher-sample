package handler

import (
	"github.com/labstack/echo/v4"
	schoolDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/user/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/user/usecase"
	"net/http"
)

type CreateTeacherRequest struct {
	FirstName       string `json:"first_name"`
	FamilyName      string `json:"last_name"`
	SchoolID        int64  `json:"school_id"`
	Email           string `json:"email"`
	PhoneNumber     string `json:"phone_number"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type CreateTeacherResponse struct {
}

type CreateTeacherHandler struct {
	useCase   *usecase.CreateTeacherUseCase
	presenter *CreateTeacherPresenter
}

func NewCreateTeacherHandler(useCase *usecase.CreateTeacherUseCase) *CreateTeacherHandler {
	return &CreateTeacherHandler{useCase: useCase}
}

func (h *CreateTeacherHandler) HandleCreateTeacher(c echo.Context) error {
	var req CreateTeacherRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	phoneNumber := sharedDomain.PhoneNumber(req.PhoneNumber)

	if err := h.useCase.Execute(
		usecase.CreateTeacherUseCaseInput{
			FirstName:  domain.TeacherName(req.FirstName),
			FamilyName: domain.TeacherName(req.FamilyName),
			Email:      sharedDomain.EmailAddress(req.Email),
			SchoolID:   schoolDomain.SchoolID(req.SchoolID),
			// TODO このまま
			PhoneNumber:     &phoneNumber,
			Password:        sharedDomain.Password(req.Password),
			ConfirmPassword: sharedDomain.Password(req.ConfirmPassword),
		},
	); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, CreateTeacherResponse{})
}

type CreateTeacherPresenter struct{}

func NewCreateTeacherPresenter() CreateTeacherPresenter {
	return CreateTeacherPresenter{}
}

func (p CreateTeacherPresenter) Present(c echo.Context, teacher domain.Teacher) error {
	return nil
}
