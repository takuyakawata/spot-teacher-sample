package main

import (
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/product/inject"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/router"

	"net/http"
	"sync" // 並行処理での競合を防ぐために使用

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	echoOnce sync.Once
	e        *echo.Echo
)

// initEcho は Echo インスタンスと必要な設定を初期化します。
// sync.Once を使って、複数のリクエストが同時に来ても一度だけ実行されるようにします。
func initEcho() {
	echoOnce.Do(func() {
		// --- 1. Echo インスタンスの作成 ---
		e = echo.New()

		// --- 2. ミドルウェアの設定 ---
		// Serverless 環境では全てのミドルウェアが必要とは限りません。
		// ロガーは Vercel の Functions ログに出力されます。
		// Recover ミドルウェアはパニックからの回復に有用です。
		// その他のミドルウェア (CORS, BodyLimit など) は必要に応じて設定してください。
		e.Use(middleware.Recover())
		// e.Use(middleware.Logger()) // 必要なら有効化

		// --- 3. DI によるハンドラーの生成 ---
		// cmd/main.go と同様に DI を行います。
		// エラーハンドリングも適切に追加してください。
		prodH := inject.InitializeProductHandler()
		// err != nil { e.Logger.Fatal(err) } のようなエラー処理は Serverless Function では適切ではありません。
		// 初期化に失敗した場合は、関数の実行を中断するか、エラーレスポンスを返すなどの対応が必要です。
		// 簡単のためここではエラー処理を省略しますが、プロダクションでは必須です。

		// --- 4. ルーティングの定義 ---
		// cmd/main.go と同様にルーティングを登録します。
		// ただし、Vercel の Serverless Function が `/api` のようなパスにデプロイされる場合、
		// 元の `/products` のようなルートは Vercel 上では `/api/products` となります。
		// router.RegisterAll が全てのルートを定義しているとして、そのまま登録します。
		// 必要であれば、ここで `/api` グループを作成してその中にルートを登録するなど、
		// Vercel のデプロイパスに合わせてルーティングを調整してください。
		// apiGroup := e.Group("/api")
		// router.RegisterRoutesForApiGroup(apiGroup, prodH) // router パッケージにグループを受け取る関数を追加するなど
		router.RegisterAll(e, prodH) // この例ではそのまま登録
	})
}

// Handler is the main entry point for the Vercel Serverless Function.
// It initializes the Echo instance (if not already initialized) and
// forwards the incoming request to it.
func Handler(w http.ResponseWriter, r *http.Request) {
	// Echo インスタンスを初期化します（初回リクエスト時のみ）
	initEcho()

	// 受信した HTTP リクエストを初期化済みの Echo インスタンスに処理させます。
	// Echo の ServeHTTP メソッドが、ルーティングとハンドラーの実行を行います。
	e.ServeHTTP(w, r)
}

// main 関数は Go プログラムの開始地点です。
func main() {
	// --- 1. Echo インスタンスの作成 ---
	e := echo.New()

	// --- 2. ミドルウェアの設定 (推奨) ---
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// --- 3. ルートの定義 ---
	e.GET("/healthz", func(c echo.Context) error {
		// ステータスコード 200 OK と "OK" という文字列を返す
		return c.String(http.StatusOK, "OK")
	})
	// ここで、特定のURLパスとHTTPメソッドに対する処理を結びつけます
	/* DI で各ハンドラセットを生成 routing */
	prodH := inject.InitializeProductHandler()
	router.RegisterAll(e, prodH)
	/// --- 4. サーバーの起動 ---
	port := ":8080"
	e.Logger.Infof("Starting server on port %s", port)

	err := e.Start(port)
	if err != nil {
		e.Logger.Fatal(err)
	}
}
