package main

import (
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/product/inject"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/router"

	"net/http"
	"sync" // 並行処理での競合を防ぐために使用

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Echo インスタンスは、Serverless Function のライフサイクル内で再利用したい場合があります。
// コールドスタート時に一度だけ初期化されるようにします。
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

// ローカル開発用の main 関数は Vercel デプロイ時には不要です。
// cmd/main.go で Echo サーバー全体を起動してローカル開発することも可能ですし、
// Vercel CLI (vercel dev) を使ってこの api/index.go をテストすることも可能です。
/*
func main() {
	// Vercel CLI でローカルテストする場合、この main 関数は実行されません。
	// 標準的な Go サーバーとして起動したい場合は cmd/main.go を実行してください。
	// もし Vercel CLI を使わず、このファイルを単体でテストしたい場合は、
	// initEcho() を呼び出し、http.ListenAndServe で Handler をラップして起動します。
	// 例:
	// initEcho()
	// port := ":8080"
	// fmt.Printf("Local server listening on %s\n", port)
	// if err := http.ListenAndServe(port, http.HandlerFunc(Handler)); err != nil {
	// 	log.Fatalf("Error starting server: %v", err)
	// }
}
*/

// main 関数は Go プログラムの開始地点です。
func main() {
	//	// --- 1. Echo インスタンスの作成 ---
	//	// まず、Echo フレームワークの本体となるインスタンスを作成します。
	//	// 慣習的に 'e' という変数名がよく使われます。
	e := echo.New()
	//
	//	// --- 2. ミドルウェアの設定 (推奨) ---
	//	// これらは必須ではありませんが、開発や運用に非常に役立ちます。
	//
	//	// Logger ミドルウェア:
	//	// サーバーへのすべてのリクエストに関する情報（メソッド、パス、ステータスコード、処理時間など）を
	//	// ターミナル (標準出力) にログとして記録します。デバッグに便利です。
	e.Use(middleware.Logger())
	//
	//	// Recover ミドルウェア:
	//	// リクエスト処理中に予期せぬエラー（パニック）が発生した場合に、
	//	// サーバー全体を停止させる代わりに、自動的に回復処理を行い、
	//	// クライアントには HTTP 500 Internal Server Error を返します。サーバーの安定性が向上します。
	e.Use(middleware.Recover())
	//
	//	// --- 3. ルートの定義 ---
	//	// ここで、特定のURLパスとHTTPメソッドに対する処理を結びつけます。
	//
	//	// e.GET("/", ...) は、
	//	// HTTP の GET メソッドで、ルートパス ("/") にアクセスがあった場合に、
	//	// 後ろに続く関数 (ハンドラー関数) を実行するように Echo に指示します。
	//	//e.GET("/", func(c echo.Context) error {
	//	//	// この関数が、実際のリクエストを処理するハンドラーです。
	//	//	// 引数の 'c' (echo.Context) は、リクエスト情報へのアクセスや、
	//	//	// レスポンスをクライアントに送信するためのメソッドを提供します。
	//	//
	//	//	// c.String(ステータスコード, 返す文字列) を使ってレスポンスを返します。
	//	//	// http.StatusOK は、HTTPステータスコード 200 (成功) を表す定数です。
	//	//	message := "Hello, World!" // 返すメッセージを変数に入れてみました
	//	return c.String(http.StatusOK, message)
	//	//})
	//
	//	/* DI で各ハンドラセットを生成 routing */
	prodH := inject.InitializeProductHandler()
	router.RegisterAll(e, prodH)
	//
	//	// --- 4. サーバーの起動 ---
	//	// 作成した Echo インスタンスに、指定したポート番号で HTTP リクエストの待機を開始させます。
	//	// ":1323" はポート番号 1323 でリクエストを待つことを意味します。
	//	// ポート番号は、他のプログラムが使用していなければ基本的に自由ですが、1024未満は特別な権限が必要な場合があります。
	//	// 1323 は Echo のドキュメントでよく使われる例です。
	port := ":8080"
	e.Logger.Infof("Starting server on port %s", port) // 起動ポートをログに出力 (オプション)
	//
	//	// e.Start(ポート番号) でサーバーが起動します。
	//	// この関数は処理をブロックし、サーバーが動き続ける間は戻ってきません。
	//	// もしサーバーの起動に失敗した場合（例：ポートが既に使用中など）、エラーが返されます。
	//	// e.Logger.Fatal() は、エラーが発生した場合にログを出力し、プログラムを終了させます。
	err := e.Start(port)
	if err != nil {
		e.Logger.Fatal(err)
	}
}
