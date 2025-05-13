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
	userDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/user/domain"
)

func TestSchoolUsecase_DeleteSchool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックリポジトリ生成
	mockSchoolRepo := mock.NewMockSchoolRepository(ctrl)
	mockTeacherRepo := mock.NewMockTeacherRepository(ctrl)

	// テストデータ
	schoolID := domain.SchoolID(1)
	school := &domain.School{ID: schoolID}

	// 成功ケース：FindByID, FindBySchoolID, Delete の期待設定
	mockSchoolRepo.EXPECT().
		FindByID(gomock.Any(), schoolID).
		Return(school, nil)

	mockTeacherRepo.EXPECT().
		FindBySchoolID(gomock.Any(), schoolID).
		Return([]*userDomain.Teacher{}, nil)

	mockSchoolRepo.EXPECT().
		Delete(gomock.Any(), schoolID).
		Return(nil)

	uc := usecase.NewDeleteSchoolUseCase(mockSchoolRepo, mockTeacherRepo)
	err := uc.DeleteSchool(context.Background(), schoolID)

	assert.NoError(t, err)
}

func TestSchoolUsecase_DeleteSchool_WithTeachers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックリポジトリ生成
	mockSchoolRepo := mock.NewMockSchoolRepository(ctrl)
	mockTeacherRepo := mock.NewMockTeacherRepository(ctrl)

	// テストデータ
	schoolID := domain.SchoolID(1)
	school := &domain.School{ID: schoolID}
	teachers := []*userDomain.Teacher{
		{ID: userDomain.TeacherID(1), SchoolID: schoolID},
	}

	// 失敗ケース：学校に関連する教師がいる場合
	mockSchoolRepo.EXPECT().
		FindByID(gomock.Any(), schoolID).
		Return(school, nil)

	mockTeacherRepo.EXPECT().
		FindBySchoolID(gomock.Any(), schoolID).
		Return(teachers, nil)

	uc := usecase.NewDeleteSchoolUseCase(mockSchoolRepo, mockTeacherRepo)
	err := uc.DeleteSchool(context.Background(), schoolID)

	assert.Error(t, err)
	assert.Equal(t, "cannot delete school with associated teachers", err.Error())
}

func TestSchoolUsecase_DeleteSchool_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックリポジトリ生成
	mockSchoolRepo := mock.NewMockSchoolRepository(ctrl)
	mockTeacherRepo := mock.NewMockTeacherRepository(ctrl)

	// テストデータ
	schoolID := domain.SchoolID(1)
	notFoundErr := errors.New("school not found")

	// 失敗ケース：学校が見つからない場合
	mockSchoolRepo.EXPECT().
		FindByID(gomock.Any(), schoolID).
		Return(nil, notFoundErr)

	uc := usecase.NewDeleteSchoolUseCase(mockSchoolRepo, mockTeacherRepo)
	err := uc.DeleteSchool(context.Background(), schoolID)

	assert.Error(t, err)
	assert.Equal(t, notFoundErr, err)
}
