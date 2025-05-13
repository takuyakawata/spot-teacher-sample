package infra_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/user/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/user/infra"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
)

// Using the setupInMemoryClient function from teacher_repository_impl_test.go
// since we're in the same package (infra_test)

func TestAdminUserRepoImpl_Create(t *testing.T) {
	entClient := setupInMemoryClient(t)
	repo := infra.NewAdminUserRepositoryImpl(entClient)

	// ドメインオブジェクトを作成
	adminUser := createTestAdminUser(t)

	// Create 実行
	err := repo.Create(context.Background(), adminUser)
	require.NoError(t, err)

	// ent 側に正しく保存されているか直接確認
	entUser, err := entClient.User.Query().First(context.Background())
	require.NoError(t, err)
	assert.Equal(t, adminUser.FamilyName.Value(), entUser.FamilyName)
	assert.Equal(t, adminUser.FirstName.Value(), entUser.FirstName)
	assert.Equal(t, adminUser.Email.Value(), entUser.Email)
	assert.Equal(t, adminUser.Password.Value(), *entUser.Password)
	assert.Equal(t, "admin", string(entUser.UserType))
}

func TestAdminUserRepositoryImpl_FindByID(t *testing.T) {
	entClient := setupInMemoryClient(t)
	repo := infra.NewAdminUserRepositoryImpl(entClient)

	// テスト用のデータを作成
	familyName := "管理"
	firstName := "太郎"
	email := "admin@example.com"
	password := "password"

	// ent 側に直接データを作成
	entUser, err := entClient.User.Create().
		SetFamilyName(familyName).
		SetFirstName(firstName).
		SetEmail(email).
		SetPassword(password).
		SetPhoneNumber("").
		SetUserType("admin").
		Save(context.Background())
	require.NoError(t, err)

	// FindByID 実行
	adminUserID, err := domain.NewAdminUserID(int64(entUser.ID))
	require.NoError(t, err)
	adminUser, err := repo.FindByID(context.Background(), adminUserID)
	require.NoError(t, err)
	assert.NotNil(t, adminUser)
	assert.Equal(t, familyName, adminUser.FamilyName.Value())
	assert.Equal(t, firstName, adminUser.FirstName.Value())
	assert.Equal(t, email, adminUser.Email.Value())
	assert.Equal(t, password, adminUser.Password.Value())
}

func TestAdminUserRepositoryImpl_FindByID_NotFound(t *testing.T) {
	entClient := setupInMemoryClient(t)
	repo := infra.NewAdminUserRepositoryImpl(entClient)

	// 存在しないIDで検索
	nonExistentID, err := domain.NewAdminUserID(999)
	require.NoError(t, err)
	adminUser, err := repo.FindByID(context.Background(), nonExistentID)
	require.Error(t, err)    // エラーが発生することを確認
	assert.Nil(t, adminUser) // 結果がnilであることを確認
}

func TestAdminUserRepositoryImpl_FindByEmail(t *testing.T) {
	entClient := setupInMemoryClient(t)
	repo := infra.NewAdminUserRepositoryImpl(entClient)

	// テスト用のデータを作成
	familyName := "管理"
	firstName := "太郎"
	email := "admin2@example.com"
	password := "password"

	// ent 側に直接データを作成
	_, err := entClient.User.Create().
		SetFamilyName(familyName).
		SetFirstName(firstName).
		SetEmail(email).
		SetPassword(password).
		SetPhoneNumber("").
		SetUserType("admin").
		Save(context.Background())
	require.NoError(t, err)

	// FindByEmail 実行
	emailAddress, err := sharedDomain.NewEmailAddress(email)
	require.NoError(t, err)
	adminUser, err := repo.FindByEmail(context.Background(), emailAddress)
	require.NoError(t, err)
	assert.NotNil(t, adminUser)
	assert.Equal(t, familyName, adminUser.FamilyName.Value())
	assert.Equal(t, firstName, adminUser.FirstName.Value())
	assert.Equal(t, email, adminUser.Email.Value())
	assert.Equal(t, password, adminUser.Password.Value())
}

func TestAdminUserRepositoryImpl_FindByEmail_NotFound(t *testing.T) {
	entClient := setupInMemoryClient(t)
	repo := infra.NewAdminUserRepositoryImpl(entClient)

	// 存在しないEmailで検索
	nonExistentEmail, err := sharedDomain.NewEmailAddress("nonexistent@example.com")
	require.NoError(t, err)
	adminUser, err := repo.FindByEmail(context.Background(), nonExistentEmail)
	require.Error(t, err)    // エラーが発生することを確認
	assert.Nil(t, adminUser) // 結果がnilであることを確認
}

func createTestAdminUser(t *testing.T) *domain.AdminUser {
	adminUserID, err := domain.NewAdminUserID(1)
	require.NoError(t, err)

	firstName, err := sharedDomain.NewUserName("太郎")
	require.NoError(t, err)

	familyName, err := sharedDomain.NewUserName("管理")
	require.NoError(t, err)

	email, err := sharedDomain.NewEmailAddress("admin@example.com")
	require.NoError(t, err)

	password, err := sharedDomain.NewPassword("password")
	require.NoError(t, err)

	return domain.NewAdminUser(
		adminUserID,
		firstName,
		familyName,
		email,
		password,
	)
}
