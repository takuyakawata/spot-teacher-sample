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

func TestSchoolUsecase_CreateSchool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックリポジトリ生成
	mockSchoolRepo := mock.NewMockSchoolRepository(ctrl)

	// テストデータ
	school := &domain.School{
		Name: domain.SchoolName("Test School"),
	}

	// 成功ケース：FindByName と Create の期待設定
	mockSchoolRepo.EXPECT().
		FindByName(gomock.Any(), school.Name).
		Return(nil, nil)

	mockSchoolRepo.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, s *domain.School) (*domain.School, error) {
			// 実際の実装同様に ID を入れて返す
			s.ID = domain.SchoolID(1)
			return s, nil
		})

	uc := usecase.NewCreateSchoolUseCase(mockSchoolRepo)
	got, err := uc.CreateSchool(context.Background(), school)

	assert.NoError(t, err)
	assert.Equal(t, domain.SchoolID(1), got.ID)
	assert.Equal(t, domain.SchoolName("Test School"), got.Name)
}
