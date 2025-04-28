# Project Guidelines

ディレクトリ構成
```
├── cmd/
│   └── app/                 エントリポイント（main関数など）、WireによるDI設定もここ
│       └── main.go
├── internal/                アプリケーション内部のパッケージ（※外部からimportされない想定）
│   ├── domain/              ドメイン層：エンティティやリポジトリインターフェース等&#8203;:contentReference[oaicite:48]{index=48}
│   │   ├── entities/        ドメインエンティティ（集約ごとにサブパッケージに分けることも）
│   │   ├── values/          値オブジェクトや汎用ドメイン型
│   │   └── repositories/    リポジトリのインターフェース定義
│   ├── usecase/             ユースケース層：アプリケーションのビジネスロジック&#8203;:contentReference[oaicite:49]{index=49}
│   │   ├── interactor/      ユースケースの実装（入力を受け取りドメイン/インフラを呼び出す）
│   │   └── dto/             ユースケース間で受け渡すDTOやレスポンスモデル
│   ├── interfaces/          インターフェース層：外部とのI/O（コントローラやプレゼンター）&#8203;:contentReference[oaicite:50]{index=50}
│   │   ├── controllers/     WebAPIのコントローラ（Echoのハンドラなど）
│   │   └── presenters/      UIへのプレゼンテーション変換（RESTではJSONにシリアライズ等）
│   └── infrastructure/      インフラ層：DBや外部サービス接続の実装&#8203;:contentReference[oaicite:51]{index=51}
│       ├── database/        DB接続やトランザクション管理（例：sql.DB保持や設定）
│       ├── persistence/     リポジトリ実装（例：GORMやsqlxを使った実装）&#8203;:contentReference[oaicite:52]{index=52}
│       ├── logger/          ロギングの実装（Zapのラッパーなど）
│       └── router/          ルーティングの設定（Echoのルート定義など）
├── config/                  設定ファイル（例：config.yaml や .env ファイル）
├── migrations/              マイグレーションSQLファイル置き場
├── pkg/                     外部公開可能な汎用パッケージ（必要なら）
└── go.mod, go.sum           Goモジュール定義（依存関係）
```


# Go Backend Project README (仮)

## プロジェクト概要
Go 1.23 / REST API サーバー。
クリーンアーキテクチャ + DDD思想でレイヤーを明確に分離して開発。
Docker Composeでローカル環境を構築し、デプロイはVercel（メイン）/AWS（検証）を想定。

---

## 使用技術スタック
- 言語: Go 1.23
- Webフレームワーク: Echo
- ルーティング: Echo内蔵Router
- DI: Google Wire
- バリデーション: go-playground/validator
- ロギング: Uber Zap
- コンフィグ管理: Viper
- ORM: GORM (MySQL)
- マイグレーションツール: golang-migrate
- テスト: Go testing + Testify + GoMock
- Docker: Docker, Docker Compose

---

## ローカル環境構築

### 1. 必要なもの
- Docker / Docker Compose
- Go 1.23 (ローカルでもビルドする場合)

### 2. 起動手順
```bash
# 環境変数を設定する（.envファイルを用意する）
cp .env.example .env

# Dockerイメージビルド・コンテナ起動
docker-compose up --build

# APIサーバー: http://localhost:8080
# DB: localhost:3306 (MySQL)
```

### 3. マイグレーション適用
```bash
# DBマイグレーションを実施（golang-migrate CLI使用）
migrate -path ./migrations -database "mysql://user:pass@tcp(localhost:3306)/dbname" up
```

---

## ディレクトリ構成 (予定)
```
├── cmd/
│   └── app/                 # エントリポイント（main関数など）、WireによるDI設定もここ
│       └── main.go
├── internal/                # アプリケーション内部のパッケージ（※外部からimportされない想定）
│   ├── product/             # 集約単位ごとにまとめる（例：product）
│   │   ├── domain/          # ドメイン層（エンティティ、値オブジェクト、リポジトリインターフェース）
│   │   │   ├── entity.go
│   │   │   ├── valueobject.go # enityに書いてもいい
│   │   │   └── repository.go
│   │   ├── usecase/         # ユースケース層（ビジネスロジック）
│   │   │   ├── interactor.go
│   │   │   └── dto.go
│   │   ├── infra/           # インフラ層（DB接続、リポジトリ実装）
│   │   │   ├── mysql_repository.go
│   │   │   └── db.go
│   │   └── handler/         # インターフェース層（HTTPハンドラ）
│   │       ├── controller.go
│   │       └── presenter.go
│   │── shared/             
│   │   ├── domain/          # ドメイン層（エンティティ、値オブジェクト、リポジトリインターフェース）
│   │   │   ├── entity.go
│   │   │   ├── valueobject.go # enityに書いてもいい
│   │   │   └── repository.go
│   │   ├── usecase/         # ユースケース層（ビジネスロジック）
│   │   │   ├── interactor.go
│   │   │   └── dto.go
│   │   ├── infra/           # インフラ層（DB接続、リポジトリ実装）
│   │   │   ├── mysql_repository.go
│   │   └── handler/         # インターフェース層（HTTPハンドラ）
│   │       ├── controller.go
│   │       └── presenter.go
│   ├── logger/              # ロガー（Zapラッパーなど）
│   │   └── logger.go
│   ├── config/              # コンフィグ管理（Viperを利用）
│   │   └── config.go
│   └── router/              # ルーティング設定（Echoルート管理）
│       └── router.go
├── db/
│   ├── migrations/          # マイグレーションSQL
│   ├── model/               # ORM定義（GORMの構造体など）
│   └── schema/              # （必要なら）テーブル定義をコードで表現したもの
├── scripts/                 # スクリプト（DBマイグレーションや初期データ投入など）
├── pkg/                     # 外部公開可能な汎用パッケージ（必要なら）
├── go.mod, go.sum           # Goモジュール定義（依存関係）
```

---

## デプロイ戦略

### Vercel (メイン)
- `/api/` ディレクトリに関数を配置してREST APIエンドポイントを作成。
- `vercel.json` でビルドフラグ、CORS設定などを管理。
- 環境変数はVercelのDashboardで設定。

### AWS (検証)
- DockerイメージをECRにpushし、ECS (Fargate) で実行。
- またはAWS Lambda + API Gatewayでサーバーレス運用検討。

---

## 開発方針まとめ
- ドメイン層（Entity, ValueObject）はインフラ・フレームワーク非依存にする
- 集約単位（例：product）ごとに、domain / usecase / infra / handler をまとめる
- 汎用コンポーネント（logger、config、router）はinternal直下に配置
- インターフェース層（Handler）はEchoに依存OK
- リポジトリはinterface定義 → infra層でGORM実装
- DIはWireで静的コード生成、グローバル変数禁止
- 設定情報はViper管理、環境別に切り替え可能
- ログはZapでJSON出力、フィールド付きログを積極活用
- ユニットテスト重視（モックはGoMock / Testify）

---

## TODO
- [ ] docker-compose.yml作成
- [ ] migrateツールセットアップ
- [ ] product配下の基本構成作成（domain / usecase / infra / handler）
- [ ] logger/config/router実装
- [ ] Vercelへの初回デプロイ動作確認

---

# ☕ 開発活動すすめていきましょう!


