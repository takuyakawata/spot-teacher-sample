package usecase_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/company/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/company/usecase"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/company/usecase/mock"
	"testing"
)

func TestCreateCompanyUseCase_CreateCompany(t *testing.T) {
	// テストデータの準備
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockCompanyRepository(ctrl)

	company := &domain.Company{
		Name: domain.CompanyName("Test Company"),
	}

	mockRepo.EXPECT().
		Create(gomock.Any(), gomock.Eq(company)).
		Return(company, nil).
		Times(1)

	uc := usecase.NewCreateCompanyUseCase(mockRepo)
	err := uc.CreateCompany(context.Background(), company)

	assert.NoError(t, err)
}
