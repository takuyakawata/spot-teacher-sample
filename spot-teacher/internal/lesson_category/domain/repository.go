package domain

import "context"

type GradeRepository interface {
	Create(ctx context.Context, grade *Grade) (*Grade, error)
	RetrieveAll(ctx context.Context) ([]*Grade, error)
}

type SubjectRepository interface {
	Create(ctx context.Context, subject *Subject) (*Subject, error)
	RetrieveAll(ctx context.Context) ([]*Subject, error)
}

type EducationCategoryRepository interface {
	Create(ctx context.Context, educationCategory *EducationCategory) (*EducationCategory, error)
	RetrieveAll(ctx context.Context) ([]*EducationCategory, error)
}
