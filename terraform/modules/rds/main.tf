## 注意!!: 開発環境用を想定して設定してあります。
# modules/rds/main.tf

# ==============================================================================
# Secrets Manager から DB パスワードを取得 (ReadOnly)
# ==============================================================================
resource "aws_db_instance" "main" { # リソース名は "rds" ですね (私の例では "main" でした)
  # Engine options
  engine         = "mysql" # エンジン種類が直書き
  engine_version = var.engine_version # バージョンは変数
  # Multi-AZ
  multi_az = false # Multi-AZ 設定が直書き
  # settings
  identifier = "${var.name}-${var.env}" # 識別子は変数使用
  # Authentication
  username                    = "admin" # ユーザー名が直書き！これは良くないプラクティスです
  manage_master_user_password = true # AWSにパスワードを管理させるか、ここで指定するか
  # Instance settings
  instance_class = var.instance_class # インスタンスタイプは変数
  # Storage
  storage_type      = "gp3" # ストレージタイプが直書き
  allocated_storage = var.db_allocated_storage # ストレージ容量は変数
  # Connection
  vpc_security_group_ids = var.security_group_ids # SG ID は変数
  db_subnet_group_name   = aws_db_subnet_group.main.name # DBサブネットグループは参照
  publicly_accessible    = false # パブリックアクセス設定が直書き
  # PI
  performance_insights_enabled = false # 拡張モニタリング設定が直書き
  #EM
  monitoring_interval = 0 # モニタリング間隔が直書き
  monitoring_role_arn = null # モニタリングロールARNは変数 (新しい変数ですね)
  # Additional settings
  parameter_group_name       = aws_db_parameter_group.main.name # パラメータグループは参照
  option_group_name          = "default:mysql-8-0" # オプショングループ名が直書き
  backup_retention_period    = 1 # バックアップ保持期間が直書き
  backup_window              = "17:00-17:30" # バックアップウィンドウが直書き
  storage_encrypted          = true # ストレージ暗号化が直書き
  auto_minor_version_upgrade = false # マイナーバージョン自動アップグレード設定が直書き
  maintenance_window         = "Sat:18:00-Sat:18:30" # メンテナンスウィンドウが直書き
  tags = {
    Name = "${var.name}-${var.env}-db" # タグは変数使用
  }
  skip_final_snapshot = true # 終了スナップショット設定が直書き (私の例では変数に連動させていました)
  # Dependency
  depends_on = [
    aws_db_subnet_group.main,
    aws_db_parameter_group.main # 依存関係は参照
  ]
}
