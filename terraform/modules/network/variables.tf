# modules/network/variables.tf

# VPC の設定
variable "vpc_cidr" {
  description = "CIDR block for the VPC"
  type        = string
}

# パブリックサブネットの設定 (AZごとに複数定義)
variable "public_subnet_cidrs" {
  description = "List of CIDR blocks for public subnets (must be one per AZ)"
  type        = list(string)
}

# プライベートサブネットの設定 (AZごとに複数定義)
variable "private_subnet_cidrs" {
  description = "List of CIDR blocks for private subnets (must be one per AZ)"
  type        = list(string)
}

# サブネットを配置するアベイラビリティゾーンのリスト (リージョン内で利用可能なAZ名)
variable "azs" {
  description = "List of Availability Zones to use"
  type        = list(string)
}

# リソース名などに使うタグ
variable "name" {
  description = "Name tag for resources"
  type        = string
}

variable "env" {
  description = "Environment tag for resources"
  type        = string
}
