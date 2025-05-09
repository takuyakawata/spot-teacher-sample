package domain

import "context"

type TeacherRepository interface {
	Create(ctx context.Context, teacher *Teacher) (*Teacher, error)
}

type CompanyMemberRepository interface {
}
