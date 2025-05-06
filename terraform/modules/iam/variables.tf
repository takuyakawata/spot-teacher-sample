# modules/iam/variables.tf

# ロール名などに使うタグ (envs/dev から渡す)
variable "name" {
  description = "Name tag for resources"
  type        = string
}

variable "env" {
  description = "Environment tag for resources"
  type        = string
}

# ポリシー定義に必要なアカウント ID とリージョン (envs/dev でデータソースから取得し渡す)
variable "aws_account_id" {
  description = "The AWS Account ID"
  type        = string
}

variable "aws_region" {
  description = "The AWS Region"
  type        = string
}

# アプリケーションが読み取る Secrets Manager Secret の ID (envs/dev から渡す)
variable "db_password_secret_id" {
  description = "The Secrets Manager Secret ID that the task role needs permission to read"
  type        = string
}

# (オプション) アプリケーションがアクセスする必要がある他の AWS サービス ARN やリソース情報など
# variable "s3_bucket_arn" {
#   description = "ARN of the S3 bucket the application needs access to"
#   type        = string
#   default     = null
# }
