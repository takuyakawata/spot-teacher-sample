# modules/iam/outputs.tf

# ECS タスク実行ロールの ARN
output "task_execution_role_arn" {
  description = "The ARN of the ECS task execution role"
  value       = aws_iam_role.task_execution_role.arn
}

# アプリケーション用タスクロールの ARN
output "app_task_role_arn" {
  description = "The ARN of the application task role"
  value       = aws_iam_role.app_task_role.arn
}

# (オプション) 作成したポリシーの ARN など
# output "app_secretsmanager_policy_arn" {
#   description = "The ARN of the application Secrets Manager read policy"
#   value       = aws_iam_role_policy.app_secretsmanager_policy.arn
# }
