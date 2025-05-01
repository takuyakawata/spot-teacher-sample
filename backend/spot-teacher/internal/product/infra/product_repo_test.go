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
	price := 9.99
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
	assert.Equal(t, price, float64(saved.Price))

	// ent 側に正しく保存されているか直接確認（オプショナル）
	entRec, err := entClient.Product.Get(context.Background(), int(saved.ID))
	require.NoError(t, err)
	assert.Equal(t, name, entRec.Name)
	assert.Equal(t, price, entRec.Price)
	assert.Equal(t, desc, entRec.Description)
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
		SetPrice(float64(price)).
		Save(context.Background())
	require.NoError(t, err)

	// FindByID 実行
	domainID, err := domain.NewProductID(int64(entProduct.ID))
	require.NoError(t, err)
	p, err := repo.FindByID(context.Background(), domainID)
	require.NoError(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, name, p.Name.Value())
	assert.Equal(t, price, float64(p.Price))
	assert.Equal(t, desc, *p.Description)
}
