# ルートのvariablesから受け取る変数
variable "name" {
  type = string
}
variable "env" {
  type = string
}
variable "security_group_ids" {
  type = list(string)
}
variable "private_subnet_ids" {
  description = "List of private subnet IDs for the RDS instance (should span multiple AZs)"
  type        = list(string)
}
variable "instance_class" {
  type = string
}
variable "engine_version" {
  type = string
}

variable "db_allocated_storage" {
  type = number
  default     = 20
}
variable "monitoring_role_arn" {
  type = string
}
