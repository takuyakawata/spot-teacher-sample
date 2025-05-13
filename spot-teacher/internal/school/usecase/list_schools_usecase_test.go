package usecase_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/usecase"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/usecase/mock"
)

func TestSchoolUsecase_ListSchools(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックリポジトリ生成
	mockSchoolRepo := mock.NewMockSchoolRepository(ctrl)

	// テストデータ
	schools := []*domain.School{
		{ID: domain.SchoolID(1)},
		{ID: domain.SchoolID(2)},
	}

	// 成功ケース：FindAll の期待設定
	mockSchoolRepo.EXPECT().
		FindAll(gomock.Any()).
		Return(schools, nil)

	uc := usecase.NewListSchoolsUseCase(mockSchoolRepo)
	got, err := uc.ListSchools(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, schools, got)
}
