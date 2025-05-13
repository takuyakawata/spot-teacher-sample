package inject

import (
	"github.com/google/wire"
	"github.com/takuyakawta/spot-teacher-sample/db/ent"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/handler"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/infra"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/usecase"
	userInfra "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/user/infra"
)

var schoolSet = wire.NewSet(
	// Repositories
	infra.NewSchoolRepositoryImpl,
	userInfra.NewTeacherRepositoryImpl,

	// Usecases
	usecase.NewSchoolUsecaseImpl,

	// Presenters
	wire.Struct(new(handler.DefaultListSchoolsPresenter), "*"),
	wire.Struct(new(handler.DefaultGetSchoolPresenter), "*"),
	wire.Struct(new(handler.DefaultCreateSchoolPresenter), "*"),
	wire.Struct(new(handler.DefaultUpdateSchoolPresenter), "*"),

	// Handlers
	handler.NewListSchoolsHandler,
	handler.NewGetSchoolHandler,
	handler.NewCreateSchoolHandler,
	handler.NewUpdateSchoolHandler,
	handler.NewDeleteSchoolHandler,

	// Controller
	wire.Struct(new(handler.SchoolHandler), "*"),
)

func InitializeSchoolHandler(client *ent.Client) *handler.SchoolHandler {
	wire.Build(schoolSet)
	return &handler.SchoolHandler{}
}
