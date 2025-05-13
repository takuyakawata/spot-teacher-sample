package infra_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/company/domain"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/company/infra"
	sharedDomain "github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/shared/domain"
	"testing"

	"entgo.io/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/takuyakawta/spot-teacher-sample/db/ent"
)

func setupInMemoryClient(t *testing.T) *ent.Client {
	// SQLite のインメモリ DSN - use unique database name for each test to ensure isolation
	dbName := "file:ent_" + t.Name() + "?mode=memory&_fk=1"
	drv, err := sql.Open("sqlite3", dbName)
	require.NoError(t, err)

	c := ent.NewClient(ent.Driver(drv))
	// スキーマ作成
	require.NoError(t, c.Schema.Create(context.Background()))
	return c
}

func TestCompanyRepository_Create(t *testing.T) {
	entClient := setupInMemoryClient(t)
	repo := infra.NewCompanyRepository(entClient)

	// ドメインオブジェクトを作成
	companyName, err := domain.NewCompanyName("Test Company")
	require.NoError(t, err)

	postCode, err := sharedDomain.NewPostCode("1234567")
	require.NoError(t, err)

	street := "1-2-3 Test Street"
	streetPtr := &street

	address := sharedDomain.Address{
		Prefecture: sharedDomain.PrefectureTokyo,
		City:       "Test City",
		Street:     streetPtr,
		PostCode:   postCode,
	}

	phoneNumber := sharedDomain.PhoneNumber("03-1234-5678")

	company := &domain.Company{
		ID:          0, // 作成前はゼロ値
		Name:        companyName,
		Address:     address,
		PhoneNumber: phoneNumber,
		URL:         "",
	}

	// Create 実行
	saved, err := repo.Create(context.Background(), company)
	require.NoError(t, err)
	assert.NotNil(t, saved)
	assert.NotZero(t, saved.ID)
	assert.Equal(t, "Test Company", saved.Name.Value())
	assert.Equal(t, sharedDomain.PrefectureTokyo, saved.Address.Prefecture)
	assert.Equal(t, "Test City", saved.Address.City)
	assert.Equal(t, "1-2-3 Test Street", *saved.Address.Street)
	assert.Equal(t, "1234567", saved.Address.PostCode.Value())
	assert.Equal(t, "03-1234-5678", saved.PhoneNumber.Value())

	// ent 側に正しく保存されているか直接確認
	entRec, err := entClient.Company.Get(context.Background(), int(saved.ID))
	require.NoError(t, err)
	assert.Equal(t, "Test Company", entRec.Name)
	assert.Equal(t, int(sharedDomain.PrefectureTokyo), entRec.Prefecture)
	assert.Equal(t, "Test City", entRec.City)
	assert.Equal(t, "1-2-3 Test Street", entRec.Street)
	assert.Equal(t, "1234567", entRec.PostCode)
	assert.Equal(t, "03-1234-5678", entRec.PhoneNumber)
}

func TestCompanyRepository_Update(t *testing.T) {
	entClient := setupInMemoryClient(t)
	repo := infra.NewCompanyRepository(entClient)

	// 初期データの準備
	companyName, err := domain.NewCompanyName("Initial Company")
	require.NoError(t, err)

	postCode, err := sharedDomain.NewPostCode("1234567")
	require.NoError(t, err)

	street := "1-2-3 Initial Street"
	streetPtr := &street

	address := sharedDomain.Address{
		Prefecture: sharedDomain.PrefectureTokyo,
		City:       "Initial City",
		Street:     streetPtr,
		PostCode:   postCode,
	}

	phoneNumber := sharedDomain.PhoneNumber("03-1234-5678")

	company := &domain.Company{
		ID:          0, // 作成前はゼロ値
		Name:        companyName,
		Address:     address,
		PhoneNumber: phoneNumber,
		URL:         "",
	}

	// 初期データをリポジトリ経由で作成
	saved, err := repo.Create(context.Background(), company)
	require.NoError(t, err)
	require.NotNil(t, saved)
	require.NotEqual(t, 0, saved.ID.Value())
	initialID := saved.ID

	// Update のテスト
	t.Run("正常系: 会社情報を更新できる", func(t *testing.T) {
		// 更新用データの準備
		updatedName, err := domain.NewCompanyName("Updated Company")
		require.NoError(t, err)

		updatedPostCode, err := sharedDomain.NewPostCode("7654321")
		require.NoError(t, err)

		updatedStreet := "4-5-6 Updated Street"
		updatedStreetPtr := &updatedStreet

		updatedAddress := sharedDomain.Address{
			Prefecture: sharedDomain.PrefectureOsaka,
			City:       "Updated City",
			Street:     updatedStreetPtr,
			PostCode:   updatedPostCode,
		}

		updatedPhoneNumber := sharedDomain.PhoneNumber("06-8765-4321")

		updateData := &domain.Company{
			ID:          initialID,
			Name:        updatedName,
			Address:     updatedAddress,
			PhoneNumber: updatedPhoneNumber,
			URL:         "",
		}

		// Update メソッドの実行
		updatedCompany, err := repo.Update(context.Background(), updateData)

		// 結果の検証
		require.NoError(t, err)
		require.NotNil(t, updatedCompany)
		require.Equal(t, initialID.Value(), updatedCompany.ID.Value())
		require.Equal(t, "Updated Company", updatedCompany.Name.Value())
		require.Equal(t, sharedDomain.PrefectureOsaka, updatedCompany.Address.Prefecture)
		require.Equal(t, "Updated City", updatedCompany.Address.City)
		require.NotNil(t, updatedCompany.Address.Street)
		require.Equal(t, "4-5-6 Updated Street", *updatedCompany.Address.Street)
		require.Equal(t, "7654321", updatedCompany.Address.PostCode.Value())
		require.Equal(t, "06-8765-4321", updatedCompany.PhoneNumber.Value())

		// DBから再取得して確認
		fetchedAgain, err := repo.FindByID(context.Background(), initialID)
		require.NoError(t, err)
		require.NotNil(t, fetchedAgain)
		require.Equal(t, "Updated Company", fetchedAgain.Name.Value())
		require.Equal(t, sharedDomain.PrefectureOsaka, fetchedAgain.Address.Prefecture)
		require.Equal(t, "Updated City", fetchedAgain.Address.City)
		require.Equal(t, "4-5-6 Updated Street", *fetchedAgain.Address.Street)
		require.Equal(t, "7654321", fetchedAgain.Address.PostCode.Value())
		require.Equal(t, "06-8765-4321", fetchedAgain.PhoneNumber.Value())
	})
}

func TestCompanyRepository_FindByID(t *testing.T) {
	entClient := setupInMemoryClient(t)
	repo := infra.NewCompanyRepository(entClient)

	// テスト用のデータを作成
	company := createTestCompany(t, "Test Company", sharedDomain.PrefectureTokyo, "Test City")
	saved, err := repo.Create(context.Background(), company)
	require.NoError(t, err)
	require.NotNil(t, saved)

	// FindByID のテスト
	found, err := repo.FindByID(context.Background(), saved.ID)
	require.NoError(t, err)
	require.NotNil(t, found)
	assert.Equal(t, saved.ID.Value(), found.ID.Value())
	assert.Equal(t, saved.Name.Value(), found.Name.Value())
	assert.Equal(t, saved.Address.Prefecture, found.Address.Prefecture)
	assert.Equal(t, saved.Address.City, found.Address.City)
	assert.Equal(t, saved.PhoneNumber.Value(), found.PhoneNumber.Value())
}

func TestCompanyRepository_FindAll(t *testing.T) {
	entClient := setupInMemoryClient(t)
	repo := infra.NewCompanyRepository(entClient)

	// 最初は空のはず
	companies, err := repo.FindAll(context.Background())
	require.NoError(t, err)
	assert.Empty(t, companies)

	// テスト用のデータを複数作成
	company1 := createTestCompany(t, "Company 1", sharedDomain.PrefectureTokyo, "City 1")
	company2 := createTestCompany(t, "Company 2", sharedDomain.PrefectureOsaka, "City 2")
	company3 := createTestCompany(t, "Company 3", sharedDomain.PrefectureKyoto, "City 3")

	_, err = repo.Create(context.Background(), company1)
	require.NoError(t, err)
	_, err = repo.Create(context.Background(), company2)
	require.NoError(t, err)
	_, err = repo.Create(context.Background(), company3)
	require.NoError(t, err)

	// FindAll のテスト
	allCompanies, err := repo.FindAll(context.Background())
	require.NoError(t, err)
	assert.Len(t, allCompanies, 3)
}

func TestCompanyRepository_Delete(t *testing.T) {
	entClient := setupInMemoryClient(t)
	repo := infra.NewCompanyRepository(entClient)

	// テスト用のデータを作成
	company := createTestCompany(t, "Company to Delete", sharedDomain.PrefectureTokyo, "Delete City")
	saved, err := repo.Create(context.Background(), company)
	require.NoError(t, err)
	require.NotNil(t, saved)

	// 削除前に存在確認
	found, err := repo.FindByID(context.Background(), saved.ID)
	require.NoError(t, err)
	require.NotNil(t, found)

	// Delete のテスト
	err = repo.Delete(context.Background(), saved.ID)
	require.NoError(t, err)

	// 削除後に存在確認（エラーになるはず）
	_, err = repo.FindByID(context.Background(), saved.ID)
	require.Error(t, err)
}

func createTestCompany(t *testing.T, name string, prefecture sharedDomain.Prefecture, city string) *domain.Company {
	companyName, err := domain.NewCompanyName(name)
	require.NoError(t, err)

	postCode, err := sharedDomain.NewPostCode("1234567")
	require.NoError(t, err)

	street := "1-2-3 Test Street"
	streetPtr := &street

	address := sharedDomain.Address{
		Prefecture: prefecture,
		City:       city,
		Street:     streetPtr,
		PostCode:   postCode,
	}

	phoneNumber := sharedDomain.PhoneNumber("03-1234-5678")

	return &domain.Company{
		ID:          0, // 作成前はゼロ値
		Name:        companyName,
		Address:     address,
		PhoneNumber: phoneNumber,
		URL:         "",
	}
}
