package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/product/inject"
	"github.com/takuyakawta/spot-teacher-sample/spot-teacher/internal/router"
	"log"
	"net/http"
	"sync"
)

var (
	initializedEcho *echo.Echo
	echoOnce        sync.Once
)

func initEcho() {
	echoOnce.Do(func() {
		// --- 1. Echo インスタンスの作成 ---
		log.Println("api/index.go: init() started")
		e := echo.New()

		// --- 2. ミドルウェアの設定 ---
		e.Use(middleware.Recover())
		e.Use(middleware.Logger())

		// --- 3. DI によるハンドラーの生成 ---
		prodH := inject.InitializeProductHandler()

		// --- 4. ルーティングの定義 ---
		router.RegisterAll(e, prodH)
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("api/index.go: Handler() started")
	initEcho()
	if initializedEcho == nil {
		log.Println("api/index.go: Error: initializedEcho is nil in Handler!")
		http.Error(w, "Internal Server Error: API not initialized", http.StatusInternalServerError)
		return
	}

	log.Printf("api/index.go: Serving request for path: %s", r.URL.Path)
	initializedEcho.ServeHTTP(w, r)
}
