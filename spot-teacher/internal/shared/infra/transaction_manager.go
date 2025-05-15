package infra

import (
	"context"
	domain3 "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/company/domain"
	domain5 "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson/domain"
	domain4 "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson_category/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
	domain2 "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/user/domain"
)

type Repositories struct {
	SchoolRepo            domain.SchoolRepository
	TeacherRepo           domain2.TeacherRepository
	CompanyRepo           domain3.CompanyRepository
	GradeRepo             domain4.GradeRepository
	SubjectRepo           domain4.SubjectRepository
	EducationCategoryRepo domain4.EducationCategoryRepository
	LessonPlanRepo        domain5.LessonPlanRepository
}

type TransactionManager interface {
	Do(ctx context.Context, fn func(ctx context.Context, r *Repositories) error) error
}
