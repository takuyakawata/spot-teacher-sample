package infra_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/user/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/user/infra"
	"testing"

	"entgo.io/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/takuyakawta/spot-teacher-sample/db/ent"
	schoolDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/school/domain"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
)

func setupInMemoryClient(t *testing.T) *ent.Client {
	// SQLite のインメモリ DSN
	drv, err := sql.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)

	c := ent.NewClient(ent.Driver(drv))
	// スキーマ作成
	require.NoError(t, c.Schema.Create(context.Background()))
	return c
}

func TestTeacherRepoImpl_Create(t *testing.T) {
	entClient := setupInMemoryClient(t)
	repo := infra.NewTeacherRepoImpl(entClient)

	// ドメインオブジェクトを作成
	teacher := createTestTeacher(t)

	// Create 実行
	err := repo.Create(context.Background(), teacher)
	require.NoError(t, err)

	// ent 側に正しく保存されているか直接確認
	entUser, err := entClient.User.Query().First(context.Background())
	require.NoError(t, err)
	assert.Equal(t, teacher.FamilyName.Value(), entUser.FamilyName)
	assert.Equal(t, teacher.FirstName.Value(), entUser.FirstName)
	assert.Equal(t, teacher.Email.Value(), entUser.Email)
	assert.Equal(t, teacher.Password.Value(), entUser.Password)
}

func TestTeacherRepoImpl_FindByID(t *testing.T) {
	entClient := setupInMemoryClient(t)
	repo := infra.NewTeacherRepoImpl(entClient)

	// テスト用のデータを作成
	familyName := "山田"
	firstName := "太郎"
	email := "yamada2@example.com"
	password := "password"

	// ent 側に直接データを作成
	entUser, err := entClient.User.Create().
		SetFamilyName(familyName).
		SetFirstName(firstName).
		SetEmail(email).
		SetPassword(password).
		Save(context.Background())
	require.NoError(t, err)

	// FindByID 実行
	teacherID, err := domain.NewTeacherID(int64(entUser.ID))
	require.NoError(t, err)
	teacher, err := repo.FindByID(context.Background(), teacherID)
	require.NoError(t, err)
	assert.NotNil(t, teacher)
	assert.Equal(t, familyName, teacher.FamilyName.Value())
	assert.Equal(t, firstName, teacher.FirstName.Value())
	assert.Equal(t, email, teacher.Email.Value())
	assert.Equal(t, password, teacher.Password.Value())
}

func TestTeacherRepoImpl_FindByID_NotFound(t *testing.T) {
	entClient := setupInMemoryClient(t)
	repo := infra.NewTeacherRepoImpl(entClient)

	// 存在しないIDで検索
	nonExistentID, err := domain.NewTeacherID(999)
	require.NoError(t, err)
	teacher, err := repo.FindByID(context.Background(), nonExistentID)
	require.Error(t, err)  // エラーが発生することを確認
	assert.Nil(t, teacher) // 結果がnilであることを確認
}

func createTestTeacher(t *testing.T) *domain.Teacher {
	schoolID, err := schoolDomain.NewSchoolID(1)
	require.NoError(t, err)

	firstName, err := domain.NewTeacherName("太郎")
	require.NoError(t, err)

	familyName, err := domain.NewTeacherName("山田")
	require.NoError(t, err)

	email, err := sharedDomain.NewEmailAddress("yamada@example.com")
	require.NoError(t, err)

	password, err := sharedDomain.NewPassword("password")
	require.NoError(t, err)

	return &domain.Teacher{
		SchoolID:   schoolID,
		FirstName:  firstName,
		FamilyName: familyName,
		Email:      email,
		Password:   password,
	}
}
