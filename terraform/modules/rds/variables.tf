# ルートのvariablesから受け取る変数
# modules/rds/variables.tf

# RDS を配置する VPC の ID (network モジュールの出力を使用)
variable "vpc_id" {
  description = "The ID of the VPC to create the RDS instance in"
  type        = string
}

# RDS を配置するプライベートサブネットの ID のリスト (network モジュールの出力を使用)
# 複数指定することで Multi-AZ 配置の基盤となります。
variable "private_subnet_ids" {
  description = "List of private subnet IDs for the RDS instance (should span multiple AZs)"
  type        = list(string)
}

# RDS に適用するセキュリティグループ ID のリスト (security_groups モジュールの出力を使用)
variable "security_group_ids" {
  description = "List of security group IDs to apply to the RDS instance"
  type        = list(string)
}

# データベースエンジンの設定 (envs/dev から渡す)
variable "db_engine" {
  description = "RDS database engine (e.g., postgres, mysql)"
  type        = string
}

variable "db_engine_version" {
  description = "RDS database engine version"
  type        = string
}

variable "db_instance_type" {
  description = "RDS instance type"
  type        = string
}

variable "db_allocated_storage" {
  description = "RDS allocated storage (GB)"
  type        = number
}

# データベースの接続情報 (envs/dev から渡す)
variable "db_name" {
  description = "Database name"
  type        = string
}

variable "db_username" {
  description = "Database username"
  type        = string
}

# データベースパスワード (Secrets Manager から取得する場合の設定)
# envs/dev/terraform.tfvars で指定した Secrets Manager の Secret ID を受け取ります。
variable "db_password_secret_id" {
  description = "The Secrets Manager Secret ID containing the DB master user password"
  type        = string
}

# Multi-AZ 配置にするか (高可用性のため本番環境では推奨)
variable "multi_az" {
  description = "Whether to deploy as a Multi-AZ instance"
  type        = bool
  default     = false # 開発環境では false にすることが多い
}

# 削除保護を有効にするか (誤削除防止のため本番環境では推奨)
variable "deletion_protection" {
  description = "Whether to enable deletion protection for the RDS instance"
  type        = bool
  default     = false # 開発環境では false にすることが多い
}

# リソース名などに使うタグ (envs/dev から渡す)
variable "name" {
  description = "Name tag for resources"
  type        = string
}

variable "env" {
  description = "Environment tag for resources"
  type        = string
}
