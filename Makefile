# 環境変数
ENV_FILE=.env
-include $(ENV_FILE)

# ループして各変数が定義されているか(空でないか)チェック
$(foreach var,$(REQUIRED_VARS),\
  $(if $(value $(var)),,$(error Required environment variable $(var) is not set. Please define it in $(ENV_FILE) or export it.))\
)

# --- 変数定義 ---
DATABASE_URL="mysql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?multiStatements=true"

MIGRATIONS_PATH=./db/migrations
GORM_GEN_TOOL=./db/tools/gormgen/main.go
ATLAS_MIGR_DIR=file://db/ent/migrate/migrations
ATLAS_TO_SCHEMA=ent://db/ent/schema
#ATLAS_TO_SCHEMA="ent://github.com/takuyakawta/spot-teacher-sample/db/ent/schema"
ATLAS_DEV_URL="docker://mysql/8/ent"

# --- ターゲット定義 ---
.PHONY: help db-create db-drop migrate-up migrate-down migrate-force gorm-gen ent-generate migrate-diff migrate-apply


help: ## このヘルプメッセージを表示
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'


migrate-force: ## 特定バージョンのマイグレーション状態に強制変更 (開発時、失敗した場合など)
	@echo "Forcing migration version (use with caution)..."
	@read -p "強制変更したいバージョン番号を入力してください: " version && \
	migrate -database "${DATABASE_URL}" -path ${MIGRATIONS_PATH} force $${version}
	@echo "Migration version forced."


db-drop: ## データベースを削除 (注意して実行！)
	@read -p "本当にデータベース ${DB_NAME} を削除しますか？ (y/N): " confirm && [ $${confirm:-N} = y ] || exit 1
	@echo "Dropping database ${DB_NAME}..."
	@mysql -h ${DB_HOST} -P ${DB_PORT} -u ${DB_USER} --password=${DB_PASSWORD} -e "DROP DATABASE IF EXISTS ${DB_NAME};"
	@echo "Database dropped."


migrate-diff: ## ent スキーマと DB の差分から SQL マイグレーションファイルを生成します (例: make migrate-diff MIGRATION_NAME=create_users)
	@if [ -z "${MIG}" ] || [ "${MIG}" = "migration" ]; then \
		echo "エラー: MIGRATIONのNAME を指定してください (例: make migrate-diff MIG=my_migration)"; \
		exit 1; \
	fi
	@echo "Generating migration file named '${MIG}'..."
	@#cd ./db & GOWORK=off atlas migrate diff ${MIG} --dir ${ATLAS_MIGR_DIR} --to ${ATLAS_TO_SCHEMA} --dev-url ${ATLAS_DEV_URL}
	@cd ./db & GOWORK=off atlas migrate diff ${MIG} --dir ${ATLAS_MIGR_DIR} --to ${ATLAS_TO_SCHEMA} --dev-url ${ATLAS_DEV_URL}
	@echo "Migration file generated."


migrate-apply: ## 保留中のマイグレーションをデータベースに適用します
	@echo "Applying migrations to ${DB_NAME}..."
	@atlas migrate apply --url ${DATABASE_URL} --dir ${ATLAS_MIGR_DIR}
	@echo "Migrations applied."


#ent-generate: ## ent/schema/*.go の定義から Go コードを生成します
#	@echo "Generating ent code..."
#	@go generate ./db/ent
#	@echo "ent code generated."


ent-describe: ## ent スキーマ定義の内容をテキストで表示します
	@echo "Describing ent schema from ${ENT_SCHEMA_PATH}..."
	@go run -mod=mod entgo.io/ent/cmd/ent describe ${ENT_SCHEMA_PATH}
	@echo "Schema description finished."


migrate-status: ## マイグレーションの状態を表示します
	@echo "Checking migration status..."
	@atlas migrate status --url "${DATABASE_URL}" --dir ${ATLAS_MIGR_DIR}
	@echo "Migration status checked."


migrate-hash: ## atalas.sumの更新
	@echo "Updating atlas.sum..."
	@atlas migrate hash --dir ${ATLAS_MIGR_DIR}
	@echo "atlas.sum updated."



#db-create: ## データベースを作成 (MySQLサーバーに接続して実行、存在する場合はエラーになる可能性あり)
#	@echo "Creating database ${DB_NAME}..."
#	@mysql -h ${DB_HOST} -P ${DB_PORT} -u ${DB_USER} --password=${DB_PASSWORD} -e "CREATE DATABASE IF NOT EXISTS ${DB_NAME};"
#	@echo "Database created (if not existed)."

#migrate-up: ## マイグレーションを実行 (最新まで適用)
#	@echo "Applying migrations..."
#	@migrate -database ${DATABASE_URL} -path ${MIGRATIONS_PATH} up
#	@echo "Migrations applied."

#migrate-down: ## マイグレーションを1つロールバック
#	@echo "Rolling back last migration..."
#	@migrate -database "${DATABASE_URL}" -path ${MIGRATIONS_PATH} down 1
#	@echo "Last migration rolled back."

#gorm-gen: ## GORM Gen を実行してコードを生成
#	@echo "Running GORM Gen..."
#	@go run ${GORM_GEN_TOOL}
#	@echo "GORM Gen finished."


# === Docker Commands ===
up: ## Dockerコンテナをバックグラウンドで起動します (docker-compose up -d)
	@echo "Starting Docker containers..."
	@docker-compose up -d
	@echo "Docker containers started."

down: ## Dockerコンテナを停止し、削除します (docker-compose down)
	@echo "Stopping and removing Docker containers..."
	@docker-compose down
	@echo "Docker containers stopped and removed."

docker-logs: ## Dockerコンテナのログを表示します (Ctrl+C で停止)
	@echo "Following Docker logs (Ctrl+C to stop)..."
	@docker-compose logs -f db # 'db' は docker-compose.yml 内のサービス名に合わせてください

docker-rebuild: ## Dockerコンテナを再ビルドして起動します
	@echo "Rebuilding and restarting Docker containers..."
	@docker-compose down
	@docker-compose up -d --build
	@echo "Docker containers rebuilt and started."


# === Go Application Commands ===
run: run-a ## (デフォルト) APIサーバーA を実行します (go run)
run-a:
	@echo "Running API server A..."
	@go run ${CMD_PATH_A}

run-b: ## APIサーバーB を実行します (go run)
	@echo "Running API server B..."
	@go run ${CMD_PATH_B}

build: build-a build-b ## 全てのAPIサーバーのバイナリをビルドします
	@echo "All application binaries built in ${OUTPUT_DIR}/"

build-a: tidy ent-generate ## APIサーバーA のバイナリをビルドします
	@echo "Building ${APP_NAME_A}..."
	@mkdir -p ${OUTPUT_DIR}
	@go build -o ${OUTPUT_DIR}/${APP_NAME_A} ${CMD_PATH_A}
	@echo "Built: ${OUTPUT_DIR}/${APP_NAME_A}"

build-b: tidy ent-generate ## APIサーバーB のバイナリをビルドします
	@echo "Building ${APP_NAME_B}..."
	@mkdir -p ${OUTPUT_DIR}
	@go build -o ${OUTPUT_DIR}/${APP_NAME_B} ${CMD_PATH_B}
	@echo "Built: ${OUTPUT_DIR}/${APP_NAME_B}"

test: ## Go のテストを実行します (./... は全パッケージ対象)
	@echo "Running Go tests..."
	@go test ./... -v

tidy: ## go.mod と go.sum ファイルを整理します
	@echo "Running go mod tidy..."
	@go mod tidy
	@echo "go mod tidy finished."
