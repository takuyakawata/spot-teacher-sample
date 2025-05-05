package infra_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/takuyakawta/spot-teacher-sample/backend/spot-teacher/internal/product/domain"
	"github.com/takuyakawta/spot-teacher-sample/backend/spot-teacher/internal/product/infra"
	"testing"

	"entgo.io/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/takuyakawta/spot-teacher-sample/backend/db/ent"
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

func TestProductRepository_Create(t *testing.T) {
	entClient := setupInMemoryClient(t)
	repo := infra.NewProductRepository(entClient)

	// ドメインオブジェクトを作成
	name := "Notebook"
	desc := "250 pages"
	price := 999
	p := &domain.Product{
		ID:          0, // 作成前はゼロ値
		Name:        domain.ProductName(name),
		Description: &desc,
		Price:       domain.ProductPrice(price),
	}

	// Create 実行
	saved, err := repo.Create(context.Background(), p)
	require.NoError(t, err)
	assert.NotNil(t, saved)
	assert.NotZero(t, saved.ID)
	assert.Equal(t, name, saved.Name.Value())
	assert.Equal(t, price, saved.Price.Value())

	// ent 側に正しく保存されているか直接確認（オプショナル）
	entRec, err := entClient.Product.Get(context.Background(), int(saved.ID))
	require.NoError(t, err)
	assert.Equal(t, name, entRec.Name)
	assert.Equal(t, price, entRec.Price)
	assert.Equal(t, desc, entRec.Description)
}

func TestProductRepository_Update(t *testing.T) {
	// --- 準備 (Provided Code) ---
	entClient := setupInMemoryClient(t)
	repo := infra.NewProductRepository(entClient)
	// // ドメインオブジェクトを作成するための元データ
	initialName := "Notebook"
	initialDesc := "250 pages"
	initialPrice := 999

	// // ★推奨: ドメインコンストラクタを使う場合
	nameVO, err := domain.NewProductName(initialName)
	require.NoError(t, err) // testify を使う場合
	priceVO, err := domain.NewProductPrice(initialPrice)
	require.NoError(t, err)

	p := &domain.Product{
		Name:        nameVO,
		Description: &initialDesc,
		Price:       priceVO,
	}

	// // 初期データをリポジトリ経由で作成
	saved, err := repo.Create(context.Background(), p)
	require.NoError(t, err)
	require.NotNil(t, saved)
	require.NotEqual(t, 0, saved.ID.Value())
	initialID := saved.ID

	// --- Update のテスト ---
	t.Run("正常系: 商品情報を更新できる", func(t *testing.T) {
		// // 1. 更新用データの準備
		updatedName := "Premium Notebook"
		updatedPrice := 1299
		updatedDesc := "300 pages, high quality paper"

		// ★推奨: 更新用データもドメインコンストラクタで生成
		updatedNameVO, err := domain.NewProductName(updatedName)
		require.NoError(t, err)
		updatedPriceVO, err := domain.NewProductPrice(updatedPrice)
		require.NoError(t, err)

		updateData := &domain.Product{ // ★注意: ポインタではなく値型かも (Updateメソッドの引数に合わせる)
			ID:          initialID, // ★更新対象のIDを指定
			Name:        updatedNameVO,
			Description: &updatedDesc,
			Price:       updatedPriceVO,
		}

		// // 2. Update メソッドの実行
		updatedProduct, err := repo.Update(context.Background(), updateData) // ★引数が値型の場合

		// // 3. 結果の検証
		require.NoError(t, err)                                        // Update でエラーが発生しないこと
		require.NotNil(t, updatedProduct)                              // 更新後のオブジェクトが返されること
		require.Equal(t, initialID.Value(), updatedProduct.ID.Value()) // IDが変わっていないこと
		require.Equal(t, updatedName, updatedProduct.Name.Value())     // 名前が更新されていること
		require.NotNil(t, updatedProduct.Description)                  // Description が nil でないこと
		require.Equal(t, updatedDesc, *updatedProduct.Description)     // Description が更新されていること
		require.Equal(t, updatedPrice, updatedProduct.Price.Value())   // 価格が更新されていること

		// // 4. (念のため) DBから再取得して確認
		fetchedAgain, err := repo.FindByID(context.Background(), initialID)
		require.NoError(t, err)
		require.NotNil(t, fetchedAgain)
		require.Equal(t, updatedName, fetchedAgain.Name.Value())
		require.Equal(t, updatedPrice, fetchedAgain.Price.Value())
		require.Equal(t, updatedDesc, *fetchedAgain.Description)
	})

	t.Run("異常系: 存在しないIDを指定するとエラー", func(t *testing.T) {
		// // 存在しないであろうIDを生成 (例: 負のIDやUUIDなど)
		nonExistentID, _ := domain.NewProductID(999)
		nonExistentName, _ := domain.NewProductName("ghost")
		nonExistentPrice, _ := domain.NewProductPrice(0)

		updateData := &domain.Product{
			ID:    nonExistentID,
			Name:  nonExistentName,
			Price: nonExistentPrice,
		}

		// // Update を実行
		_, err := repo.Update(context.Background(), updateData)

		// // エラーが発生することを確認 (特定のエラー型かどうかも確認できるとより良い)
		require.Error(t, err)
	})
}

func TestProductRepository_FindByID(t *testing.T) {
	entClient := setupInMemoryClient(t)
	repo := infra.NewProductRepository(entClient)

	// テスト用のデータを作成
	name := "Notebook"
	desc := "250 pages"
	price := 99
	entProduct, err := entClient.Product.Create().
		SetName(name).
		SetDescription(desc).
		SetPrice(price).
		Save(context.Background())
	require.NoError(t, err)

	// FindByID 実行
	domainID, err := domain.NewProductID(int64(entProduct.ID))
	require.NoError(t, err)
	p, err := repo.FindByID(context.Background(), domainID)
	require.NoError(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, name, p.Name.Value())
	assert.Equal(t, price, p.Price.Value())
	assert.Equal(t, desc, *p.Description)
}

func TestProductRepository_FindAll(t *testing.T) {
	ctx := context.Background()

	//t.Run("データがない場合、空のスライスが返る", func(t *testing.T) {
	//	// --- Execute ---
	//	products, err := repo.FindAll(ctx)
	//
	//	// --- Assert ---
	//	require.NoError(t, err)     // エラーがないこと
	//	require.NotNil(t, products) // スライス自体は nil ではないこと
	//	//require.Empty(t, products)  // スライスが空であること (要素数 0)
	//	require.Len(t, products, 0) // または Len で要素数を確認
	//})

	t.Run("データが複数ある場合、全てのスライスが返る", func(t *testing.T) {
		// --- Setup ---
		entClient := setupInMemoryClient(t)
		repo := infra.NewProductRepository(entClient)

		// テストデータを複数作成
		desc1 := "Description for P1"
		product1 := createTestDomainProduct(t, "Product 1", 100, &desc1)
		product2 := createTestDomainProduct(t, "Product 2", 200, nil) // Descriptionなし
		product3 := createTestDomainProduct(t, "Product 3", 150, nil)

		// 作成したデータを保存 (Createのテストではないのでエラーチェックは簡潔に)
		createdP1, err := repo.Create(ctx, product1)
		require.NoError(t, err)
		createdP2, err := repo.Create(ctx, product2)
		require.NoError(t, err)
		createdP3, err := repo.Create(ctx, product3)
		require.NoError(t, err)

		// 期待される結果のリスト (ID は Create 後に取得)
		expectedProducts := map[int64]*domain.Product{
			createdP1.ID.Value(): createdP1,
			createdP2.ID.Value(): createdP2,
			createdP3.ID.Value(): createdP3,
		}
		expectedCount := len(expectedProducts)

		// --- Execute ---
		actualProducts, err := repo.FindAll(ctx)

		// --- Assert ---
		require.NoError(t, err)                       // エラーがないこと
		require.NotNil(t, actualProducts)             // スライスが nil でないこと
		require.Len(t, actualProducts, expectedCount) // 取得した件数が期待通りであること

		// 内容の検証 (順序が保証されないため、IDでマップを作るなどして比較)
		actualProductsMap := make(map[int64]*domain.Product)
		for _, p := range actualProducts {
			// IDが重複していないかも暗黙的にチェック
			_, exists := actualProductsMap[p.ID.Value()]
			require.False(t, exists, "Duplicate ID found in FindAll result: %v", p.ID)
			actualProductsMap[p.ID.Value()] = p
		}

		// 期待される各商品が実際に取得した結果に含まれ、内容が一致するか確認
		for expectedID, expectedProduct := range expectedProducts {
			actualProduct, found := actualProductsMap[expectedID]
			require.True(t, found, "Expected product with ID %s not found", expectedID)

			// 各フィールドを比較 (Value() などでプリミティブ値を取得して比較)
			require.Equal(t, expectedProduct.Name.Value(), actualProduct.Name.Value(), "Name mismatch for ID %s", expectedID)
			require.Equal(t, expectedProduct.Price.Value(), actualProduct.Price.Value(), "Price mismatch for ID %s", expectedID)
			if expectedProduct.Description == nil {
				require.Nil(t, actualProduct.Description, "Description should be nil for ID %s", expectedID)
			} else {
				require.NotNil(t, actualProduct.Description, "Description should not be nil for ID %s", expectedID)
				require.Equal(t, *expectedProduct.Description, *actualProduct.Description, "Description mismatch for ID %s", expectedID)
			}
			// 他に必要なフィールドがあれば同様に比較
		}
	})
}

func createTestDomainProduct(t *testing.T, name string, price int, desc *string) *domain.Product {
	nameVO, err := domain.NewProductName(name)
	require.NoError(t, err)
	priceVO, err := domain.NewProductPrice(price)
	require.NoError(t, err)

	return &domain.Product{
		// ID は Create で割り当てられる
		Name:        nameVO,
		Description: desc,
		Price:       priceVO,
	}
}
