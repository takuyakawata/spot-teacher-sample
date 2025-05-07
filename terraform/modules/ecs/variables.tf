# modules/ecs/variables.tf
variable "app_name" {
  description = "アプリケーションの名前 (リソース名などに使用)"
  type        = string
  default     = "spot-teaccher-dev"

}

variable "image_url" {
  description = "デプロイするDockerイメージのURL"
  type        = string
  default     = ""
}

variable "container_port" {
  description = "コンテナ内でGoアプリケーションがListenするポート番号"
  type        = number
  default     = 8080
}

variable "desired_count" {
  description = "実行したいタスクの数"
  type        = number
  default     = 1
}

variable "fargate_cpu" {
  description = "Fargateタスクに割り当てるCPUユニット (例: 256 (.25 vCPU), 512 (.5 vCPU), 1024 (1 vCPU), ...)"
  type        = number
  default     = 256
}

variable "fargate_memory" {
  description = "Fargateタスクに割り当てるメモリ (MB) (例: 512, 1024, 2048, ...)"
  type        = number
  default     = 512 # CPUとの組み合わせによって指定可能な値が異なります
}

variable "vpc_id" {
  description = "デプロイ先のVPC ID"
  type        = string
  default     = ""
}

variable "subnet_ids" {
  description = "Fargateタスクを実行するサブネットIDのリスト (通常はプライベートサブネット)"
  type        = list(string)
  default     = []
}

variable "security_group_ids" {
  description = "Fargateタスクに適用するセキュリティグループIDのリスト (ALBなどからのアクセスを許可する設定)"
  type        = list(string)
  default     = []
}

variable "alb_target_group_arn" {
  description = "連携するALBターゲットグループのARN"
  type        = string
  default     = ""
}

variable "db_password_secret_id" {
  description = "The Secrets Manager Secret ID containing the DB password"
  type        = string
}

variable "private_subnet_ids" {
  description = "Fargateタスクを実行するプライベートサブネットIDのリスト"
  type        = list(string)
  # default = [] # envs/dev から必ず渡すためデフォルト値は不要
}

# IAM モジュールから受け取るロール ARN
variable "task_execution_role_arn" {
  description = "The ARN of the ECS task execution role (from IAM module)"
  type        = string
}

variable "task_role_arn" {
  description = "The ARN of the application task role (from IAM module, if needed for AWS service access)"
  type        = string
  default     = null # アプリケーションにタスクロールが不要な場合 (Secrets Manager などにアクセスしない) は null
}


variable "name" {
  description = "Name tag for resources"
  type        = string
}

variable "env" {
  description = "Environment tag for resources"
  type        = string
}

# AWS リージョン (CloudWatch Logs の設定などで必要)
variable "region" {
  description = "AWS Region"
  type        = string
}

# データベース接続情報として ECS タスクの環境変数に渡す変数 (envs/dev から RDS モジュールの出力や envs/dev 変数を受け取る)
# ★ ここが重要！これらの変数定義が不足していました。
# variable "db_host" {
#   description = "Database host endpoint (from RDS module output)"
#   type        = string
# }
#
# variable "db_port" {
#   description = "Database port (from RDS module output)"
#   type        = number
# }
#
# variable "db_name" {
#   description = "Database name"
#   type        = string
# }
#
# variable "db_username" {
#   description = "Database username"
#   type        = string
# }
