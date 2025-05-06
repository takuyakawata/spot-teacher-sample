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
variable "subnet_ids" {
  type = list(string)
}
variable "instance_class" {
  type = string
}
variable "engine_version" {
  type = string
}

variable "allocated_storage" {
  type = number
}
variable "monitoring_role_arn" {
  type = string
}
