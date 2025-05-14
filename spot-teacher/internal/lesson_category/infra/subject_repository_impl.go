package infra

import (
	"context"
	"fmt"
	"github.com/takuyakawta/spot-teacher-sample/db/ent"
	entsubject "github.com/takuyakawta/spot-teacher-sample/db/ent/subject"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/lesson_category/domain"
)

type subjectRepository struct {
	client *ent.Client
}

func NewSubjectRepository(client *ent.Client) domain.SubjectRepository {
	return &subjectRepository{client: client}
}

func (r *subjectRepository) Create(ctx context.Context, subject *domain.Subject) (*domain.Subject, error) {
	// Convert domain Subject to string representation
	subjectEnum := domain.SubjectEnum(*subject)
	subjectStr := string(subjectEnum)

	// Check if subject already exists
	existingSubject, err := r.client.Subject.Query().
		Where(entsubject.Name(subjectStr)).
		First(ctx)

	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("infra.ent: failed to query subject: %w", err)
	}

	// If subject already exists, return it
	if existingSubject != nil {
		domainSubjectEnum := domain.SubjectEnum(existingSubject.Name)
		domainSubject, err := domain.NewSubject(domainSubjectEnum)
		if err != nil {
			return nil, fmt.Errorf("infra.ent: failed to create domain subject from existing subject: %w", err)
		}
		return &domainSubject, nil
	}

	// Create new subject
	createdSubject, err := r.client.Subject.Create().
		SetName(subjectStr).
		SetCode(subjectStr). // Set code to the same value as name
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to create subject: %w", err)
	}

	// Convert back to domain Subject
	domainSubjectEnum := domain.SubjectEnum(createdSubject.Name)
	domainSubject, err := domain.NewSubject(domainSubjectEnum)
	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to create domain subject from created subject: %w", err)
	}
	return &domainSubject, nil
}

func (r *subjectRepository) RetrieveAll(ctx context.Context) ([]*domain.Subject, error) {
	entSubjects, err := r.client.Subject.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("infra.ent: failed to retrieve all subjects: %w", err)
	}

	domainSubjects := make([]*domain.Subject, 0, len(entSubjects))
	for _, entSubject := range entSubjects {
		domainSubjectEnum := domain.SubjectEnum(entSubject.Name)
		domainSubject, err := domain.NewSubject(domainSubjectEnum)
		if err != nil {
			// Skip invalid subjects but log the error
			fmt.Printf("infra.ent: failed to create domain subject from retrieved subject: %v\n", err)
			continue
		}
		domainSubjects = append(domainSubjects, &domainSubject)
	}

	return domainSubjects, nil
}
