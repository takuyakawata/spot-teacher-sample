package db

import (
	"fmt" // 文字列フォーマット用
	"log" // ログ出力用
	"os"  // 環境変数アクセス用 (設定の読み込み例)

	"gorm.io/driver/mysql" // GORM MySQLドライバ
	"gorm.io/gorm"         // GORM本体
	// "gorm.io/gorm/logger" // 必要に応じてGORMのロガー設定用
)

// グローバル変数としてDB接続を保持 (他の方法もあります)
var DB *gorm.DB

// InitDB はデータベース接続を初期化する関数
func InitDB() error {
	// --- 接続情報の準備 ---
	// DSN (Data Source Name) を組み立てます。
	// 重要: 実際のアプリケーションでは、パスワードなどをコードに直接書かず、
	// 環境変数や設定ファイルから読み込むようにしてください。
	// 例: os.Getenv("DB_USER") のように環境変数から取得
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "root" // デフォルト値 (開発用)
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "password" // デフォルト値 (開発用) - 必ず変更してください！
	}
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "127.0.0.1" // デフォルト値 (ローカルホスト)
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "3306" // MySQLのデフォルトポート
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "hello_echo_db" // デフォルトのデータベース名 (事前に作成が必要な場合あり)
	}

	// DSN文字列の組み立て (文字コードやタイムゾーン設定も必要に応じて追加)
	// user:pass@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	// --- GORMを使ってデータベースに接続 ---
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 必要に応じてGORMのロガー設定などを追加
		// Logger: logger.Default.LogMode(logger.Info), // SQLログを出力する場合
	})

	if err != nil {
		log.Printf("データベースへの接続に失敗しました: %v\n", err)
		return fmt.Errorf("データベース接続エラー: %w", err)
	}

	log.Println("データベース接続成功！")

	// (オプション) 接続プールなどの設定
	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("DBインスタンスの取得に失敗しました: %v\n", err)
		return err
	}
	// sqlDB.SetMaxIdleConns(10) // アイドル状態の接続数
	// sqlDB.SetMaxOpenConns(100) // 最大接続数
	// sqlDB.SetConnMaxLifetime(time.Hour) // 接続の最大生存期間

	return nil
}

// GetDB は初期化されたDB接続インスタンスを返す関数 (例)
func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("データベース接続が初期化されていません。InitDB()を呼び出してください。")
	}
	return DB
}

// (オプション) モデルのマイグレーション関数
// func AutoMigrateModels() error {
//  if DB == nil {
// 		return fmt.Errorf("データベース接続が初期化されていません")
// 	}
// 	// テーブルを作成したいモデルの構造体を渡す
// 	err := DB.AutoMigrate(&YourModelStruct1{}, &YourModelStruct2{})
// 	if err != nil {
// 		log.Printf("モデルのマイグレーションに失敗しました: %v\n", err)
// 		return err
// 	}
// 	log.Println("モデルのマイグレーション成功！")
// 	return nil
// }
