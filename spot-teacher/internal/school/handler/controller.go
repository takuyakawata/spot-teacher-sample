package handler

import "github.com/labstack/echo/v4"

type SchoolHandler struct {
	List   *ListSchoolsHandler
	Get    *GetSchoolHandler
	Create *CreateSchoolHandler
	Update *UpdateSchoolHandler
	Delete *DeleteSchoolHandler
}

func (h SchoolHandler) RegisterRoutes(e *echo.Echo) {
	apiGroup := e.Group("/spot-teacher")
	g := apiGroup.Group("/school")
	g.GET("", h.List.HandleListSchools)
	g.GET("/:id", h.Get.HandleGetSchool)
	g.POST("", h.Create.HandleCreateSchool)
	g.PUT("/:id", h.Update.HandleUpdateSchool)
	g.DELETE("/:id", h.Delete.HandleDeleteSchool)
}
