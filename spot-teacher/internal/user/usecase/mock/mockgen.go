package mock

//go:generate mockgen -package mock -destination mock_teacher_repository.go github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/user/domain TeacherRepository
