package infra

import (
	"github.com/takuyakawta/spot-teacher-sample/db/ent"
	//domain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson/domain"
)

type LessonScheduleRepoImpl struct {
	client *ent.Client
}

//
//func NewLessonScheduleRepoImpl(client *ent.Client) *domain.LessonScheduleRepository {
//	return &LessonScheduleRepoImpl{client: client}
//}
