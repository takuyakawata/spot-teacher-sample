package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/usecase"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/usecase/mock"
)

func TestSchoolUseCase_UpdateSchool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックリポジトリ生成
	mockSchoolRepo := mock.NewMockSchoolRepository(ctrl)

	// テストデータ
	school := &domain.School{
		ID:   domain.SchoolID(1),
		Name: domain.SchoolName("Updated School"),
	}

	// 成功ケース：FindByID と Update の期待設定
	mockSchoolRepo.EXPECT().
		FindByID(gomock.Any(), school.ID).
		Return(school, nil)

	mockSchoolRepo.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, s *domain.School) (*domain.School, error) {
			return s, nil
		})

	uc := usecase.NewUpdateSchoolUseCase(mockSchoolRepo)
	got, err := uc.UpdateSchool(context.Background(), school)

	assert.NoError(t, err)
	assert.Equal(t, school.ID, got.ID)
	assert.Equal(t, school.Name, got.Name)
}

func TestSchoolUseCase_UpdateSchool_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックリポジトリ生成
	mockSchoolRepo := mock.NewMockSchoolRepository(ctrl)

	// テストデータ
	school := &domain.School{
		ID:   domain.SchoolID(1),
		Name: domain.SchoolName("Updated School"),
	}
	notFoundErr := errors.New("school not found")

	// 失敗ケース：FindByID でエラー
	mockSchoolRepo.EXPECT().
		FindByID(gomock.Any(), school.ID).
		Return(nil, notFoundErr)

	uc := usecase.NewUpdateSchoolUseCase(mockSchoolRepo)
	got, err := uc.UpdateSchool(context.Background(), school)

	assert.Error(t, err)
	assert.Equal(t, notFoundErr, err)
	assert.Nil(t, got)
}
