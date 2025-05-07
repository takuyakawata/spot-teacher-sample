# modules/iam/main.tf

data "aws_iam_policy_document" "ecs_tasks_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }
  }
}
# ==============================================================================
# IAM ロール (ECS タスク実行ロール) の作成
# Fargate がコンテナイメージのプル、CloudWatch Logs へのログ送信などを行うために必要
# これは以前 modules/ecs/main.tf に定義していたものを移動します。
# ==============================================================================
resource "aws_iam_role" "task_execution_role" {
  name = "${var.name}-${var.env}-ecsTaskExecutionRole"

  assume_role_policy = data.aws_iam_policy_document.ecs_tasks_assume_role.json

  tags = {
    Name = "${var.name}-${var.env}-ecsTaskExecutionRole"
  }
}

# タスク実行ロールに AWS 管理ポリシーをアタッチ
resource "aws_iam_role_policy_attachment" "task_execution_policy" {
  role       = aws_iam_role.task_execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

# (オプション) タスク実行ロールに Secrets Manager からパラメータを取得するための権限を追加
# アプリケーションが Secrets Manager から DB パスワードなどを取得する場合、
# アプリケーション用タスクロールだけでなく、タスク実行ロールにもこの権限が必要になることがあります。
# depends_on = [aws_iam_role.task_execution_role] # ロール作成後にポリシーをアタッチ
# resource "aws_iam_role_policy_attachment" "task_execution_secretsmanager_policy" {
#   role       = aws_iam_role.task_execution_role.name
#   policy_arn = "arn:aws:iam::aws:policy/SecretsManagerReadWrite" # より限定的なカスタムポリシーを推奨
# }


# ==============================================================================
# IAM ロール (アプリケーション用タスクロール) の作成
# アプリケーションコードが Secrets Manager などの AWS サービスにアクセスする際に引き受けるロール
# これが Secrets Manager から DB パスワードを取得する権限を持つロールです。
# ==============================================================================
resource "aws_iam_role" "app_task_role" {
  name = "${var.name}-${var.env}-appTaskRole"

  # このロールを ECS タスクが引き受けられるようにする信頼ポリシー
  assume_role_policy = data.aws_iam_policy_document.ecs_tasks_assume_role.json

  tags = {
    Name = "${var.name}-${var.env}-app-task-role"
  }
}

# ==============================================================================
# アプリケーション用タスクロールにアタッチするポリシー定義
# (Secrets Manager から特定のシークレットを読み取る権限)
# ==============================================================================
data "aws_iam_policy_document" "app_secretsmanager_read" {
  statement {
    effect = "Allow"
    actions = [
      "secretsmanager:GetSecretValue",
      "secretsmanager:DescribeSecret",
    ]

    resources = [
      # Secrets Manager のシークレットARNを指定 (アカウントIDとリージョン変数を使用)
      # ランダムサフィックスを含む ARN パターンを指定するため末尾に -* をつける
      "arn:aws:secretsmanager:${var.aws_region}:${var.aws_account_id}:secret:${var.db_password_secret_id}-*",
    ]
  }

  # (オプション) KMS 暗号化を使用している場合、KMSへの復号化権限
  # statement {
  #   effect = "Allow"
  #   actions = [ "kms:Decrypt" ]
  #   resources = [ "arn:aws:kms:${var.aws_region}:${var.aws_account_id}:key/YOUR_KMS_KEY_ID" ] # シークレットを暗号化しているKMSキーARN
  # }

  # (オプション) アプリケーションがアクセスする必要がある他の AWS サービスへの権限を追加 (S3, SQS など)
  # statement {
  #   effect = "Allow"
  #   actions = [ "s3:GetObject" ]
  #   resources = [ "${var.s3_bucket_arn}/*" ]
  # }
}

# ==============================================================================
# Secrets Manager 読み取りポリシーをアプリケーション用タスクロールにアタッチ
# ==============================================================================
resource "aws_iam_role_policy" "app_secretsmanager_policy" {
  name   = "${var.name}-${var.env}-appSecretsManagerPolicy"
  role   = aws_iam_role.app_task_role.id
  policy = data.aws_iam_policy_document.app_secretsmanager_read.json
}

# (オプション) 他に必要なポリシーをアプリケーション用タスクロールにアタッチ
# 例: SQS 送信権限、S3 書き込み権限など
# resource "aws_iam_role_policy_attachment" "app_sqs_policy" {
#   role       = aws_iam_role.app_task_role.name
#   policy_arn = "arn:aws:iam::aws:policy/AmazonSQSFullAccess" # 必要最小限のポリシーに変更
# }

# ==============================================================================
# RDS 拡張モニタリング用 IAM ロールが引き受けるための信頼ポリシー定義 (Data Source)
# RDS サービスプリンシパルからの引き受けを許可します。
# ==============================================================================
data "aws_iam_policy_document" "rds_monitoring_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["rds.amazonaws.com"] # ★ rds.amazonaws.com を信頼 ★
    }
  }
}

# ==============================================================================
# IAM ロール (RDS 拡張モニタリング用) の作成
# RDS が CloudWatch Logs にメトリクスを送信するために必要
# ==============================================================================
resource "aws_iam_role" "rds_monitoring_role" {
  name               = "${var.name}-${var.env}-rdsMonitoringRole" # ロール名は環境に合わせて設定

  # RDS サービスプリンシパルからの引き受けを許可する信頼ポリシー
  assume_role_policy = data.aws_iam_policy_document.rds_monitoring_assume_role.json

  tags = {
    Name = "${var.name}-${var.env}-rdsMonitoringRole"
  }
}

# ==============================================================================
# RDS 拡張モニタリング用ロールに AWS 管理ポリシーをアタッチ
# CloudWatch Logs への書き込み権限を付与
# ==============================================================================
resource "aws_iam_role_policy_attachment" "rds_monitoring_policy" {
  role       = aws_iam_role.rds_monitoring_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonRDSEnhancedMonitoringRole" # ★ 必須の管理ポリシー ★
}
