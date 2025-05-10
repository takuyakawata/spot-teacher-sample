package handler

import "github.com/labstack/echo/v4"

type TeacherHandler struct {
	Create *CreateTeacherHandler
}

func (h TeacherHandler) RegisterRoutes(e *echo.Echo) {
	apiGroup := e.Group("/spot-teacher")
	g := apiGroup.Group("/teacher")
	g.POST("", h.Create.HandleCreateTeacher)
}

type CompanyMemberHandler struct {
}

type AdminUserHandler struct {
}
