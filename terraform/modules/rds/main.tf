## 注意!!: 開発環境用を想定して設定してあります。
# modules/rds/main.tf

# ==============================================================================
# Secrets Manager から DB パスワードを取得 (ReadOnly)
# ==============================================================================
# このデータソースは、既存のSecrets Managerシークレットの値を読み取ります。
# シークレット自体は、Terraform外で手動で作成するか、別のTerraformコードで管理することも可能です。
data "aws_secretsmanager_secret_version" "db_password" {
  secret_id = var.db_password_secret_id
}


# ==============================================================================
# DB サブネットグループの作成
# RDSインスタンスをどのサブネットに配置可能かを定義
# ==============================================================================
resource "aws_db_subnet_group" "main" {
  name       = "${var.name}-${var.env}-rds-subnet-group"
  subnet_ids = var.private_subnet_ids # プライベートサブネットのリストを使用
  description = "DB subnet group for RDS instance"

  tags = {
    Name = "${var.name}-${var.env}-rds-subnet-group"
    Env  = var.env
  }
}

# ==============================================================================
# RDS データベースインスタンスの作成
# ==============================================================================
resource "aws_rds_instance" "main" {
  identifier           = "${var.name}-${var.env}-rds-instance" # インスタンス識別子
  engine               = var.db_engine
  engine_version       = var.db_engine_version
  instance_class       = var.db_instance_type
  allocated_storage    = var.db_allocated_storage
  db_name              = var.db_name
  username             = var.db_username
  password             = data.aws_secretsmanager_secret_version.db_password.secret_string # Secrets Manager から取得したパスワードを使用
  db_subnet_group_name = aws_db_subnet_group.main.name # 作成したサブネットグループを指定
  vpc_security_group_ids = var.security_group_ids      # Security Groups モジュールから受け取った SG ID を指定

  multi_az             = var.multi_az          # Multi-AZ 設定
  deletion_protection  = var.deletion_protection # 削除保護設定

  skip_final_snapshot = !var.deletion_protection # 削除保護が無効な場合、終了時に最終スナップショットを取得しない (開発用)
  # 本番では true にして最終スナップショットを取得すべき

  # 認証情報は Secrets Manager を推奨するため、マスターパスワードの直接入力は避ける
  # username と password は Secrets Manager に含めるのが最も安全ですが、
  # ここではシンプルに username 変数と Secrets Manager から取得した password を組み合わせています。
  # より厳密には Secrets Manager にユーザー名とパスワードの両方をまとめて保存し、
  # アプリケーションから Secrets Manager SDK で取得するのが理想的です。

  tags = {
    Name = "${var.name}-${var.env}-rds-instance"
    Env  = var.env
  }
}
