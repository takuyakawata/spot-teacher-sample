# envs/dev/variables.tf
variable "env" {
  default = "dev"
}

variable "name" {
  default = "spot-teacher-dev"
}

# AWSリージョン
variable "region" {
  description = "AWS Region"
  type        = string
  default     = "ap-northeast-1" # デフォルト値としてリージョンを指定することも可能
}

# アプリケーション名 (リソース名のプレフィックスなどに使用)
variable "app_name" {
  description = "Application name"
  type        = string
  default     = "spot-teacher-dev"
}

# デプロイするDockerイメージのURL
variable "image_url" {
  description = "Docker image URL to deploy"
  type        = string
  default     = ""
}

# ECSタスクの数
variable "desired_count" {
  description = "Number of desired ECS tasks"
  type        = number
  default     = 1
}

# FargateのCPUとメモリ
variable "fargate_cpu" {
  description = "Fargate task CPU units"
  type        = number
  default     = 256
}

variable "fargate_memory" {
  description = "Fargate task memory (MB)"
  type        = number
  default     = 512
}

# コンテナ内でアプリがListenするポート
variable "container_port" {
  description = "Port the application container listens on"
  type        = number
}

# RDSの設定 (最低限必要なもの)
variable "db_instance_type" {
  description = "RDS instance type"
  type        = string
  default     = "db.t3.micro" # 開発環境向けの小さめのインスタンスタイプ
}

variable "db_allocated_storage" {
  description = "RDS allocated storage (GB)"
  type        = number
  default     = 20
}

variable "db_engine" {
  description = "RDS database engine"
  type        = string
  default     = "postgres" # 例: postgres, mysql
}

variable "db_engine_version" {
  description = "RDS database engine version"
  type        = string
  default     = "15.5" # 使用するエンジンに合わせて適切なバージョンを指定
}

variable "db_name" {
  description = "Database name"
  type        = string
}

variable "db_username" {
  description = "Database username"
  type        = string
}

# RDS パスワードの Secrets Manager Secret ID
# この変数は、env/dev/terraform.tfvars で具体的な値を設定し、
# main.tf で他のモジュール (IAM, ECS) に渡すために使用します。
variable "db_password_secret_id" {
  description = "The Secrets Manager Secret ID containing the DB master user password"
  type        = string
}

# db_password は Sensitive な情報なので、直接 tfvars に書くより、
# 環境変数、または AWS Systems Manager Parameter Store や AWS Secrets Manager を使うべきです。
# ここでは例として変数定義だけ示し、tfvars では示しません。
# variable "db_password" {
#   description = "Database password"
#   type        = string
#   sensitive   = true
# }
# ← ここが重要！ envs/dev/main.tf から db_host, db_port, db_name, db_username, db_password_secret_id を渡すために必要です
variable "db_host" {
  description = "Database host endpoint (from RDS module output)"
  type        = string
}

variable "db_port" {
  description = "Database port (from RDS module output)"
  type        = number
}
