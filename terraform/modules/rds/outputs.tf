# modules/rds/outputs.tf

# RDS インスタンスのエンドポイント (接続先ホスト名)
output "rds_endpoint" {
  description = "The endpoint of the RDS instance"
  value       = aws_rds_instance.main.endpoint
}

# RDS インスタンスのポート番号
output "rds_port" {
  description = "The port of the RDS instance"
  value       = aws_rds_instance.main.port
}

# RDS インスタンスの識別子
output "rds_identifier" {
  description = "The identifier of the RDS instance"
  value       = aws_rds_instance.main.identifier
}

# (オプション) マスターユーザーのユーザー名 (Secrets Manager に含める場合は不要かも)
# output "db_username" {
#   description = "The master username for the RDS instance"
#   value       = aws_rds_instance.main.username
# }

# Secrets Manager のシークレットARN (必要に応じて出力)
# output "db_password_secret_arn" {
#   description = "The ARN of the Secrets Manager secret containing the DB password"
#   value       = data.aws_secretsmanager_secret_version.db_password.arn
# }
