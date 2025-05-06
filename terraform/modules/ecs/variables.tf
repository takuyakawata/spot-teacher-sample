# modules/ecs/variables.tf

variable "app_name" {
  description = "アプリケーションの名前 (リソース名などに使用)"
  type        = string
  default     = "spot-teaccher-app" # アプリケーション名

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
