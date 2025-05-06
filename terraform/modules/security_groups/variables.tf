# modules/security_groups/variables.tf

# Security Groups を作成する VPC の ID (network モジュールの出力を使用)
variable "vpc_id" {
  description = "The ID of the VPC to create security groups in"
  type        = string
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

# アプリケーションがListenするポート (ALBからECSへのインバウンド許可に使用)
variable "app_container_port" {
  description = "Port the application container listens on (for ECS SG ingress)"
  type        = number
}

# RDS がListenするポート (ECSからRDSへのインバウンド許可に使用)
variable "rds_port" {
  description = "Port the RDS database listens on (for RDS SG ingress)"
  type        = number
  default     = 5432 # PostgreSQLの場合。MySQLなど他のDBの場合は変更
}

# SSH ポート (Bastion ホストへのインバウンド許可に使用, オプション)
variable "ssh_port" {
  description = "Port for SSH access (for Bastion SG ingress)"
  type        = number
  default     = 22
}

# Bastion ホストへのアクセス元IPアドレスリスト (SSH許可に使用, オプション)
variable "bastion_ingress_cidrs" {
  description = "List of CIDR blocks allowed to SSH to the bastion host (e.g., your office IP)"
  type        = list(string)
  default     = [] # デフォルトではSSHアクセスを許可しない
}
