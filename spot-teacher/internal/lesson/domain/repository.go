package domain

import (
	"context"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/company/domain"
)

type LessonPlanRepository interface {
	Create(ctx context.Context, lessonPlan *LessonPlan) (*LessonPlan, error)
	Update(ctx context.Context, lessonPlan *LessonPlan) (*LessonPlan, error)
	FilterByCompanyID(ctx context.Context, companyID domain.CompanyID) ([]*LessonPlan, error)
	FindByID(ctx context.Context, id LessonPlanID) (*LessonPlan, error)
}

type lessonScheduleRepository interface {
	Create(ctx context.Context, lessonSchedule LessonSchedule) error
	Update(ctx context.Context, lessonSchedule LessonSchedule) error
	FilterByLessonPlanID(ctx context.Context, lessonPlanID LessonPlanID) []LessonSchedule
}
