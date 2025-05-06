# modules/alb/variables.tf

# ALB を配置する VPC の ID (network モジュールの出力を使用)
variable "vpc_id" {
  description = "The ID of the VPC to create the ALB in"
  type        = string
}

# ALB を配置するパブリックサブネットの ID のリスト (network モジュールの出力を使用)
# 複数のAZに跨るパブリックサブネットが必要です。
variable "public_subnet_ids" {
  description = "List of public subnet IDs for the ALB (should span multiple AZs)"
  type        = list(string)
}

# ALB に適用するセキュリティグループ ID のリスト (security_groups モジュールの出力を使用)
variable "security_group_ids" {
  description = "List of security group IDs to apply to the ALB"
  type        = list(string)
}

# ECS タスクがListenするポート (ターゲットグループのポートとして使用)
variable "app_container_port" {
  description = "Port the application container listens on (for ALB target group)"
  type        = number
}

# ALBのリスナーが待ち受けるポート (インターネットからのアクセスを受け付けるポート)
variable "alb_listener_port" {
  description = "Port the ALB listener listens on"
  type        = number
  default     = 80 # HTTPの場合
}

# ALBのリスナーが待ち受けるプロトコル (HTTP or HTTPS)
variable "alb_listener_protocol" {
  description = "Protocol the ALB listener uses (HTTP or HTTPS)"
  type        = string
  default     = "HTTP"
  validation {
    condition     = contains(["HTTP", "HTTPS"], upper(var.alb_listener_protocol))
    error_message = "Valid protocols are HTTP and HTTPS."
  }
}

# HTTPS を使用する場合の ACM 証明書の ARN (オプション)
# variable "ssl_certificate_arn" {
#   description = "The ARN of the ACM certificate to use for HTTPS listener"
#   type        = string
#   default     = null # HTTPのみの場合は不要
# }

# リソース名などに使うタグ (envs/dev から渡す)
variable "name" {
  description = "Name tag for resources"
  type        = string
}

variable "env" {
  description = "Environment tag for resources"
  type        = string
}


