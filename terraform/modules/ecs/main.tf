# ==============================================================================
# ECS クラスター
# ==============================================================================
resource "aws_ecs_cluster" "app_cluster" {
  name = "${var.app_name}-cluster"

  tags = {
    Name = "${var.app_name}-cluster"
  }
}

# ==============================================================================
# IAM ロール (タスク実行ロール)
# FargateでコンテナイメージのプルやCloudWatch Logsへのログ送信に必要
# ==============================================================================
resource "aws_iam_role" "task_execution_role" {
  name = "${var.app_name}-ecsTaskExecutionRole"

  # 信頼ポリシー: ECSタスクがこのロールを引き受けられるようにする
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Principal = {
        Service = "ecs-tasks.amazonaws.com"
      }
    }]
  })
}

# タスク実行ロールにポリシーをアタッチ (マネージドポリシーを使用)
resource "aws_iam_role_policy_attachment" "task_execution_policy" {
  role       = aws_iam_role.task_execution_role.name
  # Fargateに必要な基本的な権限を持つマネージドポリシー
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}


# ==============================================================================
# ECS タスク定義 (Fargate 用)
# ==============================================================================
resource "aws_ecs_task_definition" "app_task" {
  family                   = "${var.app_name}-task" # タスク定義ファミリー名
  requires_compatibilities = ["FARGATE"]            # Fargate 互換性を必須にする
  network_mode             = "awsvpc"               # Fargateではawsvpcモードが必須
  cpu                      = var.fargate_cpu        # CPU を変数から指定
  memory                   = var.fargate_memory     # メモリを変数から指定
  execution_role_arn       = aws_iam_role.task_execution_role.arn # タスク実行ロールを指定

  # コンテナ定義 (JSON形式で記述)
  container_definitions = jsonencode([{
    name      = "${var.app_name}-container" # コンテナ名
    image     = var.image_url               # 使用するDockerイメージのURL
    cpu       = var.fargate_cpu             # コンテナに割り当てるCPU (タスク定義と同じ値を指定することが多い)
    memory    = var.fargate_memory          # コンテナに割り当てるメモリ (タスク定義と同じ値を指定することが多い)
    essential = true                        # 必須コンテナ (これが停止するとタスク全体が停止)

    # ポートマッピング (コンテナ内でGoアプリがListenするポート)
    portMappings = [{
      containerPort = var.container_port
      protocol      = "tcp"
    }]

    # ヘルスチェック設定 (オプションだが推奨)
    # Goアプリが /health エンドポイントなどで応答する場合
    healthCheck = {
      command     = ["CMD-SHELL", "curl -f http://localhost:${var.container_port}/health || exit 1"]
      interval    = 30 # チェック間隔 (秒)
      timeout     = 5  # タイムアウト (秒)
      retries     = 3  # リトライ回数
      startPeriod = 0  # 起動直後の猶予期間 (秒)
    }

    # ログ設定 (CloudWatch Logsへ出力)
    logConfiguration = {
      logDriver = "awslogs"
      options = {
        "awslogs-group"         = "/ecs/${var.app_name}" # CloudWatch Logsのロググループ名 (事前に作成しておくか、Terraformで作成)
        "awslogs-region"        = var.region             # リージョン (versions.tf の region と合わせる)
        "awslogs-stream-prefix" = "ecs"                  # ログストリーム名のプレフィックス
      }
    }

    # 環境変数 (必要に応じて追加)
    environment = [
      {
        name  = "DB_HOST"
        # envs/dev/main.tf で ECS モジュール呼び出し時に RDS モジュールの出力を変数として渡す
        # 例: db_host = module.rds.rds_endpoint
        # ここでは、ECS モジュールの variables.tf に db_host を追加し、それを参照
        value = var.db_host
      },
      {
        name  = "DB_PORT"
        # 同様に envs/dev/main.tf で渡す
        value = var.db_port # または var.db_port 変数を ECS モジュール変数に追加
      },
      {
        name  = "DB_NAME"
        value = var.db_name # envs/dev の変数を使用
      },
      {
        name  = "DB_USER"
        value = var.db_username # envs/dev の変数を使用
      },
      # DB_PASSWORD は Secrets Manager から取得するように設定
      # Secrets Manager の ARN または名前を環境変数として渡し、
      # アプリケーションコード内でその環境変数を見て Secrets Manager から値を取得する
      {
        name  = "DB_PASSWORD_SECRET_ID" # アプリコードが参照する環境変数名
        value = var.db_password_secret_id # envs/dev の変数を使用
      },
      # 他のアプリケーション設定値など
      {
        name = "APP_ENV"
        value = var.env
      }
      # GITHUB_ACCESS_TOKEN や VERCEL_API_TOKEN がアプリに必要ならここに追加
      # ただし、これらの値も Secrets Manager などで管理し、環境変数として渡すことを推奨
      # {
      #   name = "GITHUB_ACCESS_TOKEN"
      #   value = "secretsmanager:your/github/token/secret/id" # ECS Exec Role に Secrets Manager 権限が必要
      # }
    ]
  }])

  tags = {
    Name = "${var.app_name}-task-definition"
  }
}

# ==============================================================================
# ECS サービス (タスクを指定数実行し、ALBと連携)
# ==============================================================================
resource "aws_ecs_service" "app_service" {
  name            = "${var.app_name}-service"     # サービス名
  cluster         = aws_ecs_cluster.app_cluster.id # 実行するクラスター
  task_definition = aws_ecs_task_definition.app_task.arn # 使用するタスク定義
  launch_type     = "FARGATE"                     # 起動タイプは Fargate
  desired_count   = var.desired_count             # 実行したいタスク数

  # ネットワーキング設定 (Fargate では必須)
  network_configuration {
    subnets         = var.subnet_ids         # 実行するサブネット
    security_groups = var.security_group_ids # 適用するセキュリティグループ
    assign_public_ip = false # ALBなどがプライベートサブネットにある場合は false。 ALBがパブリックサブネットにあり、Fargateタスクがプライベートサブネットにある構成が一般的。
  }

  # ロードバランサー連携設定
  load_balancer {
    target_group_arn = var.alb_target_group_arn # 連携するALBターゲットグループのARN
    container_name   = "${var.app_name}-container" # タスク定義で指定したコンテナ名
    container_port   = var.container_port          # タスク定義で指定したポート
  }

  # デプロイ設定 (オプション: ローリングアップデートなど)
  # deployment_controller {
  #   type = "ECS" # ECSがデプロイを制御 (デフォルト)
  # }

  # オートスケーリング設定 (オプション)
  # ... aws_appautoscaling_target, aws_appautoscaling_policy などを定義

  tags = {
    Name = "${var.app_name}-service"
  }

  # 他のリソースへの依存関係 (Terraformは自動で解決することが多いが、明示することも可能)
  # depends_on = [
  #   aws_iam_role_policy_attachment.task_execution_policy,
  #   # ALBリスナールールなど、サービスに依存する他のリソース
  # ]
}

# ==============================================================================
# 出力値 (デプロイ後に確認したい情報など)
# ==============================================================================


# ALBのDNS名などを出力する場合、ALBリソースもTerraformで管理する必要がある
# output "alb_dns_name" {
#   description = "ALB DNS Name"
#   value       = aws_lb.main_alb.dns_name # aws_lb リソースを定義している場合
# }
