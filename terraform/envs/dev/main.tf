# envs/dev/main.tf
# ==============================================================================
# Terraform および Provider の設定
# version.tf と分ける場合は、このブロックは不要です。
# version.tf に記述している場合は、envs/dev/ ディレクトリに versions.tf を置いてください。
# ==============================================================================
terraform {
  required_version = ">= 1.10.4" # Terraform version
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.87"
    }
  }
}

# Terraformの状態ファイル(tfstate)をリモートに保存する設定 (本番運用では必須)
# envs/dev/versions.tf に記述している場合も、そちらを優先してください。
# backend "s3" {
#   bucket = "your-terraform-state-bucket-name" # 事前に作成したS3バケット名
#   key    = "envs/dev/terraform.tfstate"      # tfstateファイルのパス
#   region = var.region                         # S3バケットがあるリージョン (変数参照可能)
#   dynamodb_table = "your-dynamodb-lock-table"  # 事前に作成したDynamoDBテーブル名 (状態ロック用、任意)
#   encrypt = true
# }

# AWSプロバイダーの設定
provider "aws" {
  profile = "spot-teacher-dev-account"
  region = var.region
  default_tags {
    tags = {
      team      = var.name
      env       = var.env
      managedBy = "terraform"
    }
  }
}

# ==============================================================================
# 各インフラ要素モジュールの呼び出しと連携
# 各モジュールの variables.tf で定義された入力を渡します。
# モジュールの outputs.tf で定義された出力を他のモジュールの入力として渡します。
# ==============================================================================
# --- データソース (現在のAWSアカウントIDとリージョンを取得) ---
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

module "iam" {
  source = "../../modules/iam"
  name = var.name
  env  = var.env
  aws_account_id = data.aws_caller_identity.current.account_id
  aws_region     = data.aws_region.current.name

  # DB パスワードの Secret ID を渡す
  db_password_secret_id = var.db_password_secret_id
}

module "network" {
  source = "../../modules/network" # modules/network ディレクトリへの相対パス

  # modules/network/variables.tf で定義した変数をここで指定
  vpc_cidr            = "10.0.0.0/16" # 例: 直接記述するか、envs/dev の変数にする
  public_subnet_cidrs = ["10.0.1.0/24", "10.0.2.0/24"] # 例: 直接記述するか、envs/dev の変数にする
  private_subnet_cidrs = ["10.0.10.0/24", "10.0.11.0/24"] # 例: 直接記述するか、envs/dev の変数にする
  azs                 = ["ap-northeast-1a", "ap-northeast-1c"] # 例: 直接記述するか、envs/dev の変数にする

  # タグ関連の変数 (envs/dev から受け取った変数を渡す)
  name = var.name
  env  = var.env
}

module "security_groups" {
  source = "../../modules/security_groups"
  vpc_id             = module.network.vpc_id

  # ポートやタグ関連の変数 (envs/dev の変数や値を渡す)
  app_container_port = var.container_port
  rds_port           = 5432

  name = var.name
  env  = var.env

  # Security Groups モジュールは Network モジュールに依存します
  depends_on = [module.network] # Network モジュールが作成されてから実行
}

module "alb" {
  source = "../../modules/alb"
  vpc_id            = module.network.vpc_id
  public_subnet_ids = module.network.public_subnet_ids
  security_group_ids = [module.security_groups.alb_security_group_id]
  # ポートやプロトコル、SSL 証明書など (envs/dev の変数や値を渡す)
  app_container_port      = var.container_port # ECSタスクがListenするポート
  alb_listener_port       = 80 # 例: HTTP 80番ポート
  alb_listener_protocol   = "HTTP" # 例: HTTP
  # ssl_certificate_arn   = null # HTTPS の場合、envs/dev の変数などを渡す
  name = var.name
  env  = var.env
  depends_on = [
    module.network,
    module.security_groups
  ]
}

module "ecs" {
  source = "../../modules/ecs"
  vpc_id             = module.network.vpc_id
  private_subnet_ids = module.network.private_subnet_ids
  # Security Group ID は security_groups モジュールの出力を参照！
  security_group_ids = [module.security_groups.ecs_security_group_id]
  # ALB Target Group ARN は alb モジュールの出力を参照！
  alb_target_group_arn = module.alb.alb_target_group_arn
  # Docker イメージ URL やタスク数、CPU/メモリなど (envs/dev の変数や値を渡す)
  app_name       = var.app_name
  image_url      = var.image_url # envs/dev の変数を使用
  container_port = var.container_port # envs/dev の変数を使用
  desired_count  = var.desired_count # envs/dev の変数を使用
  fargate_cpu    = var.fargate_cpu # envs/dev の変数を使用
  fargate_memory = var.fargate_memory # envs/dev の変数を使用
  name = var.name
  env  = var.env

  depends_on = [
    module.network,
    module.security_groups,
    module.alb,
    module.rds,
    module.iam,
  ]
  db_password_secret_id   = ""
  task_execution_role_arn = ""
  region                  = var.region
}

module "rds" {
  source = "../../modules/rds"
  security_group_ids = [module.security_groups.rds_security_group_id]

  # データベースの設定 (envs/dev の変数や値を渡す)

  # DB パスワードの Secret ID は envs/dev の変数や値を渡す
  # 例: 実際に作成した Secret ID に置き換えるか、envs/dev の変数にする

  # Multi-AZ や削除保護 (envs/dev の変数や値を渡す)
  # 開発環境なので false にするか、envs/dev の変数にする
  # 開発環境なので false にするか、envs/dev の変数にする

  # タグ関連の変数 (envs/dev から受け取った変数を渡す)
  name = var.name
  env  = var.env

  # RDS モジュールは Network モジュールと Security Groups モジュールに依存します
  depends_on = [
    module.network,
    module.security_groups,
    module.iam
  ]
  db_allocated_storage = var.db_allocated_storage
  engine_version      = var.db_engine_version
  monitoring_role_arn = module.iam.rds_monitoring_role_arn
  private_subnet_ids = module.network.private_subnet_ids
  instance_class      = var.db_instance_class
}
