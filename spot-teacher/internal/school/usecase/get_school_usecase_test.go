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

func TestSchoolUseCase_GetSchool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックリポジトリ生成
	mockSchoolRepo := mock.NewMockSchoolRepository(ctrl)

	// テストデータ
	schoolID := domain.SchoolID(1)
	school := &domain.School{ID: schoolID}

	// 成功ケース：FindByID の期待設定
	mockSchoolRepo.EXPECT().
		FindByID(gomock.Any(), schoolID).
		Return(school, nil)

	uc := usecase.NewGetSchoolUseCase(mockSchoolRepo)
	got, err := uc.GetSchool(context.Background(), schoolID)

	assert.NoError(t, err)
	assert.Equal(t, school, got)
}
