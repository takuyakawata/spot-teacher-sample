# Project Guidelines

ディレクトリ構成図
```
spot-teacher── api/
            │   └── index.go                 エントリポイント（Vercel用）
            ├── cmd/                         コマンドライン用ディレクトリ（現在は空）
            ├── internal/                    アプリケーション内部のパッケージ（※外部からimportされない想定）
            │   ├── company/                 会社関連の機能
            │   │   └── domain/              会社ドメインの定義
            │   ├── lesson/                  レッスン関連の機能
            │   │   └── domain/              レッスンドメインの定義
            │   ├── product/                 商品関連の機能
            │   │   ├── domain/              商品ドメインの定義
            │   │   ├── handler/             商品APIハンドラ（Echoのコントローラ）
            │   │   ├── infra/               商品インフラ層（リポジトリ実装など）
            │   │   ├── inject/              商品DIの設定（Wire用）
            │   │   └── usecase/             商品ユースケース実装
            │   │       └── mock/            ユースケースのモック（テスト用）
            │   ├── router/                  ルーティングの設定（Echoのルート定義など）
            │   ├── school/                  学校関連の機能
            │   │   └── domain/              学校ドメインの定義
            │   ├── shared/                  共有コンポーネント
            │   │   ├── domain/              共有ドメインの定義
            │   │   └── hander/              共有ハンドラ
            │   └── user/                    ユーザー関連の機能
            │       ├── domain/              ユーザードメインの定義
            │       ├── handler/             ユーザーAPIハンドラ
            │       ├── infra/               ユーザーインフラ層
            │       └── usecase/             ユーザーユースケース実装
            │           └── mock/            ユースケースのモック（テスト用）
            ├── go.mod, go.sum               Goモジュール定義（依存関係）
            └── main.go                      メインエントリポイント

   |
   DBー  
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
- ORM: ent.go (MySQL)
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
