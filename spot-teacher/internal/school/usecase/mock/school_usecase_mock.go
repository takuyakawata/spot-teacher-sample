package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
)

// SchoolUsecaseMock is a mock implementation of the SchoolUsecase interface
type SchoolUsecaseMock struct {
	mock.Mock
}

// ListSchools mocks the ListSchools method
func (m *SchoolUsecaseMock) ListSchools(ctx context.Context) ([]*domain.School, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.School), args.Error(1)
}

// GetSchool mocks the GetSchool method
func (m *SchoolUsecaseMock) GetSchool(ctx context.Context, id domain.SchoolID) (*domain.School, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.School), args.Error(1)
}

// CreateSchool mocks the CreateSchool method
func (m *SchoolUsecaseMock) CreateSchool(ctx context.Context, school *domain.School) (*domain.School, error) {
	args := m.Called(ctx, school)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.School), args.Error(1)
}

// UpdateSchool mocks the UpdateSchool method
func (m *SchoolUsecaseMock) UpdateSchool(ctx context.Context, school *domain.School) (*domain.School, error) {
	args := m.Called(ctx, school)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.School), args.Error(1)
}

// DeleteSchool mocks the DeleteSchool method
func (m *SchoolUsecaseMock) DeleteSchool(ctx context.Context, id domain.SchoolID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
